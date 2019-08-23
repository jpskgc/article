FROM golang:1.12.9-alpine AS builder
WORKDIR /app
COPY ./go.mod ./go.sum ./
RUN go mod download
COPY ./ ./
RUN CGO_ENABLED=0 go build \
    -installsuffix 'static' \
    -o /api .

FROM scratch AS final
COPY --from=builder /myapp /myapp

EXPOSE 8080

ENTRYPOINT ["/api"]
