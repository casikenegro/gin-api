package controller

import (
	"fmt"
	"strconv"

	"github.com/casikenegro/golang-gin-api/entities"
	"github.com/casikenegro/golang-gin-api/services"
	"github.com/gin-gonic/gin"
)

type ProductController interface {
	Save(ctx *gin.Context) error
	FindAll() []entities.Product
	Update(ctx *gin.Context) error
	Delete(ctx *gin.Context) error
	FindOne(ctx *gin.Context) *entities.Product
	AddImage(ctx *gin.Context) error
	DeleteImage(ctx *gin.Context) error
}

type controller struct {
	service services.ProductService
}

func New(service services.ProductService) ProductController {
	return &controller{service: service}
}

// CreateTags		godoc
// @Summary			Get Products
// @Description		GET products data in Db.
// @Param			product body []entities.Product true "Get product"
// @Produce			application/json
// @Tags			product
// @Success			200
// @Router			/products [get]
func (c *controller) FindAll() []entities.Product {
	return c.service.FindAll()
}

// CreateTags		godoc
// @Summary			Create Products
// @Description		Save product data in Db.
// @Param			product body entities.Product true "Create product"
// @Produce			application/json
// @Tags			product
// @Success			200
// @Router			/products [post]
func (c *controller) Save(ctx *gin.Context) error {
	var product entities.Product
	if err := ctx.ShouldBindJSON(&product); err != nil {
		return err
	}
	if err := c.service.Save(&product); err != nil {
		return err
	}
	return nil
}

// CreateTags		godoc
// @Summary			Update Products
// @Description		Update product data in Db.
// @Param			product body entities.Product true "Update product"
// @Produce			application/json
// @Tags			product
// @Success			200
// @Router			/products/:id [put]
func (c *controller) Update(ctx *gin.Context) error {
	var product entities.Product
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	if err := ctx.ShouldBindJSON(&product); err != nil {
		return err
	}
	product.ID = uint(id)
	if err := c.service.Update(&product); err != nil {
		return err
	}
	return nil
}

// CreateTags		godoc
// @Summary			Add image to image
// @Description		Save image data in Db.
// @Param			image body entities.ProductImage true "Add Image "
// @Produce			application/json
// @Tags			image
// @Success			200
// @Router			/products/:id/image [post]
func (c *controller) AddImage(ctx *gin.Context) error {
	var product entities.ProductImage
	if err := ctx.ShouldBindJSON(&product); err != nil {
		return err
	}
	if err := c.service.AddImage(&product); err != nil {
		return err
	}
	return nil
}

// CreateTags		godoc
// @Summary			Delete  product
// @Description		Delete product data in Db.
// @Produce			application/json
// @Tags			product
// @Success			200
// @Router			/products/:id [delete]
func (c *controller) Delete(ctx *gin.Context) error {
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	if err := c.service.Delete(uint(id)); err != nil {
		return err
	}
	return nil
}

// CreateTags		godoc
// @Summary			FindOne Products
// @Description		FindOne product data in Db.
// @Param			product body entities.Product true "FindOne product"
// @Produce			application/json
// @Tags			product
// @Success			200
// @Router			/products/:sku [get]
func (c *controller) FindOne(ctx *gin.Context) *entities.Product {
	sku := ctx.Param("sku")
	fmt.Println(sku)
	return c.service.FindOne(sku)
}

// CreateTags		godoc
// @Summary			Delete image to image
// @Description		Delete image data in Db.
// @Param			image body entities.ProductImage true "Delete Image "
// @Produce			application/json
// @Tags			image
// @Success			200
// @Router			/products/:id/image [delete]
func (c *controller) DeleteImage(ctx *gin.Context) error {
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	if err := c.service.DeleteImage(uint(id)); err != nil {
		return err
	}
	return nil
}
