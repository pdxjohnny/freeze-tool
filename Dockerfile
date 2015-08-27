FROM pdxjohnny/adb
WORKDIR /freeze-tool
ADD ./static /freeze-tool/static
ADD ./freeze-tool_linux-amd64 /freeze-tool/run
CMD ["/freeze-tool/run"]
