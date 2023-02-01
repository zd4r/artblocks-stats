// Code generated by swaggo/swag. DO NOT EDIT
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
        "/collections/{id}/holders": {
            "get": {
                "description": "Show collection holders with scores from artacle API",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Show collection holders",
                "operationId": "holders",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Collection ID from Artacle",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.collectionHoldersResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.errResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.errResp"
                        }
                    }
                }
            }
        },
        "/collections/{id}/stats": {
            "get": {
                "description": "Show collection holders distribution based on artacle scores",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Show collection stats",
                "operationId": "stats",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Collection ID from Artacle",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.collectionStatsResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.errResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.errResp"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "entity.Holder": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "commitment_score": {
                    "type": "number"
                },
                "portfolio_score": {
                    "type": "number"
                },
                "tokens_amount": {
                    "type": "integer"
                },
                "trading_score": {
                    "type": "number"
                }
            }
        },
        "entity.HoldersDistribution": {
            "type": "object",
            "properties": {
                "by_commitment_score": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "integer"
                    }
                },
                "by_portfolio_score": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "integer"
                    }
                },
                "by_trading_score": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "integer"
                    }
                }
            }
        },
        "v1.collectionHoldersResponse": {
            "type": "object",
            "properties": {
                "holders": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Holder"
                    }
                }
            }
        },
        "v1.collectionStatsResponse": {
            "type": "object",
            "properties": {
                "collection": {
                    "type": "object",
                    "properties": {
                        "holders_count": {
                            "type": "integer"
                        },
                        "holders_distribution": {
                            "$ref": "#/definitions/entity.HoldersDistribution"
                        },
                        "id": {
                            "type": "integer"
                        }
                    }
                }
            }
        },
        "v1.errResp": {
            "type": "object",
            "properties": {
                "message": {
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
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "Artblocks stats API",
	Description:      "Collection service",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
