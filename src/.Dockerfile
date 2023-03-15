FROM golang:1.18-alpine as builder

ENV APP_ENV docker
ENV PORT ${Port}

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . ./

RUN go build -v -o server ./cmd/main.go

COPY ./ ./app/server

CMD ["/app/server"]