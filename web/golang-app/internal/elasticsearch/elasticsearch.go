package elasticsearch

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"math/rand"
	"net"
	"net/http"
	"strconv"
	"time"

	elasticsearch8 "github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

type Country struct {
	Name       string `json:"ulke"`
	Population int    `json:"nufus"`
	Capital    string `json:"baskent"`
}

// Config represents the Elasticsearch configuration
type Config struct {
	Addresses []string `mapstructure:"cluster_addresses"`
}

type ES struct {
	client *elasticsearch8.Client
}

func New(config *Config) (*ES, error) {

	cfg := elasticsearch8.Config{
		Addresses: config.Addresses,
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: time.Second,
			DialContext:           (&net.Dialer{Timeout: time.Second}).DialContext,
			TLSClientConfig: &tls.Config{
				MinVersion: tls.VersionTLS12,
			},
		},
	}

	client, err := elasticsearch8.NewClient(cfg)
	if err != nil {
		return nil, err
	}

	return &ES{client}, nil
}

// Get randomly selected one document(country) in the ulkeler Index
func (es *ES) GetOneRandomCountry() (*Country, error) {

	idx := "ulkeler"
	doc_id := strconv.Itoa(rand.Intn(10) + 1)

	req := esapi.GetRequest{
		Index:      idx,
		DocumentID: doc_id,
	}

	resp, err := req.Do(context.Background(), es.client.Transport)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.IsError() {
		return nil, fmt.Errorf("error response received: %s", resp.Status())
	}

	// Parse the response body
	var body map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return nil, fmt.Errorf("error parsing the response body: %s", err)
	}

	// Extract the source data which contains the document field source.
	data, ok := body["_source"].(map[string]interface{})
	if !ok {
		// If the type assertion fails
		return nil, fmt.Errorf("error asserting response body as map")
	}

	c := &Country{
		Name:       data["ulke"].(string),
		Population: int(data["nufus"].(float64)),
		Capital:    data["baskent"].(string),
	}

	return c, nil
}
