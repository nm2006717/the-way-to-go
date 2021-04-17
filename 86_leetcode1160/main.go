package main

func main() {
	words:=[]string{"cat","bt","hat","tree"}
	chars:="atach"

	_=countCharacters(words,chars)
}



func countCharacters(words []string, chars string) int {
	maps := make(map[rune]int)

	for _,v:=range chars{
		if _,ok:=maps[v];ok{
			maps[v]++
		}else{
			maps[v] = 1
		}

	}

	sumLen:=0
	for _,v:=range words {
		mps:=make(map[rune]int)
		for _,c:=range v{
			if _,ok:=mps[c];ok{
				mps[c]++
			}else{
				mps[c] = 1
			}
		}
		isCharWord:=true
		for k,_:=range mps{
			if mps[k] > maps[k] {
				isCharWord = false
				break
			}
		}
		if isCharWord{
			sumLen = sumLen + len(v)
		}

	}

	return sumLen
}