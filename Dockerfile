FROM busybox
ADD ./freeze-tool_linux-386 /app
CMD ["/app"]
