package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCmsGroupMetricRule = `{
  "block": {
    "attributes": {
      "category": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "contact_groups": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "dimensions": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "effective_interval": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "email_subject": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "group_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "group_metric_rule_name": {
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
      "interval": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "metric_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "namespace": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "no_effective_interval": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "period": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "rule_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "silence_time": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "webhook": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "escalations": {
        "block": {
          "block_types": {
            "critical": {
              "block": {
                "attributes": {
                  "comparison_operator": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "statistics": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "threshold": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "times": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "set"
            },
            "info": {
              "block": {
                "attributes": {
                  "comparison_operator": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "statistics": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "threshold": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "times": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "set"
            },
            "warn": {
              "block": {
                "attributes": {
                  "comparison_operator": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "statistics": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "threshold": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "times": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "set"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "set"
      },
      "targets": {
        "block": {
          "attributes": {
            "arn": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "json_params": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "level": {
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

func AlicloudCmsGroupMetricRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCmsGroupMetricRule), &result)
	return &result
}
