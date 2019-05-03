---
title: Format SD-Card wit fat32 (macOS)
date: 2019-04-26T23:52:00+01:00
categories:
  - macos
url: /post/uuid/3407fb82-466c-5259-86fb-f39e35e6c22d
---

_Be sure sd-card adapter is not locked!_

```bash
sudo diskutil eraseDisk FAT32 RASPBIAN MBRFormat /dev/disk2
```
