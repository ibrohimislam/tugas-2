#TUGAS-2

##File Structures
```
server/         directory for server package
client/         directory for client package
bin/            directory for binary files
test/           testfile
README          this file
tugas-2.thrift  thrift file
file_svc.patch  patch file for autogenerated source bug
```

#How to compile

prerequisite:
- golang compiler

procedure:
- create GO working directory:
```
mkdir ~/go
```
- set environment variable:
```
export GOPATH=/home/user/go
```
- create path, for this source:
```
mkdir -p $GOPATH/src/github.com/ibrohimislam/tugas-2
```
- move to this directory
```
cd $GOPATH/src/github.com/ibrohimislam/tugas-2/bin
```
- compile client
```
go build ../client/client.go
```
- compile server
```
go build ../server/main.go
```

##How to Use

1. run server ./bin/server
2. run client ./bin/client

##Testing

Manual test
```
> DIR /
> DIR /opt/
> CREATEDIR /home/ibrohim test 
> GETFILE /opt/rar rar.txt a.txt
> PUTFILE /home/ibrohim a.txt a.txt
```