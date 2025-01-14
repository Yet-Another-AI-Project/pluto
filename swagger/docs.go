// Package swagger Code generated by swaggo/swag. DO NOT EDIT
package swagger

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "API Support"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/token/access/verify": {
            "post": {
                "description": "Verify access token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Token"
                ],
                "summary": "Verify access token",
                "parameters": [
                    {
                        "description": "Verify access token",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.VerifyAccessToken"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Reponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "body": {
                                            "$ref": "#/definitions/jwt.AccessPayload"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/v1/token/publickey": {
            "get": {
                "description": "Get public key",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Token"
                ],
                "summary": "Get public key",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Reponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "body": {
                                            "$ref": "#/definitions/v1.PublicKeyResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/v1/token/refresh": {
            "post": {
                "description": "Refresh access token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Token"
                ],
                "summary": "Refresh access token",
                "parameters": [
                    {
                        "description": "Refresh access token",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.RefreshAccessToken"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Reponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "body": {
                                            "$ref": "#/definitions/manage.GrantResult"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/v1/user/info": {
            "get": {
                "description": "Get user info",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get user info",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer \u003ctoken\u003e",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Reponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "body": {
                                            "$ref": "#/definitions/modelexts.UserInfo"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "put": {
                "description": "Update user info",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Update user info",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer \u003ctoken\u003e",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "update user info",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.UpdateUserInfo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Reponse"
                        }
                    }
                }
            }
        },
        "/v1/user/login/google/web": {
            "post": {
                "description": "google login web",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "google login web",
                "parameters": [
                    {
                        "description": "Google login web request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.GoogleWebLogin"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Reponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "body": {
                                            "$ref": "#/definitions/manage.GrantResult"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/v1/user/login/wechat/miniprogram": {
            "post": {
                "description": "Wehcat miniprogram login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Wehcat miniprogram login",
                "parameters": [
                    {
                        "description": "Wechat miniprogram login request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.WechatMiniprogramLogin"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Reponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "body": {
                                            "$ref": "#/definitions/manage.GrantResult"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "jwt.AccessPayload": {
            "type": "object",
            "properties": {
                "exp": {
                    "type": "integer"
                },
                "iat": {
                    "type": "integer"
                },
                "iss": {
                    "type": "string"
                },
                "scopes": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "sub": {
                    "type": "integer"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "manage.GrantResult": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "access_token_expire_at": {
                    "type": "integer"
                },
                "refresh_token": {
                    "type": "string"
                },
                "refresh_token_expire_at": {
                    "type": "integer"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "modelexts.Binding": {
            "type": "object",
            "properties": {
                "login_type": {
                    "type": "string"
                },
                "mail": {
                    "type": "string"
                }
            }
        },
        "modelexts.UserInfo": {
            "type": "object",
            "properties": {
                "app_id": {
                    "type": "string"
                },
                "avatar": {
                    "type": "string"
                },
                "bindings": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/modelexts.Binding"
                    }
                },
                "created_at": {
                    "type": "integer"
                },
                "is_password_set": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "sub": {
                    "type": "integer"
                },
                "update_at": {
                    "type": "integer"
                },
                "user_id": {
                    "type": "string"
                },
                "user_updated": {
                    "type": "boolean"
                },
                "verified": {
                    "type": "boolean"
                }
            }
        },
        "pluto_error.PlutoError": {
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
        "request.GoogleWebLogin": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "app_id": {
                    "type": "string"
                },
                "device_id": {
                    "type": "string"
                }
            }
        },
        "request.RefreshAccessToken": {
            "type": "object",
            "properties": {
                "app_id": {
                    "type": "string"
                },
                "refresh_token": {
                    "type": "string"
                },
                "scopes": {
                    "type": "string"
                }
            }
        },
        "request.UpdateUserInfo": {
            "type": "object",
            "properties": {
                "avatar": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "request.VerifyAccessToken": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "request.WechatMiniprogramLogin": {
            "type": "object",
            "properties": {
                "app_id": {
                    "type": "string"
                },
                "code": {
                    "type": "string"
                }
            }
        },
        "response.Reponse": {
            "type": "object",
            "properties": {
                "body": {},
                "error": {
                    "$ref": "#/definitions/pluto_error.PlutoError"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "v1.PublicKeyResponse": {
            "type": "object",
            "properties": {
                "public_key": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.0.1",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Pluto API",
	Description:      "Client-side API intended for general users.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
