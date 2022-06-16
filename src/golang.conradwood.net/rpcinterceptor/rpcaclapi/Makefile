.PHONY:	client server
all:	client server
client:
	cd client && go install rpcaclapi-client.go logsorter.go access_config.go group_optimizer.go
server:
	cd server && go install rpcaclapi-server.go `ls -1 *.go|grep -v rpcaclapi-server.go`
