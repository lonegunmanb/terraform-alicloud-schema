package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEciContainerGroups = `{
  "block": {
    "attributes": {
      "container_group_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "enable_details": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "groups": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "container_group_id": "string",
              "container_group_name": "string",
              "containers": [
                "list",
                [
                  "object",
                  {
                    "args": [
                      "list",
                      "string"
                    ],
                    "commands": [
                      "list",
                      "string"
                    ],
                    "cpu": "number",
                    "environment_vars": [
                      "list",
                      [
                        "object",
                        {
                          "key": "string",
                          "value": "string"
                        }
                      ]
                    ],
                    "gpu": "number",
                    "image": "string",
                    "image_pull_policy": "string",
                    "memory": "number",
                    "name": "string",
                    "ports": [
                      "list",
                      [
                        "object",
                        {
                          "port": "number",
                          "protocol": "string"
                        }
                      ]
                    ],
                    "ready": "bool",
                    "restart_count": "number",
                    "volume_mounts": [
                      "list",
                      [
                        "object",
                        {
                          "mount_path": "string",
                          "name": "string",
                          "read_only": "bool"
                        }
                      ]
                    ],
                    "working_dir": "string"
                  }
                ]
              ],
              "cpu": "number",
              "discount": "number",
              "dns_config": [
                "list",
                [
                  "object",
                  {
                    "name_servers": [
                      "list",
                      "string"
                    ],
                    "options": [
                      "list",
                      [
                        "object",
                        {
                          "name": "string",
                          "value": "string"
                        }
                      ]
                    ],
                    "searches": [
                      "list",
                      "string"
                    ]
                  }
                ]
              ],
              "eci_security_context": [
                "list",
                [
                  "object",
                  {
                    "sysctls": [
                      "list",
                      [
                        "object",
                        {
                          "name": "string",
                          "value": "string"
                        }
                      ]
                    ]
                  }
                ]
              ],
              "eni_instance_id": "string",
              "events": [
                "list",
                [
                  "object",
                  {
                    "count": "number",
                    "first_timestamp": "string",
                    "last_timestamp": "string",
                    "message": "string",
                    "name": "string",
                    "reason": "string",
                    "type": "string"
                  }
                ]
              ],
              "expired_time": "string",
              "failed_time": "string",
              "host_aliases": [
                "list",
                [
                  "object",
                  {
                    "hostnames": [
                      "list",
                      "string"
                    ],
                    "ip": "string"
                  }
                ]
              ],
              "id": "string",
              "init_containers": [
                "list",
                [
                  "object",
                  {
                    "args": [
                      "list",
                      "string"
                    ],
                    "commands": [
                      "list",
                      "string"
                    ],
                    "cpu": "number",
                    "environment_vars": [
                      "list",
                      [
                        "object",
                        {
                          "key": "string",
                          "value": "string"
                        }
                      ]
                    ],
                    "gpu": "number",
                    "image": "string",
                    "image_pull_policy": "string",
                    "memory": "number",
                    "name": "string",
                    "ports": [
                      "list",
                      [
                        "object",
                        {
                          "port": "number",
                          "protocol": "string"
                        }
                      ]
                    ],
                    "ready": "bool",
                    "restart_count": "number",
                    "volume_mounts": [
                      "list",
                      [
                        "object",
                        {
                          "mount_path": "string",
                          "name": "string",
                          "read_only": "bool"
                        }
                      ]
                    ],
                    "working_dir": "string"
                  }
                ]
              ],
              "instance_type": "string",
              "internet_ip": "string",
              "intranet_ip": "string",
              "ipv6_address": "string",
              "memory": "number",
              "ram_role_name": "string",
              "resource_group_id": "string",
              "restart_policy": "string",
              "security_group_id": "string",
              "status": "string",
              "succeeded_time": "string",
              "tags": [
                "map",
                "string"
              ],
              "volumes": [
                "list",
                [
                  "object",
                  {
                    "config_file_volume_config_file_to_paths": [
                      "list",
                      [
                        "object",
                        {
                          "content": "string",
                          "path": "string"
                        }
                      ]
                    ],
                    "disk_volume_disk_id": "string",
                    "disk_volume_fs_type": "string",
                    "flex_volume_driver": "string",
                    "flex_volume_fs_type": "string",
                    "flex_volume_options": "string",
                    "name": "string",
                    "nfs_volume_path": "string",
                    "nfs_volume_read_only": "bool",
                    "nfs_volume_server": "string",
                    "type": "string"
                  }
                ]
              ],
              "vpc_id": "string",
              "vswitch_id": "string",
              "zone_id": "string"
            }
          ]
        ]
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
      "limit": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
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
      "resource_group_id": {
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
      },
      "vswitch_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "with_event": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "zone_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudEciContainerGroupsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEciContainerGroups), &result)
	return &result
}
