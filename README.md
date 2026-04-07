# CitExtractor

## Usage Guide

### 1- Getting the Binaries

Download the binaries from the [releases](https://github.com/ufscar-lec/CitExtractor/releases/tag/binaries) or build the program yourself using the `build.sh` script (for building, you are required to have the Golang toolchain installed).


### 2- Running the Program

To run the program, open the terminal inside the directory you have the binaries in and run one of the following commands, depending on your operating system:

For Linux:

    ./citextractor-linux-amd64 [DOI1] [DOI2] ... [DOIX]

For Windows:

    .\citextractor-windows-x86-64.exe [DOI1] [DOI2] ... [DOIX]

Replace the `DOI` args with your DOI (e.g., `10.1234/1234`).

### 3- Checking the Results

The results are going to be saved in a `.txt` file with the name of the DOI you provided.
