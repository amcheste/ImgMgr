FROM centos:latest

ENV GOPATH /root/scratch

RUN yum install -y golang git make && \
    mkdir $HOME/scratch && \
    go get github.com/gorilla/mux

COPY src/* /root/scratch/src/ImgMgr/

WORKDIR /root/scratch/src/ImgMgr/

CMD make
