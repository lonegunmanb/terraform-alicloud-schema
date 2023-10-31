package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudLindormInstances = `{
  "block": {
    "attributes": {
      "enable_details": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
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
      "instances": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "auto_renew": "string",
              "cold_storage": "number",
              "create_time": "string",
              "deletion_proection": "bool",
              "disk_category": "string",
              "disk_usage": "string",
              "disk_warning_threshold": "string",
              "engine_type": "number",
              "expired_time": "string",
              "file_engine_node_count": "number",
              "file_engine_specification": "string",
              "id": "string",
              "instance_id": "string",
              "instance_name": "string",
              "instance_storage": "string",
              "ip_white_list": [
                "list",
                "string"
              ],
              "lts_node_count": "number",
              "lts_node_specification": "string",
              "network_type": "string",
              "payment_type": "string",
              "phoenix_node_count": "number",
              "phoenix_node_specification": "string",
              "resource_owner_id": "string",
              "search_engine_node_count": "number",
              "search_engine_specification": "string",
              "service_type": "string",
              "status": "string",
              "table_engine_node_count": "number",
              "table_engine_specification": "string",
              "time_series_engine_node_count": "number",
              "time_serires_engine_specification": "string",
              "vpc_id": "string",
              "vswitch_id": "string",
              "zone_id": "string"
            }
          ]
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
      "query_str": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "support_engine": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudLindormInstancesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudLindormInstances), &result)
	return &result
}
