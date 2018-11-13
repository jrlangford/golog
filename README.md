# GoLog

Integrates loggers of several levels in a single struct.

## Troubleshooting

__Problem:__  
`go test` fails with following error:
"missing ... in args forwarded to print-like function"  
__Cause:__  
There is a bug in the `go vet` command, which is automatically run when `go test` is run, in go 1.11.2. You can read more about it [here](https://go-review.googlesource.com/c/go/+/129575/).  
__Solution:__  
Run `go test -vet=off`  
