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
        "/segments": {
            "post": {
                "description": "Create segment with given slug",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create segment",
                "parameters": [
                    {
                        "description": "Segment",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Segment"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/resources.SegmentResource"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/resources.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/segments/{slug}": {
            "delete": {
                "description": "Delete segment by given slug",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Delete segment by slug",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Slug",
                        "name": "slug",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/resources.SegmentResource"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/resources.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/users": {
            "post": {
                "description": "Create user with segments that should be added to User and segments that should be deleted from User",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create user",
                "parameters": [
                    {
                        "description": "User",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/resources.UserResource"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/resources.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/users/{id}/segments": {
            "get": {
                "description": "Get segments that were added to user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get user segments",
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
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/resources.UserSegmentResource"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/resources.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Segment": {
            "type": "object",
            "required": [
                "slug"
            ],
            "properties": {
                "id": {
                    "type": "integer"
                },
                "slug": {
                    "type": "string"
                }
            }
        },
        "models.User": {
            "type": "object",
            "required": [
                "email"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "segmentsToAdd": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "segmentsToDelete": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "resources.ErrorResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "resources.SegmentResource": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "resources.UserResource": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "resources.UserSegmentResource": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "segments": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Segment"
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Segmentation Service API",
	Description:      "This is a Segmentation Service that can create and delete segments, assign them to users and get active segments for user",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
