# Refactoring code

Please refactor the code below to make it more concise, efficient and readable with good logic flow.

```go
func findFirstStringInBracket(str string) string {
	if len(str) > 0 {
		indexFirstBracketFound := strings.Index(str, "(")
		if indexFirstBracketFound >= 0 {
			runes := []rune(str)
			wordsAfterFirstBracket := string(runes[indexFirstBracketFound:len(str)])
			indexClosingBracketFound := strings.Index(wordsAfterFirstBracket, ")")
			if indexClosingBracketFound >= 0 {
				runes := []rune(wordsAfterFirstBracket)
				return string(runes[1 : indexClosingBracketFound-1])
			} else {
				return ""
			}
		} else {
			return ""
		}
	} else {
		return ""
	}
	return ""
}
```

## Solution

We can make this code more readable by finding a more efficient way to get the string value between bracket by slicing the string based on the bracket pair index.

```go

func findFirstStringInBracket(str string) string {
	if len(str) > 0 {
		indexFirstBracket := strings.Index(str, "(")
		indexClosingBracket := strings.Index(str, ")")
		wordsInBracket := str[indexFirstBracket+1 : indexClosingBracket-1]

		return wordsInBracket
	}
	return ""
}

```

Try to run the [main.go](./main.go) to check the output of the function

```bash
go run main.go [INPUT]
```
