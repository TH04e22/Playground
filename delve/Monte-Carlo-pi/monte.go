package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Monte Carlo method to estimate the value of pi
func main() {
	var numOfWorkers uint32 = 4
	var pointsPerWorker uint32 = 25000
	var totalPoints uint32 = numOfWorkers * pointsPerWorker
	var hitInAreaCnt uint32 = 0

	// 這邊使用 channel 來收集各個 goroutine 運算後的資料
	ch := make(chan uint32)
	var i uint32
	for i = 0; i < numOfWorkers; i++ {
		go func() {
			// rand.Rand 取得亂數的方法並不是協程安全的，這邊
			// 讓 goroutine 各自拿一個 rand 生成物件
			randSource := rand.NewSource(time.Now().Unix())
			randGen := rand.New(randSource)

			var cnt uint32 = 0

			var j uint32
			var x, y float32
			for j = 0; j < pointsPerWorker; j++ {
				x = randGen.Float32()
				y = randGen.Float32()

				if x*x+y*y < 1 {
					cnt++
				}
			}
			// 統計完後就傳到 channel
			ch <- cnt
		}()
	}

	// 取出資料並算總和
	for i := range ch {
		hitInAreaCnt += i
	}

	var pi float32 = 4 * float32(hitInAreaCnt) / float32(totalPoints)
	fmt.Println(pi)
}
