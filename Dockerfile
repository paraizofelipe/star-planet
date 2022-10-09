FROM golang:alpine as base
RUN apk add --no-cache make cmake

FROM base as build
ADD . /opt/star-planet
WORKDIR /opt/star-planet
RUN make deps && make build 

FROM scratch
COPY --from=build /opt/star-planet-api /opt/star-planet/
ENTRYPOINT ["/opt/star-planet/star-planet-api"]
