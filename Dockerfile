FROM iron/go:1.7

ADD . /app

WORKDIR /app

RUN chmod +x demo && ls | egrep -v "demo|resource-demos" | xargs rm -rf

ENTRYPOINT ["./demo"]