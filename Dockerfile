FROM alpine:edge AS build
LABEL maintainer="flekystyley@gmail.com"
WORKDIR /go/src/github.com/OkumuraShintarou/peace
RUN apk update && \
    apk upgrade && \
    apk add --update --no-cache go=1.11.1-r0 gcc g++ make openssh git libc-dev sudo

ENV GOPATH /go
ADD . . 
RUN make build

FROM alpine:edge
LABEL maintainer="flekystyley@gmail.com"
WORKDIR /app
EXPOSE 9090
ENV GOPATH /go
COPY conf ./conf
COPY assets ./assets
COPY db ./db
COPY docs/zoneinfo.zip /usr/lib/go/lib/time/zoneinfo.zip
RUN apk add --update --no-cache ca-certificates
COPY --from=build /go/src/github.com/popshootjapan/keyakizaka-server/bin/server bin/server
CMD ["bin/server"]