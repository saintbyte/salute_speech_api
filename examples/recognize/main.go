package main

import (
	salutespeech_api "github.com/saintbyte/salute_speech_api"
	"log"
)

func main() {
	s := salutespeech_api.NewSaluteSpeechApi()
	s.AudioType = salutespeech_api.AudioTypeMP3
	result_str, err := s.Recognize("test_data/ty154_3.mp3")
	if err != nil {
		panic(err)
	}
	log.Println(result_str)
}
