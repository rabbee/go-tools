package util

func TransformPrefs(in map[string]map[string]float32) (out map[string]map[string]float32) {
	out = make(map[string]map[string]float32)
	for person, items := range in {
		for item, _ := range items {
			if out[item] == nil {
				out[item] = make(map[string]float32)
			}
			out[item][person] = in[person][item]
		}
	}
	return
}
