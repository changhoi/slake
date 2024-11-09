FROM golang:1.23-alpine
ENV TZ=Asia/Seoul
WORKDIR /home/app

RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./
RUN go mod download
COPY . .

CMD ["air", "-c", ".meta/.air.toml"]