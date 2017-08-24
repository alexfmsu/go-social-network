package settings

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	key_c  string = ":LKJNHBGVwefnkwehbfbipefwqiopw;e"
	salt_c string = "koqefwmioqwefoiqew;foEFJQI"
)

func Decrypter(token string) int {
	key := []byte(key_c)

	ciphertext, _ := hex.DecodeString(token)

	block, err := aes.NewCipher(key)

	if err != nil {
		panic(err)
	}

	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	if len(ciphertext)%aes.BlockSize != 0 {
		panic("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)

	mode.CryptBlocks(ciphertext, ciphertext)

	id, err := strconv.Atoi(strings.Split(string(ciphertext[:]), ":")[0])

	return id
}

func Encrypter(id int) string {
	key := []byte(key_c)

	plaintext := []byte(strconv.Itoa(id) + ":" + salt_c)

	end := make([]byte, aes.BlockSize-(len(plaintext)%aes.BlockSize))

	plaintext = append(plaintext, end...)

	if len(plaintext)%aes.BlockSize != 0 {
		panic("plaintext is not a multiple of the block size")
	}

	block, err := aes.NewCipher(key)

	if err != nil {
		panic(err)
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))

	iv := ciphertext[:aes.BlockSize]

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	mode := cipher.NewCBCEncrypter(block, iv)

	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

	return hex.EncodeToString(ciphertext)
}

func GetCookieFromID(id int) *http.Cookie {
	expiration := time.Now().Add(365 * 24 * time.Hour)

	h := Encrypter(id)

	cookie := http.Cookie{Name: "token", Value: h, Expires: expiration}

	return &cookie
}

func CheckAuth(w http.ResponseWriter, r *http.Request) *User_obj {
	token, err := r.Cookie("token")

	if err == nil {

		id := Decrypter(token.Value)

		fmt.Print(id, "\n")

		user := Get_user_by_id(id)

		fmt.Print("User_id = ", user.Id, "\n")

		return user
	} else {
		http.Redirect(w, r, "/login", 301)
	}

	return nil
}
