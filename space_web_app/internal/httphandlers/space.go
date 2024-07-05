package httphandlers

import (
	"context"
	"encoding/base64"
	"io"
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

	resp, err := h.getRandomSpaseImageStreamProcess(ctx, &req)

	session, _ := h.store.Get(r, "flash-session")
	if err != nil {
		session.AddFlash(err.Error(), "error")
		session.Save(r, w)

		http.Redirect(w, r, "/", http.StatusSeeOther)
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

func (c *HTTPHandlers) getRandomSpaseImageStreamProcess(ctx context.Context, req *pb.RandomSpaseImageRequest) (*pb.RandomSpaseImageResponse, error) {
	stream, err := c.nasaClient.RandomSpaseImage(ctx, req)

	if err != nil {
		c.log.Error(err.Error())
	}

	response := &pb.RandomSpaseImageResponse{}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			c.log.Error(err.Error())

			return nil, err
		}

		response.Copyright = res.Copyright
		response.Date = res.Date
		response.Explanation = res.Explanation
		response.Title = res.Title
		response.Data = res.Data
	}

	return response, nil
}
