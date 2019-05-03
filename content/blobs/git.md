---
title: GIT
date: 2019-04-21T00:18:45+01:00
categories:
  - git
url: /post/uuid/29c9f2ae-e14f-520b-a0b9-93013b2a7aff
---

# Random git commands

## Tags

```bash
git show-ref --tags
```

## Stash

```bash
git stash apply stash@{1}
```

## Reset

```bash
git reset --merge ORIG_HEAD
```

> --merge
>
> Resets the index and updates the files in the working tree that are different between <commit> and HEAD, but keeps those which are different between the index and working tree (i.e. which have changes which have not been added).

### Links

https://git-scm.com/book/en/v2/Git-Tools-Reset-Demystified

## Show

```bash
git show -s --format=%ci <commit>
```

## Rebase

https://git-scm.com/book/en/v2/Git-Branching-Rebasing

## Branch

Delete all tracked no longer existing remote branches

```bash
git fetch --prune
```
