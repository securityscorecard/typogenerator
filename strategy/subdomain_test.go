package strategy_test

import (
	"testing"

	"github.com/Zenithar/typogenerator/strategy"
)

func TestSubdomain(t *testing.T) {
	out, err := strategy.SubDomain.Generate("zenithar", "")
	if err != nil {
		t.Fail()
		t.Fatal("Error should not occurs !", err)
	}

	if len(out) == 0 {
		t.FailNow()
	}

	if len(out) != 7 {
		t.FailNow()
	}
}
