package models

type TaxiData struct {
	Attrs Attributes `json:"Attributes"`
	Geo   Geometry   `json:"geometry"`
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

type TaxiRawDataV2 struct {
	Ones []struct {
		Geo struct {
			Coords []float32 `json:"coordinates"`
		} `json:"geometry"`
		Props struct {
			Attributes struct {
				Name        string `json:"Name"`
				AdmArea     string `json:"AdmArea"`
				District    string `json:"District"`
				Address     string `json:"Address"`
				CarCapacity int    `json:"CarCapacity"`
				Mode        string `json:"Mode"`
				GlobalId    int64  `json:"global_id"`
			} `json:"Attributes"`
		} `json:"properties"`
	} `json:"features"`
}
