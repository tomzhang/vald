module github.com/vdaas/vald

go 1.15

replace (
	cloud.google.com/go => cloud.google.com/go latest
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v14.2.0
	github.com/aws/aws-sdk-go => github.com/aws/aws-sdk-go latest
	github.com/boltdb/bolt => github.com/boltdb/bolt v1.3.1
	github.com/chzyer/logex => github.com/chzyer/logex master
	github.com/coreos/etcd =>  go.etcd.io/etcd latest
	github.com/docker/docker => github.com/moby/moby latest
	github.com/envoyproxy/protoc-gen-validate => github.com/envoyproxy/protoc-gen-validate latest
	github.com/go-sql-driver/mysql => github.com/go-sql-driver/mysql latest
	github.com/gocql/gocql => github.com/gocql/gocql latest
	github.com/gogo/googleapis => github.com/gogo/googleapis latest
	github.com/gogo/protobuf => github.com/gogo/protobuf latest
	github.com/google/go-cmp => github.com/google/go-cmp latest
	github.com/google/pprof => github.com/google/pprof master
	github.com/googleapis/gnostic => github.com/googleapis/gnostic v0.4.0
	github.com/gophercloud/gophercloud => github.com/gophercloud/gophercloud latest
	github.com/gorilla/websocket => github.com/gorilla/websocket latest
	github.com/hailocab/go-hostpool => github.com/monzo/go-hostpool latest
	github.com/klauspost/compress => github.com/klauspost/compress master
	github.com/tensorflow/tensorflow => github.com/tensorflow/tensorflow v2.1.0
	golang.org/x/crypto => golang.org/x/crypto latest
	google.golang.org/grpc => google.golang.org/grpc latest
	google.golang.org/protobuf => google.golang.org/protobuf latest
	k8s.io/api => k8s.io/api v0.18.8
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.18.8
	k8s.io/apimachinery => k8s.io/apimachinery v0.18.8
	k8s.io/client-go => k8s.io/client-go v0.18.8
	k8s.io/metrics => k8s.io/metrics v0.18.8
	sigs.k8s.io/controller-runtime => sigs.k8s.io/controller-runtime v0.6.2
)