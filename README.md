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

##License
(The MIT License)

Copyright (c) 2014 Alea Soluciones SLL

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the 'Software'), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED 'AS IS', WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
