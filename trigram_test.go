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
	fmt.Println(mapA, mapB, ret, ret[1], ret[2])
	if len(ret) != 1 || ret[1] == false {
		t.Errorf("Map intersect error")
	}

	ret = IntersectTwoMap(mapB, mapA)
	fmt.Println(mapA, mapB, ret, ret[1], ret[2])
	if len(ret) != 1 || ret[1] == false {
		t.Errorf("Map intersect error")
	}

	mapA[3] = true
	mapB[3] = true
	mapA[4] = true

	ret = IntersectTwoMap(mapB, mapA)
	fmt.Println(mapA, mapB, ret)
	if len(ret) != 2 || ret[1] == false {
		t.Errorf("Map intersect error")
	}
}
