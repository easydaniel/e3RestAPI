package controllers

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"reflect"

	"github.com/revel/revel"
)

const (
	hwListUrl = baseUrl + "/mService/service.asmx/GetStuHomeworkListWithAttach"
	hwInfoUrl = baseUrl + "/mService/service.asmx/GetHomeworkInfo"
)

type Homework struct {
	*revel.Controller
}

type HomeworkDataParser struct {
	HomeworkDataList []HomeworkData `xml:"HomeworkData"`
}

type HomeworkData struct {
	HomeworkId            string
	CourseId              string
	DisplayName           string
	BeginDate             string
	EndDate               string
	SubmitType            string
	UnitName              string
	Content               string
	AttachFile            string
	TimeoutProcess        string
	ReSubmitProcess       string
	IsStuEvaluate         string
	HwkSubmitId           string
	CanSubmit             bool
	NoSubmitCount         int
	SubmitCount           int
	ListType              int
	ScoreRatioType        int
	AttachFileName        []string `xml:"AttachFileName>string"`
	AttachFileURL         []string `xml:"AttachFileURL>string"`
	AttachFileMimeType    []string `xml:"AttachFileMimeType>string"`
	AttachFileFileSize    []string `xml:"AttachFileFileSize>string"`
	AttachFileDescription []string `xml:"AttachFileDescription>string"`
	AttachFileCreateTime  []string `xml:"AttachFileCreateTime>string"`
}

type StuHomeworkDataParser struct {
	StuHomeworkDataList []StuHomeworkData `xml:"StuHomeworkData"`
}

type StuHomeworkData struct {
	HomeworkId                   string
	CourseId                     string
	DisplayName                  string
	BeginDate                    string
	EndDate                      string
	SubmitType                   string
	UnitName                     string
	Content                      string
	AttachFile                   string
	TimeoutProcess               string
	ReSubmitProcess              string
	IsStuEvaluate                string
	HwkSubmitId                  string
	CanSubmit                    bool
	NoSubmitCount                int
	SubmitCount                  int
	ListType                     int
	ScoreRatioType               int
	AttachFileName               []string `xml:"AttachFileName>string"`
	AttachFileURL                []string `xml:"AttachFileURL>string"`
	AttachFileMimeType           []string `xml:"AttachFileMimeType>string"`
	AttachFileFileSize           []string `xml:"AttachFileFileSize>string"`
	AttachFileDescription        []string `xml:"AttachFileDescription>string"`
	AttachFileCreateTime         []string `xml:"AttachFileCreateTime>string"`
	SubjectAttachFileName        []string `xml:"SubjectAttachFileName>string"`
	SubjectAttachFileURL         []string `xml:"SubjectAttachFileURL>string"`
	SubjectAttachFileMimeType    []string `xml:"SubjectAttachFileMimeType>string"`
	SubjectAttachFileFileSize    []string `xml:"SubjectAttachFileFileSize>string"`
	SubjectAttachFileDescription []string `xml:"SubjectAttachFileDescription>string"`
	SubjectAttachFileCreateTime  []string `xml:"AttachSubjectAttachFileCreateTimeFileName>string"`
}

func (c Homework) List() revel.Result {

	LoginTicket := c.Params.Get("loginTicket")
	CourseId := c.Params.Get("courseId")
	AccountId := c.Params.Get("accountId")
	ListType := c.Params.Get("listType")

	resp, err := http.Get(hwListUrl + "?loginTicket=" + LoginTicket + "&courseId=" + CourseId + "&accountId=" + AccountId + "&listType=" + ListType)
	if err != nil {
		revel.WARN.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		revel.WARN.Println(err)
	}

	var parser StuHomeworkDataParser
	xml.Unmarshal(body, &parser)

	if reflect.DeepEqual(parser, (StuHomeworkDataParser{})) {
		c.Response.Status = 405
	}

	return c.RenderJson(parser.StuHomeworkDataList)
}

func (c Homework) Info() revel.Result {

	LoginTicket := c.Params.Get("loginTicket")
	AccountId := c.Params.Get("accountId")
	HwkId := c.Params.Get("hwkId")

	resp, err := http.Get(hwInfoUrl + "?loginTicket=" + LoginTicket + "&accountId=" + AccountId + "&hwkId=" + HwkId)
	if err != nil {
		revel.WARN.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		revel.WARN.Println(err)
	}

	var parser HomeworkData
	xml.Unmarshal(body, &parser)

	if reflect.DeepEqual(parser, (HomeworkData{})) {
		c.Response.Status = 405
	}

	return c.RenderJson(parser)
}
