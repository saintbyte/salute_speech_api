package main

import (
	salutespeech_api "github.com/saintbyte/salute_speech_api"
	"log"
)

func main() {
	s := salutespeech_api.NewSaluteSpeechApi()
	s.AudioType = salutespeech_api.SaluteSpeechApi_OutputAudioTypeOPUS
	resultStr, err := s.Synthesize("Привет мир")
	if err != nil {
		panic(err)
	}
	log.Println(resultStr)
}
