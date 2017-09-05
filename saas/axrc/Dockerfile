FROM debian:8.9

RUN apt-get update && \
    apt-get install -y ca-certificates && \
    rm -rf /var/lib/apt/lists/*

CMD /axrc/bin/server

COPY ./bin/* /axrc/bin/