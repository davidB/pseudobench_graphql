# GO
## benchmark https://github.com/playlyfe/go-graphql

### launch the server

```
cd golang
go run server_graphql_hello.go
```

### launch the client

```
curl "http://localhost:3003/go-graphql?query=\{hello\}"
{"data":{"hello":"world"}}%
```

```
wrk -t12 -c400 -d30s --timeout 10s "http://localhost:3003/go-graphql?query={hello}"
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

## benchmark https://github.com/graphql-go/graphql

### launch the server

```
cd golang
go run server_graphql_hello.go
```

### launch the client

```
curl "http://localhost:3003/graphql-go?query=\{hello\}"
{"data":{"hello":"world"}}%
```

```
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

# NodeJS

## benchmark express-graphql

### launch the server

```
cd nodejs
npm install
node graphql_test.js
```

### launch the client
```
curl "http://localhost:3002/graphql?query=\{hello\}"
{"data":{"hello":"world"}}%
```
```
wrk -t12 -c400 -d30s --timeout 10s "http://localhost:3002/graphql?query={hello}"
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

# Rust

## benchmark juniper-iron

### launch the server
performance of debug ~ release x 6

```
cd rust
cargo run --release
```
### launch the client
```
curl "http://localhost:3003/graphql?query=\{hello\}"
{
  "data": {
    "hello": "world"
  }
}%
```
```
wrk -t12 -c400 -d30s --timeout 10s "http://localhost:3003/graphql?query={hello}"
Running 30s test @ http://localhost:3003/graphql?query={hello}
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.11ms  290.11us  17.75ms   85.01%
    Req/Sec     5.10k     4.51k   18.81k    75.33%
  1674581 requests in 30.04s, 236.36MB read
  Socket errors: connect 0, read 530, write 0, timeout 0
Requests/sec:  55751.66
Transfer/sec:      7.87MB
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

# Scala

## akka-http + sangria

### launch the server

```
cd scala
sbt run
```

### launch the client
```
curl "http://localhost:3003/graphql?query=\{hello\}"
{"data":{"hello":"world"}}%
```
```
wrk -t12 -c400 -d30s --timeout 10s "http://localhost:3003/graphql?query={hello}"
Running 30s test @ http://localhost:3003/graphql?query={hello}
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    83.79ms  318.58ms   3.46s    95.72%
    Req/Sec     1.43k   600.86     2.68k    66.04%
  510776 requests in 30.09s, 77.94MB read
  Socket errors: connect 0, read 649, write 0, timeout 0
Requests/sec:  16976.95
Transfer/sec:      2.59MB
```
