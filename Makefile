create-protobuf:
	protoc -I . ./protobuf/say.proto --go_out=plugins=grpc:.

build:
	GOOS=linux GOARCH=amd64 go build -o app ./cmd/server
	docker build \
		--tag gcr.io/text-to-speech-408813/say \
		--platform linux/amd64 .
	rm -rf app

push:
	docker push gcr.io/text-to-speech-408813/say
