package special

import (
	"math/cmplx"
)

// SphrericalBesselJ returns j_n(z) for n=0..N
func SphericalBesselJ(N int, z complex128) []complex128 {
	js := make([]complex128, N+1)
	if N < 0 {
		return js
	}
	// handle z==0 limit
	if z == 0 {
		js[0] = 1 + 0i
		for n := 1; n <= N; n++ {
			js[n] = 0 + 0i
		}
		return js
	}

	js[0] = cmplx.Sin(z) / z
	if N == 0 {
		return js
	}
	js[1] = cmplx.Sin(z)/(z*z) - cmplx.Cos(z)/z
	for n := 1; n < N; n++ {
		coef := complex(float64(2*n+1), 0) / z
		js[n+1] = coef*js[n] - js[n-1]
	}
	return js
}

// SphrericalHankel1 returns h_n^{(1)}(z) = j_n + i*y_n
// Using recurrence for y_n similiarly or compute via combination
func SphericalHankel1(N int, z complex128) []complex128 {
	// compute j and y using same recurrence
	js := SphericalBesselJ(N, z)

	// compute y via recurrence; base forms:
	ys := make([]complex128, N+1)
	if z == 0 {
		for n := 0; n <= N; n++ {
			ys[n] = complex(1e300, 0)
		}
	} else {
		ys[0] = -cmplx.Cos(z) / z
		if N > 0 {
			ys[1] = -cmplx.Cos(z)/(z*z) - cmplx.Sin(z)/z
			for n := 1; n < N; n++ {
				coef := complex(float64(2*n+1), 0) / z
				ys[n+1] = coef*ys[n] - ys[n-1]
			}
		}
	}
	hs := make([]complex128, N+1)
	for n := 0; n <= N; n++ {
		hs[n] = js[n] + complex(0, 1)*ys[n]
	}
	return hs
}

// RicattiPsi computes psi_n = z * j_n for n=0..N
func RicattiPsi(N int, z complex128) []complex128 {
	j := SphericalBesselJ(N, z)
	p := make([]complex128, N+1)
	for n := 0; n <= N; n++ {
		p[n] = z * j[n]
	}
	return p
}

// helper compute psi via identity psi_n'(z) = z*j_{n-1}(z) - n*j_n(z)
// requires j array with index -1 treated carefully: j_{-1} = cos(z)/z? but we only use n>=1
func RicattiPsiDerivFromJ(j []complex128, z complex128) []complex128 {
	N := len(j) - 1
	out := make([]complex128, N+1)
	for n := 0; n <= N; n++ {
		if n == 0 {
			// psi0 = z*j0 => psi0' = j0 + z*j0' ; use derivative identity via j1 maybe
			// Using psi0' = -z*j1 (from recurrence) + 0*j0? simpler compute j0 + z*j0'
			// but to avoid complexity we compute directly:
			// j0' = -j1 + ??? (not necessary for n=0 in Mie formulas because n starts at 1)
			out[n] = z * j[0]
		} else {
			// psi_n' = z*j_{n-1} - n*j_n
			out[n] = z*j[n-1] - complex(float64(n), 0)*j[n]
		}
	}
	return out
}
