# Choose whatever you want, version >= 1.16
FROM golang:1.23-alpine
LABEL authors="Anophel"

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./
RUN go mod download

CMD ["air", "-c", ".air.toml"]