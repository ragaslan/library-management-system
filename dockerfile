# Build stage
FROM golang:1.23-alpine AS build
WORKDIR /app
# cache go mod
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/server ./cmd/api

# Runtime stage
FROM alpine:3.18
RUN addgroup -S app && adduser -S -G app app
COPY --from=build /app/bin/server /usr/local/bin/server
USER app
ENV PORT=8080
EXPOSE 8080
ENTRYPOINT ["/usr/local/bin/server"]
