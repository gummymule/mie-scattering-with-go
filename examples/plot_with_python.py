import json
import sys
import numpy as np
import pandas as pd
import matplotlib.pyplot as plt

def load_data(path):
    """Load JSON or CSV."""
    if path.endswith(".json"):
        with open(path, "r") as f:
            data = json.load(f)
        return data

    if path.endswith(".csv"):
        df = pd.read_csv(path)
        return {
            "wavelength_nm": df["wavelength_nm"].to_list(),
            "total": df["total"].to_list(),
            "ed": df["ed"].to_list(),
            "md": df["md"].to_list(),
            "eq": df["eq"].to_list(),
            "mq": df["mq"].to_list(),
        }

    raise ValueError("Input file must be JSON or CSV.")


def main():
    if len(sys.argv) < 2:
        print("Usage: python plot_with_python.py mie_output.json")
        sys.exit(1)

    path = sys.argv[1]
    data = load_data(path)

    # Convert lists → numpy arrays
    wl = np.array(data["wavelength_nm"])
    total = np.array(data["total"])
    ed = np.array(data["ed"])
    md = np.array(data["md"])
    eq = np.array(data["eq"])
    mq = np.array(data["mq"])

    # Save processed CSV
    df = pd.DataFrame({
        "wavelength_nm": wl,
        "total": total,
        "electric_dipole": ed,
        "magnetic_dipole": md,
        "electric_quadrupole": eq,
        "magnetic_quadrupole": mq,
    })
    df.to_csv("multipole_data_processed.csv", index=False)
    print("Saved multipole_data_processed.csv")

    # ---- Plot ----
    plt.figure(figsize=(10, 6))

    plt.plot(wl, total, "k-", linewidth=2, label="Total scattering")
    plt.plot(wl, ed, "r--", linewidth=2, label="Electric dipole")
    plt.plot(wl, md, "y:", linewidth=2, label="Magnetic dipole")
    plt.plot(wl, mq, "m-.", linewidth=2, label="Magnetic quadrupole")
    plt.plot(wl, eq, "g-", linewidth=1.5, label="Electric quadrupole")

    plt.xlabel("Wavelength (nm)")
    plt.ylabel("Scattering cross-section (a.u.)")
    plt.title("Mie Multipole Decomposition")
    plt.grid(True, alpha=0.3)
    plt.legend()
    plt.tight_layout()

    plt.savefig("multipole_plot.png", dpi=300)
    plt.savefig("multipole_plot.pdf")
    print("Saved multipole_plot.png and multipole_plot.pdf")

    plt.show()


if __name__ == "__main__":
    main()