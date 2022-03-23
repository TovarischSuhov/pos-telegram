package telegram

import (
	"bytes"
	"net/http"

	"github.com/TovarischSuhov/pos-telegram/domain"
	"github.com/TovarischSuhov/pos-telegram/telegram/models"
	"github.com/mailru/easyjson"
)

type API struct {
	token   string
	baseURL string
}

func New(token, baseURL string) *TelegramAPI {
	return &API{token: token, baseURL: baseURL}
}

func (api *API) makeUrl(method string) string {
	result := api.baseURL + "/bot" + api.token + "/" + method
	return result
}

func (api *API) GetUpdates(from domain.Offset) ([]Update, error) {
	message := models.GetUpdateRequest{Offset: from}
	body, err := easyjson.Marshal(message)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, api.makeURL("getUpdates"), bytes.NewBuffer(body))
	req.Header.Add("Context-Type", "application/json")
}

func (api *API) SendMessage() {
}
