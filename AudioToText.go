package baiduATT

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type CreateTaskRes struct {
	// {"log_id":16806262779913675,"task_status":"Created","task_id":"642c5266871a4d0001e4d2ca"}
	LogId      int64  `json:"log_id"`
	TaskStatus string `json:"task_status"`
	TaskId     string `json:"task_id"`
}

func (c *BaiduClient) CreateAttTask(fileUrl string) (*CreateTaskRes, error) {

	if c.Token == nil {
		c.GetToken()
	}

	url := "https://aip.baidubce.com/rpc/2.0/aasr/v1/create?access_token=" + c.Token.Access_token
	// 	arg := NewAudioToTextModel(fileUrl, "mp3", 1737, 16000)
	// log.Println(arg.ToValues())

	payload := strings.NewReader(fmt.Sprintf(`{"speech_url":"%s","format":"mp3","pid":1737,"rate":16000}`, fileUrl))

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, payload)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	res, err := client.Do(req)

	//res, err := http.PostForm(url, arg.ToValues())
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	dat, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	ret := &CreateTaskRes{}
	err = json.Unmarshal(dat, ret)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	//log.Println(string(dat))
	return ret, nil
}

func (c *BaiduClient) GetTaskResult(taskIDs []string) (*BaiduATTTaskResult, error) {

	if c.Token == nil {
		c.GetToken()
	}

	url := "https://aip.baidubce.com/rpc/2.0/aasr/v1/query?access_token=" + c.Token.Access_token
	// 	arg := NewAudioToTextModel(fileUrl, "mp3", 1737, 16000)
	// log.Println(arg.ToValues())

	var ids = "["
	for i := 0; i < len(taskIDs); i++ {
		ids = ids + "\"" + taskIDs[i] + "\","
	}
	ids = ids[:len(ids)-1] + "]"
	log.Println(ids)
	payload := strings.NewReader(fmt.Sprintf(`{"task_ids":%s}`, ids))

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, payload)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	res, err := client.Do(req)

	//res, err := http.PostForm(url, arg.ToValues())
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	dat, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var ttsr = BaiduATTTaskResult{}

	log.Println(string(dat))
	json.Unmarshal(dat, &ttsr)
	log.Println(ttsr)
	return &ttsr, nil
}

type AudioToTextModel struct {
	//Access_token string
	Speech_url string ///可使用百度云对象存储进行音频存储，生成云端可外网访问的url链接，音频大小不超过500MB
	Format     string //["mp3", "wav", "pcm","m4a","amr"]单声道，编码 16bits 位深
	Pid        int    //	语言类型	[80001（中文语音近场识别模型极速版）, 80006（中文音视频字幕模型，申请试用），1737（英文模型）]
	Rate       int    //	是	采样率	[16000] 固定值
}

func NewAudioToTextModel(speech_url string, format string, pid, rate int) *AudioToTextModel {
	ret := &AudioToTextModel{
		Speech_url: speech_url,
		Format:     format,
		Pid:        pid,
		Rate:       rate,
	}
	return ret
}

func (att *AudioToTextModel) ToValues() url.Values {
	values := url.Values{
		"speech_url": {att.Speech_url},
		"format":     {att.Format},
		"pid":        {fmt.Sprintf("%d", att.Pid)},
		"rate":       {fmt.Sprintf("%d", att.Rate)},
	}
	return values
}

type BaiduATTTaskDetailedResultWordsInfo struct {
	EndTime   int    `json:"end_time"`   // "end_time": 1170,
	Words     string `json:"words"`      // "words": "section",
	BeginTime int    `json:"begin_time"` // "begin_time": 580
}

type BaiduATTTaskDetailedResult struct {
	Res     []string `json:"res"`      //["观众朋友大家好，欢迎收看本期视频哦。"]
	EndTime int      `json:"end_time"` //"end_time": 6700,

	BeginTime int                                   `json:"begin_time"` //"begin_time": 4240,
	WordsInfo []BaiduATTTaskDetailedResultWordsInfo `json:"words_info"` //"words_info": [],
	Sn        string                                `json:"sn"`         //"sn": "257826606251573543780",
	CorpusNo  string                                `json:"corpus_no"`  //"corpus_no": "6758319075297447880"
}

type TextResult struct {
	CorpusNo       string                       `json:"corpus_no"`       // ++corpus_no	str	否
	Result         string                       `json:"result"`          // ++result	str	否	转写结果
	AudioDuration  int                          `json:"audio_duration"`  // ++audio_duration	int	否	音频时长（毫秒）
	DetailedResult []BaiduATTTaskDetailedResult `json:"detailed_result"` // ++detailed_result	list	否	转写详细结果
	ErrNo          int                          `json:"err_no"`          // ++err_no	int	否	转写失败错误码
	ErrMsg         string                       `json:"err_msg"`         // ++err_msg	str	否	转写失败错误信息
	Sn             string                       `json:"sn"`              // ++sn	str	否
}
type BaiduATTTaskInfo struct {
	TaskId     string     `json:"task_id"`     // +task_id	str	是	任务id
	TaskStatus string     `json:"task_status"` // +task_status	str	是	任务状态
	TaskResult TextResult `json:"task_result"` // +task_result	dict	否	转写结果的json格式
}

type BaiduATTTaskResult struct {
	LogId     int                `json:"log_id"`     //log_id	int	是	log id
	TasksInfo []BaiduATTTaskInfo `json:"tasks_info"` // tasks_info	list	否	多个任务的结果

	ErrorCode int      `json:"error_code"` // error_code	int	否	请求错误码
	ErrorMsg  string   `json:"error_msg"`  // error_msg	str	否	请求错误信息
	ErrorInfo []string `json:"error_info"` // error_info	list	否	错误的或查询不存在的taskid数组
}
