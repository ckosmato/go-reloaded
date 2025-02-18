
# Text Processing Tool in Go

## Overview  
This Go program processes a text file by applying transformations based on specific markers enclosed in parentheses. It reads an input file, modifies its contents based on these markers, and writes the modified text to an output file.

## Features
- **Hexadecimal to Decimal Conversion**: Converts the preceding hexadecimal number when followed by **`(hex)`**.
- **Binary to Decimal Conversion**: Converts the preceding binary number when followed by **`(bin)`**.
- **Case Transformations**:
  - **`(up)`**: Converts the preceding word to uppercase.
  - **`(low)`**: Converts the preceding word to lowercase.
  - **`(cap)`**: Capitalizes the first letter of the preceding word.
- **Punctuation Handling**: Ensures correct spacing and placement of punctuation marks.
- **Article Adjustment**: Changes **"a"** to **"an"** before words that start with a vowel.
- **Contraction Handling**: Properly merges apostrophes into contractions (e.g., **"do n't"** → **"don't"**).

## Installation  
Follow these steps to install and run the program:

1. Ensure you have **Go** installed.
2. Clone the repository or copy the source file.
3. Build the program:
   ```sh
   go build -o textprocessor main.go
   ```

## Usage  
The program requires two command-line arguments:
1. The path to the **input** file.
2. The path to the **output** file.

Example:
```sh
./textprocessor input.txt output.txt
```
This will read **`input.txt`**, process transformations, and save the modified text to **`output.txt`**.

## Code Explanation

### Main Function (`main`)
- Manages file input/output.
- Calls the `Replace` function to process transformations.

### Text Processing Function (`Replace`)
- Scans the text **line by line**.
- Identifies **bracketed markers**:
  - **`(hex)`** → Converts preceding hexadecimal numbers to decimal.
  - **`(bin)`** → Converts preceding binary numbers to decimal.
  - **`(up)`** → Converts the preceding word to uppercase.
  - **`(low)`** → Converts the preceding word to lowercase.
  - **`(cap)`** → Capitalizes the first letter of the preceding word.
- Adjusts **punctuation** and **articles** accordingly.

## Error Handling
- Detects missing command-line arguments and displays an error.
- Ensures correct file handling (opening, reading, and writing).
- Validates number conversions for **`(hex)`** and **`(bin)`** transformations.

## Example

### Input (`input.txt`):
```txt
hello world (up). this is a test. 1F (hex) 101 (bin) apple (cap).
```

### Output (`output.txt`):
```txt
HELLO world. This is a test. 31 5 Apple.
```

## License  
This project is open-source under the **MIT License**.
