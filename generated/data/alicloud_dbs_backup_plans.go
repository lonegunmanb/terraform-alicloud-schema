package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudDbsBackupPlans = `{
  "block": {
    "attributes": {
      "backup_plan_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "enable_details": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ids": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "name_regex": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "names": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "page_number": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "page_size": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "plans": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "backup_gateway_id": "string",
              "backup_method": "string",
              "backup_objects": "string",
              "backup_period": "string",
              "backup_plan_id": "string",
              "backup_plan_name": "string",
              "backup_retention_period": "number",
              "backup_start_time": "string",
              "backup_storage_type": "string",
              "cross_aliyun_id": "string",
              "cross_role_name": "string",
              "database_type": "string",
              "duplication_archive_period": "number",
              "duplication_infrequent_access_period": "number",
              "enable_backup_log": "bool",
              "id": "string",
              "instance_class": "string",
              "oss_bucket_name": "string",
              "payment_type": "string",
              "resource_group_id": "string",
              "source_endpoint_database_name": "string",
              "source_endpoint_instance_id": "string",
              "source_endpoint_instance_type": "string",
              "source_endpoint_region": "string",
              "source_endpoint_sid": "string",
              "source_endpoint_user_name": "string",
              "status": "string"
            }
          ]
        ]
      },
      "status": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudDbsBackupPlansSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudDbsBackupPlans), &result)
	return &result
}
