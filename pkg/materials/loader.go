package materials

import (
	"encoding/csv"
	"math"
	"os"
	"strconv"
)

type Material struct {
	WL []float64 // wavelengths in nm
	N  []float64 // real part
	K  []float64 // imaginary part
}

func LoadCSV(path string) (*Material, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	r := csv.NewReader(f)
	rows, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	mat := &Material{}
	for i, row := range rows {
		if i == 0 {
			continue // skip header
		}
		if len(row) < 3 {
			continue
		}
		wl, _ := strconv.ParseFloat(row[0], 64)
		n, _ := strconv.ParseFloat(row[1], 64)
		k, _ := strconv.ParseFloat(row[2], 64)
		mat.WL = append(mat.WL, wl)
		mat.N = append(mat.N, n)
		mat.K = append(mat.K, k)
	}
	return mat, nil
}

// Linear interpolation for refractive index n,k
func (m *Material) NK(wl float64) (float64, float64) {
	if wl <= m.WL[0] {
		return m.N[0], m.K[0]
	}
	if wl >= m.WL[len(m.WL)-1] {
		return m.N[len(m.N)-1], m.K[len(m.K)-1]
	}

	for i := 0; i < len(m.WL)-1; i++ {
		if wl >= m.WL[i] && wl <= m.WL[i+1] {
			t := (wl - m.WL[i]) / (m.WL[i+1] - m.WL[i])
			n := m.N[i] + t*(m.N[i+1]-m.N[i])
			k := m.K[i] + t*(m.K[i+1]-m.K[i])
			return n, k
		}
	}
	return math.NaN(), math.NaN()
}
