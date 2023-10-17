FROM ubuntu:22.04
# FROM node:18
RUN apt update
RUN apt install -y curl 
ENV WORKDIR /server

ARG GIT_PULL_V 
ENV GIT_PULL_V=${GIT_PULL_V:-unknown}
ARG BUILD_DATE 
ENV BUILD_DATE=${BUILD_DATE:-unknown}

EXPOSE 8000
EXPOSE 8001
VOLUME /tmp/server

ADD ./build/ $WORKDIR/
ADD ./manifest/docker/docker.sh $WORKDIR/
RUN chmod +x $WORKDIR/docker.sh
###############################################################################
#                                   START
###############################################################################
WORKDIR $WORKDIR
CMD ./docker.sh