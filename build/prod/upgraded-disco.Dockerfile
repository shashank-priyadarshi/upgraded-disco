# Build
ARG GO_VERSION=1.23

FROM golang:${GO_VERSION}-alpine AS builder

ARG APP_NAME=${APP_NAME}
ARG APP_VERSION=${APP_VERSION}
ARG USER=${USER}
ARG SOURCE_PATH=${SOURCE_PATH}

WORKDIR /usr/${USER}/${APP_NAME}.${APP_VERSION}/
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -trimpath ${SOURCE_PATH}
RUN build/pack.sh /usr/${USER}/${APP_NAME}.${APP_VERSION}/${APP_NAME} ${APP_NAME}.${APP_VERSION}

# Run
FROM alpine AS runner

ARG USER=${USER}
ARG APP_NAME=${APP_NAME}
ARG APP_VERSION=${APP_VERSION}
ARG CONFIG_SOURCE
ARG CONFIG_PATH

LABEL maintainer="https://ssnk.in"
LABEL tags="upgraded-disco,backend,portfolio,shashank priyadarshi"
LABEL version=${APP_VERSION}

STOPSIGNAL SIGTERM

WORKDIR /usr/${USER}/${APP_NAME}.${APP_VERSION}/

RUN apk add shadow && \
    groupadd -g 10001 ${USER} && \
    useradd -u 10000 -g ${USER} ${USER} \
    && chown -R ${USER}:${USER} /usr/${USER}/${APP_NAME}.${APP_VERSION}/

COPY --from=builder /usr/${USER}/${APP_NAME}.${APP_VERSION}/build/entrypoint.sh /usr/local/bin/
RUN chmod +x /usr/local/bin/entrypoint.sh

USER ${USER}:${USER}
COPY --from=builder /usr/${USER}/${APP_NAME}.${APP_VERSION}/${APP_NAME}.${APP_VERSION} /usr/${USER}/${APP_NAME}.${APP_VERSION}/
COPY --from=builder /app/${CONFIG_PATH}/${CONFIG_SOURCE} .
COPY --from=builder /app/plugins .

ENV USER=$USER
ENV APP_NAME=$APP_NAME
ENV APP_VERSION=$APP_VERSION
ENV CONFIG_SOURCE=$CONFIG_SOURCE
ENV CONFIG_PATH=$CONFIG_PATH

ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]
