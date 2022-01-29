package reqToAPI

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"taxiTestTask/internal/json_to_struct"
)

const (
	apiKey = "85a31f5108e65a7e9bbd6c0ade6ae33b"
)

func RequestJSON() error {
	url := fmt.Sprintf("https://apidata.mos.ru/v1/datasets/621/features?api_key=%s", apiKey)
	resp, err := http.Get(url)
	defer resp.Body.Close()

	if err != nil {
		return err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	data, err := json_to_struct.ParseBodyJson(body)
	if err != nil {
		return errors.New("parsing body: " + err.Error())
	}
	fmt.Println(data)
	return nil
}
