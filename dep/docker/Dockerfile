# Sets golang:1.18.3 as the bulding image
FROM golang:1.18.3 AS build

# Updates and installs the required Linux packages
RUN set -eux \
    && apt-get update -y \
    && apt-get upgrade -y \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

# Sets alpine:3.16.0 as the releasing image
FROM alpine:3.16.0 AS relase

# Updates and installs required Linux packages
RUN set -eux \
    && apk update \
    && apk upgrade \
    && rm /var/cache/apk/*
