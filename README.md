# go-fullstack

# Table of Content

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
