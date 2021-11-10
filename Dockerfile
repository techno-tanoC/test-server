FROM golang:1.16 AS build
WORKDIR /build

COPY . .
RUN CGO_ENABLED=0 go build -trimpath -ldflags "-s -w -buildid="


FROM gcr.io/distroless/static:latest
WORKDIR /app

COPY --from=build /build/test-server /app/test-server
CMD [ "/app/test-server" ]
