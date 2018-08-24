FROM golang:alpine as builder
RUN apk add --no-cache git
RUN mkdir /src
#RUN cd /src
RUN mkdir /src/github.com
RUN mkdir /src/github.com/dawallin/
RUN mkdir /src/github.com/dawallin/playgo
ADD . /src/
COPY /cmd /src/github.com/dawallin/playgo/cmd
COPY /pkg /src/github.com/dawallin/playgo/pkg
RUN mkdir /build
ADD . /build/
WORKDIR /
ENV GOPATH /
RUN go get github.com/gorilla/mux
RUN go get github.com/mongodb/mongo-go-driver/bson
RUN go get github.com/mongodb/mongo-go-driver/mongo
RUN go build -o ./build/main ./src/github.com/dawallin/playgo/cmd/playgo
FROM alpine
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /build/main /app/
WORKDIR /app
EXPOSE 8080
ENTRYPOINT  ["./main"]
CMD [""]