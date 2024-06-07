FROM golang:1.22.2-bookworm AS builder
RUN mkdir /src
ADD . /src
WORKDIR /src

RUN go env -w GO111MODULE=auto
RUN go build -o main_v3 .

FROM gcr.io/distroless/base-debian12

WORKDIR /

COPY --from=builder /src/main_v3 /main
EXPOSE 3000
ENTRYPOINT ["/main"]