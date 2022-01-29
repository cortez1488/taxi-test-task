package models

type TaxiRawData struct {
	Attrs Attributes //`json:"features.Attributes"`
	Geo   Geometry   //`json:"geometry"`
}

type Attributes struct {
	Name        string
	AdmArea     string
	District    string
	Address     string
	CarCapacity int
	Mode        string
	GlobalId    int64 `json:"global_id"`
}

type Geometry struct {
	Coords []float32 `json:"coordinates"`
}
