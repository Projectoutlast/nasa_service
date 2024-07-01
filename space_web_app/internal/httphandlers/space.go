package httphandlers

import (
	"context"
	"encoding/base64"
	"net/http"
	"text/template"

	pb "github.com/Projectoutlast/nasa_proto/gen"
)

type randomImageData struct {
	PageTitle string
	Content   string
}

func (h *HTTPHandlers) GetRandomSpaseImage(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	req := pb.RandomSpaseImageRequest{}

	resp, err := h.nasaClient.GetRandomSpaseImage(ctx, &req)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	files := []string{
		"./assets/html/random.html",
		baseSpaceLayout,
	}

	tmpl := template.Must(template.ParseFiles(files...))

	base64Image := base64.StdEncoding.EncodeToString(resp.Data)
	imageDataUrl := "data:image/jpeg;base64," + base64Image

	data := randomImageData{
		PageTitle: "Random Space Image",
		Content:   imageDataUrl,
	}

	tmpl.Execute(w, data)
}
