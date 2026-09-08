package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudApigSources = `{
  "block": {
    "attributes": {
      "enable_details": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "gateway_id": {
        "description_kind": "plain",
        "optional": true,
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
      "resource_group_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "sources": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "association_reason": "string",
              "association_status": "string",
              "create_time": "number",
              "gateway_id": "string",
              "id": "string",
              "k8s_source_info": [
                "list",
                [
                  "object",
                  {
                    "cluster_id": "string"
                  }
                ]
              ],
              "nacos_source_info": [
                "list",
                [
                  "object",
                  {
                    "address": "string",
                    "cluster_id": "string",
                    "instance_id": "string"
                  }
                ]
              ],
              "resource_group_id": "string",
              "source_id": "string",
              "source_name": "string",
              "type": "string",
              "update_time": "number"
            }
          ]
        ]
      },
      "type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudApigSourcesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudApigSources), &result)
	return &result
}
