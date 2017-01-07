// ReportService.go
package v1

import (
	"encoding/xml"
	"errors"
	// "fmt"
)

const (
	AsyncReportAccount   = "account"
	AsyncReportCampaingn = "campaign"
	AsyncReportAdgroup   = "adgroup"
	AsyncReportKeyword   = "keyword"
	AsyncReportCreative  = "creative"
)

type ReportService struct {
	*CommonService
}

func NewReportService() *ReportService {
	a := new(ReportService)
	a.CommonService = NewCommonService("ReportService")
	return a
}

//获取报表ID
func (r *ReportService) GetReportId(reportType, UnitOfTime, startDate, endDate string) (reportId string, err error) {
	GetReportIdRequestBody := new(GetReportIdRequestBody)
	GetReportIdRequestBody.GetReportIdRequest = new(GetReportIdRequest)
	GetReportIdRequestBody.GetReportIdRequest.ReportRequestType = GetAsyncReportType(NewAsyncReportRequestType(reportType, UnitOfTime, startDate, endDate))
	b, err := r.do(GetReportIdRequestBody)
	if err != nil {
		return "", err
	}
	var s SoapGetReportIdResponse
	err = xml.Unmarshal(b, &s)
	if err != nil {
		return "", err
	}
	if s.Header.ResHeader.Status != 0 {
		var errMsg = s.Header.ResHeader.Desc
		for _, v := range s.Header.ResHeader.Failures {
			errMsg += (":" + GetIntStr(v.Code) + ":" + v.Message)
		}
		return "", errors.New(errMsg)
	}
	return s.Body.GetReportIdResponse.ReportId, nil
}
func (r *ReportService) GetReportState(reportId string) (int, error) {
	GetReportStateRequestBody := new(GetReportStateRequestBody)
	GetReportStateRequestBody.GetReportStateRequest = new(GetReportStateRequest)
	GetReportStateRequestBody.GetReportStateRequest.ReportId = reportId
	b, err := r.do(GetReportStateRequestBody)
	if err != nil {
		return -1, err
	}
	var s SoapGetReportStateResponse
	err = xml.Unmarshal(b, &s)
	if err != nil {
		return -1, err
	}
	// 	1: 已完成//  0: 处理中// -1:报表生成异常
	if s.Header.ResHeader.Status != 0 {
		var errMsg = s.Header.ResHeader.Desc
		for _, v := range s.Header.ResHeader.Failures {
			errMsg += (":" + GetIntStr(v.Code) + ":" + v.Message)
		}
		return -1, errors.New(errMsg)
	}
	return s.Body.GetReportStateResponse.IsGenerated, nil
}
func (r *ReportService) GetReportPath(reportId string) (string, error) {
	GetReportPathRequestBody := new(GetReportPathRequestBody)
	GetReportPathRequestBody.GetReportPathRequest = new(GetReportPathRequest)
	GetReportPathRequestBody.GetReportPathRequest.ReportId = reportId
	b, err := r.do(GetReportPathRequestBody)
	if err != nil {
		return "", err
	}
	var s SoapGetReportPathResponse
	err = xml.Unmarshal(b, &s)
	if err != nil {
		return "", err
	}
	if s.Header.ResHeader.Status != 0 {
		var errMsg = s.Header.ResHeader.Desc
		for _, v := range s.Header.ResHeader.Failures {
			errMsg += (":" + GetIntStr(v.Code) + ":" + v.Message)
		}
		return "", errors.New(errMsg)
	}
	return s.Body.GetReportPathResponse.ReportFilePath, nil
}

type GetReportIdRequestBody struct {
	GetReportIdRequest *GetReportIdRequest `xml:"v11:getReportIdRequest"`
}
type GetReportIdRequest struct {
	ReportRequestType *ReportRequestType `xml:"reportRequestType"`
}

//定义报表请求
type ReportRequestType struct {
	PerformanceData []string `xml:"performanceData"`
	StartDate       string   `xml:"startDate"`
	EndDate         string   `xml:"endDate"`
	IdOnly          bool     `xml:"idOnly"`
	OptType         *OptType `xml:"opt"`
	Format          int      `xml:"format"`
	ReportType      int      `xml:"reportType"`
	StatIds         []int64  `xml:"statIds"`
	StatRange       int      `xml:"statRange"`
	UnitOfTime      int      `xml:"unitOfTime"`
	Platform        int      `xml:"platform"`
}

type SoapGetReportIdResponse struct {
	XMLName xml.Name                 `xml:"soapEnvelope"`
	Soap    string                   `xml:"xmlnssoap,attr"`
	Header  *HeaderResponse          `xml:"soapHeader"`
	Body    *GetReportIdResponseBody `xml:"soapBody"`
}
type GetReportIdResponseBody struct {
	Ns3                 string               `xml:"xmlnsns3,attr"`
	Ns2                 string               `xml:"xmlnsns2,attr"`
	GetReportIdResponse *GetReportIdResponse `xml:"ns2getReportIdResponse"`
}
type GetReportIdResponse struct {
	ReportId string `xml:"reportId"`
}
type GetReportStateRequestBody struct {
	GetReportStateRequest *GetReportStateRequest `xml:"v11:getReportStateRequest"`
}
type GetReportStateRequest struct {
	ReportId string `xml:"reportId"`
}
type SoapGetReportStateResponse struct {
	XMLName xml.Name                    `xml:"soapEnvelope"`
	Soap    string                      `xml:"xmlnssoap,attr"`
	Header  *HeaderResponse             `xml:"soapHeader"`
	Body    *GetReportStateResponseBody `xml:"soapBody"`
}
type GetReportStateResponseBody struct {
	Ns3                    string                  `xml:"xmlnsns3,attr"`
	Ns2                    string                  `xml:"xmlnsns2,attr"`
	GetReportStateResponse *GetReportStateResponse `xml:"ns2getReportStateResponse"`
}
type GetReportStateResponse struct {
	IsGenerated int `xml:"isGenerated"` // 	1: 已完成//  0: 处理中// -1:报表生成异常
}
type GetReportPathRequestBody struct {
	GetReportPathRequest *GetReportPathRequest `xml:"v11:getReportPathRequest"`
}
type GetReportPathRequest struct {
	ReportId string `xml:"reportId"`
}
type SoapGetReportPathResponse struct {
	XMLName xml.Name                   `xml:"soapEnvelope"`
	Soap    string                     `xml:"xmlnssoap,attr"`
	Header  *HeaderResponse            `xml:"soapHeader"`
	Body    *GetReportPathResponseBody `xml:"soapBody"`
}
type GetReportPathResponseBody struct {
	Ns3                   string                 `xml:"xmlnsns3,attr"`
	Ns2                   string                 `xml:"xmlnsns2,attr"`
	GetReportPathResponse *GetReportPathResponse `xml:"ns2getReportPathResponse"`
}
type GetReportPathResponse struct {
	ReportFilePath string `xml:"reportFilePath"`
}

type AsyncReportRequestType struct {
	ReportType string
	UnitOfTime int
	StartDate  string
	EndDate    string
	Platform   int
}

//构建请求报表结构
func NewAsyncReportRequestType(reportType, UnitOfTime, startDate, endDate string) *AsyncReportRequestType {
	return &AsyncReportRequestType{
		ReportType: reportType,
		StartDate:  startDate,
		EndDate:    endDate,
		UnitOfTime: GetUnitTime(UnitOfTime),
	}
}

//获取拉取报表
func GetAsyncReportType(asyncReportRequestType *AsyncReportRequestType) *ReportRequestType {
	r := new(ReportRequestType)
	r.PerformanceData = []string{"cost", "cpc", "click", "impression", "ctr"}
	r.StartDate = asyncReportRequestType.StartDate
	r.EndDate = asyncReportRequestType.EndDate
	r.Format = 1
	r.StatIds = []int64{}
	r.UnitOfTime = asyncReportRequestType.UnitOfTime
	r.Platform = 0
	r.StatRange = 1
	switch asyncReportRequestType.ReportType {
	case AsyncReportAccount:
		r.ReportType = 1
	case AsyncReportCampaingn:
		r.ReportType = 2
	case AsyncReportAdgroup:
		r.ReportType = 3
	case AsyncReportKeyword:
		r.ReportType = 4
	case AsyncReportCreative:
		r.ReportType = 5
	}
	return r

}
func GetUnitTime(ut string) int {
	switch ut {
	case "month":
		return 3
	case "weekday":
		return 2
	case "day":
		return 1
	default:
		return 1
	}
}
