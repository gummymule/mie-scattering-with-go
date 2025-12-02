## Description of the Mie Scattering Calculation

This program calculates the **Mie coefficients** and **scattering cross-section** for spherical particles with size parameter $x$, complex refractive index $m$, and wavenumber $k$. The workflow is as follows:

1. **Input**: the values of $x$ (size parameter), $m$ (complex refractive index), and $k$ (wavenumber in the medium).  
2. **Determining Nmax**: the maximum multipole order $N$ is computed using the Bohren-Huffman rule:
   $$N_\text{max} = \lceil x + 4 x^{1/3} + 2 \rceil$$
3. **Complex parameters**: calculate $z_1 = x$ and $z_2 = m \cdot x$.  
4. **Bessel and Hankel functions**: slices $j_n(z)$ and $h_n^{(1)}(z)$ are computed for $n = 0..N$ using `SphericalBesselJ` and `SphericalHankel1`.  
5. **Ricatti Psi**: computed as
   $$\psi_n(z) = z \cdot j_n(z), \quad
   \psi_n'(z) = z \cdot j_{n-1}(z) - n \cdot j_n(z)$$
   The derivative slices are stored in `psi1p`, `psi2p`, and `xi1p`.  
6. **Mie coefficients**: for $n = 1..N$, computed as
   $$a_n = \frac{m \psi_n^{(2)} \psi_n^{(1)'} - \psi_n^{(1)} \psi_n^{(2)'}}{m \psi_n^{(2)} \xi_n' - \xi_n \psi_n^{(2)'}}, \quad
   b_n = \frac{\psi_n^{(2)} \psi_n^{(1)'} - m \psi_n^{(1)} \psi_n^{(2)'}}{\psi_n^{(2)} \xi_n' - m \xi_n \psi_n^{(2)'}}$$
   The results are stored in slices `a[n]` and `b[n]`.  
7. **Multipole Decomposition**: from the coefficients $a_n$ and $b_n$, the scattering contributions are calculated:
   $$\text{Total} = \sum_{n=1}^{N} \frac{2\pi}{k^2} (2n+1) (|a_n|^2 + |b_n|^2)$$
   $$\text{ED} = \frac{6\pi}{k^2} |a_1|^2, \quad
   \text{MD} = \frac{6\pi}{k^2} |b_1|^2$$
   $$\text{EQ} = \frac{10\pi}{k^2} |a_2|^2, \quad
   \text{MQ} = \frac{10\pi}{k^2} |b_2|^2$$
8. **Output**: slices `a[n]`, `b[n]`, and the scattering values ED, MD, EQ, MQ, and Total.

> All main mathematical functions (Bessel, Hankel, Ricatti Psi) are placed in the `special` package, while the computation of Mie coefficients and cross-sections is in the main `mie` package.

---

## How to Run

To execute the program, simply run:

```bash
go run ./cmd/mie-run/main.go
