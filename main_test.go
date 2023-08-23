package main

import (
	"bytes"
	crand "crypto/rand"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"testing"
)

const BodySize = 256
const DbName = "bench"

func setValue(num int) {
	key := rand.Intn(num + 1)
	url := fmt.Sprintf("http://localhost:8080/dbs/%s/entry?key=%d", DbName, key)

	buffer := make([]byte, BodySize)
	_, err := crand.Read(buffer)
	if err != nil {
		panic(err)
	}

	body := bytes.NewReader(buffer)
	res, err := http.Post(url, "text/plain", body)
	if err != nil {
		panic(err)
	}

	if res.StatusCode != http.StatusOK {
		panic(errors.New("status is not ok"))
	}

}

func BenchmarkSetValue(t *testing.B) {
	for i := 0; i < t.N; i++ {
		setValue(i)
	}
}
