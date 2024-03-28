FROM golang:1.21.5 as builder

WORKDIR "/go/src/keygen"

COPY . ./

RUN make build

FROM alpine:3.19

COPY --from=builder /go/src/keygen/build/keygen /usr/local/bin/

WORKDIR "/keygen"

CMD ["help"]
ENTRYPOINT ["keygen"]
