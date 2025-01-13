package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCloudControlResourceTypes = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ids": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "product": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "types": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "create_only_properties": [
                "list",
                "string"
              ],
              "delete_only_properties": [
                "list",
                "string"
              ],
              "filter_properties": [
                "list",
                "string"
              ],
              "get_only_properties": [
                "list",
                "string"
              ],
              "get_response_properties": [
                "list",
                "string"
              ],
              "handlers": [
                "list",
                [
                  "object",
                  {
                    "create": [
                      "list",
                      [
                        "object",
                        {
                          "permissions": [
                            "list",
                            "string"
                          ]
                        }
                      ]
                    ],
                    "delete": [
                      "list",
                      [
                        "object",
                        {
                          "permissions": [
                            "list",
                            "string"
                          ]
                        }
                      ]
                    ],
                    "get": [
                      "list",
                      [
                        "object",
                        {
                          "permissions": [
                            "list",
                            "string"
                          ]
                        }
                      ]
                    ],
                    "list": [
                      "list",
                      [
                        "object",
                        {
                          "permissions": [
                            "list",
                            "string"
                          ]
                        }
                      ]
                    ],
                    "update": [
                      "list",
                      [
                        "object",
                        {
                          "permissions": [
                            "list",
                            "string"
                          ]
                        }
                      ]
                    ]
                  }
                ]
              ],
              "id": "string",
              "info": [
                "list",
                [
                  "object",
                  {
                    "charge_type": "string",
                    "delivery_scope": "string",
                    "description": "string",
                    "title": "string"
                  }
                ]
              ],
              "list_only_properties": [
                "list",
                "string"
              ],
              "list_response_properties": [
                "list",
                "string"
              ],
              "primary_identifier": "string",
              "product": "string",
              "properties": "string",
              "public_properties": [
                "list",
                "string"
              ],
              "read_only_properties": [
                "list",
                "string"
              ],
              "required": [
                "list",
                "string"
              ],
              "resource_type": "string",
              "sensitive_info_properties": [
                "list",
                "string"
              ],
              "update_only_properties": [
                "list",
                "string"
              ],
              "update_type_properties": [
                "list",
                "string"
              ]
            }
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudCloudControlResourceTypesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCloudControlResourceTypes), &result)
	return &result
}
