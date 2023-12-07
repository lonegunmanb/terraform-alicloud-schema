package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudDtsSynchronizationJob = `{
  "block": {
    "attributes": {
      "checkpoint": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "data_check_configure": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "data_initialization": {
        "description_kind": "plain",
        "required": true,
        "type": "bool"
      },
      "data_synchronization": {
        "description_kind": "plain",
        "required": true,
        "type": "bool"
      },
      "db_list": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "dedicated_cluster_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "delay_notice": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "delay_phone": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "delay_rule_time": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "destination_endpoint_database_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "destination_endpoint_engine_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "destination_endpoint_instance_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "destination_endpoint_instance_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "destination_endpoint_ip": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "destination_endpoint_oracle_sid": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "destination_endpoint_owner_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "destination_endpoint_password": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "destination_endpoint_port": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "destination_endpoint_region": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "destination_endpoint_role": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "destination_endpoint_user_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "dts_bis_label": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "dts_instance_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "dts_job_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "error_notice": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "error_phone": {
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
      "instance_class": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "reserve": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_endpoint_database_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_endpoint_engine_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "source_endpoint_instance_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_endpoint_instance_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "source_endpoint_ip": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_endpoint_oracle_sid": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_endpoint_owner_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_endpoint_password": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_endpoint_port": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_endpoint_region": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_endpoint_role": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_endpoint_user_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_endpoint_vswitch_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "structure_initialization": {
        "description_kind": "plain",
        "required": true,
        "type": "bool"
      },
      "synchronization_direction": {
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

func AlicloudDtsSynchronizationJobSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudDtsSynchronizationJob), &result)
	return &result
}
