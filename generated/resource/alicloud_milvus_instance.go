package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudMilvusInstance = `{
  "block": {
    "attributes": {
      "auto_backup": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "auto_pay": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "auto_renew": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "configuration": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "db_admin_password": {
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": "string"
      },
      "db_version": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "encrypted": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "expire_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "ha": {
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
      "instance_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "is_multi_az_storage": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "kms_key_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "load_replicas": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "multi_zone_mode": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "order_id": {
        "computed": true,
        "description_kind": "plain",
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
      "promotion_no": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "resource_group_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "running_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "security_group_ids": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
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
      "vpc_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "zone_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "components": {
        "block": {
          "attributes": {
            "cu_num": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "cu_type": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "disk_size_type": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "pay_type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "replica": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "data_disk": {
              "block": {
                "attributes": {
                  "enabled": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "performance_level": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "size": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "storage_class": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "set"
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
      },
      "vswitch_ids": {
        "block": {
          "attributes": {
            "vsw_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "zone_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudMilvusInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudMilvusInstance), &result)
	return &result
}
