## Flowchart Alur Perhitungan Mie

```mermaid
flowchart TD
    A[Input: x, m, k] --> B[Hitung Nmax]
    B --> C[Hitung z1 dan z2]
    C --> D[Hitung j_n slices]
    D --> E[Hitung h_n slices]
    E --> F[Hitung psi slices]
    F --> G[Hitung psi deriv slices]
    G --> H[Hitung koefisien a dan koefisien b]
    H --> I[Hitung Total, ED, MD, EQ, MQ]
    I --> J[Output: koefisien a, koefisien b, ED, MD, EQ, MQ, Total]
