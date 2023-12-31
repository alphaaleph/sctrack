// Package docs Code generated by swaggo/swag at 2023-09-30 23:49:26.577958005 -0600 MDT m=+0.114655462. DO NOT EDIT
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
        "/api/action/all": {
            "get": {
                "description": "Get all action entries",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "actions"
                ],
                "summary": "Get all actions",
                "responses": {}
            }
        },
        "/api/carrier": {
            "post": {
                "description": "Add a new carrier",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "carriers"
                ],
                "summary": "Add carrier",
                "parameters": [
                    {
                        "description": "The Carrier Inout",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Carrier"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/api/carrier/all": {
            "get": {
                "description": "Get the information for all carriers",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "carriers"
                ],
                "summary": "Get all carriers",
                "responses": {}
            }
        },
        "/api/carrier/{id}": {
            "get": {
                "description": "Get carrier's data details by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "carriers"
                ],
                "summary": "Get carrier's data",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            },
            "delete": {
                "description": "Delete a carrier",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "carriers"
                ],
                "summary": "Delete carrier",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/api/journal/all": {
            "get": {
                "description": "Get all entries from the journal",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "journal"
                ],
                "summary": "Get journals",
                "responses": {}
            }
        },
        "/api/journal/{uuid}": {
            "get": {
                "description": "Get a journal entry that matches the uuid",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "journal"
                ],
                "summary": "Get a journal",
                "parameters": [
                    {
                        "type": "string",
                        "description": "UUID",
                        "name": "uuid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            },
            "delete": {
                "description": "Delete an entry in the journal by UUID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "journal"
                ],
                "summary": "Delete journal",
                "parameters": [
                    {
                        "type": "string",
                        "description": "UUID",
                        "name": "uuid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/api/journal/{uuid}/{index}": {
            "delete": {
                "description": "Delete an entry in the journal by UUID and Index",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "journal"
                ],
                "summary": "Delete journal",
                "parameters": [
                    {
                        "type": "string",
                        "description": "UUID",
                        "name": "uuid",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "index",
                        "name": "index",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/api/todos": {
            "post": {
                "description": "Add a todos entry for a carrier",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "todos"
                ],
                "summary": "Add todos",
                "parameters": [
                    {
                        "description": "New Todos",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.TodosAdd"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/api/todos/all": {
            "get": {
                "description": "Get all the entries in the todos list",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "todos"
                ],
                "summary": "Get all todos",
                "responses": {}
            }
        },
        "/api/todos/carrier/{carrier_id}": {
            "get": {
                "description": "Get all the entries in the todos list that match the carrier id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "todos"
                ],
                "summary": "Get todos by carrier id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "carrier_id",
                        "name": "carrier_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            },
            "delete": {
                "description": "Delete entries in the todos list that match an carrier_id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "todos"
                ],
                "summary": "Delete todos by carrier_id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "carrier_id",
                        "name": "carrier_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/api/todos/{uuid}": {
            "get": {
                "description": "Get all the entries in the todos list that match the uuid",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "todos"
                ],
                "summary": "Get todos by carrier uuid",
                "parameters": [
                    {
                        "type": "string",
                        "description": "UUID",
                        "name": "uuid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            },
            "delete": {
                "description": "Delete an entry in the todos list that match an uuid",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "todos"
                ],
                "summary": "Delete todos by uuid",
                "parameters": [
                    {
                        "type": "string",
                        "description": "UUID",
                        "name": "uuid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/api/todos/{uuid}/completed": {
            "patch": {
                "description": "Update a todos completed flag for a carrier",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "todos"
                ],
                "summary": "Update the todos completed",
                "parameters": [
                    {
                        "type": "string",
                        "description": "UUID",
                        "name": "uuid",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Update Completed",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.TodosStatus"
                        }
                    }
                ],
                "responses": {}
            }
        }
    },
    "definitions": {
        "models.Carrier": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "telephone": {
                    "type": "string"
                }
            }
        },
        "models.TodosAdd": {
            "type": "object",
            "properties": {
                "action": {
                    "type": "string"
                },
                "carrierID": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                }
            }
        },
        "models.TodosStatus": {
            "type": "object",
            "properties": {
                "completed": {
                    "type": "boolean"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{"https", "http"},
	Title:            "sctrack",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
