package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudDatabaseGatewayGateways = `{
  "block": {
    "attributes": {
      "enable_details": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "gateways": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "create_time": "string",
              "gateway_desc": "string",
              "gateway_instances": [
                "list",
                [
                  "object",
                  {
                    "connect_endpoint_type": "string",
                    "current_daemon_version": "string",
                    "current_version": "string",
                    "end_point": "string",
                    "gateway_instance_id": "string",
                    "gateway_instance_status": "string",
                    "last_update_time": "string",
                    "local_ip": "string",
                    "message": "string",
                    "output_ip": "string"
                  }
                ]
              ],
              "gateway_name": "string",
              "hosts": "string",
              "id": "string",
              "modified_time": "string",
              "parent_id": "string",
              "status": "string",
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
      "search_key": {
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

func AlicloudDatabaseGatewayGatewaysSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudDatabaseGatewayGateways), &result)
	return &result
}
