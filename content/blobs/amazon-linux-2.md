---
title: Amazon Linux 2
date: 2019-02-13T23:42:03+01:00
categories:
  - linux
  - aws
url: /post/uuid/31000b02-e894-5daf-afe6-fa5bb9685220
---

## Instance Types

https://aws.amazon.com/de/ec2/pricing/on-demand/

## AMI extras

amazon-linux-extras

## Find versions

```bash
[ec2-user ~]$ `**cat /etc/image-id**` image_name="amzn2-ami-hvm" image_version="2" image_arch="x86_64" image_file="amzn2-ami-hvm-2.0.20180810-x86_64.xfs.gpt" image_stamp="8008-2abd" image_date="20180811020321" recipe_name="amzn2 ami" recipe_id="c652686a-2415-9819-65fb-4dee-9792-289d-1e2846bd"
```

```bash
[ec2-user ~]$ `**cat /etc/system-release** `Amazon Linux 2
```

```bash
[ec2-user ~]$ `**cat /etc/os-release**` NAME="Amazon Linux" VERSION="2" ID="amzn" ID_LIKE="centos rhel fedora" VERSION_ID="2" PRETTY_NAME="Amazon Linux 2" ANSI_COLOR="0;33" CPE_NAME="cpe:2.3:o:amazon:amazon_linux:2" HOME_URL="https://amazonlinux.com/"
```

## Sicherheits-Updates

Sicherheitsupdates werden über die Paket-Repositorys sowie über aktualisierte AMIs bereitgestellt. Sicherheitswarnungen werden im [Amazon Linux-Sicherheitszentrum](https://alas.aws.amazon.com/) veröffentlicht. Weitere Informationen zu den AWS-Sicherheitsrichtlinien finden Sie im [AWS-Sicherheitszentrum](https://aws.amazon.com/security/); dort können Sie auch ein Sicherheitsproblem melden.

Amazon Linuxs ist so konfiguriert, dass Sicherheitsupdates während des Startvorgangs heruntergeladen und installiert werden. Dies wird mit Hilfe der folgenden cloud-init-Einstellung gesteuert: `repo_upgrade` Der folgende Ausschnitt aus der cloud-init-Konfiguration zeigt, wie Sie die Einstellungen in dem Benutzerdaten-Text ändern können, den Sie an die Instance-Initialisierung übergeben:

```
#cloud-config
repo_upgrade: security
```

Die möglichen Werte für `repo_upgrade` sind wie folgt:

- `security`

  Installieren ausstehender Aktualisierungen, die Amazon als Sicherheitsupdates gekennzeichnet hat

- `bugfix`

  Installieren von Aktualisierungen, die Amazon als Fehlerbehebungen gekennzeichnet hat. Fehlerbehebungen decken eine größere Anzahl von Aktualisierungen ab; dazu gehören Sicherheitsupdates und Patches für eine Reihe von anderen, kleineren Fehlern.

- `all`

  Installieren Sie alle verfügbaren Aktualisierungen, unabhängig davon, wie sie klassifiziert werden.

- `none`

  Installieren Sie keine Updates beim Starten der Instance.

Die Standardeinstellung für `repo_upgrade` ist „security”. D. h. wenn Sie in Ihren Benutzerdaten keinen anderen Wert angeben, installiert standardmäßig starten Amazon Linux führt die Sicherheitsupdates für alle zu dem Zeitpunkt installierten Pakete. informiert Sie Amazon Linux außerdem über alle Aktualisierungen für die installierten Pakete: Bei der Anmeldung wird mithilfe der `/etc/motd` Datei die Anzahl der verfügbaren Aktualisierungen angezeigt. Sie installieren diese Aktualisierungen, indem Sie den Befehl **sudo yum upgrade**in der Instance ausführen.
