FROM debian:latest

RUN mkdir /app
WORKDIR /app
ADD shippy-service /app/shippy-service

CMD ["./shippy-service"]