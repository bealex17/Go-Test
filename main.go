package main

import (
	"math/rand"
	"strconv"
	"unicode"
)

// Given a string that is a sequence of numbers followed by dash followed by text, eg: `23-ab-48-caba-56-haha`
//    * The numbers can be assumed to be unsigned integers
//    * The strings can be assumed to be ASCII sequences of arbitrary length larger than 0 (empty strings not allowed).

// function `testValidity`  takes the string as an input, and returns boolean flag `true` if the given string complies with the format, or `false` if the string does not comply
// Difficulty: Normal
// Estimated time: 15mins
// Used time: ~13mins
func testValidity(s string) bool {
	charType := 0
	isEmpty := true
	if len(s) == 0 {
		return false
	}
	// Loop through the string
	for i, r := range s {
		if i == len(s)-1 {
			if !unicode.IsLetter(r) {
				return false
			}
		}
		//check the charType: 0 - number and 1 - text
		if charType == 0 {
			if !unicode.IsDigit(r) {
				// r is not a Digit, check if empty number
				if isEmpty {
					return false
				} else {
					// check the dash
					if r != '-' {
						return false
					} else {
						// move to check the text
						charType = 1
						isEmpty = true
					}
				}
			} else {
				// r is a Digit, continue to check the next char
				isEmpty = false
			}
		} else {
			if !unicode.IsLetter(r) {
				// r is not a letter, check if empty text
				// check empty text
				if isEmpty {
					return false
				} else {
					// check the dash
					if r != '-' {
						return false
					} else {
						// move to check the number
						charType = 0
						isEmpty = true
					}
				}
			} else {
				// r is a letter, continue to check the next char
				isEmpty = false
			}
		}
	}
	return true
}

//  function `averageNumber` takes the string, and returns the average number from all the numbers
// Difficulty: Easy
// Estimated time: 5mins
// Used time: 7mins
func averageNumber(s string) float32 {
	var count, sum float32 = 0, 0
	numberString := ""
	// loop through the string to get chars
	for _, r := range s {
		if unicode.IsDigit(r) {
			// get the string of the number
			numberString += string(r)
		} else if numberString != "" {
			// convert the string of the number to number and calculate the count, sum
			count++
			number, _ := strconv.Atoi(numberString)
			sum += float32(number)
			numberString = ""
		}
	}
	if count != 0 {
		return sum / count
	}
	return 0
}

//  function `wholeStory` takes the string as an input, and returns boolean flag `true` if the given string complies with the format, or `false` if the string does not comply
// Difficulty: Easy
// Estimated time: 5mins
// Used time: 5mins
func wholeStory(s string) string {
	story := ""
	isNewWord := true
	// Loop throught the string to get letter chars
	for _, r := range s {
		if unicode.IsLetter(r) {
			if isNewWord {
				if len(story) != 0 {
					// put a space if it's a new word and not the start of the string.
					story += " "
				}
				isNewWord = false
			}
			story += string(r)
		} else {
			isNewWord = true
		}
	}
	return story
}

// function `storyStats` takes the string and returns :
//    * the shortest word
//    * the longest word
//    * the average word length
//    * the list (or empty list) of all words from the story that have the length the same as the average length rounded up and down.
// Difficulty: Normal
// Estimated time: 10mins
// Used time: 12mins
func storyStats(s string) (shortestWord string, longestWord string, average float32, averageWords []string) {
	wordMap := make(map[int][]string)
	word := ""
	sumLength := 0
	wordCount := 0
	// Loop throught the string to get the words
	for i, r := range s {
		if unicode.IsLetter(r) {
			word += string(r)
		}
		if !unicode.IsLetter(r) || i == len(s)-1 {
			if len(word) != 0 {
				// it's a new word, compare with the shortestWord and longestWord
				wordLength := len(word)
				if len(shortestWord) > wordLength || shortestWord == "" {
					shortestWord = word
				}
				if len(longestWord) < wordLength {
					longestWord = word
				}
				// calculate the sumLength and wordCount
				sumLength += wordLength
				wordCount++
				// Append to the length -> string array map
				wordMap[len(word)] = append(wordMap[wordLength], word)
				word = ""
			}
		}
	}
	if wordCount == 0 {
		return
	}
	average = float32(sumLength) / float32(wordCount)
	floorAverage := int(average)
	averageWords = wordMap[floorAverage]
	if average > float32(floorAverage) {
		// if the average is not INT -> get both rounded up and down, append the average + 1 strings to result
		averageWords = append(averageWords, wordMap[floorAverage+1]...)
	}
	return
}

// function generate takes boolean flag and generates random correct strings if the parameter is `true` and random incorrect strings if the flag is `false`.
// Difficulty: Easy
// Estimated time: 8mins
// Used time: 8mins
func generate(valid bool) string {
	result := ""
	if valid {
		const charset = "abcdefghijklmnopqrstuvwxyz" +
			"ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		times := rand.Intn(10)
		for i := 0; i < times; i++ {
			result += strconv.Itoa(rand.Intn(10000))
			result += "-"
			text := make([]byte, rand.Intn(10)+1)
			for i := range text {
				text[i] = charset[rand.Intn(len(charset))]
			}
			result += string(text)
			if i != times-1 {
				result += "-"
			}
		}
	} else {
		const charset = "abcdefghijklmnopqrstuvwxyz" +
			"ABCDEFGHIJKLMNOPQRSTUVWXYZ-123456789"
		text := make([]byte, rand.Intn(100)+1)
		for true {
			for i := range text {
				text[i] = charset[rand.Intn(len(charset))]
			}
			if testValidity(string(text)) == false {
				break
			}
		}
		result = string(text)
	}
	return result
}

func main() {
}
