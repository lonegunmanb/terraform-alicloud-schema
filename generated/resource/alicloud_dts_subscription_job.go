package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudDtsSubscriptionJob = `{
  "block": {
    "attributes": {
      "checkpoint": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "compute_unit": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "database_count": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "db_list": {
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
      "destination_endpoint_engine_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "destination_region": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "dts_instance_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "dts_job_name": {
        "description_kind": "plain",
        "optional": true,
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
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "payment_duration": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "payment_duration_unit": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "payment_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "reserve": {
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
        "required": true,
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
      "status": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "subscription_data_type_ddl": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "subscription_data_type_dml": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "subscription_instance_network_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "subscription_instance_vpc_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "subscription_instance_vswitch_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "sync_architecture": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "synchronization_direction": {
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

func AlicloudDtsSubscriptionJobSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudDtsSubscriptionJob), &result)
	return &result
}
