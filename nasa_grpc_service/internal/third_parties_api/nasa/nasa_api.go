package nasa

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/Projectoutlast/space_service/nasa_grpc_service/internal/config"
)

type NasaAPI struct {
	log    *slog.Logger
	config *config.Config
}

func New(log *slog.Logger, config *config.Config) *NasaAPI {
	return &NasaAPI{
		log:    log,
		config: config,
	}
}

func (n *NasaAPI) GetRandomSpaseImage() (*RandomSpaseImageResponse, error) {
	var imageData *RandomSpaseImageResponse
	var err error

	for i := 0; i < n.config.NasaConfig.MaxRetries; i++ {
		imageData, err = n.getImageData()
		if err != nil {
			return nil, err
		}

		if n.checkMediaType(imageData.MediaType) {
			break
		}

		if i == n.config.NasaConfig.MaxRetries-1 {
			return nil, fmt.Errorf("за 10 запросов неудалось получить от сервиса картинку. Попробуйте снова")
		}
	}

	res, err := n.getImage(imageData)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (n *NasaAPI) getImageData() (*RandomSpaseImageResponse, error) {
	url := fmt.Sprintf("%s?api_key=%s&count=1", n.config.NasaConfig.BaseURL, n.config.NasaConfig.ApiKey)

	resp, err := http.Get(url)

	if err != nil {
		n.log.Error(err.Error())
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		n.log.Warn("некорректный статус код при обращении к %s: %d", n.config.NasaConfig.BaseURL, resp.StatusCode)
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		n.log.Error(err.Error())
		return nil, err
	}

	var imageResp *RandomSpaseImageResponse
	err = json.Unmarshal(body, &imageResp)
	if err != nil {
		n.log.Error(err.Error())
		return nil, err
	}

	return imageResp, nil
}

func (n *NasaAPI) getImage(imageData *RandomSpaseImageResponse) (*RandomSpaseImageResponse, error) {
	image, err := http.Get(imageData.Url)
	if err != nil {
		n.log.Error(err.Error())
		return nil, err
	}

	defer image.Body.Close()

	imageData.Data, err = io.ReadAll(image.Body)
	if err != nil {
		n.log.Error(err.Error())
		return nil, err
	}

	return imageData, nil
}

func (n *NasaAPI) checkMediaType(mediaType string) bool {
	return mediaType == "image"
}
