package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEhpcClusterV2 = `{
  "block": {
    "attributes": {
      "client_version": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cluster_category": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cluster_mode": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cluster_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cluster_vpc_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cluster_vswitch_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "deletion_protection": {
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
      "resource_group_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "security_group_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "addons": {
        "block": {
          "attributes": {
            "name": {
              "description_kind": "plain",
              "required": true,
              "sensitive": true,
              "type": "string"
            },
            "resources_spec": {
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": "string"
            },
            "services_spec": {
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": "string"
            },
            "version": {
              "description_kind": "plain",
              "required": true,
              "sensitive": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "cluster_credentials": {
        "block": {
          "attributes": {
            "key_pair_name": {
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": "string"
            },
            "password": {
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "manager": {
        "block": {
          "block_types": {
            "directory_service": {
              "block": {
                "attributes": {
                  "type": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "version": {
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
            "dns": {
              "block": {
                "attributes": {
                  "type": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "version": {
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
            "manager_node": {
              "block": {
                "attributes": {
                  "auto_renew": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "auto_renew_period": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "duration": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "enable_ht": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "expired_time": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "image_id": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "instance_charge_type": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "instance_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "instance_type": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "period": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "period_unit": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "spot_price_limit": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "spot_strategy": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "system_disk": {
                    "block": {
                      "attributes": {
                        "category": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "level": {
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
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "scheduler": {
              "block": {
                "attributes": {
                  "type": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "version": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
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
      "shared_storages": {
        "block": {
          "attributes": {
            "file_system_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "mount_directory": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "mount_options": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "mount_target_domain": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "nas_directory": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "protocol_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
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

func AlicloudEhpcClusterV2Schema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEhpcClusterV2), &result)
	return &result
}
