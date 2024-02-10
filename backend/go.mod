module main

go 1.20

require (
	github.com/gorilla/websocket v1.5.1
	handler v0.0.0-00010101000000-000000000000
)

require (
	github.com/lib/pq v1.10.9 // indirect
	golang.org/x/net v0.17.0 // indirect
)

replace handler => ./handler
