package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEfloNodeGroup = `{
  "block": {
    "attributes": {
      "az": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "cluster_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ignore_failed_node_tasks": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "image_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "machine_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "node_group_description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "node_group_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "node_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "user_data": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "vpd_subnets": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "vswitch_zone_id": {
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
    "block_types": {
      "ip_allocation_policy": {
        "block": {
          "block_types": {
            "bond_policy": {
              "block": {
                "attributes": {
                  "bond_default_subnet": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "bonds": {
                    "block": {
                      "attributes": {
                        "name": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "subnet": {
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
              "max_items": 1,
              "nesting_mode": "list"
            },
            "machine_type_policy": {
              "block": {
                "attributes": {
                  "machine_type": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "bonds": {
                    "block": {
                      "attributes": {
                        "name": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "subnet": {
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
              "nesting_mode": "list"
            },
            "node_policy": {
              "block": {
                "attributes": {
                  "node_id": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "bonds": {
                    "block": {
                      "attributes": {
                        "name": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "subnet": {
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
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "nodes": {
        "block": {
          "attributes": {
            "hostname": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "login_password": {
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": "string"
            },
            "node_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "vpc_id": {
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

func AlicloudEfloNodeGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEfloNodeGroup), &result)
	return &result
}
