package tokenmanager

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base32"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	clip "github.com/atotto/clipboard"
	"github.com/xlzd/gotp"
)

// Token is the atomic struct of the program. It has the following elements:
// name, number of digits, seed length, hash function algorithm and counter
// only meant for HOTP codes.
type Token struct {
	Name     string
	Digits   int
	Seed     string
	HashFunc string
	Counter  int
}

// CreateToken allows to create a new token using parameters given by the user.
// This function will write the token into the default config file, and show the
// generated seed encoded in base32 and hexadecimal, to feed into OTP servers.
func CreateToken(name string, digits int, seedLen int, hash string) {

	if name == "" {
		fmt.Println("ERROR: Add name to create token command.")
		os.Exit(1)
	}
	seed := gotp.RandomSecret(seedLen)
	t := Token{Name: name, Digits: digits, Seed: seed, HashFunc: hash, Counter: 0}
	saveTokenToConfig(t)

	src := byteSecret(t.Seed)
	dst := make([]byte, hex.EncodedLen(len(src)))
	hex.Encode(dst, src)
	fmt.Println("Token created successfully:")
	fmt.Println("Name: ", t.Name)
	fmt.Println("B32 seed: ", t.Seed)
	fmt.Printf("Hex seed: %s\n", dst)

}

// DeleteToken will remove the token from the default config file.
func DeleteToken(name string) {

	tokens := loadDefaultConfig()
	for i, token := range tokens {
		if token.Name == name {
			tokens = removeIndex(tokens, i)
		}
	}

	saveDefaultConfig(tokens)

}

// ListTokens will list the available tokens from the default config file.
func ListTokens() {
	tokens := loadDefaultConfig()
	fmt.Println("--------------------------")
	fmt.Println("      List of Tokens      ")
	fmt.Println("--------------------------")
	for _, token := range tokens {
		fmt.Println("[*] Name: ", token.Name)
	}

}

// GetTotp will generate a Time-based OTP from the selected token.
func GetTotp(name string) {
	t, err := loadTokenFromConfig(name)
	if err != nil {
		fmt.Println("Error loading token: ", err)
		os.Exit(1)
	}
	hasher := getHasher(t.HashFunc)
	totp := gotp.NewTOTP(t.Seed, t.Digits, 30, hasher)
	clip.WriteAll(totp.Now())
}

// TODO: Enable logic for HOTP
// GetHotp will generate a event-based OTP from the selected token.
func GetHotp(name string) {
	t, err := loadTokenFromConfig(name)
	if err != nil {
		fmt.Println("Error loading token: ", err)
		os.Exit(1)
	}
	hasher := getHasher(t.HashFunc)
	hotp := gotp.NewHOTP(t.Seed, t.Digits, hasher)
	clip.WriteAll(hotp.At(t.Counter))
	t.Counter = t.Counter + 1
	saveTokenToConfig(t)
}

// This function saves the token to the def
func saveTokenToConfig(t Token) {

	tokens := loadDefaultConfig()

	for _, token := range tokens {
		if t.Name == token.Name {
			fmt.Println("ERROR: The token name already exists")
			os.Exit(1)
		}
	}

	tokens = append(tokens, t)

	saveDefaultConfig(tokens)

}

func loadTokenFromConfig(name string) (t Token, e error) {

	tokens := loadDefaultConfig()

	for _, token := range tokens {
		if token.Name == name {
			t = token
			return t, nil
		}
	}
	e = os.ErrNotExist
	return t, e
}

func loadDefaultConfig() []Token {
	conf, err := ioutil.ReadFile(defaultConfig)
	if err != nil {
		fmt.Println("Loading tokens threw an error: ", err)
	}

	var tokens []Token

	err = json.Unmarshal(conf, &tokens)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	return tokens
}

func saveDefaultConfig(tokens []Token) {
	tConf, err := json.MarshalIndent(tokens, "", " ")
	if err != nil {
		fmt.Println("Marshalling the token failed due to: ", err)
		os.Exit(1)
	}

	err = ioutil.WriteFile(defaultConfig, tConf, 0666)
	if err != nil {
		fmt.Println("Writing token to default config file failed: ", err)
		os.Exit(1)
	}
}

func getHasher(hash string) *gotp.Hasher {
	hasher := &gotp.Hasher{}
	switch hash {
	case "sha1":
		hasher.HashName = hash
		hasher.Digest = sha1.New
	case "sha256":
		hasher.HashName = hash
		hasher.Digest = sha256.New
	case "sha512":
		hasher.HashName = hash
		hasher.Digest = sha512.New
	}
	return hasher
}

func removeIndex(s []Token, index int) []Token {
	return append(s[:index], s[index+1:]...)
}

func byteSecret(seed string) []byte {
	missingPadding := len(seed) % 8
	if missingPadding != 0 {
		seed = seed + strings.Repeat("=", 8-missingPadding)
	}
	bytes, err := base32.StdEncoding.DecodeString(seed)
	if err != nil {
		panic("decode secret failed")
	}

	return bytes
}
