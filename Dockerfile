# Build Stage

FROM golang:1.23-alpine AS builder

RUN apk add --no-cache git build-base

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o bin/form-pf-extractor-service . 

# Runtime Stage

FROM python:3.11-slim-bookworm AS runner

ENV TZ=Asia/Kolkata

RUN apt-get update \
 && apt-get install --no-install-recommends -y tzdata \
 && rm -rf /var/lib/apt/lists/* \
 && useradd -r formpf \
 && mkdir -p /var/log/form-pf

WORKDIR /app

COPY --from=builder /app/bin /app/bin
COPY --from=builder /app/entrypoint.sh /app/entrypoint.sh
COPY config/config.ini /app/config/config.ini

ENTRYPOINT ["/bin/sh", "entrypoint.sh"]
CMD ["./bin/form-pf-extractor-service 2>&1 | tee -a /var/log/form-pf/service.log"]


