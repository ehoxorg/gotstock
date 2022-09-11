package domain

import (
	"errors"
	"log"
	"net/http"

	"github.com/edihoxhalli/gotstock/db"
)

func AddProduct(p *db.Product) (*db.Product, int, error) {
	if p.ID != nil {
		return nil, http.StatusBadRequest, errors.New("NEW PRODUCT MUST NOT HAVE VALUE FOR ID")
	}
	if p.Name == "" {
		return nil, http.StatusBadRequest, errors.New("NEW PRODUCT MUST HAVE A NAME")
	}
	if p.ProductCode == "" {
		return nil, http.StatusBadRequest, errors.New("NEW PRODUCT MUST HAVE A PRODUCT CODE")
	}
	_, err, noRows := db.GetProduct(&p.ProductCode, nil)
	if !noRows {
		log.Println(err)
		return nil, http.StatusBadRequest, errors.New("PRODUCT CODE ALREADY EXISTS IN DB")
	}
	// save
	r, err := db.InsertProduct(*p)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return r, http.StatusCreated, nil
}

func GetProduct(pcode string) (*db.Product, int, error) {
	// get
	p, err, noRows := db.GetProduct(&pcode, nil)
	if noRows {
		return nil, http.StatusNotFound, errors.New("PRODUCT CODE DOES NOT EXISTS IN DB")
	}
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return p, http.StatusOK, nil
}

func GetAll() ([]db.Product, int, error) {
	// get ALL
	ps, err := db.GetAllProducts()
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return ps, http.StatusOK, nil
}

func UpdateProduct(p *db.Product, code string) (*db.Product, int, error) {
	p.ProductCode = code
	if p.ID != nil {
		return nil, http.StatusBadRequest, errors.New("PRODUCT MUST NOT HAVE VALUE FOR ID")
	}
	if p.Name == "" {
		return nil, http.StatusBadRequest, errors.New("PRODUCT MUST HAVE A NAME")
	}
	if p.ProductCode == "" {
		return nil, http.StatusBadRequest, errors.New("PRODUCT MUST HAVE A PRODUCT CODE")
	}
	_, err, noRows := db.GetProduct(&code, nil)
	if noRows {
		log.Println(err)
		return nil, http.StatusBadRequest, errors.New("PRODUCT CODE DOES NOT EXISTS IN DB")
	}
	// update
	r, err := db.UpdateProduct(*p)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return r, http.StatusAccepted, nil
}

func DeleteProduct(pcode string) (int, error) {
	// delete
	err := db.DeleteProduct(pcode)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusCreated, nil
}
