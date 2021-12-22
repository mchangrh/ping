FROM golang:alpine as builder

WORKDIR /go/ping

COPY . .
RUN go build 

FROM scratch
COPY --from=builder /go/ping/ping .

ENTRYPOINT ["/ping"]
EXPOSE 80