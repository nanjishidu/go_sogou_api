// Sogou_API_Core_test.go
package v1

import (
	"fmt"
	"testing"
	// "time"
)

func TestGetAccountFile(t *testing.T) {
	m := NewAccountDownloadService()
	m.AuthHeader = authHeader
	b, err := m.GetAccountFile()
	if err != nil {
		fmt.Println(b)
		return
	}
	var serverFail bool
	for {
		s, err := m.GetAccountFileStatus(b)
		if err != nil {
			t.Log(err)
			return
		}
		if s == 1 {
			break
		}
		if s == -1 {
			serverFail = true
			break
		}
		fmt.Println(s, err)

	}
	var p = ""
	if !serverFail {
		p, err := m.GetAccountFilePath(b)
		fmt.Println(p, err)
	}
	if p == "" {
		t.Log("GetReportPath failed")
	}
}
