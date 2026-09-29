FROM golang:1.12.0-alpine3.9

RUN mkdir /go/src/fins-api-services

ADD . /go/src/fins-api-services

WORKDIR /go/src/fins-api-services

RUN sed -i 's/dl-cdn.alpinelinux.org/mirror.neolabs.kz/g' /etc/apk/repositories \
    && apk --no-cache add git

RUN go get

CMD ["go", "run", "main.go"]
