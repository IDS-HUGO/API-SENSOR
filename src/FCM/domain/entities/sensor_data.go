package entities

type SensorData struct {
	DeviceID           string  `json:"device_id"`
	CalidadAire        float64 `json:"calidad_aire"`
	GasInflamable      float64 `json:"gas_inflamable"`
	Humedad            float64 `json:"humedad"`
	Presion            float64 `json:"presion"`
	TemperaturaBMP     float64 `json:"temperatura_bmp"`
	TemperaturaDHT     float64 `json:"temperatura_dht"`
}



