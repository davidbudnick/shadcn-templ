package ui

import (
	"encoding/json"
	"time"
)

// calendarMonthModel is the serialized representation of a single month used to
// drive client-side (Alpine) month navigation without a full page reload.
type calendarMonthModel struct {
	Year  int     `json:"year"`
	Month int     `json:"month"` // 1-12
	Name  string  `json:"name"`  // e.g. "May 2026"
	Weeks [][]int `json:"weeks"` // 0 = padding day from an adjacent month
}

// calendarMonthRange builds a contiguous slice of month models centered on
// (year, month), spanning `span` months on each side. The returned index is the
// position of the requested (year, month) within the slice so the Alpine state
// can start there. This lets month navigation be entirely client-side: stepping
// the index swaps the precomputed grid instead of reloading the page.
func calendarMonthRange(year int, month time.Month, span int) ([]calendarMonthModel, int) {
	if span < 0 {
		span = 0
	}
	start := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC).AddDate(0, -span, 0)
	models := make([]calendarMonthModel, 0, span*2+1)
	for i := 0; i <= span*2; i++ {
		m := start.AddDate(0, i, 0)
		models = append(models, calendarMonthModel{
			Year:  m.Year(),
			Month: int(m.Month()),
			Name:  calendarMonthName(m.Year(), m.Month()),
			Weeks: calendarDays(m.Year(), m.Month()),
		})
	}
	return models, span
}

// calendarModelJSON serializes the month range to JSON for embedding in the
// Alpine x-data attribute.
func calendarModelJSON(models []calendarMonthModel) string {
	b, err := json.Marshal(models)
	if err != nil {
		return "[]"
	}
	return string(b)
}
