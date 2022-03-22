package telegram

type API struct {
	token   string
	baseURL string
}

func New(token, baseURL string) *TelegramAPI {
	return &API{token: token, baseURl: baseURL}
}

func (api *API) GetUpdates() ([]Update, error) {
}

func (api *API) SendMessage() {
}
