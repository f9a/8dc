---
title: Install vncserver ubuntu
date: 2019-04-12T23:45:18+01:00
categories:
  - linux
  - ubuntu
url: /post/uuid/99991184-c7d7-5d19-8c1d-a9437e754cbf
---

## Sources

- https://www.digitalocean.com/community/tutorials/how-to-install-and-configure-vnc-on-ubuntu-18-04
- https://wiki.archlinux.org/index.php/TigerVNC
- https://www.realvnc.com/en/connect/download/viewer/

## Connect via ssh

```bash
ssh -L 5901:127.0.0.1:5901 -C -N -l gitlab-runner 195.201.56.187
```

## Misc

display-port and system-port is related the basic port is 5900 + (display-port)
