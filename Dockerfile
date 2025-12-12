# Build Stage

FROM golang:1.23-alpine AS builder

RUN apk add --no-cache git build-base

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o bin/form-pf-extractor-service . 

# Runtime Stage

FROM alpine:3.18 As runner


RUN apk add -U tzdata 
ENV TZ=Asia/Kolkata

RUN adduser -D -g '' formpf \
 && mkdir -p $STORAGE_PATH /var/log/form-pf

WORKDIR /app

COPY --from=builder /app/bin /app/bin
COPY --from=builder /app/entrypoint.sh /app/entrypoint.sh

ENTRYPOINT ["/bin/sh", "entrypoint.sh"]
CMD ["./bin/form-pf-extractor-service 2>&1 | tee -a /var/log/form-pf/service.log"]


