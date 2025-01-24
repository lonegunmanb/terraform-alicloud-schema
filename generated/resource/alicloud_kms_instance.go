package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudKmsInstance = `{
  "block": {
    "attributes": {
      "ca_certificate_chain_pem": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "end_date": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "force_delete_without_backup": {
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
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "key_num": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "log": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "log_storage": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "payment_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "period": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "product_version": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "renew_period": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "renew_status": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "secret_num": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "spec": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "vpc_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "vpc_num": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "vswitch_ids": {
        "description_kind": "plain",
        "required": true,
        "type": [
          "set",
          "string"
        ]
      },
      "zone_ids": {
        "description_kind": "plain",
        "required": true,
        "type": [
          "set",
          "string"
        ]
      }
    },
    "block_types": {
      "bind_vpcs": {
        "block": {
          "attributes": {
            "region_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "vpc_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "vpc_owner_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "vswitch_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
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

func AlicloudKmsInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudKmsInstance), &result)
	return &result
}
