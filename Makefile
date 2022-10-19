build-rt:
	GOOS=linux go build -o RT ./recenttrades/main/main.go

build-aq:
	GOOS=linux go build -o AQ ./quotes/main/main.go

build-processor:
	GOOS=linux go build -o MP ./processor/main/main.go

gen-proto:
	protoc --go_out=. --go_opt=Mproto/crypto.proto=github.com/grpc/crypto --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/crypto.proto