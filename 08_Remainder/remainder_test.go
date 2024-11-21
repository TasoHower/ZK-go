package remainder_test

import (
	"testing"

	remainder "ZK-go/08_Remainder"
)

func TestRemainder(t *testing.T) {
	var list []remainder.Congruences
	list = append(list, remainder.Congruences{
		A: 2,
		N: 3,
	})
	list = append(list, remainder.Congruences{
		A: 3,
		N: 5,
	})
	list = append(list, remainder.Congruences{
		A: 2,
		N: 7,
	})

	a := remainder.ChineseRemainderTheorem(list)
	t.Log(a)
}
