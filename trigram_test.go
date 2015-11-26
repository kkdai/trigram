package trigram_test

import (
	"testing"

	. "github.com/kkdai/trigram"
)

func TestTrigramlize(t *testing.T) {
	ret := ExtractStringToTrigram("Cod")
	if ret[0] != 4419428 {
		t.Errorf("Trigram failed, expect 4419428 get %u\n", ret[0])
	}

	//string length longer than 3
	ret = ExtractStringToTrigram("Code")
	if ret[0] != 4419428 && ret[1] != 7300197 {
		t.Errorf("Trigram failed on longer string")
	}
}
