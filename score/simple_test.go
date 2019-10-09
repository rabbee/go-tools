package score_test

import (
	"github.com/rabbee/rlib/score"
	"testing"
)

var (
	p1 = score.Vector{
		Values: []float64{
			4.5,
			8,
		},
	}

	p2 = score.Vector{
		Values: []float64{
			4,
			4,
		},
	}
)

func TestEuclidean(t *testing.T) {
	d := score.EuclideanDistance(p1, p2)
	t.Log(d)
}

func TestPearsonCorrelation(t *testing.T) {
	d := score.PearsonCorrelation(p1, p2)
	t.Log(d)
}
