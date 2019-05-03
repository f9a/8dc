---
title: Python Funktionsparameter
date: 2015-11-27T22:19:44+01:00
categories:
  - python
url: /post/uuid/c8873657-143c-5fc5-8f5a-953cfda4aa6a
---

Es gibt in Python veränderbare Typen und unveränderbare Typen. Lass uns mit den veränderbaren Typen beginnen

Übergibt man veränderbare Typen wie zum Beispiel eine Liste wird eine Variable übergeben die die Referenz auf die Liste enthält.
Das bedeutet verändert man diese Liste über die Variable hat das auswirkungen außerhalb des Funktions scope. Wird jedoch eine neues Objekt an die Variable gebunden verändert das nicht die Variable außerhalb des Funktionsscopes.

Allgemein speichert Python nur Referenzen auf Objekte in Variablen.
