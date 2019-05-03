---
title: Fix pyenv Error
date: 2019-03-12T12:45:57+01:00
categories:
  - macos
  - python
url: /post/uuid/58715ca0-4e5c-58e0-8601-74f4888392dd
---

## Error

```
python-build: use openssl 1.1 from homebrew
python-build: use readline from homebrew
Downloading Python-3.7.3.tar.xz...
-> https://www.python.org/ftp/python/3.7.3/Python-3.7.3.tar.xz
Installing Python-3.7.3...
python-build: use readline from homebrew

BUILD FAILED (OS X 10.14.3 using python-build 20180424)

Inspect or clean up the working tree at /var/folders/jn/h5x2wj7j04d1sdtdjcxltf4h0000gn/T/python-build.20190410174719.88490
Results logged to /var/folders/jn/h5x2wj7j04d1sdtdjcxltf4h0000gn/T/python-build.20190410174719.88490.log

Last 10 log lines:
  File "/private/var/folders/jn/h5x2wj7j04d1sdtdjcxltf4h0000gn/T/python-build.20190410174719.88490/Python-3.7.3/Lib/ensurepip/__main__.py", line 5, in <module>
    sys.exit(ensurepip._main())
  File "/private/var/folders/jn/h5x2wj7j04d1sdtdjcxltf4h0000gn/T/python-build.20190410174719.88490/Python-3.7.3/Lib/ensurepip/__init__.py", line 204, in _main
    default_pip=args.default_pip,
  File "/private/var/folders/jn/h5x2wj7j04d1sdtdjcxltf4h0000gn/T/python-build.20190410174719.88490/Python-3.7.3/Lib/ensurepip/__init__.py", line 117, in _bootstrap
    return _run_pip(args + [p[0] for p in _PROJECTS], additional_paths)
  File "/private/var/folders/jn/h5x2wj7j04d1sdtdjcxltf4h0000gn/T/python-build.20190410174719.88490/Python-3.7.3/Lib/ensurepip/__init__.py", line 27, in _run_pip
    import pip._internal
zipimport.ZipImportError: can't decompress data; zlib not available
make: *** [install] Error 1
```

```

sudo installer -pkg /Library/Developer/CommandLineTools/Packages/macOS_SDK_headers_for_macOS_10.14.pkg -target /
```
