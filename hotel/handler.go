package hotel

import (
	"hotel-reservation/parser"
)

type Price struct {
	Weekday int
	Weekend int
}

type PricesForCustomer = map[string]Price

type Hotel struct {
	Name   string
	Prices PricesForCustomer
	Rating int
}

var (
	weekends = map[string]bool{
		"sat": true,
		"sun": true,
	}
	weekdays = map[string]bool{
		"mon":   true,
		"tues":  true,
		"wed":   true,
		"thurs": true,
		"fri":   true,
	}
)

type Quote struct {
	name   string
	price  int
	rating int
}

func GetTotalDaysForReservation(dates []parser.Date) (totalWeekdays int, totalWeekends int) {
	for _, date := range dates {
		if weekdays[date.WeekDay] {
			totalWeekdays++
		} else {
			totalWeekends++
		}
	}

	return
}

func CalculateTotalPriceForHotel(totalWeekdays, totalWeekends int, customerType string, hotel Hotel) int {
	priceForWeekdays := hotel.Prices[customerType].Weekday * totalWeekdays
	priceForWeekends := hotel.Prices[customerType].Weekend * totalWeekends

	return priceForWeekdays + priceForWeekends
}

func GetQuotes(hotels []Hotel, reservation parser.Reservation) (quotes []Quote) {
	totalWeekdays, totalWeekends := GetTotalDaysForReservation(reservation.Dates)

	for _, hotel := range hotels {
		totalPrice := CalculateTotalPriceForHotel(totalWeekdays, totalWeekends, reservation.CustomerType, hotel)
		quotes = append(quotes, Quote{
			name:  hotel.Name,
			price: totalPrice,
			rating: hotel.Rating,
		})
	}

	return
}

func GetBetQuote(quotes []Quote) (bestQuote Quote) {
	for i, quote := range quotes {
		isCurrCheaper := quote.price < bestQuote.price
		isCurrTheSamePrice := quote.price == bestQuote.price

		if i == 0 || isCurrCheaper {
			bestQuote = quote
		} else if isCurrTheSamePrice {
			hasCurrABetterRating := quote.rating > bestQuote.rating

			if hasCurrABetterRating {
				bestQuote = quote
			}
		}
	}

	return
}
