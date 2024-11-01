package salutespeech_api

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
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
	AudioType         string
	Voice             *Voice
	Rate              int
	ValidateVoiceName bool
	Debug             bool
}

func NewSaluteSpeechApi() *SaluteSpeechApi {
	// Cоздает новый обьект SaluteSpeechApi
	return &SaluteSpeechApi{
		AudioType: "",
		Voice:     nil,
		Rate:      0,
		Debug:     false,
	}
}

func (s *SaluteSpeechApi) needRate(InputAudioType string) bool {
	if InputAudioType == SaluteSpeechApi_InputAudioTypePCM_S16LE {
		return true
	}
	if InputAudioType == SaluteSpeechApi_InputAudioTypeALAW {
		return true
	}
	if InputAudioType == SaluteSpeechApi_InputAudioTypeMULAW {
		return true
	}
	return false
}

func (s *SaluteSpeechApi) getContentTypeByInputAudioType(InputAudioType string, rate int) string {
	if InputAudioType == SaluteSpeechApi_InputAudioTypePCM_S16LE {
		return "audio/x-pcm;bit=16;rate=" + string(rate)
	}
	if InputAudioType == SaluteSpeechApi_InputAudioTypeOPUS {
		return "audio/ogg;codecs=opus"
	}
	if InputAudioType == SaluteSpeechApi_InputAudioTypeMP3 {
		return "audio/mpeg"
	}
	if InputAudioType == SaluteSpeechApi_InputAudioTypeFLAC {
		return "audio/flac"
	}
	if InputAudioType == SaluteSpeechApi_InputAudioTypeALAW {
		return "audio/pcma;rate=" + string(rate)
	}
	if InputAudioType == SaluteSpeechApi_InputAudioTypeMULAW {
		return "audio/pcmu;rate=" + string(rate)
	}
	return ""
}

func (s *SaluteSpeechApi) getFormatByOutputAudioType(OutputAudioType string) string {
	if OutputAudioType == SaluteSpeechApi_OutputAudioTypePCM16 {
		return "pcm16"
	} else if OutputAudioType == SaluteSpeechApi_OutputAudioTypeWAV16 {
		return "wav16"
	} else if OutputAudioType == SaluteSpeechApi_OutputAudioTypeOPUS {
		return "opus"
	} else if OutputAudioType == SaluteSpeechApi_OutputAudioTypeALAW {
		return "alaw"
	}
	return ""
}

func (s *SaluteSpeechApi) getExpiresFilename() string {
	value, ok := os.LookupEnv(SaluteSpeechExpiresFileEnv)
	if ok {
		return value
	}
	return ".salute_speech_expires"
}
func (s *SaluteSpeechApi) getTokenFilename() string {
	value, ok := os.LookupEnv(SaluteSpeechTokenFileEnv)
	if ok {
		return value
	}
	return ".salute_speech_token"
}

func (s *SaluteSpeechApi) isSSML(src_string string) bool {
	// Определяем по текст ssml это
	if src_string[0:6] == "<speak" {
		return true
	}
	if src_string[0:5] == "<?xml" {
		return true
	}
	return false
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

func (s *SaluteSpeechApi) getUuid() string {
	u, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}
	return u.String()
}
func (s *SaluteSpeechApi) GetVoiceById(voice_id string) *Voice {
	// Получить голос по id. ID смотри фокуметации к параметру voice у synthesize
	voices := SaluteSpeechVoices()
	for _, oneVoice := range voices {
		if oneVoice.Id == voice_id {
			return &oneVoice
		}
	}
	return nil
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
	//Получить токен для авторизации. В основном происходит автоматически - но вдруг кто-то захочет делать это сам.
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	request, _ := http.NewRequest("POST", SaluteSpeechOauthUrl, bytes.NewBufferString("scope=SALUTE_SPEECH_PERS"))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("RqUID", s.getUuid())
	request.Header.Set("Authorization", "Basic "+os.Getenv("SALUTE_SPEECH_AUTH_DATA"))
	client := &http.Client{}
	if s.Debug {
		log.Println(request)
	}
	response, e := client.Do(request)
	if e != nil {
		if s.Debug {
			log.Fatal(e)
		}
		return 0, ""
	}
	if response.StatusCode != http.StatusOK {
		if s.Debug {
			log.Println(response.StatusCode)
			log.Println(response)
		}
		return 0, ""
	}
	if response.StatusCode != http.StatusOK {
		return 0, ""
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		if s.Debug {
			log.Fatal(err)
		}
		return 0, ""
	}
	if s.Debug {
		log.Println(string(body))
	}
	defer response.Body.Close()

	var result TokenResponse
	err2 := json.Unmarshal(body, &result)
	if err2 != nil {
		if s.Debug {
			log.Fatal(err2)
		}
		return 0, ""
	}
	os.Setenv("SALUTE_SPEECH_TOKEN", result.AccessToken)
	return result.ExpiresAt, result.AccessToken
}

func (s *SaluteSpeechApi) Recognize(data io.Reader) (*SpeechRecognizeAnswer, error) {
	// Распознать данные из аудио.
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	request, _ := http.NewRequest("POST", SaluteSpeechApiRestURL+"speech:recognize", data)
	if s.needRate(s.AudioType) && s.Rate == 0 {
		return nil, errors.New("Need set Rate for this audiotype")
	}
	contentType := s.getContentTypeByInputAudioType(s.AudioType, s.Rate)
	if contentType == "" {
		return nil, errors.New("AudioType is not valid")
	}
	request.Header.Set("Content-Type", contentType)
	request.Header.Set("Authorization", "Bearer "+s.getCurrentToken())
	request.Header.Set("X-Request-ID", s.getUuid())
	client := &http.Client{}
	response, e := client.Do(request)

	if e != nil {
		return nil, e
	}
	if response.StatusCode != http.StatusOK {
		return nil, errors.New("Error. Http code: " + string(response.StatusCode))
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var result SpeechRecognizeAnswer
	err2 := json.Unmarshal(body, &result)
	if err2 != nil {
		return nil, err2
	}
	defer response.Body.Close()
	return &result, nil
}
func (s *SaluteSpeechApi) RecognizeFile(filename string) (*SpeechRecognizeAnswer, error) {
	// Распознать данные из файла. Руками надо выставить AudioType, Rate
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return s.Recognize(file)
}

func (s *SaluteSpeechApi) Synthesize(text2speech_or_ssml string) (io.Reader, error) {
	// Создать звук из текста ,
	// проверяем AudioType и Rate - они другие и не подходят

	format := s.getFormatByOutputAudioType(s.AudioType)
	if format == "" {
		return nil, errors.New("OutputAudioType is not valid")
	}
	url := SaluteSpeechApiRestURL + "text:synthesize?format=" + format + "&voice=" + s.Voice.Id
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	request, _ := http.NewRequest("POST", url, strings.NewReader(text2speech_or_ssml))
	if s.isSSML(text2speech_or_ssml) {
		request.Header.Set("Content-Type", "application/ssml")
	} else {
		request.Header.Set("Content-Type", "application/text")
	}
	request.Header.Set("Authorization", "Bearer "+s.getCurrentToken())
	request.Header.Set("X-Request-ID", s.getUuid())
	client := &http.Client{}
	if s.Debug {
		log.Println(request)
	}
	response, e := client.Do(request)
	if s.Debug {
		log.Println(response.StatusCode)
	}
	if response.StatusCode != http.StatusOK {
		if s.Debug {
			log.Println(e)
		}
		return nil, errors.New("Response is not ok. Status code:" + string(response.StatusCode))
	}
	return response.Body, nil
}

func (s *SaluteSpeechApi) SynthesizeToFile(filename string, text2speech_or_ssml string) error {
	// Синтезирует текст в звуковой файл.
	data, err := s.Synthesize(text2speech_or_ssml)
	if err != nil {
		return err
	}
	file, err := os.OpenFile(filename, os.O_TRUNC|os.O_WRONLY|os.O_CREATE, 0666)
	defer file.Close()
	if err != nil {
		return err
	}
	buf := make([]byte, 1024)
	for {
		n, err := data.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}
		if _, err := file.Write(buf[:n]); err != nil {
			return err
		}
	}
	return nil
}
