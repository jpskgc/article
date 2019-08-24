FROM golang:latest
ENV GO111MODULE=on
WORKDIR /go/src/github.com/jpskgc/article
COPY . .
RUN go mod download
RUN go build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
ENTRYPOINT ["/go/src/github.com/jpskgc/article/api"]
