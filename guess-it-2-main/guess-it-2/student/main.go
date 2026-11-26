package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	in := bufio.NewScanner(os.Stdin)

	var n float64 = 0
	var sumX, sumY, sumXY, sumX2, sumY2 float64 = 0, 0, 0, 0, 0

	for in.Scan() {
		line := in.Text()
		if line == "" {
			continue
		}

		y, err := strconv.ParseFloat(line, 64)
		if err != nil {
			continue
		}

		x := n
		n++
		sumX += x
		sumY += y
		sumXY += x * y
		sumX2 += x * x
		sumY2 += y * y

		var pred, delta float64

		if n <= 2 {
			pred = y
			delta = 100
		} else {
			var a, b float64

			denom := (n*sumX2 - sumX*sumX)

			if denom != 0 {
				a = (n*sumXY - sumX*sumY) / denom
				b = (sumY - a*sumX) / n
			} else {
				a = 0
				b = sumY / n
			}

			pred = a*n + b

			meanY := sumY / n
			variance := (sumY2 / n) - (meanY * meanY)
			if variance < 0 {
				variance = 0
			}
			stdDev := math.Sqrt(variance)

			delta = stdDev * 1.5

			if delta < 70 {
				delta = 70
			}
		}

		lo := int(math.Round(pred - delta))
		hi := int(math.Round(pred + delta))

		if lo < 0 {
			lo = 0
		}

		fmt.Printf("%d %d\n", lo, hi)
	}
}
