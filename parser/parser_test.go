package parser

import "testing"

func TestSplitDates(t *testing.T) {
	mockedDate := "Regular: 16Mar2009(mon), 17Mar2009(tues), 18Mar2009(wed)"
	rawDates := SplitDates(mockedDate)
	totalParsedDates := len(rawDates)

	if totalParsedDates != 3 {
		t.Errorf("Total parsed dates = %d; expected 3", totalParsedDates)
	}
}

func TestSplitDatesWithExtraDay(t *testing.T) {
	mockedDate := "Regular: 16Mar2009(mon), 17Mar2009(tues), 18Mar2009(wed), 19Mar2009(thurs)"
	rawDates := SplitDates(mockedDate)
	totalParsedDates := len(rawDates)

	if totalParsedDates != 3 {
		t.Errorf("Total parsed dates = %d; expected 3", totalParsedDates)
	}
}
