build-rt:
	rm IngestRT.zip
	rm IngestRT
	GOOS=linux go build -o IngestRT ./recenttrades/main/main.go
	zip IngestRT IngestRT

build-aq:
	rm IngestAQ.zip
	rm IngestAQ
	GOOS=linux go build -o IngestAQ ./quotes/main/main.go
	zip IngestAQ IngestAQ

build-processor:
	rm MessageProcessor.zip
	rm MessageProcessor
	GOOS=linux go build -o MessageProcessor ./processor/main/main.go
	zip MessageProcessor MessageProcessor

gen-proto:
	protoc --go_out=. --go_opt=Mproto/crypto.proto=github.com/grpc/crypto --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/crypto.proto