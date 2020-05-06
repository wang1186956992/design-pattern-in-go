package decorator

import (
	"log"
	"math"
	"time"
)

type PiFunc func(int) float64

func LogWrapper(piFunc PiFunc, logger *log.Logger) PiFunc {
	return func(n int) float64 {
		fn := func(n int) (res float64) {
			now := time.Now()
			defer func(t time.Time) {
				logger.Printf("n值为 %v, 结果为 %v, 耗时 %v\n", n, res, time.Since(t))
			}(now)
			return piFunc(n)
		}
		return fn(n)
	}
}

func Pi(n int) float64 {
	ch := make(chan float64, n)
	for i := 0; i < n; i++ {

		go func(ch chan float64, k int) {
			ch <- 4 * math.Pow(-1, float64(k)) / float64(2*k+1)
		}(ch, i)

	}

	result := 0.0
	for i := 0; i < n; i++ {
		result += <-ch
	}
	return result
}
