FROM busybox
ADD ./static /freeze-tool/static
ADD ./freeze-tool_linux-386 /freeze-tool/run
CMD ["/freeze-tool/run"]
