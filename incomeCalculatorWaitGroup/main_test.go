package main

import (
	"os"
	"strings"

	"io"
	"testing"
)

func Test_Calculate_Year_income(t *testing.T) {
	stdout := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	w.Close()

	result, _ := io.ReadAll(r)
	outuput := string(result)

	os.Stdout = stdout

	if !strings.Contains(outuput, "32240.00") {
		t.Errorf("Expected 32240.00, but got %s", outuput)
	}

}
