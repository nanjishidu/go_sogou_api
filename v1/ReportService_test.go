// Sogou_API_Core_test.go
package v1

import (
	"fmt"
	"testing"
	"time"
)

var authHeader = &AuthHeader{
	Username: "",
	Password: "",
	Token:    "",
}

//定义 获取报表
var reportLists = []string{
	AsyncReportAccount,
	AsyncReportCampaingn,
	AsyncReportAdgroup,
	AsyncReportKeyword,
	AsyncReportCreative,
}

func TestGetReportId(t *testing.T) {
	m := NewReportService()
	m.AuthHeader = authHeader
	for _, v := range reportLists {
		b, err := m.GetReportId(v, "day", "2016-08-26T00:00:00", "2016-08-27T00:00:00")
		fmt.Println(b, err)
		if err != nil {
			fmt.Println(b)
			return
		}
		var serverFail bool
		for {
			time.Sleep(1 * time.Minute)
			s, err := m.GetReportState(b)
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
			p, err := m.GetReportPath(b)
			fmt.Println(p, err)
		}
		if p == "" {
			t.Log("GetReportPath failed")
		}
	}
}
