FROM centos:centos7

MAINTAINER lonnie.baldwin@gmail.com

RUN yum -y install epel-release golang

COPY . /src

RUNC cd /src ; go run *.go

EXPOSE 8080



