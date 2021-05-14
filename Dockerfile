FROM golang:1.16-alpine

WORKDIR /srv/app
COPY . .

RUN go mod download && \
    go build

EXPOSE 3000
CMD ["asgard-auth"]