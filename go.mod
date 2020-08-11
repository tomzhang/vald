module github.com/vdaas/vald

go 1.14

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v14.2.0+incompatible
	github.com/aws/aws-sdk-go => github.com/aws/aws-sdk-go v1.33.21
	github.com/boltdb/bolt => github.com/boltdb/bolt v1.3.1
	github.com/cockroachdb/errors => github.com/cockroachdb/errors v1.5.1-0.20200617111016-cc0024f9c4d3
	github.com/coreos/etcd => go.etcd.io/etcd v0.5.0-alpha.5.0.20200425165423-262c93980547
	github.com/docker/docker => github.com/moby/moby v17.12.0-ce-rc1.0.20200618181300-9dc6525e6118+incompatible
	github.com/envoyproxy/protoc-gen-validate => github.com/envoyproxy/protoc-gen-validate v0.4.0
	github.com/go-sql-driver/mysql => github.com/go-sql-driver/mysql v1.5.1-0.20200720071143-73dc904a9ece
	github.com/gocql/gocql => github.com/gocql/gocql v0.0.0-20200624222514-34081eda590e
	github.com/gogo/protobuf => github.com/gogo/protobuf v1.3.2-0.20200807193113-deb6fe8ca7c6
	github.com/gophercloud/gophercloud => github.com/gophercloud/gophercloud v0.12.0
	github.com/gorilla/mux => github.com/gorilla/mux v1.7.5-0.20200711200521-98cb6bf42e08
	github.com/gorilla/websocket => github.com/gorilla/websocket v1.4.2
	github.com/tensorflow/tensorflow => github.com/tensorflow/tensorflow v2.1.0+incompatible
	golang.org/x/crypto => golang.org/x/crypto v0.0.0-20200728195943-123391ffb6de
	k8s.io/api => k8s.io/api v0.18.6
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.18.6
	k8s.io/apimachinery => k8s.io/apimachinery v0.18.6
	k8s.io/client-go => k8s.io/client-go v0.18.6
	sigs.k8s.io/controller-runtime => sigs.k8s.io/controller-runtime v0.6.2
)

require (
	cloud.google.com/go v0.63.1-0.20200810190638-6b447de85572
	code.cloudfoundry.org/bytefmt v0.0.0-20200131002437-cf55d5288a48
	contrib.go.opencensus.io/exporter/jaeger v0.2.1
	contrib.go.opencensus.io/exporter/prometheus v0.2.1-0.20200609204449-6bcf6f8577f0
	contrib.go.opencensus.io/exporter/stackdriver v0.13.3-0.20200811013819-19b2ae76216d
	github.com/aws/aws-sdk-go v1.33.17
	github.com/cespare/xxhash/v2 v2.1.1
	github.com/cockroachdb/errors v1.5.1-0.20200617111016-cc0024f9c4d3
	github.com/danielvladco/go-proto-gql/pb v0.6.1
	github.com/envoyproxy/protoc-gen-validate v0.4.1-0.20200810204539-e020f3aef8ae
	github.com/fsnotify/fsnotify v1.4.10-0.20200417215612-7f4cf4dd2b52
	github.com/go-redis/redis/v7 v7.2.1-0.20200519055202-64bb0b7f3af4
	github.com/go-sql-driver/mysql v1.5.1-0.20200720071143-73dc904a9ece
	github.com/gocql/gocql v0.0.0-20200131111108-92af2e088537
	github.com/gocraft/dbr/v2 v2.7.1-0.20200218045517-f487ccffc6d0
	github.com/gogo/protobuf v1.3.2-0.20200807193113-deb6fe8ca7c6
	github.com/google/go-cmp v0.5.2-0.20200729152727-036ffc7f24c0
	github.com/google/gofuzz v1.1.0
	github.com/gorilla/mux v1.7.5-0.20200711200521-98cb6bf42e08
	github.com/hashicorp/go-version v1.2.1
	github.com/json-iterator/go v1.1.11-0.20200806011408-6821bec9fa5c
	github.com/klauspost/compress v1.10.11-0.20200811083703-3fdfbd5997ec
	github.com/kpango/fastime v1.0.17-0.20200719090826-d1de9e8e3c85
	github.com/kpango/fuid v0.0.0-20190507064958-80435564606b
	github.com/kpango/gache v1.2.2-0.20200709224359-34beea72198c
	github.com/kpango/glg v1.5.2-0.20200624074227-79f879e27c67
	github.com/lucasb-eyer/go-colorful v1.0.3
	github.com/pierrec/lz4/v3 v3.3.2
	github.com/scylladb/gocqlx v1.5.1-0.20200423154401-507391a34cf0
	github.com/tensorflow/tensorflow v0.0.0-00010101000000-000000000000
	github.com/yahoojapan/gongt v0.0.0-20190517050727-966dcc7aa5e8
	github.com/yahoojapan/ngtd v0.0.0-20200424071638-9872bbae3700
	go.opencensus.io v0.22.5-0.20200719225510-d7677d6af595
	go.uber.org/automaxprocs v1.3.1-0.20200415073007-b685be8c1c23
	go.uber.org/goleak v1.1.10
	golang.org/x/net v0.0.0-20200707034311-ab3426394381
	golang.org/x/sync v0.0.0-20200625203802-6e8e738ad208
	golang.org/x/sys v0.0.0-20200810151505-1b9f1253b3ed
	golang.org/x/tools v0.0.0-20200811032001-fd80f4dbb3ea // indirect
	gonum.org/v1/hdf5 v0.0.0-20200504100616-496fefe91614
	gonum.org/v1/plot v0.7.1-0.20200803120916-6a037fda5e90
	google.golang.org/api v0.30.1-0.20200811142409-9e2b1a1ddef6
	google.golang.org/genproto v0.0.0-20200808173500-a06252235341
	google.golang.org/grpc v1.32.0-dev.0.20200811135751-6aaac03d175a
	gopkg.in/yaml.v2 v2.3.0
	k8s.io/api v0.18.6
	k8s.io/apimachinery v0.18.6
	k8s.io/client-go v0.18.6
	k8s.io/metrics v0.18.6
	sigs.k8s.io/controller-runtime v0.0.0-00010101000000-000000000000
)
