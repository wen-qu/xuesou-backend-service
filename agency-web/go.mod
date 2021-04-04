module github.com/wen-qu/xuesou-backend-service/agency-web

go 1.15

require (
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/golang/protobuf v1.5.2
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/google/uuid v1.2.0 // indirect
	github.com/hashicorp/go-version v1.3.0 // indirect
	github.com/improbable-eng/grpc-web v0.14.0 // indirect
	github.com/jinzhu/copier v0.2.8
	github.com/klauspost/compress v1.11.13 // indirect
	github.com/micro/micro/v3 v3.2.0
	github.com/miekg/dns v1.1.41 // indirect
	github.com/rhysd/go-github-selfupdate v1.2.3 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/soheilhy/cmux v0.1.5 // indirect
	github.com/stretchr/objx v0.3.0 // indirect
	github.com/stretchr/testify v1.7.0 // indirect
	github.com/ulikunitz/xz v0.5.10 // indirect
	github.com/urfave/cli/v2 v2.3.0 // indirect
	github.com/wen-qu/xuesou-backend-service/agency-srv v0.0.0-20210404043350-8a597c03b499
	github.com/wen-qu/xuesou-backend-service/basic v0.0.0-20210404042406-e36ea8aec1e9 // indirect
	github.com/wen-qu/xuesou-backend-service/class-srv v0.0.0-20210325152105-ea9f2abd9285
	golang.org/x/crypto v0.0.0-20210322153248-0c34fe9e7dc2 // indirect
	golang.org/x/net v0.0.0-20210331212208-0fccb6fa2b5c // indirect
	golang.org/x/oauth2 v0.0.0-20210402161424-2e8d93401602 // indirect
	golang.org/x/sys v0.0.0-20210403161142-5e06dd20ab57 // indirect
	golang.org/x/text v0.3.6 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20210402141018-6c239bbf2bb1 // indirect
	google.golang.org/protobuf v1.26.0
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
	nhooyr.io/websocket v1.8.6 // indirect

)

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
