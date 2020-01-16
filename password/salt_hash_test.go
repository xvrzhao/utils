package password

import "testing"

func TestVerifySaltHashPwd(t *testing.T) {
	p := "my_pass"
	sp, s := SaltHashPwd(p, 8)

	t.Log(p, s, sp)

	if b := VerifySaltHashPwd(p, s, sp); b != true {
		t.Fail()
		return
	}
}
