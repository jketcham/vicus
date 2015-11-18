FROM google/golang:stable

MAINTAINER Jack Ketcham <jack@jackketcham.com>

RUN go get github.com/tools/godep
RUN CGO_ENABLED=0 go install -a std

ENV APP_DIR $GOPATH/src/github.com/jketcham/vicus

ENTRYPOINT ["/vicus"]
ADD . $APP_DIR

RUN cd $APP_DIR && CGO_ENABLED=0 godep go build -o /vicus -ldflags '-d -w -s'
