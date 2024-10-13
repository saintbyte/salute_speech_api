package main

import (
	salutespeech_api "github.com/saintbyte/salute_speech_api"
	"io"
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
	data, err := s.Synthesize("Привет мир")
	file, err := os.OpenFile(os.Args[1], os.O_TRUNC|os.O_WRONLY|os.O_CREATE, 0666)
	defer file.Close()
	if err != nil {
		return
	}
	buf := make([]byte, 1024)
	for {
		n, err := data.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}
		if _, err := file.Write(buf[:n]); err != nil {
			panic(err)
		}
	}
	log.Println("Complete write to: " + os.Args[1])
}
