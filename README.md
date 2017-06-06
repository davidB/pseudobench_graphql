My pseudo benchmark was made to QUICKLY & ROUGHTLY learn and compare few graphql lib.

WARNINGS:
* I didn't compare the featureset, the 2 quickest framework are also the less featured.
* I compare default code, not super-tuned one


# Projects'rules

* layout : `<language>/<http-lib>-<graphql-lib>``
* port: 3003
* every server should reply to the same curl request:
  ```
  curl "http://localhost:3003/go-graphql?query=\{hello\}"
  {"data":{"hello":"world"}}%
  ```
* every server are test on the same machine (server and client run on the same host)
* no logging on server-side

# Pseudo benchmark's Results

## GO : echo-go-graphql
see https://github.com/playlyfe/go-graphql

```
# launch the server
cd golang/echo-go-graphql
go run server.go

# launch the client
wrk -t12 -c400 -d30s --timeout 10s "http://localhost:3003/graphql?query={hello}"
Running 30s test @ http://localhost:3003/go-graphql?query={hello}
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     7.04ms    2.92ms 162.30ms   93.47%
    Req/Sec     4.67k   529.24     7.73k    67.61%
  1676165 requests in 30.04s, 228.59MB read
  Socket errors: connect 0, read 253, write 0, timeout 0
Requests/sec:  55799.27
Transfer/sec:      7.61MB
```

## GO : echo-graphql-go
see https://github.com/graphql-go/graphql

```
# launch the server
cd golang/echo-graphql-go
go run server.go

# launch the client
wrk -t12 -c400 -d30s --timeout 10s "http://localhost:3003/graphql-go?query={hello}"
Running 30s test @ http://localhost:3003/graphql-go?query={hello}
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    26.49ms   20.05ms 211.83ms   76.65%
    Req/Sec     1.32k   161.02     1.96k    74.94%
  473699 requests in 30.05s, 64.60MB read
  Socket errors: connect 0, read 254, write 0, timeout 0
Requests/sec:  15761.45
Transfer/sec:      2.15MB
```

## NodeJS: express-graphql

```
# launch the server
cd nodejs/express-graphql
npm install
node graphql_test.js

# launch the client
wrk -t12 -c400 -d30s --timeout 10s "http://localhost:3003/graphql?query={hello}"
Running 30s test @ http://localhost:3002/graphql?query={hello}
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    97.62ms   18.44ms 290.30ms   84.83%
    Req/Sec   320.42     40.18   464.00     84.35%
  115035 requests in 30.10s, 21.50MB read
  Socket errors: connect 0, read 375, write 0, timeout 0
Requests/sec:   3822.38
Transfer/sec:    731.63KB
```

## Rust: iron-juniper

* performance of debug ~ release x 6
* output is json prettified

```
# launch the server
cd rust/iron-juniper
cargo run --release

# launch the client
wrk -t12 -c400 -d30s --timeout 10s "http://localhost:3003/graphql?query={hello}"
Running 30s test @ http://localhost:3003/graphql?query={hello}
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.28ms  467.84us  40.20ms   91.16%
    Req/Sec     4.89k     2.23k   24.64k    62.99%
  1461612 requests in 30.10s, 206.30MB read
  Socket errors: connect 0, read 530, write 0, timeout 0
Requests/sec:  48553.76
Transfer/sec:      6.85MB
```

with debug mode (for memory and reference)
```
wrk -t12 -c400 -d30s --timeout 10s "http://localhost:3003/graphql?query={hello}"
Running 30s test @ http://localhost:3003/graphql?query={hello}
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     8.62ms    1.20ms  97.13ms   82.64%
    Req/Sec   677.19    356.88     1.85k    58.58%
  222658 requests in 30.09s, 31.43MB read
  Socket errors: connect 0, read 533, write 0, timeout 0
Requests/sec:   7398.63
Transfer/sec:      1.04MB
```

## Scala: akka-http-sangria

```
# launch the server
cd scala/akka-http-sangria
sbt run

# launch the client
wrk -t12 -c400 -d30s --timeout 10s 'http://localhost:3003/graphql?query={hello}'
Running 30s test @ http://localhost:3003/graphql?query={hello}
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   205.25ms  662.75ms   5.08s    92.91%
    Req/Sec     1.02k   478.02     3.90k    64.36%
  350211 requests in 30.03s, 53.44MB read
  Socket errors: connect 0, read 945, write 3, timeout 0
Requests/sec:  11661.75
Transfer/sec:      1.78MB
```
