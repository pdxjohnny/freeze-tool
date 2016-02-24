logger
---

A new logger should be started for every device. It can be started manually
or the manager can start it.

Loggers will kill logging processes and restart them as a device connects and
disconnects.

How it works
---

When logger is started it is given a device id to listen to. Commands to log
are sent to the logger. These commands are run in adb shell like logcat or
dmesg. When the device connects it runs those commands on the device and pipes
the output to a file. When a client connects requesting the logs for that
command logger runs `tail -f` on the file to send the previous output plus the
new output to the client.

API
---

All API requests must have the variable `command` set.
For example `http://url/&command=_self`. `_self` is the command to use when you
want to talk to the logger server and not the commands that it is running.

* GET /status/?
* GET /create/?
* GET /logger/?
