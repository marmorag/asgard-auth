FROM golang:1.16-alpine as build

WORKDIR /srv/app
COPY . .

RUN go mod download && \
    go build

FROM alpine
LABEL "org.opencontainers.image.source"="https://github.com/OWNER/REPO"

WORKDIR /srv/app
COPY --from=build /srv/app/asgard-auth asgard-auth

EXPOSE 3000
CMD ["/srv/app/asgard-auth"]