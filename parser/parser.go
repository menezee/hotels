package parser

import (
	"regexp"
)

type Date struct {
	Day string
	Month string
	Year string
	WeekDay string
}

type Reservation struct {
	CustomerType string
	Dates []Date
}

func GetNamedMatches(matches []string, namedCaptureGroups []string) (namedMatches map[string]string) {
	namedMatches = make(map[string]string)
	for i, name := range namedCaptureGroups {
		if i != 0 && name != "" {
			namedMatches[name] = matches[i]
		}
	}

	return
}

// ExtractCustomerType returns the customer type Regular|Rewards
func ExtractCustomerType(input string) (customerType string) {
	extractCustomerType := regexp.MustCompile(`(?P<CustomerType>\w+):`)
	customerType = extractCustomerType.FindStringSubmatch(input)[1]

	return
}

// SplitDates returns a list of raw dates
func SplitDates(input string) []string {
	splitDates := regexp.MustCompile(`(?P<Date>\d{2}[A-Z][a-z]{2}\d{4})\([a-z]{3,4}\)`)
	return splitDates.FindAllString(input, 3)
}

// ExtractDates transforms a list of raw dates into a list of Dates
func ExtractDates(input string) (listOfParsedDates []Date) {
	rawDates := SplitDates(input)

	for _, rawDate := range rawDates {
		parsedDate := ParseDate(rawDate)
		listOfParsedDates = append(listOfParsedDates, parsedDate)
	}

	return
}

// ParseDate returns a formatted Date
func ParseDate(rawDate string) Date {
	extractDateInfo := regexp.MustCompile(`(?P<Day>\d{2})(?P<Month>[A-Z][a-z]{2})(?P<Year>\d{4})\((?P<WeekDay>\w{3,4})\)`)
	matches := extractDateInfo.FindStringSubmatch(rawDate)

	namedMatches := GetNamedMatches(matches, extractDateInfo.SubexpNames())
	return Date{Day: namedMatches["Day"], Month: namedMatches["Month"], Year: namedMatches["Year"], WeekDay: namedMatches["WeekDay"]}
}

// ParseHotelReservation transforms an input into a Reservation
func ParseHotelReservation(input string) Reservation {
	customerType := ExtractCustomerType(input)
	dates := ExtractDates(input)

	return Reservation{CustomerType: customerType, Dates: dates}
}
