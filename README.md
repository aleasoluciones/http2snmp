http2snmp
=========
Scalable snmp querier with a http interface. 
It enable to send snmp (get, getnext, walk) queries to devices using http.

See [gosnmpquerier](https://github.com/aleasoluciones/gosnmpquerier) for more information about the scalability properties and other details.

```sh
Usage of http2snmp:
  -address="0.0.0.0": listen address
  -contention=4: max concurrent queries per destination
  -maxerrors=4: consecutive errors to mark destination as faulty
  -port="8080": listen port
  -resettime=30: time reset faulty state for a destination (seconds)
```


Get example
===========
```
curl 'http://localhost:8080/?cmd=get&destination=127.0.0.1&community=public&oid=1.3.6.1.2.1.1.1.0'
```
