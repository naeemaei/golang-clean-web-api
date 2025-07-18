FROM golang:1.22-bookworm AS builder

WORKDIR /app

COPY go.* ./
ENV CGO_ENABLED=0
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go mod download

COPY . ./

RUN go build -v -o server ./cmd/main.go

FROM debian:bookworm-slim 
RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/server /app/server
COPY --from=builder /app/config/config-docker.yml /app/config/config-docker.yml
COPY --from=builder /app/docs /app/docs

ENV APP_ENV=docker
ENV PORT=${Port}

CMD [ "/app/server" ]
