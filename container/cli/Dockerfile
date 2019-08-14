FROM debian:latest

RUN mkdir -p /app

WORKDIR /app

ADD container-cli /app/container-cli

CMD ["./container-cli", "--registry_address=registry:8500"]