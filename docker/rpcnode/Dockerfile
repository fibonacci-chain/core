FROM ubuntu:latest
ENV DEBIAN_FRONTEND=noninteractive
ENV TZ=Asia/Shanghai
RUN sed -i 's/archive.ubuntu.com/mirrors.tuna.tsinghua.edu.cn/g' /etc/apt/sources.list
RUN apt-get update && \
    apt-get install -y make git vim && \
    apt-get install -y jq python3 curl wget gcc && \
    wget https://dl.google.com/go/go1.20.linux-amd64.tar.gz && \
        tar -C /usr/local -xzf go1.20.linux-amd64.tar.gz && \
        rm go1.20.linux-amd64.tar.gz
SHELL ["/bin/bash", "-c"]


VOLUME [ "/sei-protocol" ]
VOLUME [ "/root/go/pkg/mod" ]
WORKDIR /fibonacci-chain/core

EXPOSE 26656 26657 26658 9090 9091

CMD ["/usr/bin/deploy.sh"]

STOPSIGNAL SIGTERM

COPY scripts/deploy.sh /usr/bin/deploy.sh
COPY scripts/step0_build.sh /usr/bin/build.sh
COPY scripts/step1_configure_init.sh /usr/bin/configure_init.sh
COPY scripts/step2_start_sei.sh /usr/bin/start_sei.sh
