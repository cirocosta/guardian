FROM concourse/concourse

ARG RUNC_URL=https://github.com/opencontainers/runc/releases/download/v1.0.0-rc8/runc.amd64

RUN apt update && apt install -y iptables curl
RUN curl -SOL ${RUNC_URL} && install -m 0755 ./runc.amd64 /usr/local/bin/runc

COPY ./bin/ /usr/local/concourse/bin/
