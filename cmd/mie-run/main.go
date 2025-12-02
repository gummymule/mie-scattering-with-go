package main

import (
	"fmt"
	"math"
	"os"
	"time"

	"mie/pkg/materials"
	"mie/pkg/mie"
)

// Progress bar helper
func printProgressBar(current, total int) {
	percent := float64(current) / float64(total)
	barWidth := 30

	filled := int(percent * float64(barWidth))
	empty := barWidth - filled

	fmt.Printf("\r[%s%s] %.2f%%",
		stringRepeat("█", filled),
		stringRepeat("░", empty),
		percent*100,
	)
}

func stringRepeat(s string, count int) string {
	out := ""
	for i := 0; i < count; i++ {
		out += s
	}
	return out
}

func main() {
	start := time.Now()

	// Load material file (Si, Au, Ag)
	mat, err := materials.LoadCSV("pkg/materials/Ag_Johnson_merged.csv")
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	radiusNm := 100.0
	radius := radiusNm * 1e-9
	m0 := 1.0

	wlStart := 400.0
	wlEnd := 1000.0
	step := 2.0

	totalSteps := int((wlEnd-wlStart)/step) + 1
	currentStep := 0

	out := [][]string{{"wavelength_nm", "ED_m2", "MD_m2", "EQ_m2", "MQ_m2", "Total_m2"}}

	for wl := wlStart; wl <= wlEnd; wl += step {
		// Update progress bar
		printProgressBar(currentStep, totalSteps)
		currentStep++

		nRe, nIm := mat.NK(wl)
		mRel := complex(nRe, -nIm) / complex(m0, 0)

		lambda := wl * 1e-9
		k := 2 * math.Pi * m0 / lambda
		x := k * radius

		a, b := mie.ComputeMieCoefficients(x, mRel)
		ED, MD, EQ, MQ, Total := mie.MultipoleDecomposition(a, b, k)

		out = append(out, []string{
			fmt.Sprintf("%.2f", wl),
			fmt.Sprintf("%.8g", ED),
			fmt.Sprintf("%.8g", MD),
			fmt.Sprintf("%.8g", EQ),
			fmt.Sprintf("%.8g", MQ),
			fmt.Sprintf("%.8g", Total),
		})
	}

	// finish bar
	printProgressBar(totalSteps, totalSteps)
	fmt.Println()

	if err := mie.SaveCSV("mie_spectrum_Ag.csv", out); err != nil {
		fmt.Fprintln(os.Stderr, "save error:", err)
	}

	elapsed := time.Since(start)
	fmt.Printf("Done! Execution time: %s\n", elapsed)
}
