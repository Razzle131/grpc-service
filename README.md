# grpc-service
______
# Table of Contents
* [About](#about)
* [Clone repo](#clone-repo)
* [Technologies](#technologies)
* [Quick start](#quick-start)

## About
This app created to observe gRPC-gateway and cobra libraries

## Clone repo
```
git clone https://github.com/Razzle131/grpc-service.git
```

## Technologies
Project is created with:
* Golang version: 1.22.2
* gRPC-gateway
* cobra

## Quick start
### Build from source
#### Without Makefile
* Ensure that Go is installed on your machine and it`s version is equal or higther than 1.22.2
```
go version
```
* Clone repo
```
git clone https://github.com/Razzle131/grpc-service.git
```
* Start server
```
go run entrypoints/server/main.go
```
* Start client with optins
```
go run entrypoints/server/main.go getDns
```
Other methods:
* addDns <ip>
* remDns <ip>
* getHostname
* setHostname <newHostname>
______
#### Using Makefile
* Ensure that Go is installed on your machine and it`s version is equal or higther than 1.22.2
```
go version
```
* Clone repo
```
git clone https://github.com/Razzle131/grpc-service.git
```
* Start server
```
make run-server
```
* Now you can use commands to start client `make run-client action=<action> arg=<arg>`
  * actions:
    * getDns
    * addDns
    * remDns
    * getHostname
    * setHostname
  * arg:
    * for host set request: new host name
    * for dns set request: ip address of dns server (example: `8.8.8.8`)
