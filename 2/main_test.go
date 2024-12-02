package main

import (
	"slices"
	"testing"
)

func TestIsReportSafeWithGradualIncrease(t *testing.T) {
	report := []int{1, 3, 6, 7, 9}
	safe := isReportSafe(report)
	if !safe {
		t.Fatalf("Report with gradual increase was not considered safe")
	}
}

func TestIsReportSafeWithGradualDecrease(t *testing.T) {
	report := []int{7, 6, 4, 2, 1}
	safe := isReportSafe(report)
	if !safe {
		t.Fatalf("Report with gradual decrease was not considered safe")
	}
}

func TestIsReportSafeWithLargeIncrease(t *testing.T) {
	report := []int{1, 2, 7, 8, 9}
	safe := isReportSafe(report)
	if safe {
		t.Fatalf("Report with large increase was considered safe")
	}
}

func TestIsReportSafeWithLargeDecrease(t *testing.T) {
	report := []int{9, 7, 6, 2, 1}
	safe := isReportSafe(report)
	if safe {
		t.Fatalf("Report with large decrease was considered safe")
	}
}

func TestIsReportSafeWithIncreaseAndDecrease(t *testing.T) {
	report := []int{1, 3, 2, 4, 5}
	safe := isReportSafe(report)
	if safe {
		t.Fatalf("Report with both increase and decrease was considered safe")
	}
}

func TestIsReportSafeWithNoIncreaseOrDecrease(t *testing.T) {
	report := []int{8, 6, 4, 4, 1}
	safe := isReportSafe(report)
	if safe {
		t.Fatalf("Report with no increase and decrease was considered safe")
	}
}

func TestRemoveFromSliceRemovesElement(t *testing.T) {
	report := []int{1, 2, 3, 4, 5}
	dampenedReport := removeFromSlice(report, 3)
	expected := []int{1, 2, 3, 5}
	if !slices.Equal(dampenedReport, expected) {
		t.Fatalf("Unsafe level not removed correctly after applying problem dampener")
	}
}
