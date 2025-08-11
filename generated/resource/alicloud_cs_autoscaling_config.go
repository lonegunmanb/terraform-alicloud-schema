package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCsAutoscalingConfig = `{
  "block": {
    "attributes": {
      "cluster_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cool_down_duration": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "daemonset_eviction_for_nodes": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "expander": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "gpu_utilization_threshold": {
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
      "max_graceful_termination_sec": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "min_replica_count": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "priorities": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "recycle_node_deletion_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "scale_down_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "scale_up_from_zero": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "scaler_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "scan_interval": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "skip_nodes_with_local_storage": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "skip_nodes_with_system_pods": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "unneeded_duration": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "utilization_threshold": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
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

func AlicloudCsAutoscalingConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCsAutoscalingConfig), &result)
	return &result
}
