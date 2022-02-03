## Ultra fast ping over the range of IP addresses

It takes `baseIP` as the starting point for the IP addresses to ping, `range` parameter is used to tell how many IP addresses should be pinged. Also `pauseAfter` paramter can be used to add delay in ping after nth packet. By default one ICMP packet with timeout=1000 msec is sent for every IP address.
The results are available in a second, and you can ping milions of IP addresses. When pinging on Linux machine you need this command
`sudo sysctl -w net.ipv4.ping_group_range="0 2147483647"`. Testing on macOS did not require this command.

```
~/wrk/goPing$ goPing -h
Usage of goPing:
  -baseIP string
    	First IP address
  -pauseAfter int
    	Pause after nth icmp packet for 2*timeout sec
  -range int
    	Incremental Range to the BaseIP (default 10)
  -timeout int
    	Ping Timeout in msec (default 1000)
```


```
$ goPing -baseIP 192.168.1.0 -range 10                                                                                                      
goPing version 1.0
    192.168.1.1 is alive
    192.168.1.8 is dead
    192.168.1.7 is dead
    192.168.1.2 is dead
    192.168.1.3 is dead
    192.168.1.4 is dead
   192.168.1.10 is dead
    192.168.1.5 is dead
    192.168.1.9 is dead
    192.168.1.6 is dead
```    
    
    
A Linux binary (called goPingLinux) is provided here as part of repo, otherwise this Go program can be used to build binaries for macOS and Windows.


```
For MacOS
go env -w GOOS=darwin
go build


For Windows
go env -w GOOS=windows
go build
```

