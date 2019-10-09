package score

import "math"

func InvalidVector(x, y Vector) bool {
	if len(x.Values) == 0 || len(y.Values) == 0 || len(x.Values) != len(y.Values) {
		return true
	}
	return false
}

//EuclideanDistance 欧几里得距离评价 [0, 1]
func EuclideanDistance(x, y Vector) float64 {
	if InvalidVector(x, y) {
		return 0
	}
	var sum float64 = 0
	for index := 0; index < len(x.Values); index++ {
		sum += math.Pow(x.Values[index]-y.Values[index], 2)
	}
	return 1 / (1 + math.Sqrt(sum))
}

//PearsonCorrelation 皮尔逊相关度评价 [-1, 1]
//和平均水平偏离很大
func PearsonCorrelation(x, y Vector) float64 {
	if InvalidVector(x, y) {
		return 0
	}
	var (
		sumX, sumY, sumXSq, sumYSq, pSum float64 = 0, 0, 0, 0, 0
		length                                   = len(x.Values)
	)
	for index := 0; index < length; index++ {
		//简单和
		sumX += x.Values[index]
		sumY += y.Values[index]
		//平方和
		sumXSq += math.Pow(x.Values[index], 2)
		sumYSq += math.Pow(y.Values[index], 2)
		//乘积和
		pSum += x.Values[index] * y.Values[index]
	}
	den := math.Sqrt(
		(sumXSq - math.Pow(sumX, 2)/float64(length)) *
			(sumYSq - math.Pow(sumY, 2)/float64(length)),
	)
	if den == 0 {
		return 0
	}
	return (pSum - (sumX * sumY / float64(length))) / den
}
