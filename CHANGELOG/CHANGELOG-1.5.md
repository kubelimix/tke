
# [1.5.0](https://github.com/tkestack/tke/compare/v1.4.0...v1.5.0) (2020-12-07)


### Bug Fixes

* **application:** enable to disable registry client ([#944](https://github.com/tkestack/tke/issues/944)) ([688fe23](https://github.com/tkestack/tke/commit/688fe23090e9d160aec54083020c2433e8611348))
* **ci:** fix build failed issue ([#917](https://github.com/tkestack/tke/issues/917)) ([ec82532](https://github.com/tkestack/tke/commit/ec8253234dca8d792ddff4ffac07fa7062450f35))
* **cluster:** check port ([#910](https://github.com/tkestack/tke/issues/910)) ([627ed57](https://github.com/tkestack/tke/commit/627ed5794e2179e5efedd91cb865a682ee7cbf93))
* **console:** 1.5.0 alpha 2020 12 03 ([#972](https://github.com/tkestack/tke/issues/972)) ([522ed60](https://github.com/tkestack/tke/commit/522ed60da5ecb0672dd18707c745c3b46320d7e3))
* **console:** 1.5.0-alpha 2020-12-04 ([#974](https://github.com/tkestack/tke/issues/974)) ([e003a53](https://github.com/tkestack/tke/commit/e003a53ef619a7fd5eb9740f8fe6e6a5d467c517))
* **console:** 1.5.0-alpha bug fix ([#959](https://github.com/tkestack/tke/issues/959)) ([f11ef5e](https://github.com/tkestack/tke/commit/f11ef5ecb954a8b1f54b0d0dadea9b9f03d96122))
* **console:** add cluster version is empty case ([#957](https://github.com/tkestack/tke/issues/957)) ([6a4002e](https://github.com/tkestack/tke/commit/6a4002e1812badfd26d5944c468328e75f23399c))
* **console:** change metadata.labels to spec.selector.matchlabels ([#955](https://github.com/tkestack/tke/issues/955)) ([5211c56](https://github.com/tkestack/tke/commit/5211c56f2a2e2a9397ab96361a3f45a5a5a74167))
* **console:** clear pod name query not reset pagesize ([#983](https://github.com/tkestack/tke/issues/983)) ([bd11ad4](https://github.com/tkestack/tke/commit/bd11ad4410defdfbfb1d76c5665b4a00d965e66f))
* **console:** enable to delete node taint ([#931](https://github.com/tkestack/tke/issues/931)) ([fdda1ad](https://github.com/tkestack/tke/commit/fdda1ade4d4958d500f7467f2e879f6d1eafed9c))
* **console:** fix config promethus params error ([#924](https://github.com/tkestack/tke/issues/924)) ([58bb539](https://github.com/tkestack/tke/commit/58bb539092e1f4dda3770d03279e27e334317237))
* **console:** fix webtty tips no hidden ([#940](https://github.com/tkestack/tke/issues/940)) ([8c692dc](https://github.com/tkestack/tke/commit/8c692dcf8b966f433935d025c7e36d7aa7993a0a))
* **console:** master update ;pod search ;worker update ([#968](https://github.com/tkestack/tke/issues/968)) ([311bf67](https://github.com/tkestack/tke/commit/311bf67ccccd8b5117bce9572afbfad46d972173))
* **console:** nfs path enable use url and memory use no-cache ([#934](https://github.com/tkestack/tke/issues/934)) ([d7b7a46](https://github.com/tkestack/tke/commit/d7b7a46d2748295d42af681776c7adbf5df295b6))
* **console:** spec for notification channel creation fixed ([#979](https://github.com/tkestack/tke/issues/979)) ([2e45a9e](https://github.com/tkestack/tke/commit/2e45a9ee4c8fcdc637b6f4bb4cc9a76756bf8901))
* **console:** support terminating status and evicted reason show ([#977](https://github.com/tkestack/tke/issues/977)) ([9ce96c3](https://github.com/tkestack/tke/commit/9ce96c365a7a81025cd57d64b476533e8030a446))
* **console:** update the doc of webkubectl ([#933](https://github.com/tkestack/tke/issues/933)) ([20b6651](https://github.com/tkestack/tke/commit/20b6651b3e0c25ac20443d13d298f8b897eec9cd))
* **console:** use pod name query can't find next page resource ([#982](https://github.com/tkestack/tke/issues/982)) ([a35f015](https://github.com/tkestack/tke/commit/a35f0152d12f336005a8023602d74f9139b53d24))
* **console:** webhook notification template fixed ([#976](https://github.com/tkestack/tke/issues/976)) ([afeb2da](https://github.com/tkestack/tke/commit/afeb2da812a0bee9bd5a3e357603825829bbef4b))
* **console:** when update notify disable qudao ([#984](https://github.com/tkestack/tke/issues/984)) ([e7f1d6f](https://github.com/tkestack/tke/commit/e7f1d6f5c41563afec883428c9b11e2a6f69db2a))
* **doc:** refine control plane enable scale doc ([99ee25f](https://github.com/tkestack/tke/commit/99ee25f057ead082e1af07caa5879acd9f2aa69e))
* **installer:** get flag from shell ([#965](https://github.com/tkestack/tke/issues/965)) ([1937ecf](https://github.com/tkestack/tke/commit/1937ecfc3436d0308a5d93a4f13698ebc89bd906))
* **monitor:** fixed request and limit statistics ([#916](https://github.com/tkestack/tke/issues/916)) ([6c5d109](https://github.com/tkestack/tke/commit/6c5d109a01f667ad1d1573daff34b08f5c826d24))
* **platform:** avoid duplicated and out of order iptables rule ([#973](https://github.com/tkestack/tke/issues/973)) ([75bc233](https://github.com/tkestack/tke/commit/75bc233d5d9cc987c0f391d95d5275cb21feab52))
* **platform:** enhance cluster controller ([#937](https://github.com/tkestack/tke/issues/937)) ([8a393c8](https://github.com/tkestack/tke/commit/8a393c8e9d830db388642208ef02460943f2a30b))
* **platform:** fix master scale issue ([#919](https://github.com/tkestack/tke/issues/919)) ([b6301d9](https://github.com/tkestack/tke/commit/b6301d9c9d4951ed151df7c683ae9e271491086d))
* **platform:** idempotent problem ([#925](https://github.com/tkestack/tke/issues/925)) ([7d9eeb8](https://github.com/tkestack/tke/commit/7d9eeb87888b88b8da48d5db91dc769e0ffc379a))
* **platform:** only validate ha for mater scale case ([#943](https://github.com/tkestack/tke/issues/943)) ([f033396](https://github.com/tkestack/tke/commit/f0333962e52025c42e1ef7b445bace948d50189a))
* **platform:** patch annotation according to k8s version ([#969](https://github.com/tkestack/tke/issues/969)) ([7ac0421](https://github.com/tkestack/tke/commit/7ac04217f6c89927e1e86e7d10b3b54ddad5cb58))
* **platform:** path annotation to all master ([#945](https://github.com/tkestack/tke/issues/945)) ([885050e](https://github.com/tkestack/tke/commit/885050e323a47ecf033e0b0c99944fa03ccdd428))
* **platform:** persistent envent flaw ([#960](https://github.com/tkestack/tke/issues/960)) ([7588fb4](https://github.com/tkestack/tke/commit/7588fb437ad593f4b45ee33f0ea15bcd3e4fee5c))
* **platform:** set gpu-manager and gpu-quota-admission log to stderr ([#967](https://github.com/tkestack/tke/issues/967)) ([203f5c1](https://github.com/tkestack/tke/commit/203f5c13fcc548ce949d24ffc41c389b1fac18ab))
* **platform:** upgrade reconcile time out ([#954](https://github.com/tkestack/tke/issues/954)) ([333673d](https://github.com/tkestack/tke/commit/333673daab18806a8d352ede3f9762c4e76625c2))
* ensure instances delete operations to be executed ([#938](https://github.com/tkestack/tke/issues/938)) ([7dcd2c5](https://github.com/tkestack/tke/commit/7dcd2c57c2fd15c3853b221a2f975ca5b58dacbd))
* **auth:** ha not work for auth webhook ([51d67a1](https://github.com/tkestack/tke/commit/51d67a1b1b510ee8d451aa43ea70d7aa0b9dfe5f))
* **auth:** update console reminder for ha ([f7fe3d2](https://github.com/tkestack/tke/commit/f7fe3d2f090076692dee501f4e9ea3d56e7a4fac))
* **auth:** update doc for ha ([00bd983](https://github.com/tkestack/tke/commit/00bd983f9de3fdea58fbf4c46a0a5c1acf32642f))
* **builder:** push arm64/v8 to docker hub ([#841](https://github.com/tkestack/tke/issues/841)) ([d8eb0b0](https://github.com/tkestack/tke/commit/d8eb0b069a595d85ba8cbc6c7511a91830498c8c))
* **business:** remove first member from projects' members in portal ([#821](https://github.com/tkestack/tke/issues/821)) ([3b67f2d](https://github.com/tkestack/tke/commit/3b67f2d3b525705ebfaeecc54f5114d0287525c9))
* **cluster:** anno or label modified but not updated ([#824](https://github.com/tkestack/tke/issues/824)) ([c28e89a](https://github.com/tkestack/tke/commit/c28e89a9db260ef97df1a7490cb44e7c6170b4ed))
* **cluster:** enqueue to workqueque no immedutely ([#874](https://github.com/tkestack/tke/issues/874)) ([abb5256](https://github.com/tkestack/tke/commit/abb5256fca899bfa15e4a8d4ec865bd0353160df))
* **cluster:** fix cronhpa edit problem ([#798](https://github.com/tkestack/tke/issues/798)) ([c41fbab](https://github.com/tkestack/tke/commit/c41fbab3b1ea6e068ebfd0c3420d82dc65f0ab44))
* **cluster:** fix hpa and cronhpa problem ([34d9355](https://github.com/tkestack/tke/commit/34d935542f7379cc7c330d610d5db73af5a1fa3b))
* **cluster:** hpa details can not show bug ([#834](https://github.com/tkestack/tke/issues/834)) ([1368f27](https://github.com/tkestack/tke/commit/1368f2772fe20c9322f602fef1cb18070738c705))
* **cluster:** lost apiserver cert ([#875](https://github.com/tkestack/tke/issues/875)) ([#888](https://github.com/tkestack/tke/issues/888)) ([c4fd4b3](https://github.com/tkestack/tke/commit/c4fd4b3feed0724b91fb4599673927f14d421785))
* **cluster:** resourceversion update not need to enqueue ([#847](https://github.com/tkestack/tke/issues/847)) ([682e263](https://github.com/tkestack/tke/commit/682e263f5f031f5e1ba202568fadfa3fb4930f40))
* **console:** kubeconfig api server path display error ([#830](https://github.com/tkestack/tke/issues/830)) ([7881ef2](https://github.com/tkestack/tke/commit/7881ef2e86d3ec192c4790c8cceeff2285cbec9e))
* **console:** logstash link path fixed ([#858](https://github.com/tkestack/tke/issues/858)) ([d6967ec](https://github.com/tkestack/tke/commit/d6967ec9e0a74da20ecd936ca1ceb76281701a23))
* **doc:** use correct name ([#803](https://github.com/tkestack/tke/issues/803)) ([2e08b9d](https://github.com/tkestack/tke/commit/2e08b9d0cfc525824bc6e9e98b2600a9f8a6505b))
* **gateway:** fix null-pointer panic when no notify registered ([#817](https://github.com/tkestack/tke/issues/817)) ([ad0cf63](https://github.com/tkestack/tke/commit/ad0cf63a7d6e015432758d5d80792fa333b6abf8))
* **installer:** fix babel runtime error ([#849](https://github.com/tkestack/tke/issues/849)) ([9b8eaff](https://github.com/tkestack/tke/commit/9b8eaff01334e9f4fc9af54db7efdf96e9e43797))
* **installer:** fix empty pointer ([#801](https://github.com/tkestack/tke/issues/801)) ([dec8c7c](https://github.com/tkestack/tke/commit/dec8c7c1962d81a028d258e22e8c21e2f5352d5d))
* **installer:** fix installer log can't auto scroll than 4170 line ([#831](https://github.com/tkestack/tke/issues/831)) ([fa7b0df](https://github.com/tkestack/tke/commit/fa7b0dfbf187d64b092a535ebe1f504e803c074c))
* **installer:** mark status as retrying rather failed when a step failed ([#650](https://github.com/tkestack/tke/issues/650)) ([82f3ef9](https://github.com/tkestack/tke/commit/82f3ef9b6008366ce160406b2dfd9b261b75d05a))
* **installer:** tke-gateway pods pull image failed by all-one ([#825](https://github.com/tkestack/tke/issues/825)) ([19fac0d](https://github.com/tkestack/tke/commit/19fac0d020e3254ca4ef9ff8df57fc976016789b))
* **monitor:** send alert to correct webhook addr ([#794](https://github.com/tkestack/tke/issues/794)) ([9e212bd](https://github.com/tkestack/tke/commit/9e212bd510d4723eb9f243e3e10bed3cf0a01046))
* **notify:** fix notify x509 certificate error ([#843](https://github.com/tkestack/tke/issues/843)) ([bd19ccf](https://github.com/tkestack/tke/commit/bd19ccf3131a80c966b81616f8f4a07c54d23d44))
* **platform:** allow user specify node name for cloud provider case ([#827](https://github.com/tkestack/tke/issues/827)) ([8489974](https://github.com/tkestack/tke/commit/8489974b91a0f3de8953b2e274766eec31df9f8b))
* **platform:** compatible webhook's certificate and private key ([#891](https://github.com/tkestack/tke/issues/891)) ([1fffb2f](https://github.com/tkestack/tke/commit/1fffb2f398fc8d274686529693eff44085db0a3c))
* **platform:** don't allow service cidr is empty when enable ipvs ([#846](https://github.com/tkestack/tke/issues/846)) ([62ea1d5](https://github.com/tkestack/tke/commit/62ea1d53473b551e087c57a8030daa4e1a6f1ebe))
* **platform:** remove flag when renew certs ([#802](https://github.com/tkestack/tke/issues/802)) ([9ed4ab9](https://github.com/tkestack/tke/commit/9ed4ab9833c6c4f67ee9efd582051580d5937954))
* **platform:** set owner reference leverage GC delete object ([#850](https://github.com/tkestack/tke/issues/850)) ([ca4c664](https://github.com/tkestack/tke/commit/ca4c66469c5b11fe8236c751a77223ff92aa233b))
* **platform:** support devcloud ([#882](https://github.com/tkestack/tke/issues/882)) ([d2e7353](https://github.com/tkestack/tke/commit/d2e7353309748766246c6299a8564009e18df221))
* **platform,monitor:** check prometheus status until it recovers ([#873](https://github.com/tkestack/tke/issues/873)) ([094d91e](https://github.com/tkestack/tke/commit/094d91e6ea3d143eff365058ca6f9307e01fcd31))
* **platform,monitor:** use extension-apiserver-authentication cm ([#877](https://github.com/tkestack/tke/issues/877)) ([e71406a](https://github.com/tkestack/tke/commit/e71406a82f7a3c568beedd6fd68551124ed256a3))
* **registry:** fix prepare registry certificate panic ([#923](https://github.com/tkestack/tke/issues/923)) ([aa5da78](https://github.com/tkestack/tke/commit/aa5da78b20b7d0fd55a64f594fdd47d8cb1260b2))
* **uam:** open admin pwd modify ([#880](https://github.com/tkestack/tke/issues/880)) ([2c96d0e](https://github.com/tkestack/tke/commit/2c96d0ea30674c408e898fbc2cc1a7c04758638a))
* comment out one job ([#890](https://github.com/tkestack/tke/issues/890)) ([1c2cfc7](https://github.com/tkestack/tke/commit/1c2cfc716b67fcdab191d778bccf8a6353b9c5a5))
* remove the condition to build image ([#881](https://github.com/tkestack/tke/issues/881)) ([5ba7897](https://github.com/tkestack/tke/commit/5ba7897011a0adb11fe64bbccdc6085bf9e2b7dc))
* update e2e workflow platform ([#895](https://github.com/tkestack/tke/issues/895)) ([d16f368](https://github.com/tkestack/tke/commit/d16f368fc5685724638bc39bc2aac47abb29d663))


### Features

* **auth:** limit action of changing user's password ([#928](https://github.com/tkestack/tke/issues/928)) ([bbf5123](https://github.com/tkestack/tke/commit/bbf5123e73d066b4e5c40d22b15752605846fa14))
* **auth:** optimize the performance of localidentity list interface ([#728](https://github.com/tkestack/tke/issues/728)) ([fac27db](https://github.com/tkestack/tke/commit/fac27db6c30741a19d486187d2b3ae9240196d9d))
* **auth:** update load casbin model from rule ([#767](https://github.com/tkestack/tke/issues/767)) ([9f3c6b5](https://github.com/tkestack/tke/commit/9f3c6b505cecd118acd656fd6bb935527f8754e7))
* **ci:** add Node.js version check ([#845](https://github.com/tkestack/tke/issues/845)) ([d58640e](https://github.com/tkestack/tke/commit/d58640ea1db54641b4849430d12bcf65f51ff56f))
* **ci:** build provider res in ci ([#912](https://github.com/tkestack/tke/issues/912)) ([45b2fcd](https://github.com/tkestack/tke/commit/45b2fcdffb740187d5fcac1c2054763764904f64))
* **console:** add alarm node memory used ([#956](https://github.com/tkestack/tke/issues/956)) ([a33d9cc](https://github.com/tkestack/tke/commit/a33d9cc1345df5ba669497020a0ff9fc1d097d53))
* **console:** add readme ([#889](https://github.com/tkestack/tke/issues/889)) ([0beb2ee](https://github.com/tkestack/tke/commit/0beb2ee67ac1853c8b8d7e30a026d4a9616ee1ad))
* **console:** adjust kubectl dialog ([#953](https://github.com/tkestack/tke/issues/953)) ([ca31bc3](https://github.com/tkestack/tke/commit/ca31bc3004d1c0d93c0dfef5a4bdede973041227))
* **console:** can config when open promethus ([#905](https://github.com/tkestack/tke/issues/905)) ([25dca34](https://github.com/tkestack/tke/commit/25dca3428c192bf90442cfc236b4d93506ea6350))
* **console:** enable tapp monitor ([#963](https://github.com/tkestack/tke/issues/963)) ([f518b01](https://github.com/tkestack/tke/commit/f518b016c426ab9a70df3900e886f5b35f318875))
* **console:** fix show the status of webconsole ([#922](https://github.com/tkestack/tke/issues/922)) ([5ceb4f2](https://github.com/tkestack/tke/commit/5ceb4f2dc96fd25452e4f83f911115871cb5f944))
* **console:** support cluster and worker upgrade ([#952](https://github.com/tkestack/tke/issues/952)) ([f81ffed](https://github.com/tkestack/tke/commit/f81ffed9c9cf82ded92c95edf2e5c14ec0562e34))
* **console:** support import cluster from kubeconfig ([#929](https://github.com/tkestack/tke/issues/929)) ([a5bd688](https://github.com/tkestack/tke/commit/a5bd6882387b1ec82fbab3042f4c284cb122c4fe))
* **console:** support pod paging and query ([#921](https://github.com/tkestack/tke/issues/921)) ([b917170](https://github.com/tkestack/tke/commit/b917170339e5883143547aa867e07d354773baef))
* **doc:** control plane enable scale ([#962](https://github.com/tkestack/tke/issues/962)) ([48d6a0f](https://github.com/tkestack/tke/commit/48d6a0f7a81ae3f2e2b901ce3c6787455b62a9f7))
* **doc:** ipv6 on business cluster ([#826](https://github.com/tkestack/tke/issues/826)) ([c51c5d6](https://github.com/tkestack/tke/commit/c51c5d6b7d51fda5ccee9cc78f3aef2cd91b3ef3))
* **installer:** limit the permissions of the certificates in webhook ([#829](https://github.com/tkestack/tke/issues/829)) ([ab74cf9](https://github.com/tkestack/tke/commit/ab74cf9aae135f27b75b22f4556bcab00f2af33e))
* **installer:** patch k8s versions ([#913](https://github.com/tkestack/tke/issues/913)) ([251eb03](https://github.com/tkestack/tke/commit/251eb03a52356d66ca815f245e5799e0c13784e6))
* **installer:** support certificate chain ([#975](https://github.com/tkestack/tke/issues/975)) ([22b87fe](https://github.com/tkestack/tke/commit/22b87fe1d943e807e89cb008b4fc3fd80d07b755))
* **installer:** support upgrade through installer ([#939](https://github.com/tkestack/tke/issues/939)) ([10e2ac1](https://github.com/tkestack/tke/commit/10e2ac15cbb522f3d811df2b616128109a891d97))
* **monitor:** add not ready data in cacher ([#935](https://github.com/tkestack/tke/issues/935)) ([c7338fc](https://github.com/tkestack/tke/commit/c7338fc2b53a8c0476af4ff405ac730619103922))
* **persistentevent:** allow es user and passowrd; update image version ([#903](https://github.com/tkestack/tke/issues/903)) ([02b5b3b](https://github.com/tkestack/tke/commit/02b5b3baa5be060405c318e81695db3e1660d003))
* **platform:** add label value for frontend ([#946](https://github.com/tkestack/tke/issues/946)) ([fc7c820](https://github.com/tkestack/tke/commit/fc7c8208dfab5714095f93b6a19c83ccdd62adc9))
* **platform:** ipv6 single stack support ([#818](https://github.com/tkestack/tke/issues/818)) ([3ab88a0](https://github.com/tkestack/tke/commit/3ab88a08eb3971d45ec4654931029bb5db7276f3))
* **platform:** ipv6 support ([#799](https://github.com/tkestack/tke/issues/799)) ([0a83821](https://github.com/tkestack/tke/commit/0a8382155a5265b04b16025f253e400351196534))
* **platform:** support customize flannel backend type ([fb9dec1](https://github.com/tkestack/tke/commit/fb9dec1035d237a0e7b323004e72ff0f40ff3bf4))
* **platform:** support k8s 1.17.13 ([#902](https://github.com/tkestack/tke/issues/902)) ([ea7b26f](https://github.com/tkestack/tke/commit/ea7b26fbe6f38dadf870e7586194571867aca87c))
* **platform:** support master scale ([#908](https://github.com/tkestack/tke/issues/908)) ([cb475bf](https://github.com/tkestack/tke/commit/cb475bf82e5e60a67d3b8861861003cb2cc58e55))
* **platform:** update cluster upgrade process ([#867](https://github.com/tkestack/tke/issues/867)) ([fa61791](https://github.com/tkestack/tke/commit/fa61791d57fa039d3148847a6f2c77ba2582ee08))
* **platform,monitor:** add more metrics for k8s node ([b62fc41](https://github.com/tkestack/tke/commit/b62fc4129fc791bbc7c44a7371e40cd2cd6260e4))
* **platform,monitor:** install prometheus-adapter via prometheus addon ([#820](https://github.com/tkestack/tke/issues/820)) ([6dc9211](https://github.com/tkestack/tke/commit/6dc9211e3ae7c559897bbbd97789e3519d6d7e9b))
* **registry:** add harbor chart proxy and inject tenant ([#887](https://github.com/tkestack/tke/issues/887)) ([57f6e5c](https://github.com/tkestack/tke/commit/57f6e5c0a0d9664de8f0ec654a7b25a4822b02c8))
* **registry:** add harbor proxy ([#852](https://github.com/tkestack/tke/issues/852)) ([e96f318](https://github.com/tkestack/tke/commit/e96f31879fd2ddf9af0a6a5617f7877388c470ce))
* **registry:** support harbor image management with tke api ([#906](https://github.com/tkestack/tke/issues/906)) ([991eb95](https://github.com/tkestack/tke/commit/991eb958dfe3eaeeb8a9bc954b50b670f91a0bae))
* **registry:** support helm chart management with harbor backend ([#918](https://github.com/tkestack/tke/issues/918)) ([533b095](https://github.com/tkestack/tke/commit/533b0953b7f236d2891cccb87fe4f747553d0586))
* add smoke test ([#872](https://github.com/tkestack/tke/issues/872)) ([40a8975](https://github.com/tkestack/tke/commit/40a8975819638f07edefed991def454f0191672f))


### Reverts

* Revert "fix: comment out one job (#890)" (#892) ([c560e77](https://github.com/tkestack/tke/commit/c560e7771fd622d21f112dbf254be0b71456881c)), closes [#890](https://github.com/tkestack/tke/issues/890) [#892](https://github.com/tkestack/tke/issues/892)
