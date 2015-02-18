# wol
Send a Wake-on-LAN magic packet on the local subnet - useful for waking up a sleeping server in another room.

Build using [go](http://golang.org), I built it using go 1.4.2.

##Usage

```
wol mac-address
```

For example:
```
wol 00-16-cb-a4-c4-00
```

Exit code is 0 when sucessful, 1 when it fails for some reason.
