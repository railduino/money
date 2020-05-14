// Copyright (c) 2020 Volker Wiegand <volker@railduino.de>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package money

import ()

type Currency struct {
	Code   string
	Symbol string `json:"-"`
	Sample string `json:"-"`
	Point  string `json:"-"`
	Group  string `json:"-"`
}

type Money struct {
	Currency
	Value int64
}

var Currencies = map[string]Currency{
	"EUR": Currency{
		Code:   "EUR",
		Symbol: "€",
		Sample: "$1",
		Point:  ",",
		Group:  ".",
	},
	// Add more currencies here
}

func New(code string, value int64) *Money {
	for _, curr := range Currencies {
		if code == curr.Code || code == curr.Symbol {
			return &Money{
				Currency: curr,
				Value:    value,
			}
		}
	}

	curr := Currencies["EUR"]
	return &Money{
		Currency: curr,
		Value:    value,
	}
}
