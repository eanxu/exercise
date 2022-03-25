FROM centos:latest
# 环境变量
ENV LANG="en_US.UTF-8"

RUN cp -r -f /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN mkdir /etc/drem
#COPY build/drem /usr/bin/drem
#COPY cuttile_demo/internal/utils_demo.toml /etc/drem
ENTRYPOINT ["/usr/bin/drem"]