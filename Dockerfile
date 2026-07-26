FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

RUN go build -o ewallet-ums .

# ---------- final stage ----------
FROM alpine:3.20

WORKDIR /app

COPY --from=builder /app/ewallet-ums .
COPY --from=builder /app/.env .

RUN chmod +x ewallet-ums

EXPOSE 8080
EXPOSE 7000

CMD ["./ewallet-ums"]
