package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudAlikafkaInstance = `{
  "block": {
    "attributes": {
      "config": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "deploy_type": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "disk_size": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "disk_type": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "domain_endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "eip_max": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "end_point": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "group_left": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "group_used": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "io_max": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "io_max_spec": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "is_partition_buy": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "kms_key_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "paid_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "partition_left": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "partition_num": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "partition_used": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "resource_group_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "sasl_domain_endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "security_group": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "selected_zones": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "service_version": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "spec_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ssl_domain_endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "ssl_endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "topic_left": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "topic_num_of_buy": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "topic_quota": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "topic_used": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
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

func AlicloudAlikafkaInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudAlikafkaInstance), &result)
	return &result
}
