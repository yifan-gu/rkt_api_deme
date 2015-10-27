rkt api-service demo
====================

### Setup your rkt dirs using `rkt install`

```shell
$ sudo rkt install
rkt directory structure successfully created.
```

### Launch your rkt API service
```shell
$ rkt api-service
2015/10/27 16:12:30 API service starting...
2015/10/27 16:12:30 API service running on localhost:15441...
```

### Build and run the example
```shell
$ godep go build
$ ./rkt_api_demo
Get info
info:<rkt_version:"0.9.0+gitf5897ef-dirty" appc_version:"0.7.1" api_version:"1.0.0-alpha" >
...
```
