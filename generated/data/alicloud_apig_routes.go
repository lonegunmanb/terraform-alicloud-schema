package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudApigRoutes = `{
  "block": {
    "attributes": {
      "http_api_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
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
      "name_regex": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "names": {
        "computed": true,
        "description_kind": "plain",
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
      "route_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "routes": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "backend": [
                "list",
                [
                  "object",
                  {
                    "scene": "string",
                    "services": [
                      "list",
                      [
                        "object",
                        {
                          "name": "string",
                          "port": "number",
                          "protocol": "string",
                          "service_id": "string",
                          "version": "string",
                          "weight": "number"
                        }
                      ]
                    ]
                  }
                ]
              ],
              "builtin": "string",
              "create_time": "string",
              "description": "string",
              "domain_infos": [
                "list",
                [
                  "object",
                  {
                    "domain_id": "string",
                    "name": "string",
                    "protocol": "string"
                  }
                ]
              ],
              "environment_info": [
                "list",
                [
                  "object",
                  {
                    "alias": "string",
                    "environment_id": "string",
                    "gateway_info": [
                      "list",
                      [
                        "object",
                        {
                          "gateway_edition": "string",
                          "gateway_id": "string",
                          "name": "string"
                        }
                      ]
                    ],
                    "name": "string",
                    "sub_domains": [
                      "list",
                      [
                        "object",
                        {
                          "domain_id": "string",
                          "name": "string",
                          "network_type": "string",
                          "protocol": "string"
                        }
                      ]
                    ]
                  }
                ]
              ],
              "gateway_status": [
                "map",
                "string"
              ],
              "id": "string",
              "match": [
                "list",
                [
                  "object",
                  {
                    "headers": [
                      "list",
                      [
                        "object",
                        {
                          "name": "string",
                          "type": "string",
                          "value": "string"
                        }
                      ]
                    ],
                    "ignore_uri_case": "bool",
                    "methods": [
                      "list",
                      "string"
                    ],
                    "path": [
                      "list",
                      [
                        "object",
                        {
                          "type": "string",
                          "value": "string"
                        }
                      ]
                    ],
                    "query_params": [
                      "list",
                      [
                        "object",
                        {
                          "name": "string",
                          "type": "string",
                          "value": "string"
                        }
                      ]
                    ]
                  }
                ]
              ],
              "route_id": "string",
              "route_name": "string",
              "status": "string",
              "update_time": "string"
            }
          ]
        ]
      },
      "status": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "environment_info": {
        "block": {
          "attributes": {
            "alias": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "environment_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "gateway_info": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                [
                  "object",
                  {
                    "gateway_edition": "string",
                    "gateway_id": "string",
                    "name": "string"
                  }
                ]
              ]
            },
            "name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "sub_domains": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                [
                  "object",
                  {
                    "domain_id": "string",
                    "name": "string",
                    "network_type": "string",
                    "protocol": "string"
                  }
                ]
              ]
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudApigRoutesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudApigRoutes), &result)
	return &result
}
