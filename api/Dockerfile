# FROM golang:latest
# WORKDIR '/app'
# RUN go get github.com/gin-contrib/cors
# RUN go get github.com/gin-gonic/gin
# RUN go get github.com/go-sql-driver/mysql
# RUN go get github.com/google/uuid
# RUN go get github.com/aws/aws-sdk-go/aws
# RUN go get github.com/aws/aws-sdk-go/service/s3
# RUN go get github.com/joho/godotenv
# COPY . .
# CMD ["go", "run", "main.go"]

# FROM golang:1.12.9-alpine AS builder
# WORKDIR /app
# COPY ./go.mod ./go.sum ./
# RUN go mod download
# COPY ./ ./
# RUN CGO_ENABLED=0 go build \
#     -installsuffix 'static' \
#     -o /api .

# FROM scratch AS final
# COPY --from=builder /myapp /myapp

# ENTRYPOINT ["/api"]

FROM golang:latest

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

ENTRYPOINT ["/app/api"]
