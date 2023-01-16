.PHONY:	client server rpcaclapi
all:	client server rpcaclapi
client:
	cd client && go install ${LDFLAGS} rpcinterceptor-client.go
server:
	cd server && go install ${LDFLAGS} rpcinterceptor-server.go `ls -1 *.go|grep -v rpcinterceptor-server.go`
#access-checks.go
rpcaclapi:
	cd rpcaclapi/server && go install ${LDFLAGS} rpcaclapi-server.go `ls -1 *.go|grep -v rpcaclapi-server.go`
	cd rpcaclapi/client && go install ${LDFLAGS} rpcaclapi-client.go
