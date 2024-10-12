package salutespeech_api

/*
	1. Подробнее о форматах аудио для преобразования в текст:
                    https://developers.sber.ru/docs/ru/salutespeech/recognition/encodings

    2. Преобразование текста в речь:
                     https://developers.sber.ru/docs/ru/salutespeech/synthesis/synthesis-http#parametry-zaprosa9
*/

const (
	SaluteSpeechApi_InputAudioTypePCM_S16LE = "PCM_S16LE"
	SaluteSpeechApi_InputAudioTypeOPUS      = "OPUS"
	SaluteSpeechApi_InputAudioTypeMP3       = "MP3"
	SaluteSpeechApi_InputAudioTypeFLAC      = "FLAC"
	SaluteSpeechApi_InputAudioTypeALAW      = "ALAW"
	SaluteSpeechApi_InputAudioTypeMULAW     = "MULAW"

	SaluteSpeechApi_OutputAudioTypeWAV16 = "wav16"
	SaluteSpeechApi_OutputAudioTypePCM16 = "pcm16"
	SaluteSpeechApi_OutputAudioTypeOPUS  = "opus"
	SaluteSpeechApi_OutputAudioTypeALAW  = "alaw"
)

func NeedRate(InputAudioType string) bool {
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
func GetContentTypeByInputAudioType(InputAudioType string, rate int) string {
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

func GetFormatByOutputAudioType(OutputAudioType string) string {
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
