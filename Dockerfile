FROM golang:latest as build

RUN mkdir -p $GOPATH/src/gitlab.com/bokjo/test_edo

COPY . /go/src/gitlab.com/bokjo/test_edo

WORKDIR /go/src/gitlab.com/bokjo/test_edo

RUN go get -v -d ./...

RUN CGO_ENABLED=0 GOOS=linux go build -o edo-api


FROM alpine:3.8

WORKDIR /root/

COPY --from=build /go/src/gitlab.com/bokjo/test_edo/edo-api .

#TODO: Obsilete - moved to docker-compose
ENV EDOAPI_USERNAME=${EDOAPI_USERNAME}
ENV EDOAPI_PASSWORD=${EDOAPI_PASSWORD}
ENV EDOAPI_DB=${EDOAPI_PASSWORD}
ENV EDOAPI_HOST=${EDOAPI_PASSWORD}
ENV EDOAPI_PORT=${EDOAPI_PORT}

EXPOSE 1234

ENTRYPOINT [ "./edo-api"]