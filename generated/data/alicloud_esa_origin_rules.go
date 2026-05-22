package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEsaOriginRules = `{
  "block": {
    "attributes": {
      "config_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "config_type": {
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
      "rule_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "rules": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "config_id": "string",
              "config_type": "string",
              "dns_record": "string",
              "follow302_enable": "string",
              "follow302_max_tries": "string",
              "follow302_retain_args": "string",
              "follow302_retain_header": "string",
              "follow302_target_host": "string",
              "id": "string",
              "origin_host": "string",
              "origin_http_port": "string",
              "origin_https_port": "string",
              "origin_mtls": "string",
              "origin_read_timeout": "string",
              "origin_scheme": "string",
              "origin_sni": "string",
              "origin_verify": "string",
              "range": "string",
              "range_chunk_size": "string",
              "rule": "string",
              "rule_enable": "string",
              "rule_name": "string",
              "sequence": "number",
              "site_version": "number"
            }
          ]
        ]
      },
      "site_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "site_version": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudEsaOriginRulesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEsaOriginRules), &result)
	return &result
}
