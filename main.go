package main

import (
	"fmt"
	"hotel-reservation/hotel"
	"hotel-reservation/parser"
	"io/ioutil"
)

var hotels = []hotel.Hotel{
	{
		Name: "Sheraton",
		Prices: hotel.PricesForCustomer{
			"Regular": hotel.Price{
				Weekday: 110,
				Weekend: 90,
			},
			"Rewards": hotel.Price{
				Weekday: 80,
				Weekend: 80,
			},
		},
		Rating: 3,
	},
	{
		Name: "Laghetto",
		Prices: hotel.PricesForCustomer{
			"Regular": hotel.Price{
				Weekday: 160,
				Weekend: 60,
			},
			"Rewards": hotel.Price{
				Weekday: 110,
				Weekend: 50,
			},
		},
		Rating: 4,
	},
	{
		Name: "Ritter",
		Prices: hotel.PricesForCustomer{
			"Regular": hotel.Price{
				Weekday: 220,
				Weekend: 150,
			},
			"Rewards": hotel.Price{
				Weekday: 100,
				Weekend: 40,
			},
		},
		Rating: 5,
	},
}

func main() {
	data, _ := ioutil.ReadFile("./input")
	input := string(data)

	reservation := parser.ParseHotelReservation(input)

	quotes := hotel.GetQuotes(hotels, reservation)
	fmt.Printf("Quotes: %#v\n", quotes)

	bestQuote := hotel.GetBetQuote(quotes)
	fmt.Printf("Best quote: %#v\n", bestQuote)
}
