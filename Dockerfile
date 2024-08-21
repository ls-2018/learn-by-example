FROM registry.cn-hangzhou.aliyuncs.com/acejilam/ubuntu:22.04
ENV PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/root/.eunomia/bin
#FROM registry.cn-hangzhou.aliyuncs.com/acejilam/mygo:v1.22.2

COPY resources/init.sh /tmp/init.sh
RUN bash /tmp/init.sh

RUN apt install curl clang-format cmake -y

ENTRYPOINT ["bash"]
