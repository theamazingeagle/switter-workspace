# STEP 1 build executable binary
FROM golang:alpine as builder
COPY ../src/switter-back $GOPATH/src/mypackage/myapp/
WORKDIR $GOPATH/src/mypackage/myapp/
#get dependencies
RUN go get -d -v
#build the binary
#RUN go build -o "/go/bin/main"
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o "/go/bin/main" .

# STEP 2 build a small image
# start from scratch
FROM scratch
# Copy our static executable
#COPY --from=builder "/go/bin/main" "/opt/app/"
#COPY "conf.json" "/opt/app/"
#ENTRYPOINT ["/opt/app/main"]

COPY --from=builder "/go/bin/main" "/"
COPY "../src/switter-back/conf.json" "/"
ENTRYPOINT ["/main"]