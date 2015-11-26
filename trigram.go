package trigram

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

	maxDocID int
}

func NewTrigramIndex([]string) *TrigramIndex {
	t := new(TrigramIndex)
	t.TrigramMap = make(map[Trigram]IndexResult)
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
			mapRet = IndexResult{}
			mapRet.DocIDs = make(map[int]bool)
			mapRet.Freq = make(map[int]int)
			mapRet.DocIDs[newDocID] = true
			mapRet.Freq[newDocID] = 1
		} else {
			//trigram already exist on this doc
			if _, docExist := mapRet.DocIDs[newDocID]; docExist {
				mapRet.Freq[newDocID] = mapRet.Freq[newDocID] + 1
			}
		}
		//Store or Add  result
		t.TrigramMap[tg] = mapRet
	}

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

func (t *TrigramIndex) Query(doc string) []int {

	return nil
}

func (t *TrigramIndex) getAllDocIDs() []int {
	var retIDs []int
	for i := 0; i <= t.maxDocID; i++ {
		retIDs = append(retIDs, i)
	}
	return retIDs
}
