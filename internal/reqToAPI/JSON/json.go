package reqToAPI

import (
	"fmt"
	"io"
	"net/http"
)

const (
	apiKey = "85a31f5108e65a7e9bbd6c0ade6ae33b"
	uri    = "https://apidata.mos.ru/v1/datasets/621/features?api_key=%s"
)

func GetJSONFromAPIRequest() ([]byte, error) {
	url := fmt.Sprintf(uri, apiKey)
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	return io.ReadAll(resp.Body)
	//if err != nil {
	//	return nil, err
	//}
	//
	//err = json_to_struct.ParseBodyJson(body, input)
	//if err != nil {
	//	return nil,  errors.New("parsing body: " + err.Error())
	//}
	//
	//return body, nil
}
