FROM alpine:latest

WORKDIR /app

RUN apk add --no-cache curl unzip

ENV PB_VERSION=0.22.8
ENV PB_ARCH=linux_amd64

RUN curl -L https://github.com/pocketbase/pocketbase/releases/download/v${PB_VERSION}/pocketbase_${PB_VERSION}_${PB_ARCH}.zip -o pb.zip \
    && unzip pb.zip && rm pb.zip \
    && chmod +x pocketbase

EXPOSE 8090

CMD ["./pocketbase", "serve", "--http=0.0.0.0:8090", "--dir", "/pb_data"]