// Package tripService Code generated by swaggo/swag. DO NOT EDIT
package tripService

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "Vacaition API Support",
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
        "/stop": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "This endpoint can be used to add a stop to an existing trip. Requirements: authenticated",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Stop"
                ],
                "summary": "Create a new stop",
                "operationId": "create-stop",
                "parameters": [
                    {
                        "description": "User Input for creating a new stop",
                        "name": "CreateStopInput",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateStopInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.TripStop"
                        }
                    },
                    "400": {
                        "description": "In case of invalid CreateStop DTO",
                        "schema": {
                            "$ref": "#/definitions/jsonHelper.HTTPError"
                        }
                    },
                    "401": {
                        "description": "In case of unauthenticated request",
                        "schema": {
                            "$ref": "#/definitions/jsonHelper.HTTPError"
                        }
                    },
                    "404": {
                        "description": "In case of unknown trip id",
                        "schema": {
                            "$ref": "#/definitions/jsonHelper.HTTPError"
                        }
                    },
                    "500": {
                        "description": "In case of persistence error",
                        "schema": {
                            "$ref": "#/definitions/jsonHelper.HTTPError"
                        }
                    }
                }
            }
        },
        "/stops": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "This endpoint can be used to multiple stops to an existing trip. Requirements: authenticated",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Stop"
                ],
                "summary": "Create multiple stops",
                "operationId": "create-stops",
                "parameters": [
                    {
                        "description": "User Input for creating new stops",
                        "name": "CreateStopsInput",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateStopsInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.TripStop"
                            }
                        }
                    },
                    "400": {
                        "description": "In case of invalid CreateStops DTO",
                        "schema": {
                            "$ref": "#/definitions/jsonHelper.HTTPError"
                        }
                    },
                    "401": {
                        "description": "In case of unauthenticated request",
                        "schema": {
                            "$ref": "#/definitions/jsonHelper.HTTPError"
                        }
                    },
                    "404": {
                        "description": "In case of unknown trip id",
                        "schema": {
                            "$ref": "#/definitions/jsonHelper.HTTPError"
                        }
                    },
                    "500": {
                        "description": "In case of persistence error",
                        "schema": {
                            "$ref": "#/definitions/jsonHelper.HTTPError"
                        }
                    }
                }
            }
        },
        "/trip": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "This endpoint can be used to create a trip. Requirements: authenticated",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Trip"
                ],
                "summary": "Create a new trip",
                "operationId": "create-trip",
                "parameters": [
                    {
                        "description": "User Input creating a new trip",
                        "name": "createTripInput",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateTripQueryInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Trip"
                        }
                    },
                    "400": {
                        "description": "In case of invalid createTrip DTO",
                        "schema": {
                            "$ref": "#/definitions/jsonHelper.HTTPError"
                        }
                    },
                    "401": {
                        "description": "In case of unauthenticated request",
                        "schema": {
                            "$ref": "#/definitions/jsonHelper.HTTPError"
                        }
                    },
                    "500": {
                        "description": "In case of persistence error",
                        "schema": {
                            "$ref": "#/definitions/jsonHelper.HTTPError"
                        }
                    }
                }
            }
        },
        "/trip/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "This endpoint is used to query a trip. Requirements: authenticated and requested trip is assigned to user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Trip"
                ],
                "summary": "Get trip",
                "operationId": "get-trip",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Trip ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Trip"
                        }
                    },
                    "400": {
                        "description": "In case of missing path param",
                        "schema": {
                            "$ref": "#/definitions/jsonHelper.HTTPError"
                        }
                    },
                    "401": {
                        "description": "In case of unauthenticated request",
                        "schema": {
                            "$ref": "#/definitions/jsonHelper.HTTPError"
                        }
                    },
                    "404": {
                        "description": "In case non existing trip for given user",
                        "schema": {
                            "$ref": "#/definitions/jsonHelper.HTTPError"
                        }
                    },
                    "500": {
                        "description": "In case of\tpersistence error",
                        "schema": {
                            "$ref": "#/definitions/jsonHelper.HTTPError"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "This endpoint can be used update a trip. This action deletes existing stops and initiates a new calculation. Requirements: authenticated",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Trip"
                ],
                "summary": "Update a trip",
                "operationId": "update-trip",
                "parameters": [
                    {
                        "description": "User Input for updating a trip",
                        "name": "updateTripInput",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UpdateTripQueryInput"
                        }
                    },
                    {
                        "type": "integer",
                        "description": "Trip ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Trip"
                        }
                    },
                    "400": {
                        "description": "In\tcase of invalid updateTrip DTO or missing path param",
                        "schema": {
                            "$ref": "#/definitions/jsonHelper.HTTPError"
                        }
                    },
                    "401": {
                        "description": "In\tcase of unauthenticated\trequest",
                        "schema": {
                            "$ref": "#/definitions/jsonHelper.HTTPError"
                        }
                    },
                    "500": {
                        "description": "In\tcase of persistence\t\terror",
                        "schema": {
                            "$ref": "#/definitions/jsonHelper.HTTPError"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "This endpoint can be used delete a trip including its stops. Requirements: authenticated",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Trip"
                ],
                "summary": "Delete a trip",
                "operationId": "delete-trip",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Trip ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Trip"
                        }
                    },
                    "400": {
                        "description": "In case of missing path param",
                        "schema": {
                            "$ref": "#/definitions/jsonHelper.HTTPError"
                        }
                    },
                    "401": {
                        "description": "In case of unauthenticated\trequest",
                        "schema": {
                            "$ref": "#/definitions/jsonHelper.HTTPError"
                        }
                    },
                    "500": {
                        "description": "In case of persistence\t\terror",
                        "schema": {
                            "$ref": "#/definitions/jsonHelper.HTTPError"
                        }
                    }
                }
            }
        },
        "/trips": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "This endpoint is used to query all trips of a given user. Requirements: authenticated and requested trip is assigned to user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Trip"
                ],
                "summary": "Get all trips",
                "operationId": "get-trips",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Trip"
                            }
                        }
                    },
                    "401": {
                        "description": "In\tcase of\tunauthenticated\trequest",
                        "schema": {
                            "$ref": "#/definitions/jsonHelper.HTTPError"
                        }
                    },
                    "404": {
                        "description": "In\tcase non existing trip for given user",
                        "schema": {
                            "$ref": "#/definitions/jsonHelper.HTTPError"
                        }
                    },
                    "500": {
                        "description": "In\tcase of\tpersistence error",
                        "schema": {
                            "$ref": "#/definitions/jsonHelper.HTTPError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "jsonHelper.HTTPError": {
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
        "models.CreateStopInput": {
            "type": "object",
            "required": [
                "stop"
            ],
            "properties": {
                "stop": {
                    "$ref": "#/definitions/models.TripStopInput"
                },
                "trip_id": {
                    "type": "integer",
                    "example": 1
                }
            }
        },
        "models.CreateStopsInput": {
            "type": "object",
            "required": [
                "stops"
            ],
            "properties": {
                "stops": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.TripStopInput"
                    }
                },
                "trip_id": {
                    "type": "integer",
                    "example": 1
                }
            }
        },
        "models.CreateTripQueryInput": {
            "type": "object",
            "required": [
                "country",
                "duration",
                "focus",
                "maximum_distance",
                "secrets"
            ],
            "properties": {
                "country": {
                    "type": "string",
                    "example": "Germany"
                },
                "duration": {
                    "type": "integer",
                    "example": 10
                },
                "focus": {
                    "type": "string",
                    "enum": [
                        "Cities",
                        "Nature",
                        "Mixed"
                    ],
                    "example": "Mixed"
                },
                "maximum_distance": {
                    "type": "integer",
                    "example": 100
                },
                "secrets": {
                    "type": "boolean",
                    "example": true
                }
            }
        },
        "models.StopHighlight": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string",
                    "example": "2023-12-01T12:37:59.008583Z"
                },
                "description": {
                    "type": "string",
                    "example": "The landmark of Berlin"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "latitude": {
                    "type": "number",
                    "example": 13.404954
                },
                "longitude": {
                    "type": "number",
                    "example": 52.520008
                },
                "name": {
                    "type": "string",
                    "example": "Brandenburger Tor"
                },
                "stop_id": {
                    "type": "integer"
                },
                "updated_at_at": {
                    "type": "string",
                    "example": "2023-12-01T12:37:59.008583Z"
                }
            }
        },
        "models.StopHighlightInput": {
            "type": "object",
            "required": [
                "description",
                "name"
            ],
            "properties": {
                "description": {
                    "type": "string",
                    "example": "The landmark of Berlin"
                },
                "latitude": {
                    "type": "number",
                    "example": 13.404954
                },
                "longitude": {
                    "type": "number",
                    "example": 52.520008
                },
                "name": {
                    "type": "string",
                    "example": "Brandenburger Tor"
                }
            }
        },
        "models.Trip": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string",
                    "example": "2023-12-01T12:37:59.008583Z"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "query": {
                    "$ref": "#/definitions/models.TripQuery"
                },
                "stops": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.TripStop"
                    }
                },
                "updated_at_at": {
                    "type": "string",
                    "example": "2023-12-01T12:37:59.008583Z"
                },
                "user_id": {
                    "type": "string",
                    "example": "1"
                }
            }
        },
        "models.TripQuery": {
            "type": "object",
            "properties": {
                "country": {
                    "type": "string",
                    "example": "Germany"
                },
                "created_at": {
                    "type": "string",
                    "example": "2023-12-01T12:37:59.008583Z"
                },
                "duration": {
                    "type": "integer",
                    "example": 10
                },
                "focus": {
                    "type": "string",
                    "enum": [
                        "Cities",
                        "Nature",
                        "Mixed"
                    ],
                    "example": "Mixed"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "maximum_distance": {
                    "type": "integer",
                    "example": 1000
                },
                "secrets": {
                    "type": "boolean"
                },
                "updated_at_at": {
                    "type": "string",
                    "example": "2023-12-01T12:37:59.008583Z"
                }
            }
        },
        "models.TripStop": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string",
                    "example": "2023-12-01T12:37:59.008583Z"
                },
                "days": {
                    "type": "integer"
                },
                "highlights": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.StopHighlight"
                    }
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "latitude": {
                    "type": "number",
                    "example": 13.404954
                },
                "longitude": {
                    "type": "number",
                    "example": 52.520008
                },
                "stopName": {
                    "type": "string",
                    "example": "Berlin"
                },
                "trip_id": {
                    "type": "integer"
                },
                "updated_at_at": {
                    "type": "string",
                    "example": "2023-12-01T12:37:59.008583Z"
                }
            }
        },
        "models.TripStopInput": {
            "type": "object",
            "required": [
                "highlights"
            ],
            "properties": {
                "days": {
                    "type": "integer",
                    "example": 10
                },
                "highlights": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.StopHighlightInput"
                    }
                },
                "latitude": {
                    "type": "number",
                    "example": 13.404954
                },
                "longitude": {
                    "type": "number",
                    "example": 52.520008
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "models.UpdateTripQueryInput": {
            "type": "object",
            "required": [
                "country",
                "duration",
                "focus",
                "id",
                "maximum_distance",
                "secrets"
            ],
            "properties": {
                "country": {
                    "type": "string",
                    "example": "Germany"
                },
                "duration": {
                    "type": "integer",
                    "example": 10
                },
                "focus": {
                    "type": "string",
                    "enum": [
                        "Cities",
                        "Nature",
                        "Mixed"
                    ],
                    "example": "Mixed"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "maximum_distance": {
                    "type": "integer",
                    "example": 1000
                },
                "secrets": {
                    "type": "boolean",
                    "example": true
                }
            }
        }
    },
    "securityDefinitions": {
        "OAuth2Application": {
            "type": "oauth2",
            "flow": "implicit",
            "authorizationUrl": "http://localhost:8080/userservice/api/oauth?provider=google",
            "scopes": {
                "admin": "\t\t\t\t\t\t\tGrants read and write access to administrative information",
                "write": "\t\t\t\t\t\t\tGrants write access"
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/tripservice/api",
	Schemes:          []string{},
	Title:            "Trip API",
	Description:      "This server is used for creating new trips",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
