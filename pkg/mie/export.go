package mie

import (
	"encoding/csv"
	"os"
	"strconv"
)

func SaveCSV(path string, rows [][]string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	w := csv.NewWriter(f)
	defer w.Flush()
	return w.WriteAll(rows)
}

// helper to format float64 -> string with chosen precision
func ff64(v float64) string { return strconv.FormatFloat(v, 'g', 8, 64) }
