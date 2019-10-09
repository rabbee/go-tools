package score

import "math"

//Pearson 皮尔逊相关系数
//和平均水平偏离很大
type Pearson struct {
}

//Compare 皮尔逊相关度评价 [-1, 1]
func (a *Pearson) Compare(x, y ItemScore) float64 {
	items := GetSimilarityItem(x, y)
	var (
		sumX, sumY, sumXSq, sumYSq, pSum float64 = 0, 0, 0, 0, 0
		length                                   = len(items)
	)
	for item, _ := range items {
		xScore, xExist := x[item]
		yScore, yExist := y[item]
		if xExist && yExist {
			//简单和
			sumX += xScore
			sumY += yScore
			//平方和
			sumXSq += math.Pow(xScore, 2)
			sumYSq += math.Pow(yScore, 2)
			//乘积和
			pSum += xScore * yScore
		}
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
