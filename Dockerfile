FROM golang:alpine as builder

WORKDIR /go/ping

COPY . .
RUN CGO_ENABLED=0 go build -a --trimpath --installsuffix cgo

FROM scratch
COPY --from=builder /go/ping/ping .

ENTRYPOINT ["/ping"]
EXPOSE 80