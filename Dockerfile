FROM golang:1.16.5 AS development
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go install github.com/cespare/reflex
EXPOSE 4000
CMD reflex -g '*go' go run cmd/main/main.go --start-service