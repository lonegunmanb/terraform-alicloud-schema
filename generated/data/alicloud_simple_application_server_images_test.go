package data_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-alicloud-schema/generated/data"
	"github.com/stretchr/testify/assert"
)

func TestAlicloudSimpleApplicationServerImagesSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := data.AlicloudSimpleApplicationServerImagesSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
