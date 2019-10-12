FROM alpine:3.10.2

RUN mkdir /app

COPY ./build/dav-proxy /app/dav-proxy
COPY ./dav /app/dav
RUN ls -la /app/dav

ENTRYPOINT [ "/app/dav-proxy" ] 