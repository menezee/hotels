package hotel

import (
	"fmt"
	"hotel-reservation/parser"
	"testing"
)

var (
	Mon   = parser.Date{WeekDay: "mon"}
	Tues  = parser.Date{WeekDay: "tues"}
	Wed   = parser.Date{WeekDay: "wed"}
	Thurs = parser.Date{WeekDay: "thurs"}
	Fri   = parser.Date{WeekDay: "fri"}
	Sat   = parser.Date{WeekDay: "sat"}
	Sun   = parser.Date{WeekDay: "sun"}
)

func TestGetTotalDaysForReservation(t *testing.T) {
	tests := []struct {
		dates            []parser.Date
		expectedWeekdays int
		expectedWeekends int
	}{
		{dates: []parser.Date{Mon, Tues}, expectedWeekdays: 2, expectedWeekends: 0},
		{dates: []parser.Date{}, expectedWeekdays: 0, expectedWeekends: 0},
		{dates: []parser.Date{Sat}, expectedWeekdays: 0, expectedWeekends: 1},
		{dates: []parser.Date{Wed, Fri, Sat}, expectedWeekdays: 2, expectedWeekends: 1},
		{dates: []parser.Date{Thurs}, expectedWeekdays: 1, expectedWeekends: 0},
		{dates: []parser.Date{Mon, Sun}, expectedWeekdays: 1, expectedWeekends: 1},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("TestGetTotalDaysForReservation %d", i), func(t *testing.T) {
			answerWeekdays, answerWeekends := GetTotalDaysForReservation(test.dates)
			if answerWeekdays != test.expectedWeekdays {
				t.Errorf("got %d, expected %d", answerWeekdays, test.expectedWeekdays)
			}
			if answerWeekends != test.expectedWeekends {
				t.Errorf("got %d, expected %d", answerWeekends, test.expectedWeekends)
			}
		})
	}
}

var (
	hotel1 = Hotel{
		Name: "Hotel Foo",
		Prices: PricesForCustomer{
			"Regular": Price{
				Weekday: 10,
				Weekend: 8,
			},
			"Rewards": Price{
				Weekday: 5,
				Weekend: 5,
			},
		},
	}
	hotel2 = Hotel{
		Name: "Hotel Baz",
		Prices: PricesForCustomer{
			"Regular": Price{
				Weekday: 8,
				Weekend: 10,
			},
			"Rewards": Price{
				Weekday: 1,
				Weekend: 1,
			},
		},
	}
)

func TestCalculateTotalPriceForHotel(t *testing.T) {
	tests := []struct {
		customerType  string
		weekdays      int
		weekends      int
		expectedPrice int
	}{
		{customerType: "Regular", weekdays: 2, weekends: 0, expectedPrice: 20},
		{customerType: "Regular", weekdays: 2, weekends: 3, expectedPrice: 44},
		{customerType: "Regular", weekdays: 0, weekends: 1, expectedPrice: 8},
		{customerType: "Regular", weekdays: 0, weekends: 0, expectedPrice: 0},

		{customerType: "Rewards", weekdays: 2, weekends: 0, expectedPrice: 10},
		{customerType: "Rewards", weekdays: 2, weekends: 3, expectedPrice: 25},
		{customerType: "Rewards", weekdays: 0, weekends: 1, expectedPrice: 5},
		{customerType: "Rewards", weekdays: 0, weekends: 0, expectedPrice: 0},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("TestCalculateTotalPriceForHotel %d", i), func(t *testing.T) {
			answerTotalPrice := CalculateTotalPriceForHotel(test.weekdays, test.weekends, test.customerType, hotel1)
			if answerTotalPrice != test.expectedPrice {
				t.Errorf("got %d, expected %d", answerTotalPrice, test.expectedPrice)
			}
		})
	}
}

func TestGetQuotes(t *testing.T) {
	tests := []struct {
		hotels      []Hotel
		reservation parser.Reservation
		expected    []Quote
	}{
		{hotels: []Hotel{hotel1, hotel2}, reservation: parser.Reservation{CustomerType: "Regular", Dates: []parser.Date{Mon, Tues, Wed}}, expected: []Quote{{name: "Hotel Foo", price: 30}, {name: "Hotel Baz", price: 24}}},
		{hotels: []Hotel{hotel1, hotel2}, reservation: parser.Reservation{CustomerType: "Rewards", Dates: []parser.Date{Fri, Sat}}, expected: []Quote{{name: "Hotel Foo", price: 10}, {name: "Hotel Baz", price: 2}}},
		{hotels: []Hotel{hotel1}, reservation: parser.Reservation{CustomerType: "Regular", Dates: []parser.Date{}}, expected: []Quote{}},
		{hotels: []Hotel{}, reservation: parser.Reservation{CustomerType: "Regular", Dates: []parser.Date{Mon, Tues, Wed}}, expected: []Quote{}},
		{hotels: []Hotel{hotel1, hotel2}, reservation: parser.Reservation{CustomerType: "Rewards", Dates: []parser.Date{Thurs, Fri}}, expected: []Quote{{name: "Hotel Foo", price: 10}, {name: "Hotel Baz", price: 2}}},
		{hotels: []Hotel{hotel1, hotel2}, reservation: parser.Reservation{CustomerType: "Rewards", Dates: []parser.Date{Sat, Sun}}, expected: []Quote{{name: "Hotel Foo", price: 10}, {name: "Hotel Baz", price: 2}}},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("TestGetQuotes %d", i), func(t *testing.T) {
			quotes := GetQuotes(test.hotels, test.reservation)

			for i, expectedQuote := range test.expected {
				currQuote := quotes[i]
				if expectedQuote.name != currQuote.name {
					t.Errorf("got %s, expected %s", currQuote.name, expectedQuote.name)
				}
				if expectedQuote.price != currQuote.price {
					t.Errorf("got %d, expected %d", currQuote.price, expectedQuote.price)
				}
			}
		})
	}
}

func TestGetBestQuote(t *testing.T) {
	tests := []struct {
		quotes   []Quote
		expected string
	}{
		{quotes: []Quote{{name: "foo", price: 10, rating: 2}, {name: "baz", price: 12, rating: 2}, {name: "bar", price: 1, rating: 5}}, expected: "bar"},
		{quotes: []Quote{{name: "foo", price: 10, rating: 2}, {name: "baz", price: 10, rating: 3}, }, expected: "baz"},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("TestGetBestQuote %d", i), func(t *testing.T) {
			bestQuote := GetBetQuote(test.quotes)

			if bestQuote.name != test.expected {
				t.Errorf("got %s, expected %s", bestQuote.name, test.expected)
			}
		})
	}
}
