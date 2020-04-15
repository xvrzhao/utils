package crypto

import "testing"

func TestMd5(t *testing.T) {
	if Md5("xvrzhao") != "6452adf3f17e11c592f0c3e7e0afa302" {
		t.Error("md5 failed")
	}
}
