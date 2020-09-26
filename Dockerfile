FROM golang:1.14-buster AS builder

COPY fuzzers/ /root/fuzzers/
COPY build.sh /root/build.sh

RUN chmod +x /root/build.sh && \
 /root/build.sh
