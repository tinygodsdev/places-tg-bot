# Stage 1: Builder
FROM golang:1.22.2-alpine AS builder

ARG GITHUB_USERNAME
ARG GITHUB_TOKEN
ARG GITHUB_REPO_PATH

RUN apk add --no-cache git ca-certificates
RUN git config --global \
url."https://${GITHUB_USERNAME}:${GITHUB_TOKEN}@github.com/${GITHUB_REPO_PATH}".insteadOf \
"https://github.com/${GITHUB_REPO_PATH}"

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/main ./cmd/main.go

# Stage 2: Final image
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
ENV PORT=8080
EXPOSE 8080
CMD ["./main"]
