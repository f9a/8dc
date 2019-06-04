---
title: Go Blob
date: 2019-01-23T22:29:01+01:00
categories:
  - go
url: /post/uuid/47146662-f7e0-54f5-9ed6-9b32db0b7d37
---

## Important package

https://github.com/gofrs (uuid)

## HTTP Debug

GODEBUG=http2debug=2

## Password prompt

Recently I've been working with an internal Go tool that uses environment variables for accessing user credentials that are being used to authenticate. While I find that very handy, I was wondering how difficult would it be to add the functionality of providing the password manually, but in a safe (non-displaying) way. As it turned out, it cannot be any easier!

### Spoiler alert

It all comes down to a single function from `golang.org/x/crypto/ssh/terminal`whose name made me feel so embarrassed that I event doubted Go for not having this functionality. It's not a part of the language _per se_, but whenever you find something in `golang.org/x/...`, it's as good as in the box. The name I'm referring to is...

```
// golang.org/x/crypto/ssh/terminal/util.go
...
func ReadPassword(fd int) ([]byte, error) {
...
```

Yeah, it's exactly that.

### Behind the covers

What happens when you call that function? It starts and ends with a system call named `SYS_IOCTL` which is used first to get the terminal instance, then set it with its configuration altered.

How is that configuration changed? First of all, the _new_ terminal has `ECHO` flag turned off so that the input is not echoed back to the user. Then there are a few more flags:

- `ICANON` puts terminal in a _canonical_ mode where the input is not available immediately, but it has to terminated (eg. with `EOF` or new line)
- `ISIG` means that if you send eg. interrupt character, the signal is being generated (and input reading stops)
- `ICRNL` means that carriage return is interpreted as newline (so it terminates input as well)

All this sums up to create a safe environment where the user can type their password and have it secured until they decide to submit it. At no point is the value available either to the person typing it in, or the application reading it.

### Working example

In order to read secret data safely from the command line, all you need to do is execute the function mentioned in the beginning and convert a slice of bytes into a string:

```
fmt.Println("Your password: ")
bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
...
password = string(bytePassword)
fmt.Println() // it's necessary to add a new line after user's input
fmt.Printf("Your password has leaked, it is '%s', password)
```

The full source code of the more detailed example is available [on Github](https://github.com/mycodesmells/golang-examples/tree/master/misc/password-input).«

## Conventions

* https://blog.golang.org/package-names

## Testing

https://segment.com/blog/5-advanced-testing-techniques-in-go/

https://tip.golang.org/cmd/go/#hdr-Test_packages

https://github.com/golang/net/blob/master/nettest/conntest.go

Snapshots https://github.com/bradleyjkemp/cupaloy

## Internals

http://www.airs.com/blog/archives/277

## Compile Time Assertions

https://commaok.xyz/post/compile-time-assertions/

## Interfaces

https://research.swtch.com/interfaces

## Work with go code

https://github.com/golang/example/tree/master/gotypes

https://github.com/golang/tools/blob/master/cmd/stringer/stringer.go

https://godoc.org/golang.org/x/tools/go/loader

https://godoc.org/golang.org/x/tools/cmd/goimports

https://github.com/gchaincl/dotsql