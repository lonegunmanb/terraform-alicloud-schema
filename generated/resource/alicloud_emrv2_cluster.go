package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEmrv2Cluster = `{
  "block": {
    "attributes": {
      "applications": {
        "description_kind": "plain",
        "required": true,
        "type": [
          "set",
          "string"
        ]
      },
      "cluster_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "cluster_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "deletion_protection": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "deploy_mode": {
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
      "log_collect_strategy": {
        "computed": true,
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
      "release_version": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resource_group_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "security_mode": {
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
      }
    },
    "block_types": {
      "application_configs": {
        "block": {
          "attributes": {
            "application_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "config_description": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "config_file_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "config_item_key": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "config_item_value": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "config_scope": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "node_group_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "node_group_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "bootstrap_scripts": {
        "block": {
          "attributes": {
            "execution_fail_strategy": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "execution_moment": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "priority": {
              "deprecated": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "script_args": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "script_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "script_path": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "node_selector": {
              "block": {
                "attributes": {
                  "node_group_id": {
                    "deprecated": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "node_group_ids": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "node_group_name": {
                    "deprecated": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "node_group_names": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "node_group_types": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "node_names": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "node_select_type": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "set"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "node_attributes": {
        "block": {
          "attributes": {
            "data_disk_encrypted": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "data_disk_kms_key_id": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "key_pair_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "ram_role": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "security_group_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "system_disk_encrypted": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "system_disk_kms_key_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "vpc_id": {
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
          "description_kind": "plain"
        },
        "min_items": 1,
        "nesting_mode": "set"
      },
      "node_groups": {
        "block": {
          "attributes": {
            "additional_security_group_ids": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "deployment_set_strategy": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "graceful_shutdown": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "instance_types": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "set",
                "string"
              ]
            },
            "node_count": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "node_group_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "node_group_type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "node_resize_strategy": {
              "computed": true,
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
            "spot_instance_remedy": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "spot_strategy": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "vswitch_ids": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "with_public_ip": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "block_types": {
            "ack_config": {
              "block": {
                "attributes": {
                  "ack_instance_id": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "limit_cpu": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "limit_memory": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "namespace": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "node_affinity": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "pod_affinity": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "pod_anti_affinity": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "pre_start_command": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "request_cpu": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "request_memory": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "block_types": {
                  "custom_annotations": {
                    "block": {
                      "attributes": {
                        "key": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "value": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "set"
                  },
                  "custom_labels": {
                    "block": {
                      "attributes": {
                        "key": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "value": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "set"
                  },
                  "node_selectors": {
                    "block": {
                      "attributes": {
                        "key": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "value": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "set"
                  },
                  "pvcs": {
                    "block": {
                      "attributes": {
                        "data_disk_size": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        },
                        "data_disk_storage_class": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "path": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "tolerations": {
                    "block": {
                      "attributes": {
                        "effect": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "key": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "operator": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "value": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "volume_mounts": {
                    "block": {
                      "attributes": {
                        "name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "path": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "volumes": {
                    "block": {
                      "attributes": {
                        "name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "path": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "type": {
                          "description_kind": "plain",
                          "required": true,
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
            "auto_scaling_policy": {
              "block": {
                "block_types": {
                  "constraints": {
                    "block": {
                      "attributes": {
                        "max_capacity": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "min_capacity": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "scaling_rules": {
                    "block": {
                      "attributes": {
                        "activity_type": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "adjustment_type": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "adjustment_value": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        },
                        "min_adjustment_value": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "rule_name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "trigger_type": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "metrics_trigger": {
                          "block": {
                            "attributes": {
                              "condition_logic_operator": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "cool_down_interval": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "evaluation_count": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "number"
                              },
                              "time_window": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "number"
                              }
                            },
                            "block_types": {
                              "conditions": {
                                "block": {
                                  "attributes": {
                                    "comparison_operator": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "metric_name": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "statistics": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "threshold": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "number"
                                    }
                                  },
                                  "block_types": {
                                    "tags": {
                                      "block": {
                                        "attributes": {
                                          "key": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          },
                                          "value": {
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
                              "time_constraints": {
                                "block": {
                                  "attributes": {
                                    "end_time": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "start_time": {
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
                        "time_trigger": {
                          "block": {
                            "attributes": {
                              "end_time": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "launch_expiration_time": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "launch_time": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "recurrence_type": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "recurrence_value": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "start_time": {
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
                    "nesting_mode": "list"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "cost_optimized_config": {
              "block": {
                "attributes": {
                  "on_demand_base_capacity": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "on_demand_percentage_above_base_capacity": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "spot_instance_pools": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "set"
            },
            "data_disks": {
              "block": {
                "attributes": {
                  "category": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "count": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "performance_level": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "size": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "description_kind": "plain"
              },
              "min_items": 1,
              "nesting_mode": "set"
            },
            "private_pool_options": {
              "block": {
                "attributes": {
                  "match_criteria": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "private_pool_ids": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "set"
            },
            "spot_bid_prices": {
              "block": {
                "attributes": {
                  "bid_price": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "instance_type": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "set"
            },
            "subscription_config": {
              "block": {
                "attributes": {
                  "auto_pay_order": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "auto_renew": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "auto_renew_duration": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "auto_renew_duration_unit": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "payment_duration": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "payment_duration_unit": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "set"
            },
            "system_disk": {
              "block": {
                "attributes": {
                  "category": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "count": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "performance_level": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "size": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "set"
            }
          },
          "description_kind": "plain"
        },
        "min_items": 1,
        "nesting_mode": "list"
      },
      "subscription_config": {
        "block": {
          "attributes": {
            "auto_pay_order": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "auto_renew": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "auto_renew_duration": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "auto_renew_duration_unit": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "payment_duration": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "payment_duration_unit": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
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

func AlicloudEmrv2ClusterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEmrv2Cluster), &result)
	return &result
}
