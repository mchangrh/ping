default: build

build:
	CGO_ENABLED=0 go build -a --trimpath --installsuffix cgo -o mchangrh-ping