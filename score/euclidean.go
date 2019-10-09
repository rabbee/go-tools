package score

import "math"

//Euclidean 欧几里得距离
type Euclidean struct {
}

//Compare 欧几里得距离评价 [0, 1]
func (e *Euclidean) Compare(x, y ItemScore) float64 {
	items := GetSimilarityItem(x, y)
	var sum float64 = 0
	for item, _ := range items {
		xScore, xExist := x[item]
		yScore, yExist := y[item]
		if xExist && yExist {
			sum += math.Pow(xScore-yScore, 2)
		}
	}
	return 1 / (1 + math.Sqrt(sum))
}
