FROM golang:1.14.6-alpine3.12 as builder
COPY go.mod go.sum /go/src/github.com/vitormcid/getnet-challange/
WORKDIR /go/src/github.com/vitormcid/getnet-challange
RUN go mod download
COPY . /go/src/github.com/vitormcid/getnet-challange
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/getnet-challange github.com/vitormcid/getnet-challange

FROM alpine
RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder /go/src/github.com/vitormcid/getnet-challange/build/getnet-challange /usr/bin/getnet-challange
EXPOSE 8010 8010
ENTRYPOINT ["/usr/bin/getnet-challange"]
