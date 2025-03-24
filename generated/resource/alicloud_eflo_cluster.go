package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEfloCluster = `{
  "block": {
    "attributes": {
      "cluster_description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cluster_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cluster_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description_kind": "plain",
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
      "ignore_failed_node_tasks": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "nimiz_vswitches": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "open_eni_jumbo_frame": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "resource_group_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
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
      }
    },
    "block_types": {
      "components": {
        "block": {
          "attributes": {
            "component_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "component_config": {
              "block": {
                "attributes": {
                  "basic_args": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "node_units": {
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
        "nesting_mode": "list"
      },
      "networks": {
        "block": {
          "attributes": {
            "security_group_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "tail_ip_version": {
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
            },
            "vswitch_zone_id": {
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
            "new_vpd_info": {
              "block": {
                "attributes": {
                  "cen_id": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "cloud_link_cidr": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "cloud_link_id": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "monitor_vpc_id": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "monitor_vswitch_id": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "vpd_cidr": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "vpd_subnets": {
                    "block": {
                      "attributes": {
                        "subnet_cidr": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "subnet_type": {
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
              "max_items": 1,
              "nesting_mode": "list"
            },
            "vpd_info": {
              "block": {
                "attributes": {
                  "vpd_id": {
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
      "node_groups": {
        "block": {
          "attributes": {
            "image_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "machine_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "node_group_description": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "node_group_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "user_data": {
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

func AlicloudEfloClusterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEfloCluster), &result)
	return &result
}
