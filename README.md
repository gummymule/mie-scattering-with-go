## Flowchart Detail Alur Perhitungan Mie

```mermaid
flowchart TD
    %% Input
    A[Input: size parameter x, refractive index m, wavenumber k] --> B[Hitung Nmax NmaxForX(x)]

    %% Parameter kompleks
    B --> C[Hitung z1 = x, z2 = m * x]

    %% Fungsi Bessel
    C --> D[Hitung Spherical Bessel j_n(z1), j_n(z2), slices j1[n], j2[n]]
    D --> E[Hitung Spherical Hankel h_n(1)(z1), slice h1[n]]

    %% Ricatti Psi
    E --> F[Hitung Ricatti Psi: psi1[n] = z1*j1[n], psi2[n] = z2*j2[n], slices psi1[n], psi2[n]]
    
    %% Turunan Ricatti Psi
    F --> G[Hitung turunan Ricatti Psi: psi1p[n] = z1*j1[n-1]-n*j1[n], psi2p[n] = z2*j2[n-1]-n*j2[n], xi1p[n] = z1*h1[n-1]-n*h1[n], slices psi1p[n], psi2p[n], xi1p[n]]

    %% Koefisien Mie
    G --> H[Hitung a[n], b[n] untuk n=1..N, simpan slices a[n], b[n]]

    %% Multipole decomposition
    H --> I[Hitung Total, ED, MD, EQ, MQ dari a[n], b[n]]

    %% Output
    I --> J[Output: slices a[n], b[n], ED, MD, EQ, MQ, Total Scattering Cross-section]
```