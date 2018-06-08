FROM ubuntu:16.04

ENV UBUNTU_RELEASE=xenial \
    SOURCE_DIR=/var/lib/s2i \
    LANG=en_US.UTF-8 \
    LANGUAGE=en_US:en \
    LC_ALL=en_US.UTF-8 \
    DEBIAN_FRONTEND=noninteractive \
    ACCEPT_EULA=y \
    STI_SCRIPTS_PATH=/usr/libexec/s2i \
    USER=ubuntu

RUN apt-get -y update && \
    apt-get -y install apt-transport-https locales &&\
    locale-gen en_US.UTF-8 &&\
    apt-get clean all

COPY files/apt/* /etc/apt/
COPY files/apt.conf.d/* /etc/apt/apt.conf.d/

# remove tsuru?
RUN apt-key add /etc/apt/tsuru.asc &&\ 
    apt-key add /etc/apt/microsoft.asc

RUN apt-get update &&\
    apt-get install -y \
        dnsutils \
        bind9utils \
        bzip2 \
        curl \
        htop \
        jq \
        lsof \
        net-tools \
        netcat-openbsd \
        rsync \
        sudo \
        telnet \
        unzip \
        vim-tiny \
        xz-utils &&\
    apt-get clean all

RUN mkdir -p ${SOURCE_DIR}/base
COPY . ${SOURCE_DIR}/base

RUN chmod 0755 ${SOURCE_DIR}/base/rc/config &&\
    chmod 0755 ${SOURCE_DIR}/base/rc/os_dependencies
 
RUN echo "export DEBIAN_FRONTEND=noninteractive" >> /etc/profile

COPY ./s2i/ $STI_SCRIPTS_PATH

LABEL io.openshift.s2i.scripts-url=image:///usr/libexec/s2i \
      io.s2i.scripts-url=image:///usr/libexec/s2i

RUN chmod 0755 -R ${STI_SCRIPTS_PATH}

RUN useradd -m ${USER} -s /bin/bash &&\
    mkdir -p /home/application/current/ &&\
    chown -R ${USER}:${USER} /home/application

RUN echo "${USER} ALL=(ALL) NOPASSWD:ALL" >> /etc/sudoers

USER 1000

RUN echo "PATH=${PATH}:/home/ubuntu/.local/bin" >> /home/ubuntu/.profile