package it

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/edihoxhalli/gotstock/db"
)

var (
	test_p = db.Product{
		ProductCode:   "10001",
		StockQuantity: 1001,
		Name:          "Aspirin",
	}
	ApiClient *http.Client = &http.Client{Timeout: time.Duration(1) * time.Second}
)

const (
	host = "http://localhost:8080/products"
)

func TestCreate(t *testing.T) {
	req, _ := http.NewRequest(http.MethodPost, host, nil)
	b, _ := json.Marshal(test_p)
	req.Body = ioutil.NopCloser(bytes.NewReader(b))
	res, _ := ApiClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	res_p := &db.Product{}
	json.Unmarshal(body, res_p)

	if res.StatusCode != http.StatusCreated {
		t.Error("incorrect status")
	}
	if !reflect.DeepEqual(*res_p, test_p) {
		t.Errorf("expected created product response (%+v), got (%+v)", spew.Sdump(test_p), spew.Sdump(*res_p))
	}
}

func TestCreateDuplicateShouldFail(t *testing.T) {
	req, _ := http.NewRequest(http.MethodPost, host, nil)
	b, _ := json.Marshal(test_p)
	req.Body = ioutil.NopCloser(bytes.NewReader(b))
	res, _ := ApiClient.Do(req)
	defer res.Body.Close()

	if res.StatusCode != http.StatusBadRequest {
		t.Error("incorrect status")
	}
}

func TestReadSuccessfully(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, host+"/"+test_p.ProductCode, nil)
	b, _ := json.Marshal(test_p)
	req.Body = ioutil.NopCloser(bytes.NewReader(b))
	res, _ := ApiClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	res_p := &db.Product{}
	json.Unmarshal(body, res_p)

	if res.StatusCode != http.StatusOK {
		t.Error("incorrect status")
	}
	if !reflect.DeepEqual(*res_p, test_p) {
		t.Errorf("expected read product response (%+v), got (%+v)", spew.Sdump(test_p), spew.Sdump(*res_p))
	}
}

func TestReadCodeNotExists(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, host+"/not-exists", nil)
	b, _ := json.Marshal(test_p)
	req.Body = ioutil.NopCloser(bytes.NewReader(b))
	res, _ := ApiClient.Do(req)
	defer res.Body.Close()
	if res.StatusCode != http.StatusNotFound {
		t.Error("incorrect status")
	}
}

func TestUpdateSuccess(t *testing.T) {
	req, _ := http.NewRequest(http.MethodPut, host+"/"+test_p.ProductCode, nil)
	test_p.Name = "Paracetamol"
	test_p.StockQuantity = 10
	b, _ := json.Marshal(test_p)
	req.Body = ioutil.NopCloser(bytes.NewReader(b))
	res, _ := ApiClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	res_p := &db.Product{}
	json.Unmarshal(body, res_p)

	if res.StatusCode != http.StatusAccepted {
		t.Error("incorrect status")
	}
	if !reflect.DeepEqual(*res_p, test_p) {
		t.Errorf("expected created product response (%+v), got (%+v)", spew.Sdump(test_p), spew.Sdump(*res_p))
	}
}

func TestUpdateCodeNotExists(t *testing.T) {
	req, _ := http.NewRequest(http.MethodPut, host+"/not-exists", nil)
	b, _ := json.Marshal(test_p)
	req.Body = ioutil.NopCloser(bytes.NewReader(b))
	res, _ := ApiClient.Do(req)
	defer res.Body.Close()
	if res.StatusCode != http.StatusBadRequest {
		t.Error("incorrect status")
	}
}

func TestReadAllSuccessfully(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, host, nil)
	b, _ := json.Marshal(test_p)
	req.Body = ioutil.NopCloser(bytes.NewReader(b))
	res, _ := ApiClient.Do(req)
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Error("incorrect status")
	}
}

func TestDelete(t *testing.T) {
	req, _ := http.NewRequest(http.MethodDelete, host+"/"+test_p.ProductCode, nil)
	b, _ := json.Marshal(test_p)
	req.Body = ioutil.NopCloser(bytes.NewReader(b))
	res, _ := ApiClient.Do(req)
	defer res.Body.Close()
	if res.StatusCode != http.StatusNoContent {
		t.Error("incorrect status")
	}
}

func TestReadAllEmpty(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, host, nil)
	b, _ := json.Marshal(test_p)
	req.Body = ioutil.NopCloser(bytes.NewReader(b))
	res, _ := ApiClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	res_ps := &[]db.Product{}
	json.Unmarshal(body, res_ps)

	if res.StatusCode != http.StatusOK {
		t.Error("incorrect status")
	}
	if !reflect.DeepEqual(*res_ps, []db.Product{}) {
		t.Errorf("expected created product response (%+v), got (%+v)", spew.Sdump([]db.Product{}), spew.Sdump(*res_ps))
	}
}
