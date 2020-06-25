
![Go CI](https://github.com/mchirico/go-pubsub/workflows/Go%20CI/badge.svg)
[![codecov](https://codecov.io/gh/mchirico/go-pubsub/branch/master/graph/badge.svg)](https://codecov.io/gh/mchirico/go-pubsub)
# go-pubsub


PubSub for ts-express




## Build with vendor
```
export GO111MODULE=on
go mod init
# Below will put all packages in a vendor folder
go mod vendor



go test -v -mod=vendor ./...

# Don't forget the "." in "./cmd/script" below
go build -v -mod=vendor ./...
```


## Don't forget golint

```

golint -set_exit_status $(go list ./... | grep -v /vendor/)

```


# mpubsub
