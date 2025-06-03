package goshellconv

import (
	"errors"
	"fmt"
	"io/ioutil"
)

func ConvertToShellcode(inputFile string, passphrase string, optimize bool, method string) ([]byte, error) {
	fileData, err := ReadFile(inputFile)
	if err != nil {
		return nil, err
	}

	var processed []byte

	switch method {
	case "aes":
		if passphrase == "" {
			return nil, errors.New("AES encryption requires passphrase")
		}
		processed, err = EncryptAES(fileData, passphrase)
	case "chacha20":
		if passphrase == "" {
			return nil, errors.New("ChaCha20 encryption requires passphrase")
		}
		processed, err = EncryptChaCha20(fileData, passphrase)
	case "xor":
		if passphrase == "" {
			return nil, errors.New("XOR encryption requires passphrase")
		}
		processed, err = EncryptXOR(fileData, passphrase)
	case "b64":
		processed, err = EncryptBase64(fileData, "")
	default:
		return nil, fmt.Errorf("unsupported encryption method: %s", method)
	}
	if err != nil {
		return nil, err
	}
	if method == "b64" {
		return processed, nil
	}

	shellcode := GenerateShellcode(processed)

	if optimize {
		shellcode = OptimizeShellcode(shellcode)
	}

	if passphrase != "" {
		err := ioutil.WriteFile("passcode.txt", []byte("Do not forget your password: "+passphrase), 0644)
		if err != nil {
			return nil, fmt.Errorf("error writing passcode file: %w", err)
		}
	}

	return shellcode, nil
}
