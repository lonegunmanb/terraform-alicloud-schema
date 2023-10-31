package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudComputeNestServiceInstances = `{
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
      "service_instances": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "enable_instance_ops": "bool",
              "id": "string",
              "operated_service_instance_id": "string",
              "operation_end_time": "string",
              "operation_start_time": "string",
              "parameters": "string",
              "resources": "string",
              "service": [
                "list",
                [
                  "object",
                  {
                    "deploy_type": "string",
                    "publish_time": "string",
                    "service_id": "string",
                    "service_infos": [
                      "list",
                      [
                        "object",
                        {
                          "image": "string",
                          "locale": "string",
                          "name": "string",
                          "short_description": "string"
                        }
                      ]
                    ],
                    "service_type": "string",
                    "status": "string",
                    "supplier_name": "string",
                    "supplier_url": "string",
                    "version": "string",
                    "version_name": "string"
                  }
                ]
              ],
              "service_instance_id": "string",
              "service_instance_name": "string",
              "source": "string",
              "status": "string",
              "tags": [
                "map",
                "string"
              ],
              "template_name": "string"
            }
          ]
        ]
      },
      "status": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      }
    },
    "block_types": {
      "filter": {
        "block": {
          "attributes": {
            "name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudComputeNestServiceInstancesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudComputeNestServiceInstances), &result)
	return &result
}
