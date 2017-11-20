package service

import (
	"net/http"

	"github.com/unrolled/render"
)

func apiTestHandler(formatter *render.Render) http.HandlerFunc {
	
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct {
			Name  string `json:"name"`
			Price string `json:"price"`
		}{Name: "Clearlove7", Price: "4396"})
	}
}