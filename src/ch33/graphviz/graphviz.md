1.安装graphviz
 brew install graphviz
2.$GOPATH/bin 加入到$PATH中
3.安装go-torch


通过文件方式输出Profile


$ go tool pprof prof cpu.prof
File: prof
Type: cpu
Time: Jan 5, 2022 at 10:46pm (CST)
Duration: 938.15ms, Total samples = 720ms (76.75%)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top 
Showing nodes accounting for 720ms, 100% of 720ms total
      flat  flat%   sum%        cum   cum%
     690ms 95.83% 95.83%      690ms 95.83%  main.fillMatrix
      20ms  2.78% 98.61%       20ms  2.78%  main.calculate (inline)
      10ms  1.39%   100%       10ms  1.39%  runtime.asyncPreempt
         0     0%   100%      720ms   100%  main.main
         0     0%   100%      720ms   100%  runtime.main
(pprof) list fillM
Total: 720ms
ROUTINE ======================== main.fillMatrix in /Users/chenxiang/vscode/go-th/src/ch33/graphviz/prof.go
     690ms      690ms (flat, cum) 95.83% of Total
         .          .     63:
         .          .     64:func fillMatrix(m *[row][col]int) {
         .          .     65:   s := rand.New(rand.NewSource(time.Now().UnixNano()))
         .          .     66:   for i := 0; i < row; i++ {
         .          .     67:           for j := 0; j < col; j++ {
     690ms      690ms     68:                   m[i][j] = s.Intn(100000)
         .          .     69:           }
         .          .     70:   }
         .          .     71:}
(pprof) svg




web图形化查看:
$ go tool pprof -http=":9877" cpu.prof
Serving web UI on http://localhost:9877