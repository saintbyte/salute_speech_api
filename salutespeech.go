package salutespeech_api

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/nu7hatch/gouuid"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type SaluteSpeechApi struct {
	ApiHost           string
	RepetitionPenalty int
	TopP              float32
	Model             string
	MaxTokens         int
	Temperature       float32
	AuthData          string
}

func NewSaluteSpeechApi() *SaluteSpeechApi {
	// Cоздает новый обьект SaluteSpeechApi
	return &SaluteSpeechApi{}
}
func (s *SaluteSpeechApi) getExpiresFilename() string {
	return ".salute_speech_expires"
}
func (s *SaluteSpeechApi) getTokenFilename() string {
	return ".salute_speech_token"
}
func (s *SaluteSpeechApi) getExpiresAtFromFile() int64 {
	data, err := os.ReadFile(s.getExpiresFilename())
	if err != nil {
		return 0
	}
	i, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return 0
	}
	return i
}
func (s *SaluteSpeechApi) getTokenFromFile() string {
	data, err := os.ReadFile(s.getTokenFilename())
	if err != nil {
		return ""
	}
	return string(data)
}
func (s *SaluteSpeechApi) setExpiresAtToFile(value int64) {
	fh, _ := os.OpenFile(s.getExpiresFilename(), os.O_TRUNC|os.O_WRONLY|os.O_CREATE, 0666)
	fh.WriteString(strconv.FormatInt(value, 10)) // writing...
	defer fh.Close()
}

func (s *SaluteSpeechApi) setTokenToFile(value string) {
	fh, _ := os.OpenFile(s.getTokenFilename(), os.O_TRUNC|os.O_WRONLY|os.O_CREATE, 0666)
	fh.WriteString(value) // writing...
	defer fh.Close()
}
func (s *SaluteSpeechApi) getCurrentToken() string {
	expAt := s.getExpiresAtFromFile()
	token := s.getTokenFromFile()
	apochNow := time.Now().Unix()
	timeDelta := apochNow - (expAt / 1000)
	if timeDelta > 0 {
		newExpAt, token2 := s.Auth()
		s.setExpiresAtToFile(newExpAt)
		s.setTokenToFile(token2)
		token = token2
	}
	return token
}
func (s *SaluteSpeechApi) Auth() (int64, string) {
	//Получить токен для авторизации.
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	u, err := uuid.NewV4()
	request, _ := http.NewRequest("POST", SaluteSpeechOauthUrl, bytes.NewBufferString("scope=SALUTE_SPEECH_PERS"))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("RqUID", u.String())
	request.Header.Set("Authorization", "Basic "+os.Getenv("SALUTE_SPEECH_CLIENT_SECRET"))
	client := &http.Client{}
	log.Println(request)
	response, e := client.Do(request)

	if e != nil {
		log.Fatal(e)
	}
	//if response.StatusCode != http.StatusOK {
	//	return "Так что-то пошло не так на удаленной стороне. Повтори вопрос.", nil
	//}
	fmt.Println(response.StatusCode)
	if response.StatusCode != http.StatusOK {
		return 0, ""
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))
	defer response.Body.Close()

	var result TokenResponse
	err2 := json.Unmarshal(body, &result)
	if err2 != nil {
		log.Fatal(err2)
	}
	os.Setenv("SALUTE_SPEECH_TOKEN", result.AccessToken)
	return result.ExpiresAt, result.AccessToken
}
func (s *SaluteSpeechApi) Recognize(filename string) (string, error) {
	url := SaluteSpeechApiRestURL + "speech:recognize"
	file, _ := os.Open(filename)
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	request, _ := http.NewRequest("POST", url, file)
	request.Header.Set("Content-Type", "audio/ogg;codecs=opus")
	request.Header.Set("Authorization", "Bearer "+s.getCurrentToken())
	client := &http.Client{}
	log.Println(request)
	response, e := client.Do(request)

	if e != nil {
		return "", e
	}
	//if response.StatusCode != http.StatusOK {
	//	return "Так что-то пошло не так на удаленной стороне. Повтори вопрос.", nil
	//}
	fmt.Println(response.StatusCode)
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))
	var result SpeechRecognizeAnswer
	err2 := json.Unmarshal(body, &result)
	if err2 != nil {
		log.Fatal(err2)
	}
	defer response.Body.Close()
	return result.Result[0], nil

}
func (s *SaluteSpeechApi) Synthesize(text2speech string) (io.Reader, error) {
	url := SaluteSpeechApiRestURL + "text:synthesize?format=opus&voice=Ost_24000"
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	request, _ := http.NewRequest("POST", url, strings.NewReader(text2speech))
	request.Header.Set("Content-Type", "application/text")
	request.Header.Set("Authorization", "Bearer "+s.getCurrentToken())
	client := &http.Client{}
	log.Println(request)
	response, e := client.Do(request)
	log.Println(response.StatusCode)
	if response.StatusCode != http.StatusOK {
		log.Println(e)
		return nil, errors.New("Response is not ok")
	}
	return response.Body, nil
}
