package main

import (
	"bytes"
	"context"
	"embed"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

var buf *bytes.Buffer

func getData(line string) (interface{}, error) {
	var location struct {
		URL string `json:"location"`
	}

	if err := json.Unmarshal([]byte(line), &location); err != nil {
		log.Fatal(err)
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, location.URL, nil)

	if err != nil {
		return nil, err
	}

	c := &http.Client{}
	res, err := c.Do(req)

	if err != nil {
		return nil, err
	}

	if _, err := io.Copy(buf, res.Body); err != nil {
		return nil, err
	}

	var payload struct {
		Data string `bson:"data"`
	}

	if err := json.Unmarshal(buf.Bytes(), &payload); err != nil {
		return nil, err
	}

	return payload.Data, nil
}

//go:embed input.txt
var inputData embed.FS

func main() {
	data, err := inputData.ReadFile("input.txt")
	var wg sync.WaitGroup

	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		wg.Add(1)

		go func(line string) {
			defer wg.Done()

			fmt.Println("In goroutine, calling getData with line:", line)
			data, err := getData(line)

			if err != nil {
				log.Printf("unable to get status: %v", err)
				return
			}

			strData, ok := data.(string)

			if !ok {
				log.Printf("unexpected type of data: %T", data)
				return
			}

			if strData == "foo" {
				fmt.Printf("data found: %s", data)
				return
			}

		}(line)
	}

	wg.Wait()
}
