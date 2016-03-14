package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func currencyToInt(in string, a *Account) int64 {
	in = strings.Replace(in, a.CurrencyCode, "", -1)
	in = strings.Replace(in, a.CurrencySymbolLeft, "", -1)
	in = strings.Replace(in, a.CurrencySymbolRight, "", -1)

	inf, _ := strconv.ParseFloat(in, 64)

	ini := int64(inf * math.Pow10(int(a.DecimalPlaces)))

	return ini
}

func currencyToStr(in int64, a *Account) string {
	total := strconv.FormatFloat(float64(in)/math.Pow(10, float64(a.DecimalPlaces)), 'f', 2, 64)

	ret := fmt.Sprintf("%s", total)
	return ret
}
