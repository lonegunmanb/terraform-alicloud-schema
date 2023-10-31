package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudSddpInstance = `{
  "block": {
    "attributes": {
      "authed": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "dataphin": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "dataphin_count": {
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
      "instance_num": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "logistics": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "modify_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "odps_set": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "oss_bucket_set": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "oss_size": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "payment_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "period": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "rds_set": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "remain_days": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "renew_period": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "renewal_status": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "sd_cbool": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "sdc": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "sddp_version": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "ud_cbool": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "udc": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudSddpInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudSddpInstance), &result)
	return &result
}
