FROM golang:1.10 as builder
RUN go get github.com/golang/dep/cmd/dep
RUN go get -d github.com/Comradin/cybr
WORKDIR /go/src/github.com/Comradin/cybr
RUN dep ensure
RUN go install

FROM alpine:latest
LABEL maintainer Marcus Franke <marcus.franke@gmail.com>
ARG cybr_config_dir="/opt/cybr"
COPY --from=builder /go/bin/cybr /bin
ENTRYPOINT ["/bin/cybr"]
