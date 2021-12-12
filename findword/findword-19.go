package findword

type WordDictionary struct {
	Words map[string]bool
}

func Constructor() WordDictionary {
	wordsDic := WordDictionary{
		Words: make(map[string]bool),
	}
	return wordsDic
}

func (this *WordDictionary) AddWord(word string) {
	this.Words[word] = true
}

func (this *WordDictionary) Search(word string) bool {
	_, ok := this.Words[word]
	if ok {
		return true
	}
	for eachword, _ := range this.Words {
		found := true
		if len(eachword) != len(word) {
			continue
		}
		for i, ch := range word {
			if string(ch) == "." {
				continue
			} else if rune(eachword[i]) != ch {
				found = false
				break
			}
		}
		if found {
			return true
		}
	}
	return false
}
