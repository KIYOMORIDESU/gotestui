package example

import "testing"

func TestExample(t *testing.T) {
	t.Run("example test 01", func(t *testing.T) {
		t.Run("pass", func(t *testing.T) {
      t.Log("Passlog01")
      t.Log("Passlog02")
      t.Log("Passlog03")
      t.Log("Passlog04")
    })
		t.Run("fail", func(t *testing.T) {
      t.Log("Faillog01")
      t.Log("Faillog02")
      t.Log("Faillog03")
      t.Log("Faillog04")
      t.Fail()
    })
		t.Run("skip", func(t *testing.T) {
      t.Log("Skiplog01")
      t.Log("Skiplog02")
      t.Log("Skiplog03")
      t.Log("Skiplog04")
      t.Skip()
    })
	})
}
