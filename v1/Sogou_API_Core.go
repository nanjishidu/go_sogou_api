// Sogou_API_Core.go
package v1

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	url = "http://api.sogou.com/sem/common/v1"
)

type CommonService struct {
	Url        string
	Soapenv    string
	V1         string
	V11        string
	AuthHeader *AuthHeader
}

func NewCommonService(servicename string) *CommonService {
	c := new(CommonService)
	c.Soapenv = "http://schemas.xmlsoap.org/soap/envelope/"
	c.V1 = "http://api.sogou.com/sem/common/v1"
	c.V11 = "https://api.sogou.com/sem/sms/v1"
	c.Url = "http://api.agent.sogou.com/sem/sms/v1/" + servicename + "/?wsdl"
	c.AuthHeader = new(AuthHeader)
	return c
}

//执行post的请求，json 交互 获取 返回数据
func (c *CommonService) do(reuestBody interface{}) (result []byte, err error) {
	//构建URL
	client := &http.Client{}
	s := &SoapGetRequest{
		Soapenv: c.Soapenv,
		V1:      c.V1,
		V11:     c.V11,
		Header:  new(HeaderRequest),
	}
	s.Header.AuthHeader = c.AuthHeader
	s.Body = reuestBody
	b, err := xml.Marshal(s)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(b))
	res, err := client.Post(c.Url, "text/xml; charset=utf-8", bytes.NewBuffer([]byte(b)))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	result, err = ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}
	var replace = map[string]string{
		"soap:":  "soap",
		"xmlns:": "xmlns",
		"ns3:":   "ns3",
		"ns2:":   "ns2",
	}
	var sb = string(result)
	for k, v := range replace {
		sb = strings.Replace(sb, k, v, -1)
	}
	return []byte(sb), nil
}

//定义请求数据结构
type SoapGetRequest struct {
	XMLName xml.Name       `xml:"soapenv:Envelope"`
	Soapenv string         `xml:"xmlns:soapenv,attr"`
	V1      string         `xml:"xmlns:v1,attr"`
	V11     string         `xml:"xmlns:v11,attr"`
	Header  *HeaderRequest `xml:"soapenv:Header"`
	Body    interface{}    `xml:"soapenv:Body"`
}

//定义验证AuthHeader
type HeaderRequest struct {
	AuthHeader *AuthHeader `xml:"v1:AuthHeader"`
}
type AuthHeader struct {
	Agentname     string `xml:"v1:agentusername"`
	Agentpassword string `xml:"v1:agentpassword"`
	Username      string `xml:"v1:username"`
	Password      string `xml:"v1:password"`
	Token         string `xml:"v1:token"`
}

//定义返回数据结构

type HeaderResponse struct {
	Ns3       string     `xml:"xmlnsns3,attr"`
	Ns2       string     `xml:"xmlnsns2,attr"`
	ResHeader *ResHeader `xml:"ns3ResHeader"`
}
type ResHeader struct {
	Desc     string      `xml:"ns3desc"`
	Failures []*Failures `xml:"ns3failures"`
	Oprs     int         `xml:"ns3oprs"`
	Oprtime  int         `xml:"ns3oprtime"`
	Quota    int         `xml:"ns3quota"`
	Rquota   int         `xml:"ns3rquota"`
	Status   int         `xml:"ns3status"`
}
type Failures struct {
	Code     int    `xml:"ns3code"`
	Message  string `xml:"ns3message"`
	Position string `xml:"ns3position"`
	Content  string `xml:"ns3content"`
}
type OptType struct {
	optString *StringMapItemType `xml:"optString"`
	optInt    *IntMapItemType    `xml:"optInt"`
	optLong   *LongMapItemType   `xml:"optLong"`
	optFloat  *FloatMapItemType  `xml:"optFloat"`
	optDouble *DoubleMapItemType `xml:"optDouble"`
}
type StringMapItemType struct {
	Key   string `xml:"key"`
	Value string `xml:"value"`
}
type IntMapItemType struct {
	Key   string `xml:"key"`
	Value int    `xml:"value"`
}
type LongMapItemType struct {
	Key   string `xml:"key"`
	Value int64  `xml:"value"`
}
type FloatMapItemType struct {
	Key   string  `xml:"key"`
	Value float32 `xml:"value"`
}
type DoubleMapItemType struct {
	Key   string  `xml:"key"`
	Value float64 `xml:"value"`
}

//<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
// <soap:Header>
// <ns3:ResHeader xmlns:ns3="http://api.sogou.com/sem/common/v1" xmlns:ns2="https://api.sogou.com/sem/sms/v1">
// 	<ns3:desc>success</ns3:desc>
// 	<ns3:oprs>1</ns3:oprs>
// 	<ns3:oprtime>0</ns3:oprtime>
// 	<ns3:quota>0</ns3:quota>
// 	<ns3:rquota>1713800</ns3:rquota>
// 	<ns3:status>0</ns3:status>
// </ns3:ResHeader>
// </soap:Header>

// <soap:Header>
// <ns3:ResHeader xmlns:ns3="http://api.sogou.com/sem/common/v1" xmlns:ns2="https://api.sogou.com/sem/sms/v1">
// <ns3:desc>failure</ns3:desc>
// <ns3:failures>
// 	<ns3:code>6</ns3:code>
// 	<ns3:message>Username is invalid</ns3:message>
// 	<ns3:position>_user</ns3:position>
// 	<ns3:content></ns3:content>
// </ns3:failures>
// <ns3:oprs>0</ns3:oprs>
// <ns3:quota>0</ns3:quota>
// <ns3:status>2</ns3:status>
// </ns3:ResHeader>
// </soap:Header>
