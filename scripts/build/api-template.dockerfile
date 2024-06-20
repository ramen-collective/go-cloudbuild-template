# ----------------------------------------------------
# Application build stage
FROM golang:1.19-alpine as builder

WORKDIR /go/src/github.com/ramen-collective/go-cloudbuild-template
COPY ./ ./

RUN apk update && apk add make git
RUN make build BIN=. GOLANG_ARCH=amd64


# ----------------------------------------------------
# Base production image
FROM alpine AS base

WORKDIR /app/
RUN apk update && apk add ca-certificates
RUN adduser -S api-template

# ----------------------------------------------------
# api-template base image
FROM base AS api-template

COPY --from=builder /go/src/github.com/ramen-collective/go-cloudbuild-template/api_template .

USER api-template
ENTRYPOINT [ "/app/api_template" ]