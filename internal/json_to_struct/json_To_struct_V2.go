package json_to_struct

import (
	"encoding/json"
	"log"
	"taxiTestTask/models"
)

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

func Parse(body []byte, result *[]models.TaxiData) error {
	var input TaxiRawDataV2
	err := json.Unmarshal(body, &input)
	if err != nil {
		log.Fatal(err.Error())
	}
	for _, one := range input.Ones {
		*result = append(*result, models.TaxiData{
			Name:        one.Props.Attributes.Name,
			AdmArea:     one.Props.Attributes.AdmArea,
			District:    one.Props.Attributes.District,
			Address:     one.Props.Attributes.Address,
			CarCapacity: one.Props.Attributes.CarCapacity,
			Mode:        one.Props.Attributes.Mode,
			GlobalId:    one.Props.Attributes.GlobalId,
			CoordX:      one.Geo.Coords[0],
			CoordY:      one.Geo.Coords[1],
		})
	}
	return nil
}
