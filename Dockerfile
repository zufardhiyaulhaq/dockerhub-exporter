#################
# Base image
#################
FROM alpine:3.12 as dockerhub-exporter-base

USER root

RUN addgroup -g 10001 dockerhub-exporter && \
    adduser --disabled-password --system --gecos "" --home "/home/dockerhub-exporter" --shell "/sbin/nologin" --uid 10001 dockerhub-exporter && \
    mkdir -p "/home/dockerhub-exporter" && \
    chown dockerhub-exporter:0 /home/dockerhub-exporter && \
    chmod g=u /home/dockerhub-exporter && \
    chmod g=u /etc/passwd

ENV USER=dockerhub-exporter
USER 10001
WORKDIR /home/dockerhub-exporter

#################
# Builder image
#################
FROM golang:1.19-alpine AS dockerhub-exporter-builder
RUN apk add --update --no-cache alpine-sdk
WORKDIR /app
COPY . .
RUN make build

#################
# Final image
#################
FROM dockerhub-exporter-base

COPY --from=dockerhub-exporter-builder /app/bin/dockerhub-exporter /usr/local/bin

# Command to run the executable
ENTRYPOINT ["dockerhub-exporter"]
