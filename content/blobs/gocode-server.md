---
title: Vim gocode-server
date: 2018-06-01T15:32:43+01:00
categories:
  - vim
  - gocode
  - errors
url: /post/uuid/2dfeff23-65f9-5949-9953-4451e99e012e
---

gocode läuft als server im hintergrund

Kann beendet werden

```bash
$ gocode close
```

Bei mir kam es nach updates von go-vim machmal zu Problemen. Zum Beispiel wurde in vim nur noch PANIC PANIC PANIC angezeigt. Um herauszufinden welches was gocode für ein Fehler wirft kannes nützlich sein gocode im Vordergrund zu starten

```bash
$ gocode -s
```
