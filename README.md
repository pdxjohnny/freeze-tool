Freeze Tool
---

This project provides notifications when a android device freezes or locks up.

It uses docker to compile the binaries and the main Dockerfile adds the linux
binary to the busybox image to create an extremely small final image

Building
---

```bash
./script/build
```

Running
---

```bash
./freeze-tool_linux-amd64
docker run --rm -ti pdxjohnny/freeze-tool
```

Changing The Name
---

```bash
./script/change-name $GITHUB_USERNAME $PROJECT_NAME
```


- John Andersen
