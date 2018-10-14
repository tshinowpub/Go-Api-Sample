package product

import (
	"encoding/json"
	"net/http"
)

type Product struct {
	Id   int
	Name string
}

type Stock struct {
	StockId int
	Stock   int
}

type ProductHalFormat struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Links Links  `json:"_links"`
}

type Links struct {
	Self string `json:"self"`
}

func Run() {
	http.HandleFunc("/v1/product", func(w http.ResponseWriter, r *http.Request) {
		product := Product{1, "テスト商品"}
		_links := Links{"http://localhost:8000/v1/product"}

		res, err := json.Marshal(toHalResponse(product, _links))

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/hal+json")
		w.Write(res)
	})

	http.ListenAndServe(":8000", nil)
}

func toHalResponse(product Product, links Links) ProductHalFormat {
	productHalFormat := ProductHalFormat{
		Id:    product.Id,
		Name:  product.Name,
		Links: links,
	}

	return productHalFormat
}
