---
title: Archlinux add ca-certificates
date: 2015-11-25T09:07:27+01:00
categories:
  - linux
  - archlinux
url: /post/uuid/3acce0bc-f4f6-5f90-a4f1-c52bb760e703
---

Damit git, curl, etc. wissen welchen Zertifikaten Sie vertrauen können müssen benötigen Sie eine Liste der zu Vertrauenden CA Zertifikaten. Eine solche Liste kann aus den in Archlinux mitgelieferneden CA Zertifikaten.

```bash
$ trust extract --foramt=pem-bundle ca-certificates.crt
```
