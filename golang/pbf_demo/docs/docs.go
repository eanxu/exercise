// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "xuyi",
            "email": "xuyi@diit.cn"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/mvt/{z}/{x}/{y}": {
            "get": {
                "description": "mvt",
                "tags": [
                    "mvt"
                ],
                "summary": "mvt",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "z",
                        "name": "z",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "x",
                        "name": "x",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "y",
                        "name": "y",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\": \"\",\"msg\":\"success\"}",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "integer"
                            }
                        }
                    },
                    "400": {
                        "description": "{\"code\":400,\"data\":{},\"msg\":\"bind query err/params error\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/orb/{z}/{x}/{y}": {
            "get": {
                "description": "orb??? mvt demo",
                "tags": [
                    "mvt"
                ],
                "summary": "orb",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "z",
                        "name": "z",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "x",
                        "name": "x",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "y",
                        "name": "y",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\": \"\",\"msg\":\"success\"}",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "integer"
                            }
                        }
                    },
                    "400": {
                        "description": "{\"code\":400,\"data\":{},\"msg\":\"bind query err/params error\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/postgis/{z}/{x}/{y}": {
            "get": {
                "description": "postgis mvt",
                "tags": [
                    "mvt"
                ],
                "summary": "postgis",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "z",
                        "name": "z",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "x",
                        "name": "x",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "y",
                        "name": "y",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\": \"\",\"msg\":\"success\"}",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "integer"
                            }
                        }
                    },
                    "400": {
                        "description": "{\"code\":400,\"data\":{},\"msg\":\"bind query err/params error\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/pyramid/{z}/{x}/{y}": {
            "get": {
                "description": "pyramid",
                "tags": [
                    "mvt"
                ],
                "summary": "pyramid",
                "parameters": [
                    {
                        "type": "string",
                        "description": "z",
                        "name": "z",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "x",
                        "name": "x",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "y",
                        "name": "y",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\": \"\",\"msg\":\"success\"}",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "integer"
                            }
                        }
                    },
                    "400": {
                        "description": "{\"code\":400,\"data\":{},\"msg\":\"bind query err/params error\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/tegola/{z}/{x}/{y}": {
            "get": {
                "description": "tegola mvt",
                "tags": [
                    "mvt"
                ],
                "summary": "tegola",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "z",
                        "name": "z",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "x",
                        "name": "x",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "y",
                        "name": "y",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\": \"\",\"msg\":\"success\"}",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "integer"
                            }
                        }
                    },
                    "400": {
                        "description": "{\"code\":400,\"data\":{},\"msg\":\"bind query err/params error\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "",
	BasePath:    "/test",
	Schemes:     []string{},
	Title:       "mvt",
	Description: "mvt",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register("swagger", &s{})
}
