.SILENT:

run-server:
	go run entrypoints/server/main.go

action:=
arg:=
run-client:
	go run entrypoints/client/main.go $(action) $(arg)


