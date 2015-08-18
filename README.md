Freeze Tool
---

This project provides notifications when a android device freezes or locks up.

It uses docker to compile the binaries and the main Dockerfile adds the linux
binary to the busybox image to create an extremely small final image

Building
---

Make sure to run bower install before building the docker image

```bash
cd static
bower install
```

Also make sure to install go dependencies

```bash
godep restore
```

```bash
./script/build
```

Running
---

```bash
./freeze-tool_linux-amd64
docker run --rm -ti --privileged -v /dev/bus/usb/:/dev/bus/usb -p 7777:7777 pdxjohnny/freeze-tool
```

Changing The Name
---

```bash
./script/change-name $GITHUB_USERNAME $PROJECT_NAME
```


- John Andersen
