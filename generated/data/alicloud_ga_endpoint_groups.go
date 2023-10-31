package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudGaEndpointGroups = `{
  "block": {
    "attributes": {
      "accelerator_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "endpoint_group_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "groups": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "description": "string",
              "endpoint_configurations": [
                "list",
                [
                  "object",
                  {
                    "enable_clientip_preservation": "bool",
                    "endpoint": "string",
                    "probe_port": "number",
                    "probe_protocol": "string",
                    "type": "string",
                    "weight": "number"
                  }
                ]
              ],
              "endpoint_group_id": "string",
              "endpoint_group_region": "string",
              "health_check_interval_seconds": "number",
              "health_check_path": "string",
              "health_check_port": "number",
              "health_check_protocol": "string",
              "id": "string",
              "listener_id": "string",
              "name": "string",
              "port_overrides": [
                "list",
                [
                  "object",
                  {
                    "endpoint_port": "number",
                    "listener_port": "number"
                  }
                ]
              ],
              "status": "string",
              "threshold_count": "number",
              "traffic_percentage": "number"
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
      "listener_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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
      "status": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudGaEndpointGroupsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudGaEndpointGroups), &result)
	return &result
}
