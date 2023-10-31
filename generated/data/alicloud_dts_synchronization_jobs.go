package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudDtsSynchronizationJobs = `{
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
              "data_initialization": "bool",
              "data_synchronization": "bool",
              "db_list": "string",
              "destination_endpoint_data_base_name": "string",
              "destination_endpoint_engine_name": "string",
              "destination_endpoint_instance_id": "string",
              "destination_endpoint_instance_type": "string",
              "destination_endpoint_ip": "string",
              "destination_endpoint_oracle_sid": "string",
              "destination_endpoint_port": "string",
              "destination_endpoint_region": "string",
              "destination_endpoint_user_name": "string",
              "dts_instance_id": "string",
              "dts_job_id": "string",
              "dts_job_name": "string",
              "expire_time": "string",
              "id": "string",
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
              "structure_initialization": "bool",
              "synchronization_direction": "string"
            }
          ]
        ]
      },
      "name_regex": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudDtsSynchronizationJobsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudDtsSynchronizationJobs), &result)
	return &result
}
