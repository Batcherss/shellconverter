package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"github.com/Batcherss/shellconverter/goshellconv"
)

func printUsage() {
	fmt.Println("Usage: shellconverter -i <input_file> -o <output_file> -crypt <passphrase> -method <aes|chacha20|xor|b64> -opt")
	fmt.Println()
	fmt.Println("Arguments:")
	fmt.Println("-i <input_file>    Input file (e.g., .exe or .dll)")
	fmt.Println("-o <output_file>   Output file (e.g., .bin)")
	fmt.Println("-crypt <passphrase> Encryption passphrase (required for aes, chacha20, xor)")
	fmt.Println("-method <method>   Encryption method: aes, chacha20, xor, b64 (default aes)")
	fmt.Println("-opt               Optional flag to optimize shellcode (removes commas)")
	fmt.Println()
	fmt.Println("Example:")
	fmt.Println(`shellconverter -i input.exe -o output.bin -crypt="mypassword" -method=aes -opt`)
}

func brainrootamobea() {
	banner := []string{
		"  _________.__           .__  .__                                ",
		" /   _____/|  |__   ____ |  | |  |       ____  ____   _______  __",
		" \\_____  \\ |  |  \\_/ __ \\|  | |  |     _/ ___\\/  _ \\ /    \\  \\/ /",
		" /        \\|   Y  \\  ___/|  |_|  |__   \\  \\__(  <_> )   |  \\   / ",
		"/_______  /|___|  /\\___  >____/____/ /\\ \\___  >____/|___|  /\\_/  ",
		"         \\/      \\/     \\/            \\/     \\/           \\/      ",
		"   -brainroot amobea f#cked our brain.",
	}

	for _, line := range banner {
		fmt.Println(line)
	}
}

func main() {
	brainrootamobea()
	inputFile := flag.String("i", "", "Input file (.exe or .dll)")
	outputFile := flag.String("o", "", "Output file (.bin)")
	passphrase := flag.String("crypt", "", "Encryption passphrase")
	method := flag.String("method", "aes", "Encryption method: aes, chacha20, xor, b64")
	optimize := flag.Bool("opt", false, "Optimize the shellcode")
	flag.Parse()
	if *inputFile == "" || *outputFile == "" {
		printUsage()
		os.Exit(1)
	}
	if (*method != "b64") && (*passphrase == "") {
		fmt.Println("[!] passphrase is required for method:", *method)
		os.Exit(1)
	}
	fmt.Println("[+] received file conversion in progress..")
	shellcode, err := goshellconv.ConvertToShellcode(*inputFile, *passphrase, *optimize, *method)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("[#] converted to shellcode..")
	err = os.WriteFile(*outputFile, shellcode, 0644)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("[#] writed file - success.")
	fmt.Println("[+] success convertation to shellcode x64.")
}
