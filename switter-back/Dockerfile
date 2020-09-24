# STEP 1 build executable binary
FROM alpine:latest as builder
RUN apk add go git
WORKDIR $GOPATH/src/switter-back/
COPY ./src/switter-back/ $GOPATH/src/switter-back/
#RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o "/go/bin/main" .

# STEP 2 build a small image
FROM scratch
COPY --from=builder "/go/bin/main" "/"
COPY "./src/switter-back/conf.json" "/"
ENTRYPOINT ["/main"]