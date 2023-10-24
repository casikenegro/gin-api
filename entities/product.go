package entities

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ID             uint           `gorm:"primary_key;auto_incremet"`
	SKU            string         `json:"sku" binding:"required,startswith=FAL-,min=11,max=12" gorm:"uniqueIndex;type:varchar(50)"`
	Name           string         `json:"name" binding:"required,min=3,max=50" gorm:"type:varchar(50)"`
	Brand          string         `json:"brand" binding:"required,min=3,max=50" gorm:"type:varchar(50)"`
	Size           string         `json:"size" gorm:"type:varchar(100)"`
	Price          float64        `json:"price" binding:"required,gte=1,lte=99999999" gorm:"type:decimal(9,2)"`
	PrincipalImage string         `json:"principal_image" binding:"required,url" gorm:"type:varchar(50)"`
	ProductImages  []ProductImage `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type ProductImage struct {
	gorm.Model
	ID        uint   `gorm:"primary_key;auto_incremet" json:"id"`
	ProductID uint   `json:"product_id"`
	Url       string `json:"url" binding:"required,url" gorm:"type:varchar(50)"`
}
