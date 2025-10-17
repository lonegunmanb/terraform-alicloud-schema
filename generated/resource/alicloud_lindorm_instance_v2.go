package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudLindormInstanceV2 = `{
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
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "auto_renewal": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "cloud_storage_size": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "cloud_storage_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "deletion_protection": {
        "computed": true,
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
      "instance_alias": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "payment_type": {
        "description_kind": "plain",
        "required": true,
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
      "region_id": {
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
      "vpc_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "vswitch_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "zone_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "engine_list": {
        "block": {
          "attributes": {
            "connect_address_list": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "set",
                [
                  "object",
                  {
                    "address": "string",
                    "port": "string",
                    "type": "string"
                  }
                ]
              ]
            },
            "engine_type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "is_last_version": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "latest_version": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "version": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "block_types": {
            "node_group": {
              "block": {
                "attributes": {
                  "category": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "cpu_core_count": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "enable_attach_local_disk": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "memory_size_gi_b": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "node_count": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "node_disk_size": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "node_disk_type": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "node_spec": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "resource_group_name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "spec_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "status": {
                    "computed": true,
                    "description_kind": "plain",
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
        "min_items": 1,
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

func AlicloudLindormInstanceV2Schema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudLindormInstanceV2), &result)
	return &result
}
