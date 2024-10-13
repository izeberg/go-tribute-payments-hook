FROM golang:1.23 AS builder

WORKDIR /build
COPY app/go.mod app/go.sum /build/
RUN cd /build; go mod download

COPY app /build/
RUN cd /build/app; ls; go build -o /usr/bin/tribute-hook

FROM golang:1.23

COPY --from=builder /usr/bin/tribute-hook /usr/bin/

CMD ["/usr/bin/tribute-hook"]