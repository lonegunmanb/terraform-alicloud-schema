package data_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-alicloud-schema/generated/data"
	"github.com/stretchr/testify/assert"
)

func TestAlicloudEciVirtualNodesSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := data.AlicloudEciVirtualNodesSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
