package services

import (
	"github.com/casikenegro/golang-gin-api/entities"
	"github.com/casikenegro/golang-gin-api/repository"
)

type ProductService interface {
	Save(product *entities.Product) error
	FindAll() []entities.Product
	Update(product *entities.Product) error
	Delete(id uint) error
	FindOne(sku string) *entities.Product
	AddImage(productImage *entities.ProductImage) error
	DeleteImage(id uint) error
}

type productService struct {
	productRepository repository.ProductRepository
}

func New(repo repository.ProductRepository) ProductService {
	return &productService{
		productRepository: repo,
	}
}

func (p *productService) Save(product *entities.Product) error {
	if err := p.productRepository.Save(product); err != nil {
		return err
	}
	return nil
}

func (p *productService) FindAll() []entities.Product {
	products := p.productRepository.FindAll()
	if len(products) < 1 {
		return []entities.Product{}
	}
	for i, _ := range products {
		products[i].ProductImages = p.productRepository.GetImages(products[i].ID)
	}
	return products
}

func (p *productService) Update(product *entities.Product) error {
	if err := p.productRepository.Update(product); err != nil {
		return err
	}
	return nil
}

func (p *productService) Delete(id uint) error {
	if err := p.productRepository.Delete(id); err != nil {
		return err
	}
	return nil
}
func (p *productService) FindOne(sku string) *entities.Product {
	product := p.productRepository.FindOne(sku)
	if product == nil {
		return nil
	}
	product.ProductImages = p.productRepository.GetImages(product.ID)
	return product
}
func (p *productService) AddImage(productImage *entities.ProductImage) error {
	if err := p.productRepository.AddImage(productImage); err != nil {
		return err
	}
	return nil
}
func (p *productService) DeleteImage(id uint) error {
	if err := p.productRepository.DeleteImage(id); err != nil {
		return err
	}
	return nil
}
