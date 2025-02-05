package dto

import "mime/multipart"

type Data struct {
	NodeID       string         `json:"nodes_id"`
	GatewayID    string         `json:"gateway_id"`
	Humidity     string         `json:"humidity"`
	Temperature  string         `json:"temperature"`
	SoilMoisture string         `json:"soil_moisture"`
	SoilPH       string         `json:"soil_ph"`
	Gas          string         `json:"gas"`
	Coordinate   string         `json:"coordinate"`
	Image        multipart.File `json:"image"`
}
