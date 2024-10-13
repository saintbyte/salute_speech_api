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
