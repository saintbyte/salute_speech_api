package main

import (
	salutespeech_api "github.com/saintbyte/salute_speech_api"
	"log"
)

func main() {

	s := salutespeech_api.NewSaluteSpeechApi()
	s.AudioType = salutespeech_api.SaluteSpeechApi_InputAudioTypeMP3
	resultStr, err := s.RecognizeFile("test_data/ty154_3.mp3")
	if err != nil {
		panic(err)
	}
	log.Println(resultStr)
}
