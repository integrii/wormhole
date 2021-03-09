# wormhole

![Wormhole Icon](https://raw.githubusercontent.com/integrii/wormhole/master/icon.png)

Very fast and simple [go](https://golang.org) program that transparently proxies incoming TCP connections to a specified ip:port.  The TCP destination is the same every time and specified by flags at startup time.  Fun for all kinds of things.

## installation (go 1.16+)
`go install github.com/integrii/wormhole`

## usage

Just execute the binary from your terminal with a `-?` flag to see help that looks like this:

```
./wormhole -?
Usage of ./wormhole:
  -from string
    	The address and port that wormhole should listen on.  Connections enter here. (default "0.0.0.0:443")
  -to string
    	Specifies the address and port that wormhole should redirect TCP connections to.  Connections exit here. (default "127.0.0.1:80")
```


### example

This opens a wormhole from all local interfaces on port `22` to `8.8.8.8` on port `2222`.

`./wormhole -from 0.0.0.0:22 -to 8.8.8.8:2222`

You can now SSH to `127.0.0.1:22` and it will come out at `8.8.8.8:2222`.


### Docker example

```docker
docker run -d --restart=always -p 8000:8000 integrii/wormhole -f 0.0.0.0:8000 -t google.com:80
```
