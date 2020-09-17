package client

import (
	"io/ioutil"
	"log"
	"net/http"
)

// GenericClient as the name suggests
type GenericClient struct {
	HTTPClient http.Client
	Endpoint   string
}

// GetResponse in general []byte
// input: *GenericClient
// output: []byte
func (c *GenericClient) GetResponse() []byte {
	req, _ := http.NewRequest(http.MethodGet, c.Endpoint, nil)
	// any non-default "User-Agent", to resolve empty response bug
	req.Header.Set("User-Agent", "")
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		log.Println("error HTTPClient.Do(req):", err)
		return nil
	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	return bodyBytes
}
