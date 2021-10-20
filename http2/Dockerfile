FROM golang as builder

ADD . /src

WORKDIR /src

RUN go build ./cmd/server && go build ./cmd/client

FROM debian:bullseye 

RUN apt-get update \
        && apt-get install -y curl ca-certificates \
  	&& apt-get clean autoclean \
	&& apt-get autoremove --yes \
	&& rm -rf /var/lib/{apt,dpkg,cache,log}

WORKDIR /src/

COPY --from=builder /src/server /usr/bin/server
COPY --from=builder /src/client /usr/bin/client

CMD ["server"]
