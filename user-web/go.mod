module github.com/wen-qu/xuesou-backend-service/user-web

go 1.16

require (
	github.com/golang/protobuf v1.4.3
	github.com/google/uuid v1.2.0
	github.com/improbable-eng/grpc-web v0.14.0 // indirect
	github.com/klauspost/compress v1.11.12 // indirect
	github.com/micro/micro/v3 v3.1.1
	github.com/miekg/dns v1.1.40 // indirect
	github.com/rhysd/go-github-selfupdate v1.2.3 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/stretchr/objx v0.3.0 // indirect
	github.com/stretchr/testify v1.7.0 // indirect
	github.com/ulikunitz/xz v0.5.10 // indirect
	github.com/urfave/cli/v2 v2.3.0 // indirect
	github.com/wen-qu/xuesou-backend-service/basic v0.0.0-20210311093833-4286b09a3392 // indirect
	github.com/wen-qu/xuesou-backend-service/user-srv v0.0.0-20210311100442-8cc8f7b972c0
	golang.org/x/crypto v0.0.0-20210220033148-5ea612d1eb83 // indirect
	golang.org/x/net v0.0.0-20210226172049-e18ecbb05110 // indirect
	golang.org/x/oauth2 v0.0.0-20210220000619-9bb904979d93 // indirect
	golang.org/x/sys v0.0.0-20210309074719-68d13333faf2 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20210310155132-4ce2db91004e // indirect
	google.golang.org/grpc v1.36.0 // indirect
	google.golang.org/protobuf v1.25.0
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
	nhooyr.io/websocket v1.8.6 // indirect
)

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
