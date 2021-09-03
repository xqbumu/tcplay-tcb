LDFLAGS=-ldflags "-s -w"

buildHelloWrold:
	GOOS=linux GO111MODULE=on go build $(LDFLAGS) -o functions/HelloWorld/main functions/HelloWorld/main.go

deployHelloWrold: buildHelloWrold
	tcb fn deploy HelloWorld

buildHelloApiGw:
	GOOS=linux GO111MODULE=on go build $(LDFLAGS) -o functions/HelloApiGw/main functions/HelloApiGw/main.go

deployHelloApiGw: buildHelloApiGw
	tcb fn deploy HelloApiGw
