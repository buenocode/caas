# build caas
FROM golang:1.20 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY pkg/ ./pkg/
RUN CGO_ENABLED=0 GOOS=linux go build -o caas ./pkg/main.go

# add caas to headless-shell image
FROM chromedp/headless-shell:latest
COPY --from=builder /app/caas ./
ENTRYPOINT ["./caas"]  
