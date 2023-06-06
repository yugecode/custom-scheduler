FROM debian:stretch-slim

WORKDIR /

COPY custom-scheduler /usr/local/bin

CMD ["custom-scheduler"]