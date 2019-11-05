FROM golang:apline as builder

RUN apk --no-cache add git

WORKDIR /app/vessel-service

COPY . .

RUN go mod downlonad
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffinx cgo -o shippy-service-vessel

FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /app/vessel-service .

CMD ["./shippy-service-vessel"]