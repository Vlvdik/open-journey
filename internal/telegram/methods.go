package bot

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"strings"
	"time"
)

type ResponseData struct {
	Data struct {
		Translations []struct {
			TranslatedText string `json:"translatedText"`
		} `json:"translations"`
	} `json:"data"`
}

func getPrompt(text string) string {
	return strings.Join(strings.Split(text, " ")[1:], "")
}

func translate(key string, text string) (string, error) {
	if IsPrompt(text) {
		text = getPrompt(text)

		url := "https://google-translate1.p.rapidapi.com/language/translate/v2"

		payload := strings.NewReader("source=ru&target=en&q=" + text)

		req, _ := http.NewRequest("POST", url, payload)

		req.Header.Add("content-type", "application/x-www-form-urlencoded")
		req.Header.Add("Accept-Encoding", "application/gzip")
		req.Header.Add("X-RapidAPI-Key", key)
		req.Header.Add("X-RapidAPI-Host", "google-translate1.p.rapidapi.com")

		res, _ := http.DefaultClient.Do(req)

		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)

		var convertedResponse ResponseData
		err := json.Unmarshal(body, &convertedResponse)
		if err != nil {
			return "", errors.New(ErrTranslationFailed)
		}

		return convertedResponse.Data.Translations[0].TranslatedText, nil
	} else {
		return "", errors.New(ErrInvalidPrompt)
	}
}

func useImagine(imgChn chan string, prompt string) {
	script := "from app.model import imagine; print(imagine('" + prompt + "'))"
	cmd := exec.Command("python", "-c", script)

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	imgChn <- string(out)
}

func IsPrompt(text string) bool {
	if len(text) > 1 {
		return true
	} else {
		return false
	}
}

func GetPromptURL(prompt string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 150*time.Second)
	defer cancel()

	imgChn := make(chan string)

	go useImagine(imgChn, prompt)

	select {
	case URL := <-imgChn:
		defer close(imgChn)

		return URL, nil
	case <-ctx.Done():
		return "", errors.New(ErrImagineTimeOut)
	}
}
