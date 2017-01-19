ROM golang:1.7.4-alpine
MAINTAINER Khoi Nguyen <minhkhoi@siliconstraits.com>
ENV APP_HOME $GOPATH/src/app
WORKDIR $APP_HOME
RUN apk update && apk upgrade && apk add --no-cache postgresql-client bash git openssh openjdk7-jre
RUN go get github.com/tools/godep