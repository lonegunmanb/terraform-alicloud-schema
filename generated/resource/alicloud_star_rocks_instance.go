package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudStarRocksInstance = `{
  "block": {
    "attributes": {
      "admin_password": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "auto_renew": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "cluster_zone_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "duration": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "encrypted": {
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
      "kms_key_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "oss_accessing_role_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "package_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "pay_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "pricing_cycle": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "promotion_option_no": {
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
      "run_mode": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
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
      "version": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "vpc_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "backend_node_groups": {
        "block": {
          "attributes": {
            "cu": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "disk_number": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "local_storage_instance_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "resident_node_number": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "spec_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "storage_performance_level": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "storage_size": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
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
      },
      "frontend_node_groups": {
        "block": {
          "attributes": {
            "cu": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "disk_number": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "local_storage_instance_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "resident_node_number": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "spec_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "storage_performance_level": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "storage_size": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
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
      },
      "observer_node_groups": {
        "block": {
          "attributes": {
            "cu": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "disk_number": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "local_storage_instance_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "resident_node_number": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "spec_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "storage_performance_level": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "storage_size": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
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
      "vswitches": {
        "block": {
          "attributes": {
            "vswitch_id": {
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
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudStarRocksInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudStarRocksInstance), &result)
	return &result
}
