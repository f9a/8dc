---
title: netcat File Server
date: 2019-04-28T23:44:43+01:00
categories:
  - linux
  - netcat
url: /post/uuid/323da666-e4cf-5f1d-a252-1f0ff661de17
---

```bash
while true; do { echo -e 'HTTP/1.1 200 OK\r\n'; cat Dateiname; } | nc -l 8080; done
```
