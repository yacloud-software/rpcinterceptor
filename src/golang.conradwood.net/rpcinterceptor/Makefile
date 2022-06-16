.PHONY:	client server rpcaclapi
all:	client server rpcaclapi
client:
	cd client && go install ${LDFLAGS} rpcinterceptor-client.go
server:
	cd server && go install ${LDFLAGS} rpcinterceptor-server.go log.go requestid_creator.go audit.go error-log.go
#access-checks.go
rpcaclapi:
	cd rpcaclapi/server && go install ${LDFLAGS} rpcaclapi-server.go
	cd rpcaclapi/client && go install ${LDFLAGS} rpcaclapi-client.go
