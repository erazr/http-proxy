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
        "/proxy": {
            "post": {
                "description": "Proxy client request",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Proxy"
                ],
                "summary": "Proxy client request",
                "parameters": [
                    {
                        "description": "Request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/Request"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "Request": {
            "type": "object",
            "required": [
                "headers",
                "method",
                "url"
            ],
            "properties": {
                "headers": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    },
                    "example": {
                        "Authorization": "Bearer token",
                        "Content-Type": "application/json"
                    }
                },
                "method": {
                    "type": "string",
                    "example": "GET"
                },
                "url": {
                    "type": "string",
                    "example": "http://google.com"
                }
            }
        },
        "Response": {
            "type": "object",
            "properties": {
                "headers": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    },
                    "example": {
                        "Content-Type": "application/json"
                    }
                },
                "id": {
                    "type": "string",
                    "example": "1"
                },
                "length": {
                    "type": "integer",
                    "example": 100
                },
                "status": {
                    "type": "integer",
                    "example": 200
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
