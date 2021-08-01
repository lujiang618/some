package utils_test

import (
	"fmt"
	"some/api/pkg/utils"
	"testing"
)

func TestGetCurrentYearMonth(t *testing.T) {
	ym := utils.GetCurrentYearMonth()
	fmt.Println("ym:", ym)
}

func TestGetLastYearMonth(t *testing.T) {
	ym := utils.GetLastYearMonth()
	fmt.Println("ym:", ym)
}

// go test -v ./pkg/utils -run TestGetFirstDateOfWeek
func TestGetFirstDateOfWeek(t *testing.T) {
	first, end := utils.GetCurrentWeekScope()
	fmt.Println("first:", first, end)
}

func TestGetLastWeekScope(t *testing.T) {
	first, end := utils.GetLastWeekScope()
	fmt.Println("first:", first, end)
}
