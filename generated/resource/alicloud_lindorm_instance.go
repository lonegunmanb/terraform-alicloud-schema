package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudLindormInstance = `{
  "block": {
    "attributes": {
      "arbiter_vswitch_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "arbiter_zone_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "arch_version": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "auto_renew": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "auto_renew_period": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cold_storage": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "core_num": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "core_single_storage": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "core_spec": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "deletion_proection": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "disk_category": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "duration": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "enabled_file_engine": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "enabled_lts_engine": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "enabled_search_engine": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "enabled_stream_engine": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "enabled_table_engine": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "enabled_time_serires_engine": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "file_engine_node_count": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "file_engine_specification": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "group_name": {
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
      "instance_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instance_storage": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ip_white_list": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "log_disk_category": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "log_num": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "log_single_storage": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "log_spec": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "lts_node_count": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "lts_node_specification": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "multi_zone_combination": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "payment_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "phoenix_node_count": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "phoenix_node_specification": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "pricing_cycle": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "primary_vswitch_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "primary_zone_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_group_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "search_engine_node_count": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "search_engine_specification": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "service_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "standby_vswitch_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "standby_zone_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "stream_engine_node_count": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "stream_engine_specification": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "table_engine_node_count": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "table_engine_specification": {
        "computed": true,
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
      },
      "time_series_engine_node_count": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "time_series_engine_specification": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "time_serires_engine_specification": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "upgrade_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "vpc_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "vswitch_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "zone_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "timeouts": {
        "block": {
          "attributes": {
            "create": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "delete": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "single"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudLindormInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudLindormInstance), &result)
	return &result
}
