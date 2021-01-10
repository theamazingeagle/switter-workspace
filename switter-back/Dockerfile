# STEP 1 build executable binary
FROM alpine:latest as builder
RUN apk add go git
WORKDIR /switter-back/
COPY ./ ./
#RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o "/go/bin/main" ./cmd/

# STEP 2 build a small image
FROM scratch
COPY --from=builder "/go/bin/main" "/"
COPY "./conf.json" "/"
EXPOSE 8080
ENTRYPOINT ["/main"]