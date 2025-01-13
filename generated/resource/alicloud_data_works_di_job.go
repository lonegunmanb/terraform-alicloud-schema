package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudDataWorksDiJob = `{
  "block": {
    "attributes": {
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "destination_data_source_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "di_job_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "job_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "migration_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "source_data_source_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "destination_data_source_settings": {
        "block": {
          "attributes": {
            "data_source_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "min_items": 1,
        "nesting_mode": "list"
      },
      "job_settings": {
        "block": {
          "attributes": {
            "channel_settings": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "column_data_type_settings": {
              "block": {
                "attributes": {
                  "destination_data_type": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "source_data_type": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "cycle_schedule_settings": {
              "block": {
                "attributes": {
                  "cycle_migration_type": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "schedule_parameters": {
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
            "ddl_handling_settings": {
              "block": {
                "attributes": {
                  "action": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "type": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "runtime_settings": {
              "block": {
                "attributes": {
                  "name": {
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
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "resource_settings": {
        "block": {
          "block_types": {
            "offline_resource_settings": {
              "block": {
                "attributes": {
                  "requested_cu": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "resource_group_identifier": {
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
            "realtime_resource_settings": {
              "block": {
                "attributes": {
                  "requested_cu": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "resource_group_identifier": {
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
            "schedule_resource_settings": {
              "block": {
                "attributes": {
                  "requested_cu": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "resource_group_identifier": {
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
        "min_items": 1,
        "nesting_mode": "list"
      },
      "source_data_source_settings": {
        "block": {
          "attributes": {
            "data_source_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "data_source_properties": {
              "block": {
                "attributes": {
                  "encoding": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "timezone": {
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
        "min_items": 1,
        "nesting_mode": "list"
      },
      "table_mappings": {
        "block": {
          "block_types": {
            "source_object_selection_rules": {
              "block": {
                "attributes": {
                  "action": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "expression": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "expression_type": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "object_type": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "transformation_rules": {
              "block": {
                "attributes": {
                  "rule_action_type": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "rule_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "rule_target_type": {
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
        "min_items": 1,
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
      "transformation_rules": {
        "block": {
          "attributes": {
            "rule_action_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "rule_expression": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "rule_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "rule_target_type": {
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

func AlicloudDataWorksDiJobSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudDataWorksDiJob), &result)
	return &result
}
