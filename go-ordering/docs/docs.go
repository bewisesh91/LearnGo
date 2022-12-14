// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
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
        "/oos/order/changeOrder": {
            "post": {
                "description": "ChangeOrder 주문 변경 - 주문자",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "call ChangeOrder, return ok by json.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Controller"
                        }
                    }
                }
            }
        },
        "/oos/order/createReview": {
            "post": {
                "description": "CreateReview 리뷰 등록 - 주문자",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "call CreateReview, return ok by json.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Controller"
                        }
                    }
                }
            }
        },
        "/oos/order/newOrder": {
            "post": {
                "description": "NewOrder 주문 등록 - 주문자",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "call NewOrder, return ok by json.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Controller"
                        }
                    }
                }
            }
        },
        "/oos/order/searchMenu": {
            "post": {
                "description": "SearchMenu 메뉴 검색 - 주문자, 피주문자",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "call SearchMenu, return ok by json.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Controller"
                        }
                    }
                }
            }
        },
        "/oos/order/searchOrder": {
            "post": {
                "description": "SearchOrder 주문 내역 조회 기능 - 주문자",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "call SearchOrder, return ok by json.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Controller"
                        }
                    }
                }
            }
        },
        "/oos/order/viewMenu": {
            "post": {
                "description": "ViewMenu 메뉴 상세 - 주문자, 피주문자",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "call ViewMenu, return ok by json.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Controller"
                        }
                    }
                }
            }
        },
        "/oos/seller/createMenu": {
            "post": {
                "description": "CreateMenu 메뉴 등록 - 피주문자",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "call CreateMenu, return ok by json.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Controller"
                        }
                    }
                }
            }
        },
        "/oos/seller/deleteMenu": {
            "post": {
                "description": "DeleteMenu 메뉴 삭제 - 피주문자",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "call DeleteMenu, return ok by json.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Controller"
                        }
                    }
                }
            }
        },
        "/oos/seller/orderStates": {
            "post": {
                "description": "OrderStates 주문 내역 조회 - 피주문자",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "call OrderStates, return ok by json.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Controller"
                        }
                    }
                }
            }
        },
        "/oos/seller/updateMenu": {
            "post": {
                "description": "UpdateMenu 메뉴 수정 - 피주문자",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "call UpdateMenu, return ok by json.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Controller"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controller.Controller": {
            "type": "object"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "WBA [Backend Final Project]",
	Description:      "띵동주문이요, 온라인 주문 시스템(Online Ordering System)",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
