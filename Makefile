# Quickly to command the app
# Author: Teerapong, Singthong

install:
	dep ensure

update:
	dep ensure -update

status:
	dep status

status-img:
	dep status -dot | dot -T png | open -f -a /Applications/Preview.app	

start:
	go run main.go

version:
	go version
	dep version

build:
	# go build 

test:
	go test **/*.test.go
