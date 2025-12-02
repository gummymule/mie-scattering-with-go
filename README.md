## Flowchart Detail Alur Perhitungan Mie

```mermaid
flowchart TD
    %% Input
    A[Input: size parameter x, refractive index m, wavenumber k] --> B[Hitung Nmax: N = NmaxForX(x)]

    %% Parameter kompleks
    B --> C[Hitung z1 = x, z2 = m * x]

    %% Fungsi Bessel
    C --> D[Hitung Spherical Bessel j_n(z1), j_n(z2) → slices j1[n], j2[n]]
    D --> E[Hitung Spherical Hankel h_n^{(1)}(z1) → slice h1[n]]

    %% Ricatti Psi
    E --> F[Hitung Ricatti Psi: 
        psi1[n] = z1 * j1[n], 
        psi2[n] = z2 * j2[n] → slices psi1[n], psi2[n]]
    
    %% Turunan Ricatti Psi
    F --> G[Hitung turunan Ricatti Psi:
        psi1'[n] = z1*j1[n-1] - n*j1[n],
        psi2'[n] = z2*j2[n-1] - n*j2[n],
        xi1'[n] = z1*h1[n-1] - n*h1[n] → slices psi1p[n], psi2p[n], xi1p[n]]

    %% Koefisien Mie
    G --> H[Hitung koefisien Mie a[n], b[n] untuk n=1..N:
        a[n] = (m*psi2[n]*psi1p[n] - psi1[n]*psi2p[n]) / (m*psi2[n]*xi1p[n] - xi1[n]*psi2p[n])
        b[n] = (psi2[n]*psi1p[n] - m*psi1[n]*psi2p[n]) / (psi2[n]*xi1p[n] - m*xi1[n]*psi2p[n]) → slices a[n], b[n]]

    %% Multipole decomposition
    H --> I[Hitung Multipole Decomposition:
        Total = sum_{n=1}^{N} (2π/k²) * (2n+1) * (|a[n]|² + |b[n]|²),
        ED = 3*(2π/k²)*|a[1]|²,
        MD = 3*(2π/k²)*|b[1]|²,
        EQ = 5*(2π/k²)*|a[2]|²,
        MQ = 5*(2π/k²)*|b[2]|²]

    %% Output
    I --> J[Output: slices a[n], b[n], ED, MD, EQ, MQ, Total Scattering Cross-section]
