package json_to_struct

import (
	"encoding/json"
	"errors"
	"taxiTestTask/models"
)

func ParseBodyJson(byt []byte, input *[]models.TaxiRawData) error {
	ones, _ := getOnes(byt)
	for _, data := range ones {
		taxiData, err := dataToTaxiData(data)
		if err != nil {
			return err
		}
		*input = append(*input, *taxiData)
	}
	return nil
}

func dataToTaxiData(data interface{}) (*models.TaxiRawData, error) {
	var taxiData models.TaxiRawData
	attrs := getAttrs(data)
	geo := getGeo(data)

	attrsByte, _ := json.Marshal(attrs)
	geoBytes, _ := json.Marshal(geo)

	err := unmarshallingToStruct(attrsByte, geoBytes, &taxiData)
	if err != nil {
		return nil, err
	}
	return &taxiData, nil
}

func unmarshallingToStruct(attrsBytes, geoBytes []byte, taxiData *models.TaxiRawData) error {
	err := json.Unmarshal(attrsBytes, &taxiData.Attrs)
	if err != nil {
		return err
	}

	err = json.Unmarshal(geoBytes, &taxiData.Geo)
	if err != nil {
		return err
	}
	return nil
}

func getAttrs(data interface{}) interface{} {
	return toMap(toMap(data)["properties"])["Attributes"]
}

func getGeo(data interface{}) interface{} {
	return toMap(data)["geometry"]
}

func getOnes(byt []byte) ([]interface{}, error) {
	var rawMap map[string]interface{}
	if err := json.Unmarshal(byt, &rawMap); err != nil {
		return nil, errors.New("unmarshalling raw data: " + err.Error())
	}
	return rawMap["features"].([]interface{}), nil
}

func toMap(v interface{}) map[string]interface{} {
	return v.(map[string]interface{})
}
