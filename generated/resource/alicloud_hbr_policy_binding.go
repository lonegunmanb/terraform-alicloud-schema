package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudHbrPolicyBinding = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "cross_account_role_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cross_account_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cross_account_user_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "data_source_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "disabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "exclude": {
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
      "include": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "policy_binding_description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "policy_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "speed_limit": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "advanced_options": {
        "block": {
          "block_types": {
            "oss_detail": {
              "block": {
                "attributes": {
                  "ignore_archive_object": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "inventory_cleanup_policy": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "inventory_id": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "udm_detail": {
              "block": {
                "attributes": {
                  "destination_kms_key_id": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "disk_id_list": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "exclude_disk_id_list": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
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

func AlicloudHbrPolicyBindingSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudHbrPolicyBinding), &result)
	return &result
}
