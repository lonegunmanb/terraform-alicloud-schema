package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCmsHybridMonitorSlsTask = `{
  "block": {
    "attributes": {
      "collect_interval": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "collect_target_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "description": {
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
      "namespace": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "task_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "attach_labels": {
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
        "nesting_mode": "set"
      },
      "sls_process_config": {
        "block": {
          "block_types": {
            "express": {
              "block": {
                "attributes": {
                  "alias": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "express": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "set"
            },
            "filter": {
              "block": {
                "attributes": {
                  "relation": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "filters": {
                    "block": {
                      "attributes": {
                        "operator": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "sls_key_name": {
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
                    "nesting_mode": "set"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "set"
            },
            "group_by": {
              "block": {
                "attributes": {
                  "alias": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "sls_key_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "set"
            },
            "statistics": {
              "block": {
                "attributes": {
                  "alias": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "function": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "parameter_one": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "parameter_two": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "sls_key_name": {
                    "description_kind": "plain",
                    "optional": true,
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
        "max_items": 1,
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

func AlicloudCmsHybridMonitorSlsTaskSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCmsHybridMonitorSlsTask), &result)
	return &result
}
