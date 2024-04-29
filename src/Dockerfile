FROM golang:1.19-buster AS builder

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . ./

RUN go build -v -o server ./cmd/main.go

FROM debian:buster-slim 
RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive \
    rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/server /app/server
COPY --from=builder /app/config/config-docker.yml /app/config/config-docker.yml
COPY --from=builder /app/docs /app/docs

ENV APP_ENV docker
ENV PORT ${Port}

CMD [ "/app/server" ]
 