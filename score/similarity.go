package score

//ItemScore 项目评分集合
type ItemScore map[string]float64

//相同项目的集合
type SharedItems map[string]bool

//Similarity 相似度算法接口
type Similarity interface {
	Compare(x, y ItemScore) float64
}

//GetSimilarityItem 获取两个评分集合的相似项
func GetSimilarityItem(x, y ItemScore) SharedItems {
	var result = make(SharedItems)
	for xItem, _ := range x {
		if _, ok := y[xItem]; ok {
			result[xItem] = true
		}
	}
	return result
}
