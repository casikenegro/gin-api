package services

import (
	"github.com/casikenegro/golang-gin-api/entities"
	"github.com/casikenegro/golang-gin-api/repository"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const (
	ID             = 1
	Sku            = "FAL-8406270"
	Name           = "Test"
	Brand          = "test"
	Size           = "m"
	Price          = 1.1
	PrincipalImage = "https://image.com"
)

var testProduct entities.Product = entities.Product{}

var _ = Describe("Product Service", func() {

	var (
		productRepository repository.ProductRepository
		productService    ProductService
	)

	BeforeSuite(func() {
		productRepository = repository.NewProductRepository()
		productService = New(productRepository)
	})

	Describe("Fetching all existing products", func() {

		Context("If there is a product in the database", func() {

			BeforeEach(func() {
				productService.Save(&testProduct)
			})

			It("should return at least one element", func() {
				productList := productService.FindAll()

				Ω(productList).ShouldNot(BeEmpty())
			})

			It("should map the fields correctly", func() {
				firstProduct := productService.FindAll()[0]

				Ω(firstProduct.Name).Should(Equal(Name))
				Ω(firstProduct.SKU).Should(Equal(Sku))
				Ω(firstProduct.Brand).Should(Equal(Brand))
				Ω(firstProduct.Size).Should(Equal(Size))
				Ω(firstProduct.Price).Should(Equal(Price))
				Ω(firstProduct.PrincipalImage).Should(Equal(PrincipalImage))
			})

			AfterEach(func() {
				product := productService.FindAll()[0]
				productService.Delete(product.ID)
			})

		})

		Context("If there are no products in the database", func() {

			It("should return an empty list", func() {
				products := productService.FindAll()

				Ω(products).Should(BeEmpty())
			})

		})
	})

})
