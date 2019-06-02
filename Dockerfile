FROM concourse/concourse

	ARG RUNC_URL=https://github.com/opencontainers/runc/releases/download/v1.0.0-rc8/runc.amd64

	RUN apt update && apt install -y iptables curl
	RUN curl -SOL ${RUNC_URL} && install -m 0755 ./runc.amd64 /usr/local/concourse/bin/runc

	ENV CONCOURSE_GARDEN_BIN 	      /usr/local/concourse/bin/gdn
	ENV CONCOURSE_GARDEN_INIT_BIN 	      /usr/local/concourse/bin/init
	ENV CONCOURSE_GARDEN_RUNTIME_PLUGIN   /usr/local/concourse/bin/runc
	ENV CONCOURSE_GARDEN_DADOO_BIN 	      /usr/local/concourse/bin/dadoo

	COPY ./bin/ /usr/local/concourse/bin/
