# goimport
The test for import on golang.

## Environment
```ShellSession
$ go version
go version go1.14.2 darwin/amd64

$ echo $GO111MODULE
on
```

## Tree
```ShellSession
$ tree $GOPATH/src/github.com/kotaoue/goimport -L 3
/$GOPATH/src/github.com/kotaoue/goimport
├── LICENSE
├── README.md
├── go.mod
├── go.sum
├── main.go
└── packages
    └── car
        ├── car.go
        └── go.mod
```

## cf.
* https://github.com/golang/go/wiki/Modules#quick-start-example
* https://github.com/golang/go/wiki/Modules#when-should-i-use-the-replace-directive

## Command logs
```ShellSession
$ cd $GOPATH/src/github.com/kotaoue/goimport/packages/car
$ go mod init

$ cd $GOPATH/src/github.com/kotaoue/goimport/
$ go mod init github.com/kotaoue/goimport
$ go mod edit -replace github.com/kotaoue/goimport/packages/car=./packages/car

# When build completed a require for "packages/car" are attached on "go.mod".
$ go build main.go

$ go list -m all
github.com/kotaoue/goimport
github.com/kotaoue/goimport/packages/car v0.0.0-00010101000000-000000000000 => ./packages/car
golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c
rsc.io/quote v1.5.2
rsc.io/sampler v1.3.0
```

## Result
```ShellSession
$ go run main.go 
2020/06/24 10:51:27 Start
2020/06/24 10:51:27 こんにちは世界。
2020/06/24 10:51:27 vroom!!
2020/06/24 10:51:27 End
```