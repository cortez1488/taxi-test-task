package json_to_struct

import (
	"encoding/json"
	"log"
	"taxiTestTask/models"
)

func Parse(body []byte, result *[]models.TaxiData) error {
	var input models.TaxiRawDataV2
	err := json.Unmarshal(body, &input)
	if err != nil {
		log.Fatal(err.Error())
	}
	for _, one := range input.Ones {
		*result = append(*result, models.TaxiData{
			Geo: models.Geometry{
				Coords: one.Geo.Coords,
			},
			Attrs: models.Attributes{
				Name:        one.Props.Attributes.Name,
				AdmArea:     one.Props.Attributes.AdmArea,
				District:    one.Props.Attributes.District,
				Address:     one.Props.Attributes.Address,
				CarCapacity: one.Props.Attributes.CarCapacity,
				Mode:        one.Props.Attributes.Mode,
				GlobalId:    one.Props.Attributes.GlobalId,
			}})
	}
	return nil
}
