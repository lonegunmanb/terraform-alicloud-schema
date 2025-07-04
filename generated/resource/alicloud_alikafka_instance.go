package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudAlikafkaInstance = `{
  "block": {
    "attributes": {
      "config": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "default_topic_partition_num": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "deploy_type": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "disk_size": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "disk_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "domain_endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "eip_max": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "enable_auto_group": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "enable_auto_topic": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "end_point": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "group_left": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "group_used": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instance_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "io_max": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "io_max_spec": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "is_partition_buy": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "kms_key_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "paid_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "partition_left": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "partition_num": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "partition_used": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "password": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_group_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "sasl_domain_endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "security_group": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "selected_zones": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "service_version": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "spec_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ssl_domain_endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "ssl_endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "topic_left": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "topic_num_of_buy": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "topic_quota": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "topic_used": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "vpc_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "vswitch_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "vswitch_ids": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "zone_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "confluent_config": {
        "block": {
          "attributes": {
            "connect_cu": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "connect_replica": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "control_center_cu": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "control_center_replica": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "control_center_storage": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "kafka_cu": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "kafka_replica": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "kafka_rest_proxy_cu": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "kafka_rest_proxy_replica": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "kafka_storage": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "ksql_cu": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "ksql_replica": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "ksql_storage": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "schema_registry_cu": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "schema_registry_replica": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "zookeeper_cu": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "zookeeper_replica": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "zookeeper_storage": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "serverless_config": {
        "block": {
          "attributes": {
            "reserved_publish_capacity": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "reserved_subscribe_capacity": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
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

func AlicloudAlikafkaInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudAlikafkaInstance), &result)
	return &result
}
