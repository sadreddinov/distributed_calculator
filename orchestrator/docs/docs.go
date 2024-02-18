// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Фарид",
            "url": "https://t.me/M00nfI0wer",
            "email": "faridsadreddinov@yandex.ru"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/computing_resources/": {
            "get": {
                "description": "Get computing resources",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "computing resources"
                ],
                "summary": "Get computing resources",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.ComputingResource"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
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
        },
        "/expression/": {
            "patch": {
                "description": "Update operation time",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "operations"
                ],
                "summary": "Update operation time",
                "parameters": [
                    {
                        "description": "Time of operations",
                        "name": "operations",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Operation"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Operation"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/expressions/": {
            "get": {
                "description": "Get expressions",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "expressions"
                ],
                "summary": "Get expressions",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Page num",
                        "name": "page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Num of record per page",
                        "name": "recordPerPage",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.ExpressionToRead"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Add new expression",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "expressions"
                ],
                "summary": "Add expression",
                "parameters": [
                    {
                        "description": "Expression info",
                        "name": "expression",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Expression"
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
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/expressions/{id}": {
            "get": {
                "description": "Get expression by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "expressions"
                ],
                "summary": "Get expression",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Expression ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ExpressionToRead"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/operations/": {
            "get": {
                "description": "Get time of operations in seconds",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "operations"
                ],
                "summary": "Get time of operations",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Operation"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.ComputingResource": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "last_ping_at": {
                    "type": "string"
                },
                "work_state": {
                    "type": "string"
                }
            }
        },
        "models.Expression": {
            "type": "object"
        },
        "models.ExpressionToRead": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "expression": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "result": {
                    "type": "string"
                },
                "solved_at": {
                    "type": "string"
                },
                "work_state": {
                    "type": "string"
                }
            }
        },
        "models.Operation": {
            "type": "object",
            "properties": {
                "divide": {
                    "type": "string"
                },
                "minus": {
                    "type": "string"
                },
                "multiply": {
                    "type": "string"
                },
                "plus": {
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
	Title:            "Distributed Calculator",
	Description:      "Это простой распределенный калькулятор.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
