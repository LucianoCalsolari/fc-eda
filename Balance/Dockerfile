# Build stage
FROM golang:1.18

WORKDIR /app/

RUN apt-get update && apt-get install -y librdkafka-dev

CMD ["go", "run", "cmd/balance/main.go"]