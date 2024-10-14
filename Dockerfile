FROM golang:alpine AS builder
WORKDIR $GOPATH/src/product-api
COPY . .
ARG cmd
ENV CMDDIR=$cmd
RUN cp $GOPATH/src/product-api/cmd/$CMDDIR/main.go .
RUN go mod tidy
RUN go mod vendor
RUN go build -o /go/bin/app

FROM golang:alpine
ARG cmd
ENV CMDBIN=$cmd
COPY --from=builder /go/bin/app /bin/$CMDBIN
ENTRYPOINT $CMDBIN
