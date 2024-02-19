package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Decrypt(encryptedText string, n int64) string {
	var decryptedText string

	for _, char := range encryptedText {
		if char >= 'a' && char <= 'z' {
			decryptedText += string((char-'a'+rune(26-n))%26 + 'a')
		} else if char >= 'A' && char <= 'Z' {
			decryptedText += string((char-'A'+rune(26-n))%26 + 'A')
		} else {
			decryptedText += string(char)
		}
	}

	return decryptedText
}

func Encrypt(plaintText string, n int64) string {
	var encryptedText string

	for _, char := range plaintText {
		if char >= 'a' && char <= 'z' {
			encryptedText += string((char-'a'+rune(n))%26 + 'a')
		} else if char >= 'A' && char <= 'Z' {
			encryptedText += string((char-'A'+rune(n))%26 + 'A')
		} else {
			encryptedText += string(char)
		}
	}

	return encryptedText
}

func main() {

	var choice string
	var text string
	var n int64

	fmt.Println("Substitution Cipher: ")
	fmt.Println("[1] Decrypt Text")
	fmt.Println("[2] Encrypt Text")
	fmt.Print("Masukkan pilihan: ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	choice = scanner.Text()

	fmt.Print("Masukkan teks: ")
	scanner2 := bufio.NewScanner(os.Stdin)
	scanner2.Scan()
	err = scanner2.Err()

	if err != nil {
		log.Fatal(err)
	}

	text = scanner2.Text()

	fmt.Print("Masukkan nilai shift (n): ")

	scanner3 := bufio.NewScanner(os.Stdin)
	scanner3.Scan()
	err = scanner3.Err()

	if err != nil {
		log.Fatal(err)
	}

	n, err = strconv.ParseInt(scanner3.Text(), 10, 10)

	if err != nil {
		log.Fatal(err)
	}

	switch choice {
	case "1":
		decryptedText := Decrypt(text, n)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Text Decrypted:", string(decryptedText))
	case "2":
		encryptedText := Encrypt(text, n)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Text Encrypted:", string(encryptedText))
	default:
		fmt.Println("Masukkan pilihan yang benar.")
	}

}
