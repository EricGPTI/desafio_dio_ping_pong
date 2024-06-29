module main

go 1.22.3

require (
    ping_pong/ping v0.0.0
    ping_pong/pong v0.0.0
)
replace ping_pong/ping => ../ping
replace ping_pong/pong => ../pong
