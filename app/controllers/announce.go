package controllers

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"

	"github.com/revel/revel"
)

const (
	annListUrl      = baseUrl + "/mService/service.asmx/GetAnnouncementListWithAttach"
	annInfoUrl      = baseUrl + "/mService/service.asmx/GetAnnouncementDetail"
	annLoginListUrl = baseUrl + "/mService/service.asmx/GetAnnouncementList_LoginWithAttach"
)

type Announce struct {
	*revel.Controller
}

type BulletinDataParser struct {
	BulletinDataList []BulletinData `xml:"BulletinData"`
}

type BulletinData struct {
	BulletinId            string
	CourseId              string
	BulType               int
	Caption               string
	Content               string
	BeginDate             string
	EndDate               string
	NotifyType            int
	NeedVote              int
	ExpiredStatus         int
	CourseName            string
	FileName              string
	AttachFileName        []string `xml:"AttachFileName>string"`
	AttachFileURL         []string `xml:"AttachFileURL>string"`
	AttachFileMimeType    []string `xml:"AttachFileMimeType>string"`
	AttachFileFileSize    []string `xml:"AttachFileFileSize>string"`
	AttachFileDescription []string `xml:"AttachFileDescription>string"`
	AttachFileCreateTime  []string `xml:"AttachFileCreateTime>string"`
}

func (c Announce) List() revel.Result {

	LoginTicket := c.Params.Get("loginTicket")
	CourseId := c.Params.Get("courseId")
	BulType := c.Params.Get("bulType")

	resp, err := http.Get(annListUrl + "?loginTicket=" + LoginTicket + "&courseId=" + CourseId + "&bulType=" + BulType)
	if err != nil {
		revel.WARN.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		revel.WARN.Println(err)
	}

	var parser BulletinDataParser
	xml.Unmarshal(body, &parser)

	return c.RenderJson(parser.BulletinDataList)
}

func (c Announce) Info() revel.Result {

	LoginTicket := c.Params.Get("loginTicket")
	BulletinId := c.Params.Get("bulletinId")

	resp, err := http.Get(annInfoUrl + "?loginTicket=" + LoginTicket + "&bulletinId=" + BulletinId)
	if err != nil {
		revel.WARN.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		revel.WARN.Println(err)
	}

	var parser BulletinData
	xml.Unmarshal(body, &parser)

	return c.RenderJson(parser)
}

func (c Announce) LoginList(count string) revel.Result {

	LoginTicket := c.Params.Get("loginTicket")
	StudentId := c.Params.Get("studentId")

	getParam := "?loginTicket=" + LoginTicket + "&studentId=" + StudentId

	if count != "" {
		getParam += "&ShowCount=" + count
	}

	resp, err := http.Get(annLoginListUrl + getParam)
	if err != nil {
		revel.WARN.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		revel.WARN.Println(err)
	}

	var parser BulletinDataParser
	xml.Unmarshal(body, &parser)

	return c.RenderJson(parser.BulletinDataList)
}
