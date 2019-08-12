FROM golang AS builder

WORKDIR /go/src
COPY revprox.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags '-w' revprox.go

FROM scratch
LABEL maintainer "Waret Li <waret87@163.com>"

COPY --from=builder /go/src/revprox /

EXPOSE 5000

env REMOTE_SVC localhost:5000
env LOCAL_PORT 5000

ENTRYPOINT ["/revprox"]
