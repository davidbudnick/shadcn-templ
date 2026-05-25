package main

import (
	"math"
	"time"
)

// visitorSeries returns labelled, wavy sample data (two overlapping series)
// for the "Total Visitors" area chart, deterministic so the static export is
// stable.
func visitorSeries(days int) ([]string, [][]float64) {
	labels := make([]string, days)
	a := make([]float64, days)
	b := make([]float64, days)
	start := time.Date(2026, time.April, 5, 0, 0, 0, 0, time.UTC)
	for i := 0; i < days; i++ {
		x := float64(i)
		base := 240 + 95*math.Sin(x/4.0) + 55*math.Sin(x/1.7) + 40*math.Sin(x/9.0)
		a[i] = math.Round(math.Max(40, base+35*math.Sin(x/2.3)))
		b[i] = math.Round(math.Max(20, base*0.55+30*math.Sin(x/1.3+1)))
		labels[i] = start.AddDate(0, 0, i).Format("Jan 2")
	}
	return labels, [][]float64{a, b}
}
