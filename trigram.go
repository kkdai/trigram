package trigram

import "fmt"

type Trigram uint32

//The trigram indexing result include all Document IDs and its Frequence in that document
type IndexResult struct {
	//Save all trigram mapping docID
	DocIDs map[int]bool

	//Save all trigram appear time for trigram deletion
	Freq map[int]int
}

// Extract one string to trigram list
// Note the Trigram is a uint32 for ascii code
func ExtractStringToTrigram(str string) []Trigram {
	if len(str) == 0 {
		return nil
	}

	var result []Trigram
	for i := 0; i < len(str)-2; i++ {
		var trigram Trigram
		trigram = Trigram(uint32(str[i])<<16 | uint32(str[i+1])<<8 | uint32(str[i+2]))
		result = append(result, trigram)
	}

	return result
}

type TrigramIndex struct {
	//To store all current trigram indexing result
	TrigramMap map[Trigram]IndexResult

	//it represent and document incremental index
	maxDocID int

	//it include currently all the doc list, it will be used when query string length less than 3
	docIDsMap map[int]bool
}

//Create a new trigram indexing
func NewTrigramIndex() *TrigramIndex {
	t := new(TrigramIndex)
	t.TrigramMap = make(map[Trigram]IndexResult)
	t.docIDsMap = make(map[int]bool)
	return t
}

//Add new document into this trigram index
func (t *TrigramIndex) Add(doc string) int {
	newDocID := t.maxDocID + 1
	trigrams := ExtractStringToTrigram(doc)
	for _, tg := range trigrams {
		var mapRet IndexResult
		var exist bool
		if mapRet, exist = t.TrigramMap[tg]; !exist {
			//New doc ID handle
			fmt.Println("tg=", tg, " not exist!")
			mapRet = IndexResult{}
			mapRet.DocIDs = make(map[int]bool)
			mapRet.Freq = make(map[int]int)
			mapRet.DocIDs[newDocID] = true
			mapRet.Freq[newDocID] = 1
		} else {
			fmt.Println("tg=", tg, " exist!")
			//trigram already exist on this doc
			if _, docExist := mapRet.DocIDs[newDocID]; docExist {
				fmt.Println("docID:", newDocID, " eixst")
				mapRet.Freq[newDocID] = mapRet.Freq[newDocID] + 1
			} else {
				//tg eixist but new doc id is not exist, add it
				mapRet.DocIDs[newDocID] = true
				fmt.Println("docID:", newDocID, " not eixst")
				mapRet.Freq[newDocID] = 1
			}
		}
		//Store or Add  result
		t.TrigramMap[tg] = mapRet
	}

	t.maxDocID = newDocID
	return newDocID
}

//Delete a doc from this trigram indexing
func (t *TrigramIndex) Delete(doc string, docID int) {
	trigrams := ExtractStringToTrigram(doc)
	for _, tg := range trigrams {
		if obj, exist := t.TrigramMap[tg]; exist {
			if freq, docExist := obj.Freq[docID]; docExist && freq > 1 {
				obj.Freq[docID] = obj.Freq[docID] - 1
			} else {
				//need remove trigram from such docID
				delete(obj.Freq, docID)
				delete(obj.DocIDs, docID)
			}
			t.TrigramMap[tg] = obj
		} else {
			//trigram not exist in map, leave
			return
		}
	}
}

//This function help you to intersect two map
func IntersectTwoMap(IDsA, IDsB map[int]bool) map[int]bool {
	var retIDs map[int]bool   //for traversal it is smaller one
	var checkIDs map[int]bool //for checking it is bigger one
	if len(IDsA) >= len(IDsB) {
		retIDs = IDsB
		checkIDs = IDsA

	} else {
		retIDs = IDsA
		checkIDs = IDsB
	}

	for id, _ := range retIDs {
		if _, exist := checkIDs[id]; !exist {
			delete(retIDs, id)
		}
	}
	return retIDs
}

//Query a target string to return the doc ID
func (t *TrigramIndex) Query(doc string) []int {
	trigrams := ExtractStringToTrigram(doc)
	if len(trigrams) == 0 {
		return nil
	}

	//Find first trigram as base for intersect
	retObj, exist := t.TrigramMap[trigrams[0]]
	if !exist {
		return nil
	}
	retIDs := retObj.DocIDs

	//Remove first one and do intersect with other trigram
	trigrams = trigrams[1:]
	for _, tg := range trigrams {
		checkObj, exist := t.TrigramMap[tg]
		if !exist {
			return nil
		}
		checkIDs := checkObj.DocIDs
		retIDs = IntersectTwoMap(retIDs, checkIDs)
	}

	return getMapToSlice(retIDs)
}

//Transfer map to slice for return result
func getMapToSlice(inMap map[int]bool) []int {
	var retSlice []int
	for k, _ := range inMap {
		retSlice = append(retSlice, k)
	}
	return retSlice
}

func (t *TrigramIndex) getAllDocIDs() []int {
	return getMapToSlice(t.docIDsMap)
}
