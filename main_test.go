package main

import "testing"

// function TestTestValidity - unit test for testValidity
// 10 testcases
// Difficulty: Easy
// Estimated time: 5mins
// Used time: 9mins
func TestTestValidity(t *testing.T) {
	emptyStringResult := testValidity("")
	if emptyStringResult != false {
		t.Errorf(`testValidity(""), failed, output expect false instead of %v`, emptyStringResult)
	}

	missingDashStringResult := testValidity("23aa-49-bb-33-aa")
	if missingDashStringResult != false {
		t.Errorf(`testValidity("23aa-49-bb-33-aa") failed, output expect false instead of %v`, missingDashStringResult)
	}

	emptyTextStringResult := testValidity("23--48-aa-33")
	if emptyTextStringResult != false {
		t.Errorf(`testValidity("23--48") failed, output expect false instead of %v`, emptyTextStringResult)
	}

	emptyNumberStringResult := testValidity("23-ab--caba-56-haha")
	if emptyNumberStringResult != false {
		t.Errorf(`testValidity("23-ab--caba-56-haha") failed, output expect false instead of %v`, emptyNumberStringResult)
	}

	invalidTextStringResult := testValidity("23-a3b-33-caba-56-haha")
	if invalidTextStringResult != false {
		t.Errorf(`testValidity("23-a3b-bb-caba-56-haha") failed, output expect false instead of %v`, invalidTextStringResult)
	}

	invalidNumberStringResult := testValidity("2a3-ab-44-caba-56-haha")
	if invalidNumberStringResult != false {
		t.Errorf(`testValidity("23-ab-3a3-caba-56-haha") failed, output expect false instead of %v`, invalidNumberStringResult)
	}

	invalidFormatStringResult := testValidity("23-ab-cc-caba-56-haha")
	if invalidFormatStringResult != false {
		t.Errorf(`testValidity("23-ab-bb-caba-56-haha") failed, output expect false instead of %v`, invalidFormatStringResult)
	}

	invalidLastChracterStringResult := testValidity("23-ab-33-caba-56-haha3")
	if invalidLastChracterStringResult != false {
		t.Errorf(`testValidity("23-ab-cc-caba-56-haha3") failed, output expect false instead of %v`, invalidLastChracterStringResult)
	}

	invalidFirstChracterStringResult := testValidity("ab-33-caba-56-haha")
	if invalidFirstChracterStringResult != false {
		t.Errorf(`testValidity("23-ab-cc-caba-56-haha3") failed, output expect false instead of %v`, invalidFirstChracterStringResult)
	}

	validStringResult := testValidity("23-ab-33-caba-56-haha")
	if validStringResult != true {
		t.Errorf(`testValidity("23-ab-33-caba-56-haha") failed, output expect true instead of %v`, validStringResult)
	}
}

// function TestAverageNumber - unit test for averageNumber
// 2 testcases
// Difficulty: Easy
// Estimated time: 5mins
// Used time: 3mins
func TestAverageNumber(t *testing.T) {
	emptyStringResult := averageNumber("")
	if emptyStringResult != 0 {
		t.Errorf(`averageNumber("") failed, output expect 0 instead of %v`, emptyStringResult)
	}
	validStringResult := averageNumber("22-ab-33-caba-44-haha")
	if validStringResult != 33 {
		t.Errorf(`averageNumber("22-ab-33-caba-44-haha") failed, output expect 33 instead of %v`, validStringResult)
	}
}

// function TestWholeStory - unit test for wholeStory
// 2 testcases
// Difficulty: Easy
// Estimated time: 5mins
// Used time: 3mins
func TestWholeStory(t *testing.T) {
	emptyStringResult := wholeStory("")
	if emptyStringResult != "" {
		t.Errorf(`wholeStory("") failed, output expect "" instead of %v`, emptyStringResult)
	}
	validStringResult := wholeStory("22-ab-33-caba-44-haha")
	if validStringResult != "ab caba haha" {
		t.Errorf(`wholeStory("22-ab-33-caba-44-haha") failed, output expect "ab caba haha" instead of %v`, validStringResult)
	}
}

// function TestStoryStat - unit test for storyStats
// 2 testcases
// Difficulty: Easy
// Estimated time: 5mins
// Used time: 7mins
func TestStoryStat(t *testing.T) {

	shortestWord, longestWord, average, averageWords := storyStats("")
	if shortestWord != "" || longestWord != "" || average != 0 || len(averageWords) != 0 {
		t.Errorf(`storyStats("") failed, output expect "","", 0, [] instead of %v, %v, %v, %v`, shortestWord, longestWord, average, averageWords)
	}

	shortestWord, longestWord, average, averageWords = storyStats("22-abcd-33-abc-44-def-5-abcdedf-67-tugh-33-abcde")
	var result [3]string
	copy(result[:], averageWords)
	if shortestWord != "abc" || longestWord != "abcdedf" || average != float32(26)/6 || result != [3]string{"abcd", "tugh", "abcde"} {
		t.Errorf(`storyStats("22-abcd-33-abc-44-def-5-abcdedf-67-tugh") failed, output expect  abc, abcdedf, 4.3333335, [abcd tugh abcde] instead of %v, %v, %v, %v`, shortestWord, longestWord, average, averageWords)
	}
}
