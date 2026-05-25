package main

import (
	"context"
	"strings"
	"testing"
)

func TestDashboardPageRenders(t *testing.T) {
	var sb strings.Builder
	if err := dashboardPage().Render(context.Background(), &sb); err != nil {
		t.Fatalf("render: %v", err)
	}
	out := sb.String()

	for _, want := range []string{"<!doctype html>", "Dashboard", "Total Revenue", "Recent Sales", "/static/css/styles.css", "alpinejs"} {
		if !strings.Contains(strings.ToLower(out), strings.ToLower(want)) {
			t.Errorf("output missing %q", want)
		}
	}
	for _, bad := range []string{"%!", "<nil>", "ZgotmplZ", "{{"} {
		if strings.Contains(out, bad) {
			t.Errorf("output contains render artifact %q", bad)
		}
	}
}
