package services

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func PushService(message, pushoverUser, pushoverToken string) error {
	if pushoverUser == "" || pushoverToken == "" {
		log.Println("Pushover credentials are missing!")
		return errors.New("credentails missing")
	}

	fmt.Printf("Push: %s\n", message)

	form := url.Values{}
	form.Add("user", pushoverUser)
	form.Add("token", pushoverToken)
	form.Add("message", message)

	resp, err := http.PostForm("https://api.pushover.net/1/messages.json", form)
	if err != nil {
		log.Println("Error sending push:", err)
		return err
	}
	defer resp.Body.Close()

	fmt.Println("Pushover status:", resp.Status)
	return nil
}