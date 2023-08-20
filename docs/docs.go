// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "email": "info@mangi.org"
        },
        "license": {
            "name": "MIT License",
            "url": "https://opensource.org/license/mit/"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/jwt/check": {
            "post": {
                "description": "Checks the validity of a token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "JWT"
                ],
                "summary": "Check token",
                "parameters": [
                    {
                        "description": "Token object",
                        "name": "requestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/jwt.CheckTokenJSON"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/schema.Response"
                        }
                    }
                }
            }
        },
        "/api/jwt/delete": {
            "delete": {
                "description": "Deletes a user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "JWT"
                ],
                "summary": "Delete user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/schema.Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/schema.Response"
                        }
                    }
                }
            }
        },
        "/api/jwt/login": {
            "post": {
                "description": "Logs in a user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "JWT"
                ],
                "summary": "User login",
                "parameters": [
                    {
                        "description": "User object",
                        "name": "requestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/jwt.UserJSON"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/jwt.RefreshAndAccessTokens"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/schema.Response"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/schema.Response"
                        }
                    }
                }
            }
        },
        "/api/jwt/logout": {
            "get": {
                "description": "Logs out a user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "JWT"
                ],
                "summary": "User logout",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/jwt.RefreshAndAccessTokens"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/schema.Response"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/schema.Response"
                        }
                    }
                }
            }
        },
        "/api/jwt/refresh": {
            "get": {
                "description": "Refreshes the JWT token using the refresh token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "JWT"
                ],
                "summary": "Refresh JWT token",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer {refresh_token}",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/jwt.AccessToken"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/schema.Response"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/schema.Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/schema.Response"
                        }
                    }
                }
            }
        },
        "/api/jwt/register": {
            "post": {
                "description": "Registers a new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "JWT"
                ],
                "summary": "Register user",
                "parameters": [
                    {
                        "description": "User object",
                        "name": "requestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/jwt.UserJSON"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/jwt.RefreshAndAccessTokens"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/schema.Response"
                        }
                    }
                }
            }
        },
        "/api/urls/": {
            "post": {
                "description": "Создает URL с предоставленным исходным URL.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Параметры URL"
                ],
                "summary": "Создать URL",
                "parameters": [
                    {
                        "description": "Тело запроса",
                        "name": "c",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/urls.CreateURLBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/urls.URLResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/schema.Response"
                        }
                    }
                }
            }
        },
        "/api/urls/{shorturl}": {
            "get": {
                "description": "Обрабатывает HTTP-запрос для получения параметров URL.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Параметры URL"
                ],
                "summary": "Получить параметры URL",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Короткий URL",
                        "name": "shorturl",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/urls.URLResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/schema.Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/schema.Response"
                        }
                    }
                }
            },
            "delete": {
                "description": "Удалить URL с предоставленным коротким URL.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Параметры URL"
                ],
                "summary": "Удалить URL",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Короткий URL",
                        "name": "shorturl",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/urls.URLResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/schema.Response"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/schema.Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/schema.Response"
                        }
                    }
                }
            },
            "patch": {
                "description": "Обновить URL с предоставленным коротким и существующим URL.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Параметры URL"
                ],
                "summary": "Обновить URL",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Короткий URL",
                        "name": "shorturl",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Original URL",
                        "name": "bodyJson",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/urls.ShortURLBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/urls.URLResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/schema.Response"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/schema.Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/schema.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "jwt.AccessToken": {
            "type": "object",
            "properties": {
                "access": {
                    "type": "string"
                }
            }
        },
        "jwt.CheckTokenJSON": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "jwt.RefreshAndAccessTokens": {
            "type": "object",
            "properties": {
                "access": {
                    "type": "string"
                },
                "refresh": {
                    "type": "string"
                }
            }
        },
        "jwt.UserJSON": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "schema.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "urls.CreateURLBody": {
            "type": "object",
            "properties": {
                "original_url": {
                    "type": "string"
                }
            }
        },
        "urls.ShortURLBody": {
            "type": "object",
            "properties": {
                "original_url": {
                    "type": "string"
                }
            }
        },
        "urls.URLResponse": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "original_url": {
                    "type": "string"
                },
                "short_url": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Fiber Example API",
	Description:      "This is a sample swagger for Fiber",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
