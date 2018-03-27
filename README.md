# go-fullstack

# TODO
- Implement the middleware to expose the STDOUT log to ELK stack.
- Implement distributed tracing system middleware
- Exposing the log to Prometheus and Grafana
- Implement strest test and load test with Gatling (Scala)

# Table of Content
- [Prerequisite](#prerequisite)
- [Installation guide](#installation-guide)
- [Update dependency](#update-dependency)
- [Project checking status](#project-checking-status)
- [Stack](#stack)
- [How to add go dependency?](#how-to-add-go-dependency)
- [Go Testing](#testing)
- [Commands](#command)
- [Example Request](#example)
- [Build](#build)

# Prerequisite
- [dep](https://github.com/golang/dep/blob/master/docs/installation.md)
- [Docker](https://docs.docker.com/samples/library/golang/)

# Installation guide
```
dep ensure
```

# Update dependency
```
dep update
```

# Project checking status
```
dep status

The command above will be shown reult like this

PROJECT                            CONSTRAINT     VERSION        REVISION  LATEST   PKGS USED
github.com/labstack/echo           ^3.3.0         3.3.0          27b5253   3.3.0    1
github.com/labstack/gommon         0.2.4          0.2.4          6fe1405   0.2.4    2

or

dep status -dot | dot -T png | open -f -a /Applications/Preview.app
```

# Stack
- [Dep](https://golang.github.io) for go package manager. [see more go package manager](https://github.com/golang/go/wiki/PackageManagementTools)

# How to add go dependency?
```
dep ensure -add github.com/labstack/echo
dep ensure
```

# Testing
You would create new test file and follows naming convension like this `<yourfile>.test.go`

```
go test
```
# Commands
**Install package** = `dep ensure`

**Update dependency** = `dep ensure -update`

**Install new package** = `go install <your package>`

**Check project status** = `dep status`

**Generate dependency graphic** = `dep status -dot | dot -T png | open -f -a /Applications/Preview.app`

**Checking version** = `go version`

**Start server** = `go run main.go`

**Build** = `env GOOS=target-OS GOARCH=target-architecture go build package-import-path` [read more](https://www.digitalocean.com/community/tutorials/how-to-build-go-executables-for-multiple-platforms-on-ubuntu-16-04)

**Build platform Windows Intel**
```sh
env GOOS=windows GOARCH=386 go build KTAXA/go-api-boilerplate
```

**Build platform Windows AMD**
```sh
env GOOS=windows GOARCH=amd64 go build KTAXA/go-api-boilerplate
```
# Alternative command
You will easily to execute any command please see `Makefile`, and execute it like `make start`

# Example Request
## [GET] /users
```sh
curl -X GET \
  http://localhost:3000/users \
  -H 'Cache-Control: no-cache' \
  -H 'Content-Type: application/json'
```

## [GET] /user/:id
```sh
curl -X GET \
  http://localhost:1323/user/5ab35fccd4989ee67ef04330 \
  -H 'Cache-Control: no-cache' \
  -H 'Content-Type: application/json'
```

## [POST] /addUser
```sh
curl -X POST \
  http://localhost:1323/addUser \
  -H 'Cache-Control: no-cache' \
  -H 'Content-Type: application/json' \
  -d '{
	"Name": "Golf",
	"Email": "golf@test.com"
}'
```