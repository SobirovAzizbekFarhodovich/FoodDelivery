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
        "/user/change-password": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Change user password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Change Password",
                "parameters": [
                    {
                        "description": "Change Password",
                        "name": "ChangePass",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.changePass"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Password Changed Successfully",
                        "schema": {
                            "type": "body"
                        }
                    },
                    "400": {
                        "description": "Error while changing password",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/delete-profil": {
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Delete an existing Profil",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Delete Profil",
                "responses": {
                    "200": {
                        "description": "Delete Successful",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Error while deleting user",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/delete/{id}": {
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Delete an existing user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Delete User",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Delete Successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Error while deleting user",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/forgot-password": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Initiate forgot password process",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Forgot Password",
                "parameters": [
                    {
                        "description": "Forgot Password",
                        "name": "ForgotPass",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/genprotos.ForgotPasswordRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Forgot Password Initiated Successfully",
                        "schema": {
                            "type": "body"
                        }
                    },
                    "400": {
                        "description": "Error while initiating forgot password",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/get-by-id/{id}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get a user by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get User By ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Get By ID Successful",
                        "schema": {
                            "$ref": "#/definitions/genprotos.GetByIdUserResponse"
                        }
                    },
                    "400": {
                        "description": "Error while retrieving user",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "User Not Found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/get-profil": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get a user Profil",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get User Profil",
                "responses": {
                    "200": {
                        "description": "Get Profil Successful",
                        "schema": {
                            "$ref": "#/definitions/genprotos.GetByIdUserResponse"
                        }
                    },
                    "400": {
                        "description": "Error while retrieving user",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "description": "Login a user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Login User",
                "parameters": [
                    {
                        "description": "Create",
                        "name": "Create",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/genprotos.LoginUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Login Successful",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Error while logging in",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/register": {
            "post": {
                "description": "Register a new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Register User",
                "parameters": [
                    {
                        "description": "Create",
                        "name": "Create",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/genprotos.RegisterUserRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Create Successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Error while creating user",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/reset-password": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Reset user password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Reset Password",
                "parameters": [
                    {
                        "description": "Reset Password",
                        "name": "ResetPass",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.resetPass"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Password Reset Successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Error while resetting password",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/update-profil": {
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Update an existing user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Update Profil",
                "parameters": [
                    {
                        "description": "Update",
                        "name": "Update",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/genprotos.UpdateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Update Successful",
                        "schema": {
                            "$ref": "#/definitions/genprotos.UpdateUserResponse"
                        }
                    },
                    "400": {
                        "description": "Error while updating user",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/update/{id}": {
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Update an existing user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Update User",
                "parameters": [
                    {
                        "description": "Update",
                        "name": "Update",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/genprotos.UpdateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Update Successful",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Error while updating user",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "genprotos.ForgotPasswordRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                }
            }
        },
        "genprotos.GetByIdUserResponse": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "password_hash": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                }
            }
        },
        "genprotos.LoginUserRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password_hash": {
                    "type": "string"
                }
            }
        },
        "genprotos.RegisterUserRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password_hash": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                }
            }
        },
        "genprotos.UpdateUserRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "password_hash": {
                    "type": "string"
                }
            }
        },
        "genprotos.UpdateUserResponse": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "password_hash": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                }
            }
        },
        "handler.changePass": {
            "type": "object",
            "properties": {
                "currentPassword": {
                    "type": "string"
                },
                "newPassword": {
                    "type": "string"
                }
            }
        },
        "handler.resetPass": {
            "type": "object",
            "properties": {
                "newPassword": {
                    "type": "string"
                },
                "resetToken": {
                    "type": "string"
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

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "auth service API",
	Description:      "auth service API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
