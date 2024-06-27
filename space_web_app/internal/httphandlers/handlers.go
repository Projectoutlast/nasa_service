package httphandlers

import (
	"context"
	"log/slog"
	"net/http"

	pb "github.com/Projectoutlast/nasa_proto/gen"
)

type NasaGRPCClient interface {
	GetRandomSpaseImage(context.Context, *pb.RandomSpaseImageRequest) (*pb.RandomSpaseImageResponse, error)
}

type HTTPHandlers struct {
	log        *slog.Logger
	nasaClient NasaGRPCClient
}

func New(log *slog.Logger, nasaClient NasaGRPCClient) *HTTPHandlers {
	return &HTTPHandlers{
		log:        log,
		nasaClient: nasaClient,
	}
}

func (h *HTTPHandlers) GetRandomSpaseImage(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	req := pb.RandomSpaseImageRequest{}

	resp, err := h.nasaClient.GetRandomSpaseImage(ctx, &req)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "image/jpg")
	w.WriteHeader(http.StatusOK)
	w.Write(resp.Data)
}
