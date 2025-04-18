package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEfloExperimentPlanTemplate = `{
  "block": {
    "attributes": {
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
      "privacy_level": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "template_description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "template_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "template_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "template_pipeline": {
        "block": {
          "attributes": {
            "pipeline_order": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "scene": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "setting_params": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "workload_id": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "workload_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "env_params": {
              "block": {
                "attributes": {
                  "cpu_per_worker": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "cuda_version": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "gpu_driver_version": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "gpu_per_worker": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "memory_per_worker": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "nccl_version": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "py_torch_version": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "share_memory": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "worker_num": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
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
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudEfloExperimentPlanTemplateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEfloExperimentPlanTemplate), &result)
	return &result
}
