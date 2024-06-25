package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/google/uuid"
	collectorDto "github.com/tudemaha/data-collector/internal/collector/dto"
	globalDto "github.com/tudemaha/data-collector/internal/global/dto"
	"github.com/tudemaha/data-collector/pkg"
)

func HandleCollectData() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var response globalDto.Response
		w.Header().Set("Content-Type", "application/json")

		if r.Method != "POST" {
			response.MethodNotAllowed()
			w.WriteHeader(response.Status)
			json.NewEncoder(w).Encode(response)
			return
		}

		if err := r.ParseMultipartForm(2 << 10); err != nil {
			response.InternalServerError()
			response.Message = err.Error()
			w.WriteHeader(response.Status)
			json.NewEncoder(w).Encode(response)
			return
		}

		var data collectorDto.Data
		data.Humidity = r.FormValue("humidity")
		data.Temperature = r.FormValue("temperature")
		data.SoilMoisture = r.FormValue("soil_moisture")
		data.SoilPH = r.FormValue("soil_ph")
		data.Gas = r.FormValue("gas")
		data.Langitude = r.FormValue("langitude")
		data.Latitude = r.FormValue("latitude")
		data.Image = uuid.New().String()

		file, header, err := r.FormFile("image")
		if err != nil {
			response.BadRequest()
			response.Message = err.Error()
			w.WriteHeader(response.Status)
			json.NewEncoder(w).Encode(response)
			return
		}
		defer file.Close()

		filenameSplit := strings.Split(header.Filename, ".")
		data.Image += "." + filenameSplit[len(filenameSplit)-1]

		dst, err := os.Create("images/" + data.Image)
		if err != nil {
			response.InternalServerError()
			response.Message = err.Error()
			w.WriteHeader(response.Status)
			json.NewEncoder(w).Encode(response)
			return
		}
		defer dst.Close()

		_, err = io.Copy(dst, file)
		if err != nil {
			response.InternalServerError()
			response.Message = err.Error()
			w.WriteHeader(response.Status)
			json.NewEncoder(w).Encode(response)
			return
		}

		db := pkg.DBConnection

		stmt, err := db.Prepare("INSERT INTO sensors" +
			"(humidity, temp, soil_moisture, soil_ph, gas, latitude, langitude, image)" +
			"VALUE (?, ?, ?, ?, ?, ?, ?, ?)")
		if err != nil {
			response.InternalServerError()
			response.Message = err.Error()
			w.WriteHeader(response.Status)
			json.NewEncoder(w).Encode(response)
			return
		}

		_, err = stmt.Exec(data.Humidity, data.Temperature, data.SoilMoisture, data.SoilPH, data.Gas, data.Langitude, data.Langitude, data.Image)
		if err != nil {
			response.InternalServerError()
			response.Message = err.Error()
			w.WriteHeader(response.Status)
			json.NewEncoder(w).Encode(response)
			return
		}

		response.OK()
		json.NewEncoder(w).Encode(response)
	}
}
