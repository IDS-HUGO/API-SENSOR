package entities

type Sensors struct {

    ID              int     `json:"id"`
    CalidadAire     float64 `json:"calidad_aire"`
    GasInflamable   float64 `json:"gas_inflamable"`
    Humedad         float64 `json:"humedad"`
    Presion         float64 `json:"presion"`
    TemperaturaBMP  float64 `json:"temperatura_bmp"`
	
}
