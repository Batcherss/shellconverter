#  Disclaimer
This tool is for educational and research purposes only.
The author is not responsible for any misuse or illegal activity.

#  goexectoshell
[![Go](https://img.shields.io/badge/Go-1.22-blue.svg)](https://golang.org)  
[![Platform](https://img.shields.io/badge/Platform-Windows-lightgrey)]()  
[![License](https://img.shields.io/badge/License-MIT-green.svg)]()  
[![Shellcode](https://img.shields.io/badge/Output-Shellcode%20(.bin)-orange)]()
Required library "golang.org/x/crypto/chacha20"!\n
A simple and fast `.exe` / `.dll` to raw `.bin` shellcode converter with optional AES encryption and shellcode optimization.

---

##  Features

-  Converts PE files (`.exe` or `.dll`) into shellcode (`.bin`)
-  Optional AES encryption (CTR mode) using a passphrase
-  Optional shellcode optimization (removes commas for cleaner output)
-  Minimal and fast, written in pure Go

---

Usage
```
shellconverter -i <input_file> -o <output_file> [-crypt <passphrase>] [-opt]
```
- What it does?:
  - `-i	Input file (.exe or .dll)`
  - `-o	Output file (.bin shellcode)`
  - `-crypt	(Optional) AES encryption passphrase`
  - `-opt	(Optional) Optimizes shellcode (removes commas)`

 Example
```
shellconverter -i input.exe -o output.bin -crypt="mysecurepassword" -opt
```
The resulting output.bin contains raw shellcode, ready for use in loaders or injection tools.
