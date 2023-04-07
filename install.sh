#!/bin/bash

# Set timezone to Japan
sed -i -e "s/ZONE=\"UTC\"/ZONE=\"Japan\"/g" /etc/sysconfig/clock
ln -sf /usr/share/zoneinfo/Japan /etc/localtime

# Update and install packages
yum update -y && yum install -y \
  sudo \
  git \
  wget \
  nano \
  vim \
  zip \
  unzip \
  tar \
  telnet \
  net-tools \
  htop \
  && yum clean all

# Install golang.
wget https://go.dev/dl/go1.20.3.linux-amd64.tar.gz
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.20.3.linux-amd64.tar.gz
echo "export PATH=$PATH:/usr/local/go/bin" >> ~/.profile
source ~/.profile

go version
