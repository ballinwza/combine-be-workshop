FROM --platform=$BUILDPLATFORM golang:1.25.0-alpine AS build
ARG TARGETPLATFORM
ARG BUILDPLATFORM
RUN echo "Building on $BUILDPLATFORM, building for $TARGETPLATFORM" > /log

WORKDIR /app

COPY . ./
RUN go mod download

RUN go build 

FROM --platform=$BUILDPLATFORM node:20.11-alpine AS runner
ARG TARGETPLATFORM
ARG BUILDPLATFORM
RUN echo "Building on $BUILDPLATFORM, building for $TARGETPLATFORM" > /log

RUN apk add --no-cache libc6-compat
WORKDIR /app
COPY --from=build /app ./
COPY --from=build /app/docker-entrypoint.sh ./
ENV HOST_PORT=8080
EXPOSE 8080

ENTRYPOINT [ "/bin/sh", "docker-entrypoint.sh" ]