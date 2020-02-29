# Build Container
FROM golang:latest as builder
WORKDIR /go/src/app
COPY . .
# Set Environment Variable
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
# Build
RUN go build -o app main.go

# Runtime Container
FROM alpine
RUN apk add --no-cache ca-certificates
COPY --from=builder /go/src/app/app /app
EXPOSE 8080
ENTRYPOINT ["/app"]