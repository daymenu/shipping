FROM debian:latest

# MICRO 配置
ENV MICRO_ADDRESS=:50051
ENV MICRO_REGISTRY=consul
ENV DB_HOST=mariadb.host
ENV DB_NAME=shipping
ENV DB_USER=root
ENV DB_PASSWORD=123456

# app 
RUN mkdir /app  
WORKDIR /app  
ADD user /app/user

# 开启container微服务
CMD ["./user",  "--registry_address=registry:8500"]