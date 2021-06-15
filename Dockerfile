FROM golang:1.16.5-buster AS builder

COPY fuzzers/ /root/fuzzers/
COPY build.sh /root/build.sh

RUN chmod +x /root/build.sh && \
 /root/build.sh
