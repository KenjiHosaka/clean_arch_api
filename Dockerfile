FROM golang:1.12

RUN mkdir -p /go/src/clean_arch_api

WORKDIR /go/src/clean_arch_api
COPY ./backend /go/src/clean_arch_api/backend
RUN go get -u github.com/golang/dep/...
RUN go get -u gopkg.in/godo.v2/cmd/godo
WORKDIR /go/src/clean_arch_api/backend

EXPOSE 8000
CMD ["/go/bin/godo", "server", "--watch"]
