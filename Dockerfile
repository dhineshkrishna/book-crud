FROM golang:1.24.2-alpine AS build
WORKDIR /app

# SQLite needs CGO + gcc
RUN apk add --no-cache gcc musl-dev

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# ❗ DO NOT disable CGO
RUN CGO_ENABLED=1 GOOS=linux go build -o app ./cmd/server

FROM alpine:latest
WORKDIR /root/

# Runtime dependency for sqlite
RUN apk add --no-cache libc6-compat

COPY --from=build /app/app .

EXPOSE 8080
CMD ["./app"]
