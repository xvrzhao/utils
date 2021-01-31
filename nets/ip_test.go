package nets

import "testing"

func TestGetIntranetIP(t *testing.T) {
	IP, err := GetIntranetIP()
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("intranet ip: %s", IP)
}
