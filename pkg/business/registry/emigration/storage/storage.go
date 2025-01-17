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
	"fmt"

	"k8s.io/apimachinery/pkg/api/errors"
	metainternal "k8s.io/apimachinery/pkg/apis/meta/internalversion"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	genericregistry "k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/registry/generic/registry"
	"k8s.io/apiserver/pkg/registry/rest"
	"tkestack.io/tke/api/business"
	businessinternalclient "tkestack.io/tke/api/client/clientset/internalversion/typed/business/internalversion"
	platformversionedclient "tkestack.io/tke/api/client/clientset/versioned/typed/platform/v1"
	"tkestack.io/tke/pkg/apiserver/authentication"
	apiserverutil "tkestack.io/tke/pkg/apiserver/util"
	"tkestack.io/tke/pkg/business/registry/emigration"
	"tkestack.io/tke/pkg/business/util"
	"tkestack.io/tke/pkg/util/log"
)

// Storage includes storage for emigration and all sub resources.
type Storage struct {
	Emigration *REST
}

// NewStorage returns a Storage object that will work against emigration sets.
func NewStorage(optsGetter genericregistry.RESTOptionsGetter, businessClient *businessinternalclient.BusinessClient,
	platformClient platformversionedclient.PlatformV1Interface, privilegedUsername string) *Storage {
	strategy := emigration.NewStrategy(businessClient, platformClient)
	store := &registry.Store{
		NewFunc:                  func() runtime.Object { return &business.NsEmigration{} },
		NewListFunc:              func() runtime.Object { return &business.NsEmigrationList{} },
		DefaultQualifiedResource: business.Resource("nsemigrations"),
		PredicateFunc:            emigration.MatchNsEmigration,

		CreateStrategy: strategy,
		UpdateStrategy: strategy,
		DeleteStrategy: strategy,

		ShouldDeleteDuringUpdate: registry.ShouldDeleteDuringUpdate,
	}
	store.TableConvertor = rest.NewDefaultTableConvertor(store.DefaultQualifiedResource)
	options := &genericregistry.StoreOptions{
		RESTOptions: optsGetter,
		AttrFunc:    emigration.GetAttrs,
	}

	if err := store.CompleteWithOptions(options); err != nil {
		log.Panic("Failed to create emigration etcd rest storage", log.Err(err))
	}

	return &Storage{Emigration: newREST(store, platformClient, privilegedUsername)}
}

// ValidateGetObjectAndTenantID validate name and tenantID, if success return NsEmigration
func ValidateGetObjectAndTenantID(ctx context.Context, store *registry.Store, name string, options *metav1.GetOptions) (runtime.Object, error) {
	obj, err := store.Get(ctx, name, options)
	if err != nil {
		return nil, err
	}

	o := obj.(*business.NsEmigration)
	if err := util.FilterNsEmigration(ctx, o); err != nil {
		return nil, err
	}
	return o, nil
}

// REST implements a RESTStorage for emigration sets against etcd.
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
	return []string{"nse"}
}

// DeleteCollection selects all resources in the storage matching given 'listOptions'
// and deletes them.
func (r *REST) DeleteCollection(ctx context.Context, deleteValidation rest.ValidateObjectFunc,
	options *metav1.DeleteOptions, listOptions *metainternal.ListOptions) (runtime.Object, error) {
	if !authentication.IsAdministrator(ctx, r.privilegedUsername) {
		return nil, errors.NewMethodNotSupported(business.Resource("nsemigrations"), "delete collection")
	}
	return r.Store.DeleteCollection(ctx, deleteValidation, options, listOptions)
}

// List selects resources in the storage which match to the selector. 'options' can be nil.
func (r *REST) List(ctx context.Context, options *metainternal.ListOptions) (runtime.Object, error) {
	wrappedOptions := apiserverutil.PredicateListOptions(ctx, options)
	return r.Store.List(ctx, wrappedOptions)
}

// Get retrieves the object from the storage. It is required to support Patch.
func (r *REST) Get(ctx context.Context, name string, options *metav1.GetOptions) (runtime.Object, error) {
	return ValidateGetObjectAndTenantID(ctx, r.Store, name, options)
}

// Update alters the object subset of an object.
func (r *REST) Update(ctx context.Context, name string, objInfo rest.UpdatedObjectInfo,
	createValidation rest.ValidateObjectFunc, updateValidation rest.ValidateObjectUpdateFunc,
	forceAllowCreate bool, options *metav1.UpdateOptions) (runtime.Object, bool, error) {
	// We are explicitly setting forceAllowCreate to false in the call to the underlying storage because
	// subresources should never allow create on update.
	_, err := ValidateGetObjectAndTenantID(ctx, r.Store, name, &metav1.GetOptions{})
	if err != nil {
		return nil, false, err
	}
	return r.Store.Update(ctx, name, objInfo, createValidation, updateValidation, false, options)
}

// Delete enforces life-cycle rules for emigration termination
func (r *REST) Delete(ctx context.Context, name string, deleteValidation rest.ValidateObjectFunc,
	options *metav1.DeleteOptions) (runtime.Object, bool, error) {
	nsObj, err := ValidateGetObjectAndTenantID(ctx, r.Store, name, &metav1.GetOptions{})
	if err != nil {
		return nil, false, err
	}

	em := nsObj.(*business.NsEmigration)

	if em.Status.Phase != business.NsEmigrationFinished && em.Status.Phase != business.NsEmigrationFailed {
		return nil, false, fmt.Errorf("emigration is in %s phase, only %s or %s emigrations can be deleted",
			em.Status.Phase, business.NsEmigrationFinished, business.NsEmigrationFailed)
	}

	// Ensure we have a UID precondition
	if options == nil {
		options = metav1.NewDeleteOptions(0)
	}
	if options.Preconditions == nil {
		options.Preconditions = &metav1.Preconditions{}
	}
	if options.Preconditions.UID == nil {
		options.Preconditions.UID = &em.UID
	} else if *options.Preconditions.UID != em.UID {
		err = errors.NewConflict(
			business.Resource("nsemigrations"),
			name,
			fmt.Errorf("precondition failed: UID in precondition: %v, UID in object meta: %v", *options.Preconditions.UID, em.UID),
		)
		return nil, false, err
	}
	if options.Preconditions.ResourceVersion != nil && *options.Preconditions.ResourceVersion != em.ResourceVersion {
		err = errors.NewConflict(
			business.Resource("nsemigrations"),
			name,
			fmt.Errorf("precondition failed: ResourceVersion in precondition: %v, ResourceVersion in object meta: %v",
				*options.Preconditions.ResourceVersion, em.ResourceVersion),
		)
		return nil, false, err
	}

	return r.Store.Delete(ctx, name, deleteValidation, options)
}
