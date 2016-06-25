package main

func median(ls []float64) float64 {
	return ls[len(ls)/2]
}

func average(ls []float64) float {
	var tot float64
	for i := range ls {
		tot += ls[i]
	}
	return tot / float64(len(ls))
}

func main() {

}
