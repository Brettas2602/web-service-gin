package usecase

import (
	"fmt"
	"web-service-gin/model"
	"web-service-gin/repository"
)

type ProductUsecase struct {
	repository repository.ProductRepository
}

func NewProductUseCase(repo repository.ProductRepository) ProductUsecase {
	return ProductUsecase{
		repository: repo,
	}
}

func (pu *ProductUsecase) GetProducts() ([]model.Product, error) {
	productList, err := pu.repository.GetProducts()
	if err != nil {
		fmt.Println(err)
	}

	return productList, err
}

func (pu *ProductUsecase) CreateProduct(product *model.Product) (*model.Product, error) {
	productId, err := pu.repository.CreateProduct(product)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	product.ID = productId

	return product, err
}

func (pu *ProductUsecase) GetProductById(id int) (*model.Product, error) {
	product, err := pu.repository.GetProductById(id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return product, err
}