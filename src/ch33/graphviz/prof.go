package main

import (
	"log"
	"math/rand"
	"os"
	"runtime/pprof"
	"time"
)

const row = 10000
const col = 10000

func main() {
	f, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal("could not create cpu profile", err)
	}
	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal("could not start cpu profile", err)
	}
	defer pprof.StopCPUProfile()

	//测试代码段
	x := [row][col]int{}
	fillMatrix(&x)
	calculate(&x)

	f1, err := os.Create("mem.prof")
	if err != nil {
		log.Fatal("could not create mem profile", err)
	}
	// runtime.GC()
	if err := pprof.WriteHeapProfile(f1); err != nil {
		log.Fatal("could not write mem profile", err)
	}
	f1.Close()

	f2, err := os.Create("goroutine.prof")
	if err != nil {
		log.Fatal("could not create goroutine profile", err)
	}
	// go help testflag
	if gProf := pprof.Lookup("goroutine"); gProf == nil {
		log.Fatal("could not write goroutine profile", err)
	} else {
		gProf.WriteTo(f2, 0)
	}
	f2.Close()

	defer pprof.StopCPUProfile()

}

func calculate(m *[row][col]int) {
	for i := 0; i < row; i++ {
		tmp := 0
		for j := 0; j < col; j++ {
			tmp += m[i][j]
		}
	}
}

func fillMatrix(m *[row][col]int) {
	s := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			m[i][j] = s.Intn(100000)
		}
	}
}
