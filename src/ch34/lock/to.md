```azure
$ go test -bench=. -cpuprofile=cpu.prof
goos: darwin
goarch: arm64
pkg: go_ph/src/ch34/lock
BenchmarkLockFree-8          430           2566145 ns/op
BenchmarkLock-8                4         320640198 ns/op
PASS
ok      go_ph/src/ch34/lock     4.579s



$ go tool pprof cpu.prof                                 
Type: cpu
Time: Jan 10, 2022 at 11:11am (CST)
Duration: 4.14s, Total samples = 20.87s (503.57%)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top
Showing nodes accounting for 20.79s, 99.62% of 20.87s total
Dropped 17 nodes (cum <= 0.10s)
Showing top 10 nodes out of 22
      flat  flat%   sum%        cum   cum%
     7.41s 35.51% 35.51%      7.41s 35.51%  sync.(*RWMutex).RLock (inline)
     7.10s 34.02% 69.53%      7.10s 34.02%  sync.(*RWMutex).RUnlock
     4.36s 20.89% 90.42%      4.36s 20.89%  runtime.pthread_cond_wait
     1.19s  5.70% 96.12%      1.19s  5.70%  runtime.mapaccess2_faststr
     0.30s  1.44% 97.56%      0.30s  1.44%  runtime.pthread_cond_signal
     0.28s  1.34% 98.90%     14.95s 71.63%  go_ph/src/ch34/lock.lockAccess.func1
     0.15s  0.72% 99.62%      1.18s  5.65%  go_ph/src/ch34/lock.lockFreeAccess.func1
         0     0% 99.62%      4.42s 21.18%  runtime.findrunnable
         0     0% 99.62%      4.42s 21.18%  runtime.goexit0
         0     0% 99.62%      0.29s  1.39%  runtime.goready.func1
(pprof)  list lockgo_ph/src/ch34/lock.lockAccess.func1
Total: 20.87s
(pprof) list go_ph/src/ch34/lock.lockAccess.func1
Total: 20.87s
ROUTINE ======================== go_ph/src/ch34/lock.lockAccess.func1 in /Users/chenxiang/GolandProjects/go-th/src/ch34/lock/lock_test.go
     280ms     14.95s (flat, cum) 71.63% of Total
         .          .     38:   var wg sync.WaitGroup
         .          .     39:   wg.Add(NUM_OF_READER)
         .          .     40:   m := new(sync.RWMutex)
         .          .     41:   for i := 0; i < NUM_OF_READER; i++ {
         .          .     42:           go func() {
     240ms      240ms     43:                   for j := 0; j < READ_TIMES; j++ {
         .      7.41s     44:                           m.RLock()
      30ms      190ms     45:                           _, ok := cache["a"]
         .          .     46:                           if !ok {
         .          .     47:                                   fmt.Println("nothing")
         .          .     48:                           }
      10ms      7.11s     49:                           m.RUnlock()
         .          .     50:                   }
         .          .     51:                   wg.Done()
         .          .     52:           }()
         .          .     53:   }
         .          .     54:   wg.Wait()
(pprof) 


```