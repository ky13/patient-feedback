package main

import "testing"

func TestVerifyRating(t *testing.T) {

	positiveTestCases := []string{
		"1",
		"5",
		"10",
	}

	for _, value := range positiveTestCases {
		result := VerifyRating(value)
		if result == false {
			t.Errorf("Value '%s' should return true, got '%t'.", value, result)
		}
	}

	negativeTestCases := []string{
		"-1",
		"0",
		"11",
		"abcdef",
	}

	for _, value := range negativeTestCases {
		result := VerifyRating(value)
		if result == true {
			t.Errorf("Value '%s' should return false, got '%t'.", value, result)
		}
	}

}

func TestIsYes(t *testing.T) {

	positiveTestCases := []string{
		"y",
		"Y",
		"yes",
		"yeS",
		"Yes",
	}

	for _, value := range positiveTestCases {
		result := IsYes(value)
		if result == false {
			t.Errorf("Value '%s' should return true, got '%t'.", value, result)
		}
	}

	negativeTestCases := []string{
		"yyyyy",
		"noyes",
		"yesno",
		"0",
		"1",
		"abcdef",
	}

	for _, value := range negativeTestCases {
		result := IsYes(value)
		if result == true {
			t.Errorf("Value '%s' should return false, got '%t'.", value, result)
		}
	}
}

func TestIsNo(t *testing.T) {
	positiveTestCases := []string{
		"n",
		"N",
		"no",
		"nO",
		"No",
	}

	for _, value := range positiveTestCases {
		result := IsNo(value)
		if result == false {
			t.Errorf("Value '%s' should return true, got '%t'.", value, result)
		}
	}

	negativeTestCases := []string{
		"nnnnn",
		"noyes",
		"yesno",
		"0",
		"1",
		"abcdef",
	}

	for _, value := range negativeTestCases {
		result := IsNo(value)
		if result == true {
			t.Errorf("Value '%s' should return false, got '%t'.", value, result)
		}
	}
}

func TestPrompt(t *testing.T) {
}

func TestAsk(t *testing.T) {
}

func TestGetUnderstandingWording(t *testing.T) {

	testCases := map[string]string{
		"true":  "do",
		"false": "do not",
		"":      "do not",
		"abcde": "do not",
	}

	for input, expect := range testCases {

		result := GetUnderstandingWording(input)
		if result != expect {
			t.Errorf("Value '%s' should return '%s', got '%s'.", input, expect, result)
		}
	}
}

func TestAskRating(t *testing.T) {
}

func TestAskUnderstanding(t *testing.T) {
}

func TestAskFeeling(t *testing.T) {
}

func TestGreeting(t *testing.T) {
}

func TestSummary(t *testing.T) {
}

func Testmain(t *testing.T) {
}
