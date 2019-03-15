package jewelsAndStone

import "testing"

func TestBadSolution(t *testing.T) {
	J := "aA"
	S := "aAAbbbb"
	if jCount := BadSolution(J, S); jCount != 3 {
		t.Errorf("Jewels count should be 3, found %d", jCount)
	}

	J = "z"
	S = "ZZ"
	if jCount := BadSolution(J, S); jCount != 0 {
		t.Errorf("Jewels count should be 0, found %d", jCount)
	}
}

func TestImprovedSolution(t *testing.T) {

	J := "aA"
	S := "aAAbbbb"
	if jCount := ImprovedSolution(J, S); jCount != 3 {
		t.Errorf("Jewels count should be 3, found %d", jCount)
	} else {
		t.Log(jCount)
	}

	J = "z"
	S = "ZZ"
	if jCount := ImprovedSolution(J, S); jCount != 0 {
		t.Errorf("Jewels count should be 0, found %d", jCount)
	} else {
		t.Log(jCount)
	}

}
