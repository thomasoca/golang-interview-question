# Logic test

Given an array of strings:

```
['kita', 'atik', 'tika', 'aku', 'kia', 'makan', 'kua']
```

Group each word based on its anagram group, i.e.

```
# Expected Outputs
[["kita", "atik", "tika"],
["aku", "kua"],
["makan"],
["kia"]]
```

## Solution

To check if a word is anagram of another word, we can sort and compared each letter of the word by creating a function to sort a string:

```go
func sortString(str string) string {
	s := strings.Split(str, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
```

Then to compare two words from the array of strings, we can use a nested loop and compare the sorted words:

```go
for i := 0; i < len(listOfAnagrams); i++ {
    for j := 1; j < len(listOfAnagrams); j++ {
	    if i != j && sortString(listOfAnagrams[i]) == sortString(listOfAnagrams[j]) {
	        //Do logic here
	    }
    }
}
```

To produce the same output as expected, we can append the anagrams into an empty slice:

```go
for i := 0; i < len(listOfAnagrams); i++ {
    keyStr := []string{}
    for j := 1; j < len(listOfAnagrams); j++ {
	    if i != j && sortString(listOfAnagrams[i]) == sortString(listOfAnagrams[j]) {
		keyStr = append(keyStr, listOfAnagrams[j])
	    }
    }
}
```

To avoid duplicate grouping do to the loop, we can implement a lookup function to check if the word is already on the anagram group:

```go
func lookupList(str string, lookup []string) bool {
	for _, val := range lookup {
		if str == val {
			return true
		}
	}
	return false
}
```

The final main code can be write as:

```go
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
```

The complete solution can be seen by running the [main.go](./main.go) file on this directory.

```bash
go run main.go
```
