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

import (
	"fmt"
)

func abs(n int64) int64 {
	if n < 0 {
		return -n
	}
	return n
}

func (m *Money) String() string {
	c := m.Currency
	major := m.Value / 100
	minor := abs(m.Value % 100)
	value := fmt.Sprintf("%d%s%02d", major, c.Point, minor)

	data := c.Sample
	data = strings.Replace(data, "$", c.Symbol, 1)
	data = strings.Replace(data, "1", value, 1)

	return data
}
