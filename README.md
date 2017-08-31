## Task
Make a command line script:	read comma separated numbers from stdin and check are this numbers numerical progression.

## Examples of usage
```
go run cmd/cmd.go 2,4,8
go run cmd/cmd.go 2,4,8,
go run cmd/cmd.go    2,4,8,
go run cmd/cmd.go "2,4,8,           "
go run cmd/cmd.go 0.1,0.2,0.3
go run cmd/cmd.go .1,.2,.3
go run cmd/cmd.go .1,0.2,.3
go run cmd/cmd.go 0001,2.000,3
go run cmd/cmd.go -1,5,-25
go run cmd/cmd.go "0001,2.000, 3"
```