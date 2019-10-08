package score

import "math"

func InvalidVector(a, b Vector) bool {
	if len(a.Values) == 0 || len(b.Values) == 0 || len(a.Values) != len(b.Values) {
		return true
	}
	return false
}

//EuclideanDistance 欧几里得距离评价 [0-1]
func EuclideanDistance(a, b Vector) float64 {
	if InvalidVector(a, b) {
		return 0
	}
	var sum float64 = 0
	for index := 0; index < len(a.Values); index++ {
		sum += math.Pow(a.Values[index]-b.Values[index], 2)
	}
	return 1 / (1 + math.Sqrt(sum))
}

//PearsonCorrelation 皮尔逊相关度评价 [0-1]
//和平均水平偏离很大
//TODO: 未完成
func PearsonCorrelation(a, b Vector) float64 {
	if InvalidVector(a, b) {
		return 0
	}
	var (
		sumA, sumB, sumASq, sumBSq, pSum float64
	)
	for index := 0; index < len(a.Values); index++ {
		sumA += a.Values[index]
		sumASq += math.Pow(a.Values[index], 2)
		sumB += b.Values[index]
		sumBSq += math.Pow(b.Values[index], 2)
		pSum += a.Values[index] * b.Values[index]
	}
	den := math.Sqrt((sumASq - math.Pow(sumA, 2)/float64(len(a.Values))) *
		(sumBSq - math.Pow(sumB, 2)/float64(len(b.Values))))
	if den == 0 {
		return 0
	}
	return (pSum - (sumA * sumB / float64(len(a.Values)))) / den
}
