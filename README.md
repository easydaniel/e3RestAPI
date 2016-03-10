# Rest API for E3

### Set environment variable for GO and get revel
    mkdir <projectname>
    cd <projectname>
    export GOPATH=`pwd`
    export PATH="$PATH:$GOPATH/bin"
    go get github.com/revel/cmd/revel

### Get project
    git clone http://github.com/easydaniel/e3RestAPI $GOPATH/src/e3RestAPI
  - Your project will be in $GOPATH/src/e3RestAPI

### Runserver
    PORT=<PORT> revel run e3RestAPI

### Todo List
- [ ] AddAnnouncement
- [ ] AddBoardMessage
- [ ] AddBoardMessage2
- [ ] AddDiscussForum
- [ ] AddDiscussReply
- [ ] AddDiscussTopic
- [ ] DeleteDiscussReply
- [ ] DeleteDiscussTopic
- [ ] DeleteQuestion
- [ ] DeleteQuiz
- [ ] ForgotPassword
- [x] GetAnnouncementDetail
- [x] GetAnnouncementList
- ~~[ ] GetAnnouncementListWithAttach (Same as GetAnnouncementList)~~
- [x] GetAnnouncementList_Login
- [x] GetAnnouncementList_LoginByCount
- ~~[ ] GetAnnouncementList_LoginByCountWithAttach (Same as GetAnnouncementList_LoginByCount)~~
- [ ] GetAttachFileList
- [ ] GetAttachFiles
- [ ] GetBoardMessage
- [ ] GetBoardMessage2
- [ ] GetCalendarData
- [ ] GetCanAnswerQuestionInfo
- [x] GetCourseInfo
- [x] GetCourseList
- [x] GetCourseTime
- [x] GetCourseUserContactInfo
- [ ] GetForumList
- [ ] GetHelpClassList
- [ ] GetHomeworkInfo
- [ ] GetHwkSubmitInfo
- [ ] GetHwkSubmitList
- [ ] GetHwkSubmitListWithAttach
- [ ] GetMaterialDocList
- [ ] GetMaterialDocSummary
- [ ] GetMemberList
- [ ] GetPersonalCoursePosition
- [ ] GetPersonalData
- [ ] GetQuestionStatistics
- [ ] GetQuizInfo
- [ ] GetQuizQuestionInfo
- [ ] GetQuizQuestionList
- [ ] GetReplyList
- [ ] GetReplyListWithAttach
- [ ] GetScoreData
- [ ] GetScoreGradeList
- [ ] GetStuAnnouncementList
- [ ] GetStuHomeworkList
- [ ] GetStuHomeworkListWithAttach
- [ ] GetStuMaterialDocList
- [ ] GetStuQuizList
- [ ] GetSystemNewsDetail
- [ ] GetSystemNewsDetail2
- [ ] GetSystemNewsList
- [ ] GetSystemNewsList2
- [ ] GetSystemNewsList3
- [ ] GetTeaHomeworkList
- [ ] GetTeaHomeworkListWithAttach
- [ ] GetTeaQuizList
- [ ] GetTopReply
- [ ] GetTopicForumList
- [ ] GetTopicInfo
- [ ] GetTopicInfoWithAttach
- [ ] GetTopicList
- [ ] GetTopicUnitList
- [ ] HwkSubmit
- [x] Login
- [ ] PushNotificationAddQueue
- [ ] PushNotificationAddQueueForWeb
- [ ] PushNotificationDeleteQueue
- [ ] PushNotificationDeleteQueueForWeb
- [ ] PushNotificationGetNotificationUpdateInfo
- [ ] PushNotificationHandleJoin
- [ ] PushNotificationHandleJoinPhone
- [ ] PushNotificationHandleLeave
- [ ] PushNotificationUpdateBadge
- [ ] RegistryHelpMessage
- [ ] SendMail
- [ ] SetAnswerQuestion
- [ ] SetDiscussForum
- [ ] SetQuestionOpenState
- [ ] SetQuizAllQuestionOpened
- [ ] SetQuizFinished
- [ ] SetQuizInfo
- [ ] SetQuizQuestionInfo
- [ ] SetStuHomeworkScore
