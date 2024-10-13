package main

import (
	salutespeech_api "github.com/saintbyte/salute_speech_api"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Println("Usage program.exe [output_file].ogg")
		return
	}
	log.Println("Start ")
	s := salutespeech_api.NewSaluteSpeechApi()
	s.Debug = false
	s.AudioType = salutespeech_api.SaluteSpeechApi_OutputAudioTypeOPUS
	s.Voice = s.GetVoiceById("Ost_24000")
	err := s.SynthesizeToFile(os.Args[1], "Привет мир.")
	if err != nil {
		panic(err)
	}
	log.Println("Complete write to: " + os.Args[1])
}
