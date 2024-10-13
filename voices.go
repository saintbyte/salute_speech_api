package salutespeech_api

type Voice struct {
	Id     string
	Name   string
	Lang   string
	Gender string
	Rate   int
}

func SaluteSpeechVoices() []Voice {
	return []Voice{
		Voice{
			Id:     "Nec_24000",
			Name:   "Наталья",
			Lang:   "ru",
			Gender: "female",
			Rate:   24000,
		},
		Voice{
			Id:     "Nec_8000",
			Name:   "Наталья",
			Lang:   "ru",
			Gender: "female",
			Rate:   8000,
		},
		Voice{
			Id:     "Bys_24000",
			Name:   "Борис",
			Lang:   "ru",
			Gender: "male",
			Rate:   24000,
		},
		Voice{
			Id:     "Bys_8000",
			Name:   "Борис",
			Lang:   "ru",
			Gender: "male",
			Rate:   8000,
		},
		Voice{
			Id:     "May_24000",
			Name:   "Марфа",
			Lang:   "ru",
			Gender: "female",
			Rate:   24000,
		},
		Voice{
			Id:     "May_8000",
			Name:   "Марфа",
			Lang:   "ru",
			Gender: "female",
			Rate:   8000,
		},
		Voice{
			Id:     "Tur_24000",
			Name:   "Тарас",
			Lang:   "ru",
			Gender: "male",
			Rate:   24000,
		},
		Voice{
			Id:     "Tur_8000",
			Name:   "Тарас",
			Lang:   "ru",
			Gender: "male",
			Rate:   8000,
		},
		Voice{
			Id:     "Ost_24000",
			Name:   "Александра",
			Lang:   "ru",
			Gender: "female",
			Rate:   24000,
		},
		Voice{
			Id:     "Ost_8000",
			Name:   "Александра",
			Lang:   "ru",
			Gender: "female",
			Rate:   8000,
		},
		Voice{
			Id:     "Pon_24000",
			Name:   "Сергей",
			Lang:   "ru",
			Gender: "male",
			Rate:   24000,
		},
		Voice{
			Id:     "Pon_8000",
			Name:   "Сергей",
			Lang:   "ru",
			Gender: "male",
			Rate:   8000,
		},
		Voice{
			Id:     "Kin_24000",
			Name:   "Kira",
			Lang:   "en",
			Gender: "female",
			Rate:   24000,
		},
		Voice{
			Id:     "Kin_8000",
			Name:   "Kira",
			Lang:   "en",
			Gender: "female",
			Rate:   8000,
		},
	}
}
