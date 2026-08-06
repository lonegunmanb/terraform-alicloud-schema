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
      "cluster_description": {
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
      "cluster_status": {
        "computed": true,
        "description_kind": "plain",
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
      "ehpc_version": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "enable_scale_in": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "enable_scale_out": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "grow_interval": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "idle_interval": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "is_enterprise_security_group": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "max_core_count": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "max_count": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "modify_time": {
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
      "security_group_id": {
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
      "additional_packages": {
        "block": {
          "attributes": {
            "name": {
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
        "nesting_mode": "list"
      },
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
      "cluster_custom_configuration": {
        "block": {
          "attributes": {
            "args": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "script": {
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
      "monitor_spec": {
        "block": {
          "attributes": {
            "enable_compute_load_monitor": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "queues": {
        "block": {
          "attributes": {
            "allocation_strategy": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "enable_scale_in": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "enable_scale_out": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "hostname_prefix": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "hostname_suffix": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "initial_count": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "inter_connect": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "keep_alive_nodes": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "max_count": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "max_count_per_cycle": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "min_count": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "queue_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "ram_role": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "reserved_node_pool_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "vswitch_ids": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "block_types": {
            "compute_nodes": {
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
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "enable_ht": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
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
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "scheduler_spec": {
        "block": {
          "attributes": {
            "enable_topology_awareness": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
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
