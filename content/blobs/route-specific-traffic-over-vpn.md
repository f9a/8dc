---
title: Route Specific Network Traffic Over VPN
date: 2019-04-30T01:01:00+01:00
categories:
  - linux
  - macos
url: /post/uuid/84cb4537-e282-5183-b162-a8534a1f1c97
---

```bash
/sbin/route add -net 192.168.1.0/16 -interface utun1
```

MacOS hint:
Tunnelblick use utun1 interface
