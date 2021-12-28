# http2snmp

[![License](https://img.shields.io/github/license/aleasoluciones/http2snmp)](https://github.com/aleasoluciones/http2snmp/blob/master/LICENSE)

Scalable snmp querier with a http interface. 
It enable to send snmp (get, getnext, walk) queries to devices using http.

See [gosnmpquerier](https://github.com/aleasoluciones/gosnmpquerier) for more information about the scalability properties and other details.

```sh
Usage of http2snmp:
  -address string
    	listen address (default "0.0.0.0")
  -contention int
    	max concurrent queries per destination (default 4)
  -maxerrors int
    	consecutive errors to mark destination as faulty (default 4)
  -port string
    	listen port (default "8080")
  -resettime int
    	time reset faulty state for a destination (seconds) (default 30)
```

## Example

```
curl http://localhost:8080/?cmd=get&destination=127.0.0.1&community=public&oid=1.3.6.1.2.1.1.1.0
```
