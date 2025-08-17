package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudArmsEnvironments = `{
  "block": {
    "attributes": {
      "environment_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "environments": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "bind_resource_id": "string",
              "bind_resource_type": "string",
              "bind_vpc_cidr": "string",
              "environment_id": "string",
              "environment_name": "string",
              "environment_type": "string",
              "grafana_datasource_uid": "string",
              "grafana_folder_uid": "string",
              "id": "string",
              "managed_type": "string",
              "prometheus_instance_id": "string",
              "region_id": "string",
              "resource_group_id": "string",
              "tags": [
                "map",
                "string"
              ],
              "user_id": "string"
            }
          ]
        ]
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
      "tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudArmsEnvironmentsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudArmsEnvironments), &result)
	return &result
}
