package data_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-alicloud-schema/generated/data"
	"github.com/stretchr/testify/assert"
)

func TestAlicloudArmsAlertContactsSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := data.AlicloudArmsAlertContactsSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
