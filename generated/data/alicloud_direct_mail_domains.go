package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudDirectMailDomains = `{
  "block": {
    "attributes": {
      "domains": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "cname_auth_status": "string",
              "cname_confirm_status": "string",
              "cname_record": "string",
              "create_time": "string",
              "default_domain": "string",
              "dkim_auth_status": "string",
              "dkim_public_key": "string",
              "dkim_rr": "string",
              "dmarc_auth_status": "string",
              "dmarc_host_record": "string",
              "dmarc_record": "string",
              "dns_dmarc": "string",
              "dns_mx": "string",
              "dns_spf": "string",
              "dns_txt": "string",
              "domain_id": "string",
              "domain_name": "string",
              "domain_record": "string",
              "domain_type": "string",
              "host_record": "string",
              "icp_status": "string",
              "id": "string",
              "mx_auth_status": "string",
              "mx_record": "string",
              "spf_auth_status": "string",
              "spf_record": "string",
              "status": "string",
              "tl_domain_name": "string",
              "tracef_record": "string"
            }
          ]
        ]
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
      "key_word": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudDirectMailDomainsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudDirectMailDomains), &result)
	return &result
}
