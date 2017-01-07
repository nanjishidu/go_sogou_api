// AccountDownloadService
package v1

//整账户下载&跨层级下载服务
// getAccountFile
import (
	"encoding/xml"
	"errors"
	"fmt"
)

type AccountDownloadService struct {
	*CommonService
}

func NewAccountDownloadService() *AccountDownloadService {
	a := new(AccountDownloadService)
	a.CommonService = NewCommonService("AccountDownloadService")
	return a
}

//获取报表ID
func (r *AccountDownloadService) GetAccountFile() (accountFileId string, err error) {
	GetAccountFileRequestBody := new(GetAccountFileRequestBody)
	GetAccountFileRequestBody.GetAccountFileRequest = new(GetAccountFileRequest)
	GetAccountFileRequestBody.GetAccountFileRequest.AccoutFileRequest = new(AccoutFileRequest)
	b, err := r.do(GetAccountFileRequestBody)
	if err != nil {
		return "", err
	}

	var s SoapGetAccountFileResponse
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
	return s.Body.GetAccountFileResponse.AccountFileId, nil
}
func (r *AccountDownloadService) GetAccountFileStatus(accountFileId string) (int, error) {
	GetAccountFileStatusRequestBody := new(GetAccountFileStatusRequestBody)
	GetAccountFileStatusRequestBody.GetAccountFileStatusRequest = new(GetAccountFileStatusRequest)
	GetAccountFileStatusRequestBody.GetAccountFileStatusRequest.AccountFileId = accountFileId
	b, err := r.do(GetAccountFileStatusRequestBody)
	fmt.Println(string(b), err)
	if err != nil {
		return -1, err
	}
	var s SoapGetAccountFileStatusResponse
	err = xml.Unmarshal(b, &s)
	if err != nil {
		return -1, err
	}
	if s.Header.ResHeader.Status != 0 {
		var errMsg = s.Header.ResHeader.Desc
		for _, v := range s.Header.ResHeader.Failures {
			errMsg += (":" + GetIntStr(v.Code) + ":" + v.Message)
		}
		return -1, errors.New(errMsg)
	}
	return s.Body.GetAccountFileStatusResponse.IsGenerated, nil
}
func (r *AccountDownloadService) GetAccountFilePath(accountFileId string) (string, error) {
	GetAccountFilePathRequestBody := new(GetAccountFilePathRequestBody)
	GetAccountFilePathRequestBody.GetAccountFilePathRequest = new(GetAccountFilePathRequest)
	GetAccountFilePathRequestBody.GetAccountFilePathRequest.AccountFileId = accountFileId
	b, err := r.do(GetAccountFilePathRequestBody)
	if err != nil {
		return "", err
	}
	var s SoapGetAccountFilePathResponse
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
	return s.Body.GetAccountFilePathResponse.AccountFilePath, nil
}

type GetAccountFileRequestBody struct {
	GetAccountFileRequest *GetAccountFileRequest `xml:"v11:getAccountFileRequest"`
}
type GetAccountFileRequest struct {
	AccoutFileRequest *AccoutFileRequest `xml:"accoutFileRequest"`
}
type AccoutFileRequest struct {
	CpcPlanIds     []int64  `xml:"cpcPlanIds"`
	IncludeQuality bool     `xml:"includeQuality"`
	IncludeTemp    bool     `xml:"includeTemp"`
	Format         int      `xml:"format"`
	OptType        *OptType `xml:"opt"`
}

type SoapGetAccountFileResponse struct {
	XMLName xml.Name            `xml:"soapEnvelope"`
	Soap    string              `xml:"xmlnssoap,attr"`
	Header  *HeaderResponse     `xml:"soapHeader"`
	Body    *GetAccountFileBody `xml:"soapBody"`
}
type GetAccountFileBody struct {
	Ns3                    string                  `xml:"xmlnsns3,attr"`
	Ns2                    string                  `xml:"xmlnsns2,attr"`
	GetAccountFileResponse *GetAccountFileResponse `xml:"ns2getAccountFileResponse"`
}
type GetAccountFileResponse struct {
	AccountFileId string `xml:"accountFileId"`
}

type GetAccountFileStatusRequestBody struct {
	GetAccountFileStatusRequest *GetAccountFileStatusRequest `xml:"v11:getAccountFileStatusRequest"`
}
type GetAccountFileStatusRequest struct {
	AccountFileId string `xml:"accountFileId"`
}

type SoapGetAccountFileStatusResponse struct {
	XMLName xml.Name                          `xml:"soapEnvelope"`
	Soap    string                            `xml:"xmlnssoap,attr"`
	Header  *HeaderResponse                   `xml:"soapHeader"`
	Body    *GetAccountFileStatusResponseBody `xml:"soapBody"`
}
type GetAccountFileStatusResponseBody struct {
	Ns3                          string                        `xml:"xmlnsns3,attr"`
	Ns2                          string                        `xml:"xmlnsns2,attr"`
	GetAccountFileStatusResponse *GetAccountFileStatusResponse `xml:"ns2getAccountFileStatusResponse"`
}

//1: 已完成
// 0: 处理中
// -1:下载异常
type GetAccountFileStatusResponse struct {
	IsGenerated int `xml:"isGenerated"`
}

type GetAccountFilePathRequestBody struct {
	GetAccountFilePathRequest *GetAccountFilePathRequest `xml:"v11:getAccountFilePathRequest"`
}
type GetAccountFilePathRequest struct {
	AccountFileId string `xml:"accountFileId"`
}

type SoapGetAccountFilePathResponse struct {
	XMLName xml.Name                        `xml:"soapEnvelope"`
	Soap    string                          `xml:"xmlnssoap,attr"`
	Header  *HeaderResponse                 `xml:"soapHeader"`
	Body    *GetAccountFilePathResponseBody `xml:"soapBody"`
}
type GetAccountFilePathResponseBody struct {
	Ns3                        string                      `xml:"xmlnsns3,attr"`
	Ns2                        string                      `xml:"xmlnsns2,attr"`
	GetAccountFilePathResponse *GetAccountFilePathResponse `xml:"ns2getAccountFilePathResponse"`
}

type GetAccountFilePathResponse struct {
	AccountFilePath string `xml:"accountFilePath"`
}
