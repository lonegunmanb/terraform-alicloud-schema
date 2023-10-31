package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudDtsSubscriptionJobs = `{
  "block": {
    "attributes": {
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
      "jobs": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "checkpoint": "string",
              "create_time": "string",
              "db_list": "string",
              "dts_instance_id": "string",
              "dts_job_id": "string",
              "dts_job_name": "string",
              "expire_time": "string",
              "id": "string",
              "payment_type": "string",
              "source_endpoint_database_name": "string",
              "source_endpoint_engine_name": "string",
              "source_endpoint_instance_id": "string",
              "source_endpoint_instance_type": "string",
              "source_endpoint_ip": "string",
              "source_endpoint_oracle_sid": "string",
              "source_endpoint_owner_id": "string",
              "source_endpoint_port": "string",
              "source_endpoint_region": "string",
              "source_endpoint_role": "string",
              "source_endpoint_user_name": "string",
              "status": "string",
              "subscription_data_type_ddl": "bool",
              "subscription_data_type_dml": "bool",
              "subscription_host": [
                "list",
                [
                  "object",
                  {
                    "private_host": "string",
                    "public_host": "string",
                    "vpc_host": "string"
                  }
                ]
              ],
              "subscription_instance_network_type": "string",
              "subscription_instance_vpc_id": "string",
              "subscription_instance_vswitch_id": "string",
              "tags": [
                "map",
                "string"
              ]
            }
          ]
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
      "status": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudDtsSubscriptionJobsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudDtsSubscriptionJobs), &result)
	return &result
}
