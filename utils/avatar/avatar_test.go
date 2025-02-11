package avatar_test

import (
	"fmt"
	"testing"

	"pluto/utils/avatar"

	"github.com/stretchr/testify/assert"
)

// this one needs network and it's very slow, turn it off
// func TestGenFromGravatar(t *testing.T) {
// 	ag := avatar.AvatarGen{}
// 	ar, err := ag.GenFromGravatar()
// 	if err != nil {
// 		assert.FailNow(t, "Except no error, but %s", err.LogError.Error())
// 	}
// 	url := ar.OriginURL

// 	assert.NotEqual(t, url, "", "avatar url should not be empty")
// }

const testImageBase64 = "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"

func TestGenFromBase64String(t *testing.T) {
	ag := avatar.AvatarGen{}
	ar, err := ag.GenFromBase64String(testImageBase64)
	if err != nil {
		assert.FailNow(t, "Except no error, but %s", err.LogError.Error())
	}
	fmt.Println("avatar image type: " + ar.Ext)
	assert.NotEqual(t, ar.Ext, "", "avatar url should not be empty")
}
