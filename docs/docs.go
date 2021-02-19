// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/student/card/balance": {
            "post": {
                "description": "校园卡余额查询",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "校园卡"
                ],
                "summary": "校园卡余额查询",
                "parameters": [
                    {
                        "description": "用户名",
                        "name": "username",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "密码",
                        "name": "password",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code\":200,\"data\":123.4,\"msg\":\"OK\"}",
                        "schema": {
                            "type": "json"
                        }
                    },
                    "400": {
                        "description": "code\":400,\"data\":null,\"msg\":\"\"}",
                        "schema": {
                            "type": "json"
                        }
                    }
                }
            }
        },
        "/student/card/history": {
            "post": {
                "description": "校园卡历史查询",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "校园卡"
                ],
                "summary": "校园卡历史查询",
                "parameters": [
                    {
                        "description": "用户名",
                        "name": "username",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "密码",
                        "name": "password",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "年份",
                        "name": "year",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "月份",
                        "name": "month",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code\":200,\"data\":[{...}],\"msg\":\"OK\"}",
                        "schema": {
                            "type": "json"
                        }
                    },
                    "400": {
                        "description": "code\":400,\"data\":null,\"msg\":\"\"}",
                        "schema": {
                            "type": "json"
                        }
                    }
                }
            }
        },
        "/student/card/today": {
            "post": {
                "description": "校园卡今日消费查询",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "校园卡"
                ],
                "summary": "校园卡今日消费查询",
                "parameters": [
                    {
                        "description": "用户名",
                        "name": "username",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "密码",
                        "name": "password",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code\":200,\"data\":[{}],\"msg\":\"OK\"}",
                        "schema": {
                            "type": "json"
                        }
                    },
                    "400": {
                        "description": "code\":400,\"data\":null,\"msg\":\"\"}",
                        "schema": {
                            "type": "json"
                        }
                    }
                }
            }
        },
        "/student/libraryService/current": {
            "post": {
                "description": "图书馆当前借书记录",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "图书馆"
                ],
                "summary": "图书馆当前借书记录",
                "parameters": [
                    {
                        "description": "用户名",
                        "name": "username",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "密码",
                        "name": "password",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code\":200,\"data\":[{...}],\"msg\":\"OK\"}",
                        "schema": {
                            "type": "json"
                        }
                    },
                    "400": {
                        "description": "code\":400,\"data\":null,\"msg\":\"\"}",
                        "schema": {
                            "type": "json"
                        }
                    }
                }
            }
        },
        "/student/libraryService/history/0": {
            "post": {
                "description": "图书馆借书记录（暂时只支持10本）",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "图书馆"
                ],
                "summary": "图书馆历史借书记录",
                "parameters": [
                    {
                        "description": "用户名",
                        "name": "username",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "密码",
                        "name": "password",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code\":200,\"data\":[{...}],\"msg\":\"OK\"}",
                        "schema": {
                            "type": "json"
                        }
                    },
                    "400": {
                        "description": "code\":400,\"data\":null,\"msg\":\"\"}",
                        "schema": {
                            "type": "json"
                        }
                    }
                }
            }
        },
        "/student/zfService/exam": {
            "post": {
                "description": "正方教务考试信息",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "正方"
                ],
                "summary": "正方教务考试信息",
                "parameters": [
                    {
                        "description": "学期",
                        "name": "term",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "年份",
                        "name": "year",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "用户名",
                        "name": "username",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "密码",
                        "name": "password",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code\":200,\"data\":{object},\"msg\":\"OK\"}",
                        "schema": {
                            "type": "json"
                        }
                    },
                    "400": {
                        "description": "code\":400,\"data\":null,\"msg\":\"\"}",
                        "schema": {
                            "type": "json"
                        }
                    }
                }
            }
        },
        "/student/zfService/program": {
            "post": {
                "description": "正方教务考试信息",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "正方"
                ],
                "summary": "正方教务考试信息",
                "parameters": [
                    {
                        "description": "用户名",
                        "name": "username",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "密码",
                        "name": "password",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code\":200,\"data\":{object},\"msg\":\"OK\"}",
                        "schema": {
                            "type": "json"
                        }
                    },
                    "400": {
                        "description": "code\":400,\"data\":null,\"msg\":\"\"}",
                        "schema": {
                            "type": "json"
                        }
                    }
                }
            }
        },
        "/student/zfService/room": {
            "post": {
                "description": "正方教务考试信息",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "正方"
                ],
                "summary": "正方教务考试信息",
                "parameters": [
                    {
                        "description": "学期",
                        "name": "term",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "年份",
                        "name": "year",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "校区",
                        "name": "campus",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "星期几 1，2，3",
                        "name": "weekday",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "第几周的2次幂的和",
                        "name": "week",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "第几节课的2次幂的和",
                        "name": "classPeriod",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "用户名",
                        "name": "username",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "密码",
                        "name": "password",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code\":200,\"data\":{object},\"msg\":\"OK\"}",
                        "schema": {
                            "type": "json"
                        }
                    },
                    "400": {
                        "description": "code\":400,\"data\":null,\"msg\":\"\"}",
                        "schema": {
                            "type": "json"
                        }
                    }
                }
            }
        },
        "/student/zfService/score": {
            "post": {
                "description": "正方教务课表",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "正方"
                ],
                "summary": "正方教务课表",
                "parameters": [
                    {
                        "description": "学期",
                        "name": "term",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "年份",
                        "name": "year",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "用户名",
                        "name": "username",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "密码",
                        "name": "password",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code\":200,\"data\":{object},\"msg\":\"OK\"}",
                        "schema": {
                            "type": "json"
                        }
                    },
                    "400": {
                        "description": "code\":400,\"data\":null,\"msg\":\"\"}",
                        "schema": {
                            "type": "json"
                        }
                    }
                }
            }
        },
        "/student/zfService/score/info": {
            "post": {
                "description": "正方教务成绩",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "正方"
                ],
                "summary": "正方教务成绩",
                "parameters": [
                    {
                        "description": "学期",
                        "name": "term",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "年份",
                        "name": "year",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "用户名",
                        "name": "username",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "密码",
                        "name": "password",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code\":200,\"data\":{object},\"msg\":\"OK\"}",
                        "schema": {
                            "type": "json"
                        }
                    },
                    "400": {
                        "description": "code\":400,\"data\":null,\"msg\":\"\"}",
                        "schema": {
                            "type": "json"
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
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
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
	swag.Register(swag.Name, &s{})
}
