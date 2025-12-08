package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEfloNode = `{
  "block": {
    "attributes": {
      "billing_cycle": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "classify": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cluster_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "computing_server": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "discount_level": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "hostname": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "hpn_zone": {
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
      "install_pai": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "login_password": {
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": "string"
      },
      "machine_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "node_group_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "node_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "payment_ratio": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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
      "product_form": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "renew_period": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "renewal_status": {
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
      "server_arch": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "stage_num": {
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
      "tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "user_data": {
        "computed": true,
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
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "zone": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "data_disk": {
        "block": {
          "attributes": {
            "category": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "performance_level": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "size": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
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
                  "hostname": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
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

func AlicloudEfloNodeSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEfloNode), &result)
	return &result
}
