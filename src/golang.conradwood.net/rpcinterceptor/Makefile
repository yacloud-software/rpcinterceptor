.PHONY:	client server
all:	client server
client:
	cd client && go install rpcinterceptor-client.go
server:
	cd server && go install rpcinterceptor-server.go log.go requestid_creator.go audit.go error-log.go
#access-checks.go
