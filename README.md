# CitExtractor

## Usage Guide

### 1- Getting the Binaries

Download the binaries from the [releases](https://github.com/ufscar-lec/CitExtractor/releases/tag/binaries) or build the program yourself using the `build.sh` script (for building, you are required to have the Golang toolchain installed).


### 2- Running the Program

To run the program, open the terminal inside the directory you have the binaries in and run one of the following commands, depending on your operating system:

For Linux:

    ./citextractor-linux-amd64 [DOI1] [DOI2] ... [DOIX] > results.txt

For Windows:

    .\citextractor-windows-x86-64.exe [DOI1] [DOI2] ... [DOIX] > results.txt

Replace the `DOI` args with your DOI (e.g., `10.1234/1234`).
