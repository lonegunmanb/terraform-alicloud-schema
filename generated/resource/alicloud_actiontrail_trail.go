package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudActiontrailTrail = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "data_event_trail_region": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "event_rw": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "event_selectors": {
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
      "is_organization_trail": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "max_compute_project_arn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "max_compute_write_role_arn": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "mns_topic_arn": {
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "oss_bucket_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "oss_key_prefix": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "oss_write_role_arn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "role_name": {
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "sls_project_arn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "sls_write_role_arn": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "trail_name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "trail_region": {
        "computed": true,
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

func AlicloudActiontrailTrailSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudActiontrailTrail), &result)
	return &result
}
