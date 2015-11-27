package trigram_test

import (
	"fmt"
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

func TestMapIntersect(t *testing.T) {
	mapA := make(map[int]bool)
	mapB := make(map[int]bool)

	mapA[1] = true
	mapA[2] = true
	mapB[1] = true

	ret := IntersectTwoMap(mapA, mapB)
	if len(ret) != 1 || ret[1] == false {
		t.Errorf("Map intersect error")
	}

	ret = IntersectTwoMap(mapB, mapA)
	if len(ret) != 1 || ret[1] == false {
		t.Errorf("Map intersect error")
	}

	mapA[3] = true
	mapB[3] = true
	mapA[4] = true

	ret = IntersectTwoMap(mapB, mapA)
	if len(ret) != 2 || ret[1] == false {
		t.Errorf("Map intersect error")
	}
}

func TestTrigramIndexAdd(t *testing.T) {
	ti := NewTrigramIndex()
	ti.Add("Code is my life")
	ti.Add("Search")
	ti.Add("I write a lot of Codes")

	ret := ti.Query("Code")
	fmt.Println(ret)
	if ret[0] != 1 || ret[1] != 3 {
		t.Errorf("Basic query is failed.")
	}
}
