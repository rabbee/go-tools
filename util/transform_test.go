package util_test

import (
	"testing"

	"github.com/rabbee/rlib/util"
)

func TestTransformPrefs(t *testing.T) {
	in := map[string]map[string]float32{
		"小明": map[string]float32{
			"汤达人": 4.5,
			"康师傅": 3.0,
			"合味道": 4.0,
		},
		"小红": map[string]float32{
			"汤达人": 4.0,
			"令麦郎": 3.0,
			"碗仔面": 4.5,
		},
	}
	out := util.TransformPrefs(in)
	t.Log(out)
}
