FROM tarantool/tarantool:2.x-centos7 as base
RUN yum install -y cmake3 make gcc

WORKDIR /opt
COPY generated .

RUN cmake3 . && cmake3 --build .
