// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth/login": {
            "post": {
                "description": "Login user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Login user",
                "parameters": [
                    {
                        "description": "request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.AuthLoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.AuthLoginResponseDoc"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/response.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.errorResponse"
                        }
                    }
                }
            }
        },
        "/auth/register": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Register user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Register user",
                "parameters": [
                    {
                        "description": "request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.AuthRegisterRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.AuthRegisterResponseDoc"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/response.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.errorResponse"
                        }
                    }
                }
            }
        },
        "/samples": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get samples",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "samples"
                ],
                "summary": "Get samples",
                "parameters": [
                    {
                        "type": "string",
                        "name": "created_at",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "created_by",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "key",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "modified_at",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "modified_by",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "name": "page_size",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "sort",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "sort_by",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "value",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.SampleGetResponseDoc"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/response.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.errorResponse"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Create samples",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "samples"
                ],
                "summary": "Create samples",
                "parameters": [
                    {
                        "description": "request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.SampleCreateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.SampleCreateResponseDoc"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/response.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.errorResponse"
                        }
                    }
                }
            }
        },
        "/samples/{id}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get samples by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "samples"
                ],
                "summary": "Get samples by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id path",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.SampleGetByIDResponseDoc"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/response.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.errorResponse"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Delete samples",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "samples"
                ],
                "summary": "Delete samples",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id path",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.SampleDeleteResponseDoc"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/response.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.errorResponse"
                        }
                    }
                }
            },
            "patch": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Update samples",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "samples"
                ],
                "summary": "Update samples",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id path",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.SampleUpdateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.SampleUpdateResponseDoc"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/response.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.errorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "abstraction.PaginationInfo": {
            "type": "object",
            "required": [
                "page",
                "page_size",
                "sort",
                "sort_by"
            ],
            "properties": {
                "count": {
                    "type": "integer"
                },
                "more_records": {
                    "type": "boolean"
                },
                "page": {
                    "type": "integer"
                },
                "page_size": {
                    "type": "integer"
                },
                "sort": {
                    "type": "string"
                },
                "sort_by": {
                    "type": "string"
                }
            }
        },
        "dto.AuthLoginRequest": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "dto.AuthLoginResponse": {
            "type": "object",
            "required": [
                "email",
                "is_active",
                "name",
                "password",
                "phone"
            ],
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "created_by": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "is_active": {
                    "type": "boolean"
                },
                "modified_at": {
                    "type": "string"
                },
                "modified_by": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "dto.AuthLoginResponseDoc": {
            "type": "object",
            "properties": {
                "body": {
                    "type": "object",
                    "properties": {
                        "data": {
                            "$ref": "#/definitions/dto.AuthLoginResponse"
                        },
                        "meta": {
                            "$ref": "#/definitions/response.Meta"
                        }
                    }
                }
            }
        },
        "dto.AuthRegisterRequest": {
            "type": "object",
            "required": [
                "email",
                "is_active",
                "name",
                "password",
                "phone"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "is_active": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                }
            }
        },
        "dto.AuthRegisterResponse": {
            "type": "object",
            "required": [
                "email",
                "is_active",
                "name",
                "password",
                "phone"
            ],
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "created_by": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "is_active": {
                    "type": "boolean"
                },
                "modified_at": {
                    "type": "string"
                },
                "modified_by": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                }
            }
        },
        "dto.AuthRegisterResponseDoc": {
            "type": "object",
            "properties": {
                "body": {
                    "type": "object",
                    "properties": {
                        "data": {
                            "$ref": "#/definitions/dto.AuthRegisterResponse"
                        },
                        "meta": {
                            "$ref": "#/definitions/response.Meta"
                        }
                    }
                }
            }
        },
        "dto.SampleCreateRequest": {
            "type": "object",
            "required": [
                "key",
                "value"
            ],
            "properties": {
                "key": {
                    "type": "string"
                },
                "value": {
                    "type": "string"
                }
            }
        },
        "dto.SampleCreateResponse": {
            "type": "object",
            "required": [
                "key",
                "value"
            ],
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "created_by": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "key": {
                    "type": "string"
                },
                "modified_at": {
                    "type": "string"
                },
                "modified_by": {
                    "type": "string"
                },
                "sample_childs": {
                    "description": "relations",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.SampleChildEntityModel"
                    }
                },
                "value": {
                    "type": "string"
                }
            }
        },
        "dto.SampleCreateResponseDoc": {
            "type": "object",
            "properties": {
                "body": {
                    "type": "object",
                    "properties": {
                        "data": {
                            "$ref": "#/definitions/dto.SampleCreateResponse"
                        },
                        "meta": {
                            "$ref": "#/definitions/response.Meta"
                        }
                    }
                }
            }
        },
        "dto.SampleDeleteResponse": {
            "type": "object",
            "required": [
                "key",
                "value"
            ],
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "created_by": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "key": {
                    "type": "string"
                },
                "modified_at": {
                    "type": "string"
                },
                "modified_by": {
                    "type": "string"
                },
                "sample_childs": {
                    "description": "relations",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.SampleChildEntityModel"
                    }
                },
                "value": {
                    "type": "string"
                }
            }
        },
        "dto.SampleDeleteResponseDoc": {
            "type": "object",
            "properties": {
                "body": {
                    "type": "object",
                    "properties": {
                        "data": {
                            "$ref": "#/definitions/dto.SampleDeleteResponse"
                        },
                        "meta": {
                            "$ref": "#/definitions/response.Meta"
                        }
                    }
                }
            }
        },
        "dto.SampleGetByIDResponse": {
            "type": "object",
            "required": [
                "key",
                "value"
            ],
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "created_by": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "key": {
                    "type": "string"
                },
                "modified_at": {
                    "type": "string"
                },
                "modified_by": {
                    "type": "string"
                },
                "sample_childs": {
                    "description": "relations",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.SampleChildEntityModel"
                    }
                },
                "value": {
                    "type": "string"
                }
            }
        },
        "dto.SampleGetByIDResponseDoc": {
            "type": "object",
            "properties": {
                "body": {
                    "type": "object",
                    "properties": {
                        "data": {
                            "$ref": "#/definitions/dto.SampleGetByIDResponse"
                        },
                        "meta": {
                            "$ref": "#/definitions/response.Meta"
                        }
                    }
                }
            }
        },
        "dto.SampleGetResponseDoc": {
            "type": "object",
            "properties": {
                "body": {
                    "type": "object",
                    "properties": {
                        "data": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.SampleEntityModel"
                            }
                        },
                        "meta": {
                            "$ref": "#/definitions/response.Meta"
                        }
                    }
                }
            }
        },
        "dto.SampleUpdateRequest": {
            "type": "object",
            "required": [
                "id",
                "key",
                "value"
            ],
            "properties": {
                "id": {
                    "type": "integer"
                },
                "key": {
                    "type": "string"
                },
                "value": {
                    "type": "string"
                }
            }
        },
        "dto.SampleUpdateResponse": {
            "type": "object",
            "required": [
                "key",
                "value"
            ],
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "created_by": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "key": {
                    "type": "string"
                },
                "modified_at": {
                    "type": "string"
                },
                "modified_by": {
                    "type": "string"
                },
                "sample_childs": {
                    "description": "relations",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.SampleChildEntityModel"
                    }
                },
                "value": {
                    "type": "string"
                }
            }
        },
        "dto.SampleUpdateResponseDoc": {
            "type": "object",
            "properties": {
                "body": {
                    "type": "object",
                    "properties": {
                        "data": {
                            "$ref": "#/definitions/dto.SampleUpdateResponse"
                        },
                        "meta": {
                            "$ref": "#/definitions/response.Meta"
                        }
                    }
                }
            }
        },
        "model.SampleChildEntityModel": {
            "type": "object",
            "required": [
                "key",
                "value"
            ],
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "created_by": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "key": {
                    "type": "string"
                },
                "modified_at": {
                    "type": "string"
                },
                "modified_by": {
                    "type": "string"
                },
                "sample_id": {
                    "description": "relations",
                    "type": "integer"
                },
                "value": {
                    "type": "string"
                }
            }
        },
        "model.SampleEntityModel": {
            "type": "object",
            "required": [
                "key",
                "value"
            ],
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "created_by": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "key": {
                    "type": "string"
                },
                "modified_at": {
                    "type": "string"
                },
                "modified_by": {
                    "type": "string"
                },
                "sample_childs": {
                    "description": "relations",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.SampleChildEntityModel"
                    }
                },
                "value": {
                    "type": "string"
                }
            }
        },
        "response.Meta": {
            "type": "object",
            "properties": {
                "info": {
                    "$ref": "#/definitions/abstraction.PaginationInfo"
                },
                "message": {
                    "type": "string",
                    "default": "true"
                },
                "success": {
                    "type": "boolean",
                    "default": true
                }
            }
        },
        "response.errorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "meta": {
                    "$ref": "#/definitions/response.Meta"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "0.0.1",
	Host:        "localhost:3030",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "codeid-boiler",
	Description: "This is a doc for codeid-boiler.",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
