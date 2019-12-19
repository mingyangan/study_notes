# local表示本地程序pprof测试
time go run main.go 
VS
time wc -w moby.txt

OR 

go test -bench=. -cpupprof=cpu.p
go test -bench=. -mempprof=mem.p

# 使用 go tool pprof $(file) 或者
# go tool pprof -http=:8080 $(file) 查看


# 阻塞分析
go test -run=XXX -bench=ClientServer$ -blockprofile=block.p net/http



## pprof流程
一、本地
1. 收集pb文件
  (1). go test -bench=. -xxxprof=xxx.p
  (2). github.com/pkg/profile

2. 查看pb内容
  (1). go tool pprof $(file)
  (2). go tool pprof -web(/-http=:8080) $file (graphviz)

二、Server
1. Web收集 + 查看
  (1). net/http/pprof
  (2). localhost:8080/debug/pprof
  (3). go tool pprof -http=:8080 $(url)