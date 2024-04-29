package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudSlsAlert = `{
  "block": {
    "attributes": {
      "alert_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "display_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "project_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "configuration": {
        "block": {
          "attributes": {
            "auto_annotation": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "dashboard": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "mute_until": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "no_data_fire": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "no_data_severity": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "send_resolved": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "tags": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "threshold": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "type": {
              "computed": true,
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
          "block_types": {
            "annotations": {
              "block": {
                "attributes": {
                  "key": {
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
            "condition_configuration": {
              "block": {
                "attributes": {
                  "condition": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "count_condition": {
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
            "group_configuration": {
              "block": {
                "attributes": {
                  "fields": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "type": {
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
            "join_configurations": {
              "block": {
                "attributes": {
                  "condition": {
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
            "labels": {
              "block": {
                "attributes": {
                  "key": {
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
            "policy_configuration": {
              "block": {
                "attributes": {
                  "action_policy_id": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "alert_policy_id": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "repeat_interval": {
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
            "query_list": {
              "block": {
                "attributes": {
                  "chart_title": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "dashboard_id": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "end": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "power_sql_mode": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "project": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "query": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "region": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "role_arn": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "start": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "store": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "store_type": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "time_span_type": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "ui": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "severity_configurations": {
              "block": {
                "attributes": {
                  "severity": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "block_types": {
                  "eval_condition": {
                    "block": {
                      "attributes": {
                        "condition": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "count_condition": {
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
            },
            "sink_alerthub": {
              "block": {
                "attributes": {
                  "enabled": {
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
            "sink_cms": {
              "block": {
                "attributes": {
                  "enabled": {
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
            "sink_event_store": {
              "block": {
                "attributes": {
                  "enabled": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "endpoint": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "event_store": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "project": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "role_arn": {
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
            "template_configuration": {
              "block": {
                "attributes": {
                  "annotations": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "lang": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "template_id": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "tokens": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  },
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
        "min_items": 1,
        "nesting_mode": "list"
      },
      "schedule": {
        "block": {
          "attributes": {
            "cron_expression": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "delay": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "interval": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "run_immdiately": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "time_zone": {
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
        "max_items": 1,
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
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudSlsAlertSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudSlsAlert), &result)
	return &result
}
