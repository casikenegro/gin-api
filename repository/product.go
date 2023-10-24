package repository

import (
	"errors"

	"github.com/casikenegro/golang-gin-api/database"
	"github.com/casikenegro/golang-gin-api/entities"
)

type ProductRepository interface {
	Save(product *entities.Product) error
	Update(product *entities.Product) error
	Delete(id uint) error
	FindOne(sku string) *entities.Product
	AddImage(productImage *entities.ProductImage) error
	DeleteImage(id uint) error
	FindAll() []entities.Product
	GetImages(id uint) []entities.ProductImage
}

type productRepository struct{}

func NewProductRepository() ProductRepository {
	return &productRepository{}
}
func (r *productRepository) Save(product *entities.Product) error {
	newProduct := database.DB.Create(product)

	if newProduct.Error != nil {
		return newProduct.Error
	}
	return nil
}

func (r *productRepository) Update(product *entities.Product) error {
	saveProduct := database.DB.Save(product)
	if saveProduct.Error != nil {
		return saveProduct.Error
	}
	return nil
}
func (r *productRepository) Delete(id uint) error {
	var product entities.Product
	database.DB.First(&product, id)

	if product.ID == 0 {
		return errors.New("product not exist")
	}

	database.DB.Unscoped().Delete(&product)
	return nil
}

func (r *productRepository) FindOne(sku string) *entities.Product {
	var product entities.Product
	database.DB.Where("sku = ?", sku).First(&product)
	if product.ID == 0 {
		return nil
	}
	return &product
}

func (r *productRepository) AddImage(image *entities.ProductImage) error {
	newImage := database.DB.Create(image)
	if newImage.Error != nil {
		return newImage.Error
	}
	return nil
}

func (r *productRepository) DeleteImage(id uint) error {
	var image entities.ProductImage
	database.DB.First(&image, id)

	if image.ID == 0 {
		return errors.New("image not exist")
	}

	database.DB.Unscoped().Delete(&image)
	return nil
}
func (r *productRepository) FindAll() []entities.Product {
	var products []entities.Product
	database.DB.Find(&products)
	return products
}

func (r *productRepository) GetImages(id uint) []entities.ProductImage {
	var images []entities.ProductImage
	database.DB.Find(&images)
	return images
}
