package main

import (
	"fmt"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"time"
)

// CPU-intensive workload for profiling
func cpuIntensive(w http.ResponseWriter, r *http.Request) {
	result := 0
	for i := 0; i < 10000000; i++ {
		result += i
	}
	fmt.Fprintf(w, "CPU work done: %d\n", result)
}

// Memory-intensive workload for profiling
var memoryStore [][]byte

func memoryIntensive(w http.ResponseWriter, r *http.Request) {
	// Allocate 100MB
	data := make([]byte, 100*1024*1024)
	for i := range data {
		data[i] = byte(rand.Intn(256))
	}
	memoryStore = append(memoryStore, data)

	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Fprintf(w, "Allocated: %d MB, Total Alloc: %d MB\n",
		m.Alloc/1024/1024, m.TotalAlloc/1024/1024)
}

// Goroutine creation for profiling
func goroutineSpawn(w http.ResponseWriter, r *http.Request) {
	for i := 0; i < 100; i++ {
		go func(n int) {
			time.Sleep(30 * time.Second)
		}(i)
	}
	fmt.Fprintf(w, "Spawned 100 goroutines, NumGoroutine: %d\n", runtime.NumGoroutine())
}

// Blocking operation for profiling
func blockingOperation(w http.ResponseWriter, r *http.Request) {
	ch := make(chan int)
	go func() {
		time.Sleep(100 * time.Millisecond)
		ch <- 1
	}()
	<-ch
	fmt.Fprintf(w, "Blocking operation completed\n")
}

func main() {
	// Register handlers
	http.HandleFunc("/cpu", cpuIntensive)
	http.HandleFunc("/memory", memoryIntensive)
	http.HandleFunc("/goroutine", goroutineSpawn)
	http.HandleFunc("/block", blockingOperation)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `pprof Demo Server

Workload endpoints:
- /cpu - CPU-intensive operation
- /memory - Memory allocation
- /goroutine - Spawn goroutines
- /block - Blocking operations

Profile endpoints (via net/http/pprof):
- /debug/pprof/ - Index of available profiles
- /debug/pprof/profile?seconds=30 - 30-second CPU profile
- /debug/pprof/heap - Heap profile
- /debug/pprof/goroutine - Goroutine stack traces
- /debug/pprof/block - Block profile
- /debug/pprof/mutex - Mutex profile
- /debug/pprof/allocs - Memory allocation profile
- /debug/pprof/threadcreate - Thread creation profile

Usage examples:
  curl http://localhost:8080/cpu
  go tool pprof http://localhost:8080/debug/pprof/profile?seconds=30
  go tool pprof http://localhost:8080/debug/pprof/heap
  curl http://localhost:8080/debug/pprof/goroutine?debug=1
`)
	})

	fmt.Println("Starting server on :8080")
	fmt.Println("Visit http://localhost:8080 for instructions")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
