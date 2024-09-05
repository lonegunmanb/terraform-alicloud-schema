package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudGpdbHadoopDataSource = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "data_source_description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "data_source_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "data_source_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "data_source_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "db_instance_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "emr_instance_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "hadoop_core_conf": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "hadoop_create_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "hadoop_hosts_address": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "hdfs_conf": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "hive_conf": {
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
      "map_reduce_conf": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "yarn_conf": {
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

func AlicloudGpdbHadoopDataSourceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudGpdbHadoopDataSource), &result)
	return &result
}
