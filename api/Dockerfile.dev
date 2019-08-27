FROM golang:latest
WORKDIR '/app'
RUN go get github.com/gin-contrib/cors
RUN go get github.com/gin-gonic/gin
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/google/uuid
RUN go get github.com/aws/aws-sdk-go/aws
RUN go get github.com/aws/aws-sdk-go/service/s3
RUN go get github.com/joho/godotenv
COPY . .
CMD ["go", "run", "main.go", "controller.go", "service.go", "dao.go", "util.go"]
