package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

var ErrNoTextToEncrypt = errors.New("tidak ada teks yang dienkripsi")
var ErrKeyMissing = errors.New("key tidak ditemukan")

const placeholder = ' '

func getKey(keyWord string) []int {
	keyWord = strings.ToLower(keyWord)
	word := []rune(keyWord)
	var sortedWord = make([]rune, len(word))
	copy(sortedWord, word)
	sort.Slice(sortedWord, func(i, j int) bool { return sortedWord[i] < sortedWord[j] })
	usedLettersMap := make(map[rune]int)
	wordLength := len(word)
	resultKey := make([]int, wordLength)
	for i := 0; i < wordLength; i++ {
		char := word[i]
		numberOfUsage := usedLettersMap[char]
		resultKey[i] = getIndex(sortedWord, char) + numberOfUsage + 1
		numberOfUsage++
		usedLettersMap[char] = numberOfUsage
	}
	return resultKey
}

func getIndex(wordSet []rune, subString rune) int {
	n := len(wordSet)
	for i := 0; i < n; i++ {
		if wordSet[i] == subString {
			return i
		}
	}
	return 0
}

func Encrypt(text []rune, keyWord string) ([]rune, error) {
	key := getKey(keyWord)
	keyLength := len(key)
	textLength := len(text)
	if keyLength <= 0 {
		return nil, ErrKeyMissing
	}
	if textLength <= 0 {
		return nil, ErrNoTextToEncrypt
	}
	if text[len(text)-1] == placeholder {
		return nil, fmt.Errorf("%w: Gagal melakukan encrypt, %q, karena berakhiran dengan karakter %q", ErrNoTextToEncrypt, text, placeholder)
	}
	n := textLength % keyLength

	for i := 0; i < keyLength-n; i++ {
		text = append(text, placeholder)
	}
	textLength = len(text)
	var result []rune
	for i := 0; i < textLength; i += keyLength {
		transposition := make([]rune, keyLength)
		for j := 0; j < keyLength; j++ {
			transposition[key[j]-1] = text[i+j]
		}
		result = append(result, transposition...)
	}
	return result, nil
}

func Decrypt(text []rune, keyWord string) ([]rune, error) {
	key := getKey(keyWord)
	textLength := len(text)
	if textLength <= 0 {
		return nil, ErrNoTextToEncrypt
	}
	keyLength := len(key)
	if keyLength <= 0 {
		return nil, ErrKeyMissing
	}
	n := textLength % keyLength
	for i := 0; i < keyLength-n; i++ {
		text = append(text, placeholder)
	}
	var result []rune
	for i := 0; i < textLength; i += keyLength {
		transposition := make([]rune, keyLength)
		for j := 0; j < keyLength; j++ {
			transposition[j] = text[i+key[j]-1]
		}
		result = append(result, transposition...)
	}
	result = []rune(strings.TrimRight(string(result), string(placeholder)))
	return result, nil
}

func main() {
	var choice string
	var text string
	var keyWord string

	fmt.Println("Transposition Cipher: ")
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

	fmt.Print("Masukkan KEY: ")
	scanner3 := bufio.NewScanner(os.Stdin)
	scanner3.Scan()
	err = scanner3.Err()
	if err != nil {
		log.Fatal(err)
	}

	keyWord = scanner3.Text()

	switch choice {
	case "1":
		decryptedText, err := Decrypt([]rune(text), keyWord)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Text Decrypted:", string(decryptedText))
	case "2":
		encryptedText, err := Encrypt([]rune(text), keyWord)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Text Encrypted:", string(encryptedText))
	default:
		fmt.Println("Masukkan pilihan yang benar.")
	}
}
