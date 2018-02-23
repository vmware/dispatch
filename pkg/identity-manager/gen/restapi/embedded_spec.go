///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

// SwaggerJSON embedded version of the swagger document used at generation time
var SwaggerJSON json.RawMessage

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http",
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "description": "VMware Dispatch Identity Manager\n",
    "title": "Identity Manager",
    "contact": {
      "email": "dispatch@vmware.com"
    },
    "version": "1.0.0"
  },
  "basePath": "/",
  "paths": {
    "/": {
      "get": {
        "summary": "an placehold root page, no authentication is required at this point",
        "operationId": "root",
        "responses": {
          "200": {
            "description": "home page",
            "schema": {
              "$ref": "#/definitions/Message"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/v1/iam/auth": {
      "get": {
        "summary": "handles authentication",
        "operationId": "auth",
        "security": [
          {
            "cookie": []
          }
        ],
        "responses": {
          "202": {
            "description": "default response if authenticated",
            "schema": {
              "$ref": "#/definitions/Message"
            }
          },
          "401": {
            "description": "Unauthorized"
          },
          "403": {
            "description": "Forbidden"
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/v1/iam/home": {
      "get": {
        "summary": "an placehold home page",
        "operationId": "home",
        "security": [
          {
            "cookie": []
          }
        ],
        "responses": {
          "200": {
            "description": "home page",
            "schema": {
              "$ref": "#/definitions/Message"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/v1/iam/policy": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "policy"
        ],
        "summary": "List all existing policies",
        "operationId": "getPolicies",
        "security": [
          {
            "cookie": []
          }
        ],
        "responses": {
          "200": {
            "description": "Successful operation",
            "schema": {
              "$ref": "#/definitions/getPoliciesOKBody"
            }
          },
          "500": {
            "description": "Internal Error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "default": {
            "description": "Unexpected Error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "policy"
        ],
        "summary": "Add a new policy",
        "operationId": "addPolicy",
        "security": [
          {
            "cookie": []
          }
        ],
        "parameters": [
          {
            "description": "Policy Object",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Policy"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "created",
            "schema": {
              "$ref": "#/definitions/Policy"
            }
          },
          "400": {
            "description": "Invalid input",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal Error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "default": {
            "description": "Generic error response",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/v1/iam/policy/{policyName}": {
      "get": {
        "description": "get an Policy by name",
        "produces": [
          "application/json"
        ],
        "tags": [
          "policy"
        ],
        "summary": "Find Policy by name",
        "operationId": "getPolicy",
        "security": [
          {
            "cookie": []
          }
        ],
        "responses": {
          "200": {
            "description": "Successful operation",
            "schema": {
              "$ref": "#/definitions/Policy"
            }
          },
          "400": {
            "description": "Invalid Name supplied",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "404": {
            "description": "Policy not found",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "put": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "policy"
        ],
        "summary": "Update a Policy",
        "operationId": "updatePolicy",
        "security": [
          {
            "cookie": []
          }
        ],
        "parameters": [
          {
            "description": "Policy object",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Policy"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successful update",
            "schema": {
              "$ref": "#/definitions/Policy"
            }
          },
          "400": {
            "description": "Invalid input",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "404": {
            "description": "Policy not found",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "delete": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "policy"
        ],
        "summary": "Deletes an Policy",
        "operationId": "deletePolicy",
        "security": [
          {
            "cookie": []
          }
        ],
        "responses": {
          "200": {
            "description": "Successful operation",
            "schema": {
              "$ref": "#/definitions/Policy"
            }
          },
          "400": {
            "description": "Invalid Name supplied",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "404": {
            "description": "Policy not found",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "parameters": [
        {
          "pattern": "^[\\w\\d\\-]+$",
          "type": "string",
          "description": "Name of Policy to work on",
          "name": "policyName",
          "in": "path",
          "required": true
        }
      ]
    },
    "/v1/iam/redirect": {
      "get": {
        "summary": "redirect to localhost for vs-cli login (testing)",
        "operationId": "redirect",
        "security": [
          {
            "cookie": []
          }
        ],
        "parameters": [
          {
            "type": "string",
            "description": "the local server url redirecting to",
            "name": "redirect",
            "in": "query"
          }
        ],
        "responses": {
          "302": {
            "description": "redirect",
            "headers": {
              "Location": {
                "type": "string",
                "description": "redirect location"
              }
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "Message": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "message": {
          "type": "string"
        }
      }
    },
    "Policy": {
      "type": "object",
      "required": [
        "name",
        "rules"
      ],
      "properties": {
        "createdTime": {
          "type": "integer",
          "readOnly": true
        },
        "id": {
          "type": "string",
          "format": "uuid"
        },
        "modifiedTime": {
          "type": "integer",
          "readOnly": true
        },
        "name": {
          "type": "string",
          "pattern": "^[\\w\\d\\-]+$"
        },
        "rules": {
          "$ref": "#/definitions/policyRules"
        },
        "status": {
          "$ref": "#/definitions/Status",
          "readOnly": true
        }
      }
    },
    "Redirect": {
      "type": "object",
      "required": [
        "location"
      ],
      "properties": {
        "location": {
          "type": "string"
        }
      }
    },
    "Rule": {
      "type": "object",
      "required": [
        "subjects",
        "resources",
        "actions"
      ],
      "properties": {
        "actions": {
          "type": "array",
          "items": {
            "type": "string",
            "pattern": "^[\\w\\d\\-]+$"
          }
        },
        "resources": {
          "type": "array",
          "items": {
            "type": "string",
            "pattern": "^[\\w\\d\\-]+$"
          }
        },
        "subjects": {
          "type": "array",
          "items": {
            "type": "string",
            "pattern": "^[\\w\\d\\-]+$"
          }
        }
      }
    },
    "Status": {
      "type": "string",
      "enum": [
        "CREATING",
        "READY",
        "UPDATING",
        "DELETED",
        "ERROR"
      ]
    },
    "getPoliciesOKBody": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/Policy"
      },
      "x-go-gen-location": "operations"
    },
    "policyRules": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/Rule"
      },
      "x-go-gen-location": "models"
    }
  },
  "securityDefinitions": {
    "cookie": {
      "description": "use cookies for authentication, when the user already logged in",
      "type": "apiKey",
      "name": "Cookie",
      "in": "header"
    }
  },
  "tags": [
    {
      "name": "authentication"
    }
  ]
}`))
}
