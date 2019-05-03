---
title: UEFI
date: 2015-11-25T09:07:27+01:00
url: /post/uuid/95a902f7-fee4-5003-82a4-80c47a8164ea
---

Man kann nun über UEFI direkt eine EFI Executables starten, so kann man z. Bsp. GRUB los werden.
Dazu auf einer Festeplatte, USB-Stick, etc. eine GPT-Schema anlegen und einer Partition, sollte min. 150Mb groß sein,
den EFI Type vergeben. Anschließend diese Partion mit fat32 formatieren.
Diese Partiton kann nun dazu verwendet werden eine um den Linux-Kernel direkt oder sagen wir direkter zu starten.
Dazu muss der Kernel allerdings als EFI BOOT STUB[*] kompiliert werden dazu muss CONFIG_EFI_STUB=y gesetzt werden diese ist im Fall von ArchLinux jedoch schon der Standard.
Möchte man nun direkt ein Kernel über denn UEFI Bootmanager starten kann man mittels efibootmgr einen neuen Eintrag im UEFI Bootmanager anlegen [*].
Alternative kann man mittels bootctl auf der EFI Partition einen Simple UEFI Bootmanager anlegen um so unterschiedliche Kernel zu laden z.Bsp. unter ArchLinux den Fallback-Kernel zu starten [*].

_) https://wiki.archlinux.org/index.php/EFISTUB
_) https://wiki.archlinux.de/title/UEFI_Installation
