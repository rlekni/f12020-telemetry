FROM ubuntu:20.04
WORKDIR /home/mongobi
RUN apt-get update
RUN apt-get install -y libssl1.0.0 libssl-dev libgssapi-krb5-2 wget
RUN wget https://info-mongodb-com.s3.amazonaws.com/mongodb-bi/v2/mongodb-bi-linux-arm64-ubuntu2004-v2.14.0.tgz
RUN tar -xvzf mongodb-bi-linux-arm64-ubuntu2004-v2.14.0.tgz
WORKDIR /home/mongobi/mongodb-bi-linux-arm64-ubuntu2004-v2.14.0
RUN mkdir /logs
RUN ls
RUN echo $PATH
RUN install -m755 bin/mongo* /usr/local/bin/
EXPOSE 3307
CMD ["mongosqld", "--config=/home/mongobi/mongosqld.conf"]