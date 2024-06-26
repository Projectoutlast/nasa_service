package nasa

import (
	"log/slog"

	"github.com/Projectoutlast/space_service/nasa_grpc_service/internal/third_parties_api/nasa"
	pb "github.com/Projectoutlast/nasa_proto/gen"

)

type NasaAPI interface {
	GetRandomSpaseImage() (*nasa.RandomSpaseImageResponse, error)
}

type NasaUsecase struct {
	log *slog.Logger
	nasaAPI NasaAPI
}

func New(log *slog.Logger, nasaAPI NasaAPI) *NasaUsecase {
	return &NasaUsecase{
		log:     log,
		nasaAPI: nasaAPI,
	}
}

func (n *NasaUsecase) GetRandomSpaseImage() (*pb.RandomSpaseImageResponse, error) {
	nasaAPIresp, err := n.nasaAPI.GetRandomSpaseImage()
	if err != nil {
		return nil, err
	}

	res := &pb.RandomSpaseImageResponse{
		Copyright:      nasaAPIresp.Copyright,
		Date:           nasaAPIresp.Date,
		Explanation:    nasaAPIresp.Explanation,
		Title: nasaAPIresp.Title,
		Data: nasaAPIresp.Data,
	}

	return res, nil
}
