package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudSaeApplications = `{
  "block": {
    "attributes": {
      "app_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "applications": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "acr_assume_role_arn": "string",
              "acr_instance_id": "string",
              "app_description": "string",
              "app_name": "string",
              "application_id": "string",
              "command": "string",
              "command_args": "string",
              "config_map_mount_desc": "string",
              "cpu": "number",
              "create_time": "string",
              "custom_host_alias": "string",
              "edas_container_version": "string",
              "envs": "string",
              "id": "string",
              "image_url": "string",
              "jar_start_args": "string",
              "jar_start_options": "string",
              "jdk": "string",
              "liveness": "string",
              "memory": "number",
              "min_ready_instances": "number",
              "mount_desc": [
                "set",
                [
                  "object",
                  {
                    "mount_path": "string",
                    "nas_path": "string"
                  }
                ]
              ],
              "mount_host": "string",
              "namespace_id": "string",
              "nas_id": "string",
              "oss_ak_id": "string",
              "oss_ak_secret": "string",
              "oss_mount_descs": "string",
              "oss_mount_details": [
                "list",
                [
                  "object",
                  {
                    "bucket_name": "string",
                    "bucket_path": "string",
                    "mount_path": "string",
                    "read_only": "bool"
                  }
                ]
              ],
              "package_type": "string",
              "package_url": "string",
              "package_version": "string",
              "php_arms_config_location": "string",
              "php_config": "string",
              "php_config_location": "string",
              "post_start": "string",
              "pre_stop": "string",
              "readiness": "string",
              "region_id": "string",
              "replicas": "number",
              "repo_name": "string",
              "repo_namespace": "string",
              "repo_origin_type": "string",
              "security_group_id": "string",
              "sls_configs": "string",
              "status": "string",
              "tags": [
                "map",
                "string"
              ],
              "termination_grace_period_seconds": "number",
              "timezone": "string",
              "tomcat_config": "string",
              "vpc_id": "string",
              "vswitch_id": "string",
              "war_start_options": "string",
              "web_container": "string"
            }
          ]
        ]
      },
      "enable_details": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "field_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "field_value": {
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
      "ids": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "namespace_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "order_by": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "reverse": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
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

func AlicloudSaeApplicationsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudSaeApplications), &result)
	return &result
}
