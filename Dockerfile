FROM golang:alpine as base
RUN apk add --no-cache make cmake

FROM base as build
ADD . /opt/star-planet-api
WORKDIR /opt/star-planet-api
RUN make deps && make build 

FROM alpine
COPY --from=build /opt/star-planet-api /opt/star-planet/
ENTRYPOINT ["/opt/star-planet/star-planet-api"]
