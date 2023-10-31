package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEdasK8SApplication = `{
  "block": {
    "attributes": {
      "application_descriotion": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "application_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "cluster_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "command": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "command_args": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "edas_container_version": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "envs": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "image_url": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "internet_slb_id": {
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "internet_slb_port": {
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "internet_slb_protocol": {
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "internet_target_port": {
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "jdk": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "limit_m_cpu": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "limit_mem": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "liveness": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "local_volume": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "logical_region_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "mount_descs": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "namespace": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "nas_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "package_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "package_url": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "package_version": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "post_start": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "pre_stop": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "readiness": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "replicas": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "requests_m_cpu": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "requests_mem": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "web_container": {
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

func AlicloudEdasK8SApplicationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEdasK8SApplication), &result)
	return &result
}
