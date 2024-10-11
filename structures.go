package salutespeech_api

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresAt   int64  `json:"expires_at"`
}

type SpeechRecognizeAnswer struct {
	Result   []string `json:"result"`
	Emotions []struct {
		Negative float64 `json:"negative"`
		Neutral  float64 `json:"neutral"`
		Positive float64 `json:"positive"`
	} `json:"emotions"`
	Status int `json:"status"`
}
