package mie

import (
	"math"
	"math/cmplx"
	"mie/internal/special"
)

// Nmax rule (Bohren-Hufman): Nmax = int(x + 4*x^(1/3) + 2)
func NmaxForX(x float64) int {
	return int(math.Ceil(x + 4*math.Pow(x, 1.0/3.0) + 2))
}

// ComputeMieCoefficients returns arrays a[0..N], b[0..N] (a[0], b[0] usused)
// given size parameter x and complex refractive index m = m_real + i*m_imag
func ComputeMieCoefficients(x float64, m complex128) (a, b []complex128) {
	if x <= 0 {
		return nil, nil
	}
	N := NmaxForX(x)
	// arguments
	z1 := complex(x, 0)
	z2 := m * z1

	// compute j_n for z1 and z2
	j1 := special.SphericalBesselJ(N+1, z1)
	j2 := special.SphericalBesselJ(N+1, z2)

	// hankel for z1
	h1 := special.SphericalHankel1(N+1, z1)

	// ricatti psi xi
	psi1 := make([]complex128, N+1)
	psi2 := make([]complex128, N+1)
	xi1 := make([]complex128, N+1)
	for n := 0; n <= N; n++ {
		psi1[n] = z1 * j1[n]
		psi2[n] = z2 * j2[n]
		xi1[n] = z1 * h1[n]
	}

	// derivatives using identity: psi_n' = z*j_{n-1} - n*j_n
	psi1p := special.RicattiPsiDerivFromJ(j1, z1)
	psi2p := special.RicattiPsiDerivFromJ(j2, z2)
	xi1p := make([]complex128, N+1)
	for n := 0; n <= N; n++ {
		// xi' = z*h_{n-1} - n*h_n
		if n == 0 {
			// rarely used, put placholder
			xi1p[n] = xi1[0]
		} else {
			xi1p[n] = z1*h1[n-1] - complex(float64(n), 0)*h1[n]
		}
	}

	a = make([]complex128, N+1)
	b = make([]complex128, N+1)

	for n := 1; n <= N; n++ {
		numA := m*psi2[n]*psi1p[n] - psi1[n]*psi2p[n]
		denA := m*psi2[n]*xi1p[n] - xi1[n]*psi2p[n]
		a[n] = numA / denA

		numB := psi2[n]*psi1p[n] - m*psi1[n]*psi2p[n]
		denB := psi2[n]*xi1p[n] - m*xi1[n]*psi2p[n]
		b[n] = numB / denB
	}
	return a, b
}

// Multipole decomposition: returns ED, MD, EQ, MQ, Total Scattering cross-sections
// supply k (wavenumber in medium) or use k = 2*pi/lambda_in_medium (units consistent)
func MultipoleDecomposition(a, b []complex128, k float64) (ED, MD, EQ, MQ, Total float64) {
	if len(a) == 0 || len(b) == 0 {
		return
	}
	pref := 2 * math.Pi / (k * k)
	N := len(a) - 1
	for n := 1; n <= N; n++ {
		factor := float64(2*n + 1)
		an2 := cmplx.Abs(a[n])
		bn2 := cmplx.Abs(b[n])
		term := pref * factor * (an2*an2 + bn2*bn2)
		Total += term
		if n == 1 {
			ED = pref * 3 * (cmplx.Abs(a[1]) * cmplx.Abs(a[1]))
			MD = pref * 3 * (cmplx.Abs(b[1]) * cmplx.Abs(b[1]))
		}
		if n == 2 {
			EQ = pref * 5 * (cmplx.Abs(a[2]) * cmplx.Abs(a[2]))
			MQ = pref * 5 * (cmplx.Abs(b[2]) * cmplx.Abs(b[2]))
		}
	}
	return
}
