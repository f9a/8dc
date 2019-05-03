---
title: Create Tun/Tap devices
date: 2016-04-01T23:00:00+01:00
categories:
  - linux
url: /post/uuid/5c8c76d4-b60b-558a-85c2-a88648b86f99
---

```console
$ ip tuntap add name tap0 mode tap
$ ip link set up dev tap0
```

Beachte, das Tun/Tap Device wird erst aktiviert (UP) wenn ein Programm sich mit dem Device verbunden hat. Damit ist nicht gemeint das ein Programm aut der IP Adresse oder ähnliche lauscht.
