# Sets golang:1.18.3 as the bulding image
FROM golang:1.18.3 AS build

# Updates and installs the required Linux packages
RUN set -eux \
    && apt-get update -y \
    && apt-get upgrade -y \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

# Sets the working directory
WORKDIR /app

# Copies the source code from the host to the container
COPY --chown=root:root . .

# Sets the required operating system from a building argument
ARG OPERARTING_SYSTEM
ENV OPERARTING_SYSTEM ${OPERARTING_SYSTEM:-linux}

# Sets the required architecture from a building argument
ARG ARCHITECTURE
ENV ARCHITECTURE ${ARCHITECTURE:-amd64}

# Generates an executable binary out of the source code
RUN go mod tidy \
    && CGO_ENABLED=0 GOOS=${OPERARTING_SYSTEM} GOARCH=${ARCHITECTURE} go build -a -o arachne .

# Sets alpine:3.16.0 as the releasing image
FROM alpine:3.16.0 AS relase

# Updates and installs required Linux packages
RUN set -eux \
    && apk update \
    && apk upgrade \
    && rm /var/cache/apk/*

# Copies the final artifactory from the building container
COPY --from=build /app/arachne /usr/local/bin/arachne

# Sets the default serving port from a building argument
ARG PORT
ENV PORT ${PORT:-8080}

# Exposes the default serving port
EXPOSE ${PORT}

# Sets the default entry point of the service
ENTRYPOINT ["./usr/local/bin/arachne"]

# Sets the default launching instruction
CMD ["-addr=0.0.0.0", "-port=8080"]
