# halfway_demo

halfway template

```text
├── Dockerfile
├── README.md
├── api
│   └── demo
│       ├── demo.pb.go
│       └── demo.proto
├── cmd
│   └── main.go
├── config
│   └── example.yaml
├── go.mod
├── go.sum
├── internal
│   ├── consts
│   │   └── consts.go
│   ├── dao
│   │   ├── dao.go
│   │   ├── mysql.go
│   │   └── redis.go
│   ├── entity
│   │   └── orm
│   │       └── demo.go
│   ├── routers
│   │   ├── grpc.go
│   │   └── http.go
│   ├── server
│   │   ├── grpc.go
│   │   └── http.go
│   └── service
│       ├── hello.go
│       └── service.go
└── utils
    └── common.go
```