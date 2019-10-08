package score_test

import (
	"github.com/rabbee/rlib/score"
	"testing"
)

func TestEuclidean(t *testing.T) {
	p1 := score.Vector{
		Values: []float64{
			1.99,
			1.01,
		},
	}
	p2 := score.Vector{
		Values: []float64{
			2,
			1,
		},
	}
	d := score.EuclideanDistance(p1, p2)
	t.Log(d)
}

func TestPearsonCorrelation(t *testing.T) {
	p1 := score.Vector{
		Values: []float64{
			1.99,
			1.01,
		},
	}
	p2 := score.Vector{
		Values: []float64{
			2,
			1,
		},
	}
	d := score.PearsonCorrelation(p1, p2)
	t.Log(d)
}
