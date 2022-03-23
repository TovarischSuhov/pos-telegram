package telegram

import (
	"bytes"
	"flag"
	"io/ioutil"
	"net/http"
	"sync"
	"time"

	"github.com/TovarischSuhov/pos-telegram/domain"
	"github.com/TovarischSuhov/pos-telegram/telegram/models"
	"github.com/mailru/easyjson"
)

type API struct {
	token   string
	baseURL string

	client *http.Client
}

var (
	token   = flag.String("telegram_token", "", "Telegram bot token")
	baseURL = flag.String("telegram_base_url", "https://api.telegram.org", "Telegram api host")

	once sync.Once
	_api *API
)

const (
	longpollTimeout = time.Second * 0
)

func GetAPI() *API {
	once.Do(func() {
		_api = &API{
			token:   *token,
			baseURL: *baseURL,
			client:  http.Client{Timeout: longpollTimeout},
		}
	})
	return _api
}

func (api *API) makeUrl(method string) string {
	result := api.baseURL + "/bot" + api.token + "/" + method
	return result
}

func (api *API) GetUpdates(from domain.Offset) (models.UpdateResponse, error) {
	message := models.GetUpdateRequest{Offset: from}
	body, err := easyjson.Marshal(message)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, api.makeURL("getUpdates"), bytes.NewBuffer(body))
	req.Header.Add("Context-Type", "application/json")
	response, err := api.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	result := UpdateResponse{}
	err := easyjson.Unmarshal(responseBody, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (api *API) SendMessage() {
}
