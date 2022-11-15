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
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/auth/users/signing": {
            "post": {
                "description": "User Login.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Login"
                ],
                "summary": "User Login.",
                "parameters": [
                    {
                        "description": "The body to login",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.Receive_Login_Data"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Auth_Token"
                        }
                    }
                }
            }
        },
        "/v1/proposal/create/{token}": {
            "post": {
                "description": "Create UsProposaler.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Proposal"
                ],
                "summary": "Create Proposal.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Token",
                        "name": "token",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "The body to create a proposal",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.Send_Proposal_Data"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Receive_Proposal_Data"
                        }
                    }
                }
            }
        },
        "/v1/users": {
            "get": {
                "description": "Get all users list.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Show all users.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Users"
                        }
                    }
                }
            }
        },
        "/v1/users/create": {
            "post": {
                "description": "Create User.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Create User.",
                "parameters": [
                    {
                        "description": "The body to create a user",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.Send_User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Users"
                        }
                    }
                }
            }
        },
        "/v1/users/delete/{id}": {
            "delete": {
                "description": "Delete User Data.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Delete User.",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Usuário deletado com sucesso!"
                    }
                }
            }
        },
        "/v1/users/update/{id}/{token}": {
            "put": {
                "description": "Update User Data.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Update User.",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Token",
                        "name": "token",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "The body to update a user",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.Send_User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Send_User"
                        }
                    }
                }
            }
        },
        "/v1/users/{id}": {
            "get": {
                "description": "Get user.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Show an user.",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Users"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "entity.Auth_Token": {
            "type": "object",
            "properties": {
                "expires": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                },
                "userID": {
                    "type": "integer"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "entity.Receive_Login_Data": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string",
                    "example": "aB@123456"
                },
                "username": {
                    "type": "string",
                    "example": "teste@gmail.com"
                }
            }
        },
        "entity.Receive_Proposal_Data": {
            "type": "object",
            "properties": {
                "proposalId": {
                    "type": "integer"
                },
                "proposalattachments": {
                    "type": "integer",
                    "example": 1
                },
                "proposaldescription": {
                    "type": "string",
                    "example": "Proposal Description"
                },
                "proposalpictures": {
                    "type": "integer",
                    "example": 2
                },
                "proposalstatus": {
                    "type": "boolean",
                    "example": true
                },
                "proposaltitle": {
                    "type": "string",
                    "example": "Proposal Title"
                },
                "proposaluserid": {
                    "type": "integer",
                    "example": 2
                },
                "token": {
                    "type": "string"
                },
                "userid": {
                    "type": "integer"
                }
            }
        },
        "entity.Send_Proposal_Data": {
            "type": "object",
            "properties": {
                "proposalattachments": {
                    "type": "integer",
                    "example": 1
                },
                "proposaldescription": {
                    "type": "string",
                    "example": "Proposal Description"
                },
                "proposalpictures": {
                    "type": "integer",
                    "example": 2
                },
                "proposalstatus": {
                    "type": "boolean",
                    "example": true
                },
                "proposaltitle": {
                    "type": "string",
                    "example": "Proposal Title"
                },
                "proposaluserid": {
                    "type": "integer",
                    "example": 2
                }
            }
        },
        "entity.Send_User": {
            "type": "object",
            "required": [
                "email",
                "name",
                "newsletter",
                "password",
                "picture"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "example": "test@test.com"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string",
                    "example": "User Name"
                },
                "newsletter": {
                    "type": "boolean",
                    "example": true
                },
                "password": {
                    "type": "string",
                    "example": "aB@123456"
                },
                "picture": {
                    "type": "integer",
                    "example": 1
                }
            }
        },
        "entity.Users": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "test@test.com"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "name": {
                    "type": "string",
                    "example": "User Name"
                },
                "newsletter": {
                    "type": "boolean",
                    "example": true
                },
                "password": {
                    "type": "string",
                    "example": "b3f8b6283fce62d85c5b6334c8ee9a611aed144c3d93d11ef2759f6baabdc3b0"
                },
                "picture": {
                    "type": "integer",
                    "example": 1
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:7500",
	BasePath:         "/",
	Schemes:          []string{"http"},
	Title:            "Echo Propostas Populares API",
	Description:      "This is a sample crud api.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
