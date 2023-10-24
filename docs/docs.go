// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/products": {
            "get": {
                "description": "GET products data in Db.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "product"
                ],
                "summary": "Get Products",
                "parameters": [
                    {
                        "description": "Get product",
                        "name": "product",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entities.Product"
                            }
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            },
            "post": {
                "description": "Save product data in Db.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "product"
                ],
                "summary": "Create Products",
                "parameters": [
                    {
                        "description": "Create product",
                        "name": "product",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entities.Product"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/products/:id": {
            "put": {
                "description": "Update product data in Db.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "product"
                ],
                "summary": "Update Products",
                "parameters": [
                    {
                        "description": "Update product",
                        "name": "product",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entities.Product"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            },
            "delete": {
                "description": "Delete product data in Db.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "product"
                ],
                "summary": "Delete  product",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/products/:id/image": {
            "post": {
                "description": "Save image data in Db.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "image"
                ],
                "summary": "Add image to image",
                "parameters": [
                    {
                        "description": "Add Image ",
                        "name": "image",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entities.ProductImage"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            },
            "delete": {
                "description": "Delete image data in Db.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "image"
                ],
                "summary": "Delete image to image",
                "parameters": [
                    {
                        "description": "Delete Image ",
                        "name": "image",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entities.ProductImage"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/products/:sku": {
            "get": {
                "description": "FindOne product data in Db.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "product"
                ],
                "summary": "FindOne Products",
                "parameters": [
                    {
                        "description": "FindOne product",
                        "name": "product",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entities.Product"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        }
    },
    "definitions": {
        "entities.Product": {
            "type": "object"
        },
        "entities.ProductImage": {
            "type": "object"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:3000",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Tag Service API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
