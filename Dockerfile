FROM busybox
WORKDIR /freeze-tool
ADD ./static /freeze-tool/static
ADD ./config.json /freeze-tool/config.json
ADD ./freeze-tool_linux-386 /freeze-tool/run
CMD ["/freeze-tool/run"]
