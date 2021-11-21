package main

import (
	"fmt"
	"sort"
	"strings"
)

func sortString(str string) string {
	s := strings.Split(str, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func lookupList(str string, lookup []string) bool {
	for _, val := range lookup {
		if str == val {
			return true
		}
	}
	return false
}

func main() {
	listOfAnagrams := []string{"kita", "atik", "tika", "aku", "kia", "makan", "kua"}
	returnList := [][]string{}
	checkList := []string{}
	for i := 0; i < len(listOfAnagrams); i++ {
		keyStr := []string{}
		if !lookupList(listOfAnagrams[i], checkList) {
			keyStr = append(keyStr, listOfAnagrams[i])
			for j := 1; j < len(listOfAnagrams); j++ {
				if i != j && sortString(listOfAnagrams[i]) == sortString(listOfAnagrams[j]) {
					keyStr = append(keyStr, listOfAnagrams[j])
					checkList = append(checkList, listOfAnagrams[j])
				}
			}
			returnList = append(returnList, keyStr)
		}
	}
	fmt.Println(returnList)
}
