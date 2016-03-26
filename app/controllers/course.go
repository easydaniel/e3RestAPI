package controllers

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"

	"reflect"

	"github.com/revel/revel"
)

const (
	crsListUrl    = baseUrl + "/mService/service.asmx/GetCourseList"
	crsInfoUrl    = baseUrl + "/mService/service.asmx/GetCourseInfo"
	crsTimeUrl    = baseUrl + "/mService/service.asmx/GetCourseTime"
	crsContactUrl = baseUrl + "/mService/service.asmx/GetCourseUserContactInfo"
)

type Course struct {
	*revel.Controller
}

type CourseDataParser struct {
	CourseList []CourseData `xml:"CourseData"`
}

type CourseData struct {
	CourseId       string
	PmtCrsNo       string
	CourseNo       string
	CourseName     string
	CrsYear        string
	CrsSemester    string
	TeacherName    string
	ForumEnable    bool
	OldCourse      bool
	UserRole       string
	PositionStatus string
}

type CourseTimeParser struct {
	TimeList []CourseTimeData `xml:"CourseTimeData"`
}

type CourseTimeData struct {
	CourseId   string
	WeekDay    string
	Section    string
	RoomNo     string
	CourseName string
}

type ContactInfoParser struct {
	InfoList []ContactInfoData `xml:"ContactInfoData"`
}

type ContactInfoData struct {
	AccountId   string
	AccountName string
	EMail       string
}

func (c Course) List() revel.Result {

	LoginTicket := c.Params.Get("loginTicket")
	AccountId := c.Params.Get("accountId")
	Role := c.Params.Get("role")

	resp, err := http.Get(crsListUrl + "?loginTicket=" + LoginTicket + "&accountId=" + AccountId + "&role=" + Role)
	if err != nil {
		revel.WARN.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		revel.WARN.Println(err)
	}

	var parser CourseDataParser
	xml.Unmarshal(body, &parser)

	if reflect.DeepEqual(parser, (CourseDataParser{})) {
		c.Response.Status = 405
	}

	return c.RenderJson(parser.CourseList)
}

func (c Course) Info() revel.Result {

	LoginTicket := c.Params.Get("loginTicket")
	CourseId := c.Params.Get("courseId")

	resp, err := http.Get(crsInfoUrl + "?loginTicket=" + LoginTicket + "&courseId=" + CourseId)
	if err != nil {
		revel.WARN.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		revel.WARN.Println(err)
	}

	var parser CourseData
	xml.Unmarshal(body, &parser)

	if reflect.DeepEqual(parser, (CourseData{})) {
		c.Response.Status = 405
	}

	return c.RenderJson(parser)
}

func (c Course) Time() revel.Result {

	LoginTicket := c.Params.Get("loginTicket")
	CourseId := c.Params.Get("courseId")

	resp, err := http.Get(crsTimeUrl + "?loginTicket=" + LoginTicket + "&courseId=" + CourseId)
	if err != nil {
		revel.WARN.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		revel.WARN.Println(err)
	}

	var parser CourseTimeParser
	xml.Unmarshal(body, &parser)

	if reflect.DeepEqual(parser, (CourseTimeParser{})) {
		c.Response.Status = 405
	}

	return c.RenderJson(parser.TimeList)
}

func (c Course) Contact() revel.Result {

	LoginTicket := c.Params.Get("loginTicket")
	CourseId := c.Params.Get("courseId")
	UserRole := c.Params.Get("userRole")

	resp, err := http.Get(crsContactUrl + "?loginTicket=" + LoginTicket + "&courseId=" + CourseId + "&userRole=" + UserRole)
	if err != nil {
		revel.WARN.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		revel.WARN.Println(err)
	}

	var parser ContactInfoParser
	xml.Unmarshal(body, &parser)

	if reflect.DeepEqual(parser, (ContactInfoParser{})) {
		c.Response.Status = 405
	}

	return c.RenderJson(parser.InfoList)
}
