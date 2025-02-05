package routes

import (
	"net/http"

	"github.com/tudemaha/data-collector/internal/collector/controller"
)

func Router() {
	http.Handle("/upload", controller.HandleCollectData())
	http.Handle("/image", controller.HandleRetrieveImage())
}
