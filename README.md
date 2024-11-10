# PRACSCANNER

A port scanner written in go. Shows whether a port is open or closed on a network

This was written to just to practice my go skills and show my understanding of port scanning


Usage:

Just running pracscanner without any arguments will scan localhost

```go
    go build pracscanner.go
    ./pracscanner
```

Running it with an ip will scan all commonports on that network and display the ones that are open
```go
    ./pracscanner -ip 192.168.0.1
```

you can tell it you want to scan all 65535 ports by doing
```go
    ./pracscanner -ip 192.168.0.1 -all
```

if you want to show both open and closed port you can do
```go
    ./pracscanner -ip 192.168.0.1 -all -show
```
