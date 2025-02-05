package controller

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
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
		data.NodeID = r.FormValue("node_id")
		data.GatewayID = r.FormValue("gateway_id")
		data.Humidity = r.FormValue("humidity")
		data.Temperature = r.FormValue("temperature")
		data.SoilMoisture = r.FormValue("soil_moisture")
		data.SoilPH = r.FormValue("soil_ph")
		data.Gas = r.FormValue("gas")
		data.Coordinate = r.FormValue("coordinate")

		var header *multipart.FileHeader
		var err error
		data.Image, header, err = r.FormFile("image")
		if err != nil {
			response.BadRequest()
			response.Message = err.Error()
			w.WriteHeader(response.Status)
			json.NewEncoder(w).Encode(response)
			return
		}
		defer data.Image.Close()

		id, _ := uuid.NewV7()

		filenameSplit := strings.Split(header.Filename, ".")
		filename := id.String() + "." + filenameSplit[len(filenameSplit)-1]

		dst, err := os.Create("images/" + filename)
		if err != nil {
			response.InternalServerError()
			response.Message = err.Error()
			w.WriteHeader(response.Status)
			json.NewEncoder(w).Encode(response)
			return
		}
		defer dst.Close()

		_, err = io.Copy(dst, data.Image)
		if err != nil {
			response.InternalServerError()
			response.Message = err.Error()
			w.WriteHeader(response.Status)
			json.NewEncoder(w).Encode(response)
			return
		}

		data.Image, _, err = r.FormFile("image")
		if err != nil {
			response.BadRequest()
			response.Message = err.Error()
			w.WriteHeader(response.Status)
			json.NewEncoder(w).Encode(response)
			return
		}
		defer data.Image.Close()

		buf := bytes.NewBuffer(nil)
		if _, err := io.Copy(buf, data.Image); err != nil {
			response.InternalServerError()
			response.Message = err.Error()
			w.WriteHeader(response.Status)
			json.NewEncoder(w).Encode(response)
			return
		}

		db := pkg.DBConnection

		stmt, err := db.Prepare("INSERT INTO sensors" +
			"(id, gateway_id, node_id, temp, humidity, soil_moisture, soil_ph, gas, coordinate, image)" +
			"VALUE (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
		if err != nil {
			response.InternalServerError()
			response.Message = err.Error()
			w.WriteHeader(response.Status)
			json.NewEncoder(w).Encode(response)
			return
		}

		_, err = stmt.Exec(id,
			data.GatewayID,
			data.NodeID,
			data.Temperature,
			data.Humidity,
			data.SoilMoisture,
			data.SoilPH,
			data.Gas,
			data.Coordinate,
			buf.Bytes())
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

func HandleRetrieveImage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		db := pkg.DBConnection

		var imgByte []byte
		_ = db.QueryRow("SELECT image FROM sensors WHERE id = ?", id).Scan(&imgByte)

		w.Header().Set("Content-Type", "image/jpeg")
		w.Write(imgByte)
	}
}
