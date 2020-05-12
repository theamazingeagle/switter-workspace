# STEP 1 build executable binary
FROM golang:alpine as builder
#ENV GOPATH /home/root/go/src
#ENV GOROOT /home/root/go/src
WORKDIR $GOPATH/src/switter-back/
COPY ./src/switter-back/ $GOPATH/src/switter-back/
#WORKDIR /var/go/src/switter-back/
#ENV GOPATH /var/go/src/switter-back/
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
COPY "./src/switter-back/conf.json" "/"
ENTRYPOINT ["/main"]