/*
 * Tencent is pleased to support the open source community by making TKEStack
 * available.
 *
 * Copyright (C) 2012-2019 Tencent. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use
 * this file except in compliance with the License. You may obtain a copy of the
 * License at
 *
 * https://opensource.org/licenses/Apache-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OF ANY KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations under the License.
 */

package storage

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"reflect"
	"strconv"
	"time"

	"k8s.io/apimachinery/pkg/api/errors"
	metainternal "k8s.io/apimachinery/pkg/apis/meta/internalversion"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/runtime"
	genericregistry "k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/registry/generic/registry"
	"k8s.io/apiserver/pkg/registry/rest"
	"k8s.io/apiserver/pkg/storage"
	storageerr "k8s.io/apiserver/pkg/storage/errors"
	"k8s.io/apiserver/pkg/util/dryrun"
	"tkestack.io/tke/api/business"
	businessinternalclient "tkestack.io/tke/api/client/clientset/internalversion/typed/business/internalversion"
	platformversionedclient "tkestack.io/tke/api/client/clientset/versioned/typed/platform/v1"
	platformv1 "tkestack.io/tke/api/platform/v1"
	"tkestack.io/tke/pkg/apiserver/authentication"
	apiserverutil "tkestack.io/tke/pkg/apiserver/util"
	"tkestack.io/tke/pkg/business/registry/namespace"
	"tkestack.io/tke/pkg/business/util"
	"tkestack.io/tke/pkg/util/log"
)

const _rsaKeyBits = 2048
const _defaultCertValidDays = 365

// Storage includes storage for namespace and all sub resources.
type Storage struct {
	Namespace   *REST
	Status      *StatusREST
	Finalize    *FinalizeREST
	Certificate *CertificateREST
}

// NewStorage returns a Storage object that will work against namespace sets.
func NewStorage(optsGetter genericregistry.RESTOptionsGetter, businessClient *businessinternalclient.BusinessClient, platformClient platformversionedclient.PlatformV1Interface, privilegedUsername string) *Storage {
	strategy := namespace.NewStrategy(businessClient, platformClient)
	store := &registry.Store{
		NewFunc:                  func() runtime.Object { return &business.Namespace{} },
		NewListFunc:              func() runtime.Object { return &business.NamespaceList{} },
		DefaultQualifiedResource: business.Resource("namespaces"),
		PredicateFunc:            namespace.MatchNamespace,

		CreateStrategy: strategy,
		UpdateStrategy: strategy,
		DeleteStrategy: strategy,

		ShouldDeleteDuringUpdate: shouldDeleteDuringUpdate,
	}
	store.TableConvertor = rest.NewDefaultTableConvertor(store.DefaultQualifiedResource)
	options := &genericregistry.StoreOptions{
		RESTOptions: optsGetter,
		AttrFunc:    namespace.GetAttrs,
	}

	if err := store.CompleteWithOptions(options); err != nil {
		log.Panic("Failed to create namespace etcd rest storage", log.Err(err))
	}

	statusStore := *store
	statusStore.UpdateStrategy = namespace.NewStatusStrategy(strategy)

	finalizeStore := *store
	finalizeStore.UpdateStrategy = namespace.NewFinalizeStrategy(strategy)

	certificateStore := *store

	return &Storage{
		Namespace:   newREST(store, platformClient, privilegedUsername),
		Status:      &StatusREST{&statusStore},
		Finalize:    &FinalizeREST{&finalizeStore},
		Certificate: &CertificateREST{&certificateStore, platformClient, privilegedUsername},
	}
}

// ValidateGetObjectAndTenantID validate name and tenantID, if success return Namespace
func ValidateGetObjectAndTenantID(ctx context.Context, store *registry.Store, name string, options *metav1.GetOptions) (runtime.Object, error) {
	obj, err := store.Get(ctx, name, options)
	if err != nil {
		return nil, err
	}

	o := obj.(*business.Namespace)
	if err := util.FilterNamespace(ctx, o); err != nil {
		return nil, err
	}
	return o, nil
}

// REST implements a RESTStorage for namespace sets against etcd.
type REST struct {
	*registry.Store
	platformClient     platformversionedclient.PlatformV1Interface
	privilegedUsername string
}

func newREST(store *registry.Store, platformClient platformversionedclient.PlatformV1Interface, privilegedUsername string) *REST {
	return &REST{
		Store:              store,
		platformClient:     platformClient,
		privilegedUsername: privilegedUsername,
	}
}

var _ rest.ShortNamesProvider = &REST{}

// ShortNames implements the ShortNamesProvider interface. Returns a list of short names for a resource.
func (r *REST) ShortNames() []string {
	return []string{"ns"}
}

// DeleteCollection selects all resources in the storage matching given 'listOptions'
// and deletes them.
func (r *REST) DeleteCollection(ctx context.Context, deleteValidation rest.ValidateObjectFunc, options *metav1.DeleteOptions, listOptions *metainternal.ListOptions) (runtime.Object, error) {
	if !authentication.IsAdministrator(ctx, r.privilegedUsername) {
		return nil, errors.NewMethodNotSupported(business.Resource("namespaces"), "delete collection")
	}
	return r.Store.DeleteCollection(ctx, deleteValidation, options, listOptions)
}

// List selects resources in the storage which match to the selector. 'options' can be nil.
func (r *REST) List(ctx context.Context, options *metainternal.ListOptions) (runtime.Object, error) {
	wrappedOptions := apiserverutil.PredicateListOptions(ctx, options)
	obj, err := r.Store.List(ctx, wrappedOptions)
	if err == nil && obj != nil {
		if err := r.patchNamespaceList(ctx, obj); err != nil {
			return nil, err
		}
	}
	return obj, err
}

// Get retrieves the object from the storage. It is required to support Patch.
func (r *REST) Get(ctx context.Context, name string, options *metav1.GetOptions) (runtime.Object, error) {
	obj, err := ValidateGetObjectAndTenantID(ctx, r.Store, name, options)
	if err == nil && obj != nil {
		if err := r.patchNamespace(ctx, obj, nil); err != nil {
			return nil, err
		}
	}
	return obj, err
}

// Update alters the object subset of an object.
func (r *REST) Update(ctx context.Context, name string, objInfo rest.UpdatedObjectInfo, createValidation rest.ValidateObjectFunc, updateValidation rest.ValidateObjectUpdateFunc, forceAllowCreate bool, options *metav1.UpdateOptions) (runtime.Object, bool, error) {
	// We are explicitly setting forceAllowCreate to false in the call to the underlying storage because
	// subresources should never allow create on update.
	_, err := ValidateGetObjectAndTenantID(ctx, r.Store, name, &metav1.GetOptions{})
	if err != nil {
		return nil, false, err
	}
	return r.Store.Update(ctx, name, objInfo, createValidation, updateValidation, false, options)
}

// Delete enforces life-cycle rules for cluster termination
func (r *REST) Delete(ctx context.Context, name string, deleteValidation rest.ValidateObjectFunc, options *metav1.DeleteOptions) (runtime.Object, bool, error) {
	nsObj, err := ValidateGetObjectAndTenantID(ctx, r.Store, name, &metav1.GetOptions{})
	if err != nil {
		return nil, false, err
	}

	ns := nsObj.(*business.Namespace)

	// Ensure we have a UID precondition
	if options == nil {
		options = metav1.NewDeleteOptions(0)
	}
	if options.Preconditions == nil {
		options.Preconditions = &metav1.Preconditions{}
	}
	if options.Preconditions.UID == nil {
		options.Preconditions.UID = &ns.UID
	} else if *options.Preconditions.UID != ns.UID {
		err = errors.NewConflict(
			business.Resource("namespaces"),
			name,
			fmt.Errorf("precondition failed: UID in precondition: %v, UID in object meta: %v", *options.Preconditions.UID, ns.UID),
		)
		return nil, false, err
	}
	if options.Preconditions.ResourceVersion != nil && *options.Preconditions.ResourceVersion != ns.ResourceVersion {
		err = errors.NewConflict(
			business.Resource("namespaces"),
			name,
			fmt.Errorf("precondition failed: ResourceVersion in precondition: %v, ResourceVersion in object meta: %v", *options.Preconditions.ResourceVersion, ns.ResourceVersion),
		)
		return nil, false, err
	}

	if ns.DeletionTimestamp.IsZero() {
		key, err := r.Store.KeyFunc(ctx, name)
		if err != nil {
			return nil, false, err
		}

		preconditions := storage.Preconditions{UID: options.Preconditions.UID, ResourceVersion: options.Preconditions.ResourceVersion}

		out := r.Store.NewFunc()
		err = r.Store.Storage.GuaranteedUpdate(
			ctx, key, out, false, &preconditions,
			storage.SimpleUpdate(func(existing runtime.Object) (runtime.Object, error) {
				existingNamespace, ok := existing.(*business.Namespace)
				if !ok {
					// wrong type
					return nil, fmt.Errorf("expected *business.Namespace, got %v", existing)
				}
				if err := deleteValidation(ctx, existingNamespace); err != nil {
					return nil, err
				}
				// Set the deletion timestamp if needed
				if existingNamespace.DeletionTimestamp.IsZero() {
					now := metav1.Now()
					existingNamespace.DeletionTimestamp = &now
				}
				// Set the namespace phase to terminating, if needed
				if existingNamespace.Status.Phase != business.NamespaceTerminating {
					existingNamespace.Status.Phase = business.NamespaceTerminating
				}

				// the current finalizers which are on namespace
				currentFinalizers := map[string]bool{}
				for _, f := range existingNamespace.Finalizers {
					currentFinalizers[f] = true
				}
				// the finalizers we should ensure on namespace
				shouldHaveFinalizers := map[string]bool{
					metav1.FinalizerOrphanDependents: apiserverutil.ShouldHaveOrphanFinalizer(options, currentFinalizers[metav1.FinalizerOrphanDependents]),
					metav1.FinalizerDeleteDependents: apiserverutil.ShouldHaveDeleteDependentsFinalizer(options, currentFinalizers[metav1.FinalizerDeleteDependents]),
				}
				// determine whether there are changes
				changeNeeded := false
				for finalizer, shouldHave := range shouldHaveFinalizers {
					changeNeeded = currentFinalizers[finalizer] != shouldHave || changeNeeded
					if shouldHave {
						currentFinalizers[finalizer] = true
					} else {
						delete(currentFinalizers, finalizer)
					}
				}
				// make the changes if needed
				if changeNeeded {
					var newFinalizers []string
					for f := range currentFinalizers {
						newFinalizers = append(newFinalizers, f)
					}
					existingNamespace.Finalizers = newFinalizers
				}
				return existingNamespace, nil
			}),
			dryrun.IsDryRun(options.DryRun),
			nil,
		)

		if err != nil {
			err = storageerr.InterpretGetError(err, business.Resource("namespaces"), name)
			err = storageerr.InterpretUpdateError(err, business.Resource("namespaces"), name)
			if _, ok := err.(*errors.StatusError); !ok {
				err = errors.NewInternalError(err)
			}
			return nil, false, err
		}

		return out, false, nil
	}

	// prior to final deletion, we must ensure that finalizers is empty
	if len(ns.Spec.Finalizers) != 0 {
		err = errors.NewConflict(business.Resource("namespaces"), ns.Name, fmt.Errorf("the system is ensuring all content is removed from this namespace.  Upon completion, this namespace will automatically be purged by the system"))
		return nil, false, err
	}
	return r.Store.Delete(ctx, name, deleteValidation, options)
}

func (r *REST) patchNamespaceList(ctx context.Context, obj runtime.Object) error {
	nl, ok := obj.(*business.NamespaceList)
	if !ok {
		return fmt.Errorf("patchNamespaceList, expect *business.NamespaceList, but got %s", reflect.TypeOf(obj))
	}

	cache := map[string]*platformv1.Cluster{}
	for idx := range nl.Items {
		if err := r.patchNamespace(ctx, &nl.Items[idx], cache); err != nil {
			return err
		}
	}
	return nil
}

func (r *REST) patchNamespace(ctx context.Context, obj runtime.Object, cache map[string]*platformv1.Cluster) error {
	ns, ok := obj.(*business.Namespace)
	if !ok {
		return fmt.Errorf("patchNamespace, expect *business.Namespace, but got %s", reflect.TypeOf(obj))
	}

	cls, has := cache[ns.Spec.ClusterName]
	if !has {
		var err error
		cls, err = r.platformClient.Clusters().Get(ctx, ns.Spec.ClusterName, metav1.GetOptions{})
		if err != nil {
			log.Errorf("patchNamespace %s: %s", ns.Name, err)
			return nil
		}
		if cache != nil {
			cache[ns.Spec.ClusterName] = cls
		}
	}
	ns.Spec.ClusterType = cls.Spec.Type
	ns.Spec.ClusterVersion = cls.Status.Version
	ns.Spec.ClusterDisplayName = cls.Spec.DisplayName
	return nil
}

// StatusREST implements the REST endpoint for changing the status of a replication controller
type StatusREST struct {
	store *registry.Store
}

// StatusREST implements Patcher
var _ = rest.Patcher(&StatusREST{})

// New returns an empty object that can be used with Create and Update after
// request data has been put into it.
func (r *StatusREST) New() runtime.Object {
	return r.store.New()
}

// Get retrieves the object from the storage. It is required to support Patch.
func (r *StatusREST) Get(ctx context.Context, name string, options *metav1.GetOptions) (runtime.Object, error) {
	return ValidateGetObjectAndTenantID(ctx, r.store, name, options)
}

// Update alters the status subset of an object.
func (r *StatusREST) Update(ctx context.Context, name string, objInfo rest.UpdatedObjectInfo, createValidation rest.ValidateObjectFunc, updateValidation rest.ValidateObjectUpdateFunc, forceAllowCreate bool, options *metav1.UpdateOptions) (runtime.Object, bool, error) {
	// We are explicitly setting forceAllowCreate to false in the call to the underlying storage because
	// subresources should never allow create on update.
	_, err := ValidateGetObjectAndTenantID(ctx, r.store, name, &metav1.GetOptions{})
	if err != nil {
		return nil, false, err
	}
	return r.store.Update(ctx, name, objInfo, createValidation, updateValidation, false, options)
}

// FinalizeREST implements the REST endpoint for finalizing a namespace.
type FinalizeREST struct {
	store *registry.Store
}

// New returns an empty object that can be used with Create and Update after
// request data has been put into it.
func (r *FinalizeREST) New() runtime.Object {
	return r.store.New()
}

// Get retrieves the object from the storage. It is required to support Patch.
func (r *FinalizeREST) Get(ctx context.Context, name string, options *metav1.GetOptions) (runtime.Object, error) {
	return ValidateGetObjectAndTenantID(ctx, r.store, name, options)
}

// Update alters the status finalizers subset of an object.
func (r *FinalizeREST) Update(ctx context.Context, name string, objInfo rest.UpdatedObjectInfo, createValidation rest.ValidateObjectFunc, updateValidation rest.ValidateObjectUpdateFunc, forceAllowCreate bool, options *metav1.UpdateOptions) (runtime.Object, bool, error) {
	_, err := ValidateGetObjectAndTenantID(ctx, r.store, name, &metav1.GetOptions{})
	if err != nil {
		return nil, false, err
	}
	return r.store.Update(ctx, name, objInfo, createValidation, updateValidation, false, options)
}

// CertificateREST implements the REST endpoint for getting a x509 certificate for namespaces.
type CertificateREST struct {
	store              *registry.Store
	platformClient     platformversionedclient.PlatformV1Interface
	privilegedUsername string
}

// New returns an empty object that can be used with Create and Update after
// request data has been put into it.
func (r *CertificateREST) New() runtime.Object {
	return r.store.New()
}

func (r *CertificateREST) NewGetOptions() (runtime.Object, bool, string) {
	return &business.NamespaceCertOptions{ValidDays: strconv.Itoa(_defaultCertValidDays)}, false, ""
}

// Get retrieves the namespace from the storage and patch a x509 certificate.
func (r *CertificateREST) Get(ctx context.Context, name string, options runtime.Object) (runtime.Object, error) {
	obj, err := newREST(r.store, r.platformClient, r.privilegedUsername).Get(ctx, name, &metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	ns := obj.(*business.Namespace)

	cluster, err := r.platformClient.Clusters().Get(ctx, ns.Spec.ClusterName, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("prj:%s, ns:%s, get cluster %s, %s", ns.Namespace, ns.Name, ns.Spec.ClusterName, err)
	}
	if cluster.Spec.Type == "Imported" {
		return nil, fmt.Errorf("prj:%s, ns:%s, cluster %s is Imported, NOT support generating certificate", ns.Namespace, ns.Name, ns.Spec.ClusterName)
	}
	if len(cluster.Status.Addresses) == 0 {
		return nil, fmt.Errorf("prj:%s, ns:%s, cluster %s has NO valid addresses", ns.Namespace, ns.Name, ns.Spec.ClusterName)
	}
	fieldSelector := fields.OneTermEqualSelector("clusterName", ns.Spec.ClusterName).String()
	list, err := r.platformClient.ClusterCredentials().List(ctx, metav1.ListOptions{FieldSelector: fieldSelector})
	if err != nil {
		return nil, fmt.Errorf("prj:%s, ns:%s, get cluster credential, %s", ns.Namespace, ns.Name, err)
	} else if len(list.Items) == 0 {
		return nil, fmt.Errorf("prj:%s, ns:%s, no cluster credential", ns.Namespace, ns.Name)
	}
	credential := list.Items[0]
	certBlock, _ := pem.Decode(credential.CACert)
	if certBlock == nil {
		return nil, fmt.Errorf("prj:%s, ns:%s, pem decode root cert error, bytes:%v", ns.Namespace, ns.Name, credential.CACert)
	}
	if certBlock.Type != "CERTIFICATE" {
		return nil, fmt.Errorf("prj:%s, ns:%s, pem decode root cert, invalid type %s", ns.Namespace, ns.Name, certBlock.Type)
	}
	rootCert, err := x509.ParseCertificate(certBlock.Bytes)
	if err != nil {
		return nil, fmt.Errorf("prj:%s, ns:%s, parse root cert, %s, bytes:%v", ns.Namespace, ns.Name, err, certBlock.Bytes)
	}
	keyBlock, _ := pem.Decode(credential.CAKey)
	if keyBlock == nil {
		return nil, fmt.Errorf("prj:%s, ns:%s, pem decode root key error, bytes:%v", ns.Namespace, ns.Name, credential.CAKey)
	}
	if keyBlock.Type != "RSA PRIVATE KEY" {
		return nil, fmt.Errorf("prj:%s, ns:%s, pem decode root key, invalid type %s", ns.Namespace, ns.Name, keyBlock.Type)
	}
	rootKey, err := x509.ParsePKCS1PrivateKey(keyBlock.Bytes)
	if err != nil {
		return nil, fmt.Errorf("prj:%s, ns:%s, parse root key, %s, bytes:%v", ns.Namespace, ns.Name, err, keyBlock.Bytes)
	}

	private, err := rsa.GenerateKey(rand.Reader, _rsaKeyBits)
	if err != nil {
		return nil, fmt.Errorf("prj:%s, ns:%s, generate prive key, %s", ns.Namespace, ns.Name, err)
	}
	validDays, err := strconv.Atoi(options.(*business.NamespaceCertOptions).ValidDays)
	if err != nil {
		return nil, fmt.Errorf("prj:%s, ns:%s, query string '%s', %s", ns.Namespace, ns.Name, business.CertOptionValidDays, err)
	}
	user, _ := authentication.UsernameAndTenantID(ctx)
	template := x509.Certificate{
		Subject: pkix.Name{
			CommonName: user,
			Organization: []string{
				fmt.Sprintf("cluster:%s", ns.Spec.ClusterName),
				fmt.Sprintf("project:%s", ns.Namespace),
				fmt.Sprintf("namespace:%s", ns.Spec.Namespace),
				fmt.Sprintf("tenant:%s", ns.Spec.TenantID),
			},
		},
		SerialNumber:          big.NewInt(rootCert.SerialNumber.Int64() + 1),
		NotBefore:             rootCert.NotBefore,
		NotAfter:              time.Now().AddDate(0, 0, validDays),
		BasicConstraintsValid: true,
		IsCA:                  false,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageDataEncipherment,
	}
	certBytes, err := x509.CreateCertificate(rand.Reader, &template, rootCert, &private.PublicKey, rootKey)
	if err != nil {
		return nil, fmt.Errorf("CreateCertificate(%+v), %s", template.Subject, err)
	}
	keyBytes := x509.MarshalPKCS1PrivateKey(private)

	address := cluster.Status.Addresses[0]
	for _, one := range cluster.Status.Addresses {
		if one.Type == "Advertise" {
			address = one
			break
		}
	}
	ns.Status.Certificate = &business.NamespaceCert{
		CertPem: pem.EncodeToMemory(&pem.Block{
			Type:  "CERTIFICATE",
			Bytes: certBytes,
		}),
		KeyPem: pem.EncodeToMemory(&pem.Block{
			Type:  "PRIVATE KEY",
			Bytes: keyBytes,
		}),
		CACertPem: credential.CACert,
		APIServer: fmt.Sprintf("https://%s:%d", address.Host, address.Port),
	}
	return ns, nil
}

// EmigrateREST implements the REST endpoint for moving a namespace.
type EmigrateREST struct {
	store *registry.Store
}

// New returns an empty object that can be used with Create and Update after
// request data has been put into it.
func (r *EmigrateREST) New() runtime.Object {
	return r.store.New()
}

// Get retrieves the object from the storage. It is required to support Patch.
func (r *EmigrateREST) Get(ctx context.Context, name string, options *metav1.GetOptions) (runtime.Object, error) {
	return ValidateGetObjectAndTenantID(ctx, r.store, name, options)
}

// Update alters the Status.Phase and Annotations of a namespace.
func (r *EmigrateREST) Update(ctx context.Context, name string, objInfo rest.UpdatedObjectInfo,
	createValidation rest.ValidateObjectFunc, updateValidation rest.ValidateObjectUpdateFunc,
	forceAllowCreate bool, options *metav1.UpdateOptions) (runtime.Object, bool, error) {
	_, err := ValidateGetObjectAndTenantID(ctx, r.store, name, &metav1.GetOptions{})
	if err != nil {
		return nil, false, err
	}
	// We are explicitly setting forceAllowCreate to false in the call to the underlying storage because
	return r.store.Update(ctx, name, objInfo, createValidation, updateValidation, false, options)
}

func shouldDeleteDuringUpdate(ctx context.Context, key string, obj, existing runtime.Object) bool {
	ns, ok := obj.(*business.Namespace)
	if !ok {
		log.Errorf("unexpected object, key:%s", key)
		return false
	}
	return len(ns.Spec.Finalizers) == 0 && registry.ShouldDeleteDuringUpdate(ctx, key, obj, existing)
}
