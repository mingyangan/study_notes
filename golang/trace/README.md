# 收集
## local
1. runtime/trace
2. go run main.go 2>trace.out
## web
1. curl localhost:6060/debug/pprof/trace

# 分析
go tool trace trace.out



