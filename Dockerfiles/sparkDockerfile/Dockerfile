FROM apache/spark:3.5.0

USER root 

RUN apt-get update && \
    apt-get install -y wget tar gcc && \
    rm -rf /var/lib/apt/lists/* && \
    wget https://golang.org/dl/go1.21.3.linux-amd64.tar.gz -O /tmp/go.tar.gz && \
    tar -C /usr/local -xzf /tmp/go.tar.gz && \
    rm /tmp/go.tar.gz


RUN python3 -m pip install --no-cache-dir pyspark

RUN mkdir -p /opt/spark/generation && \
    mkdir -p /opt/spark/input_data && \
    mkdir -p /opt/spark/jobs

ENV PATH="/usr/local/go/bin:${PATH}"
ENV GOPATH="/go"
ENV PATH="${GOPATH}/bin:${PATH}"

RUN mkdir -p /home/spark && \
    chown spark:spark /home/spark && \
    chmod 755 /home/spark

USER spark
WORKDIR /opt/spark