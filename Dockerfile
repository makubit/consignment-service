#FROM debian:latest

#RUN mkdir /app
#WORKDIR /app
#ADD shippy-service /app/shippy-service

#CMD ["./shippy-service"]

FROM golang:alpine as builder

RUN apk --no-cache add git

WORKDIR /app/shippy-service

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o shippy-service

FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app

COPY --from=builder /app/shippy-service .

CMD ["./shippy-service"]