package score_test

import (
	"github.com/rabbee/rlib/score"
	"testing"
)

var (
	Similarity score.Similarity
	Rose       = score.ItemScore{
		"Night": 3.0,
		"Lady":  2.5,
		"Luck":  3.0,
	}
	Seymour = score.ItemScore{
		"Night": 3.0,
		"Lady":  3.0,
		"Luck":  1.5,
	}
	Puig = score.ItemScore{
		"Night": 4.5,
		"Luck":  3.0,
	}
	LaSalle = score.ItemScore{
		"Night": 3.0,
		"Lady":  3.0,
		"Luck":  2.0,
	}
	Matthews = score.ItemScore{
		"Night": 3.0,
		"Lady":  3.0,
	}
)

func TestEuclidean(t *testing.T) {
	Similarity = &score.Euclidean{}
	d := Similarity.Compare(Matthews, Puig)
	t.Log(d)
}

func TestPearsonCorrelation(t *testing.T) {
	Similarity = &score.Pearson{}
	d := Similarity.Compare(Matthews, Puig)
	t.Log(d)
}
