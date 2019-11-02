//将yaml文件，转换成对象，再转换成json格式输出
package tcbadmingo
import (
    "fmt"
    //"gopkg.in/yaml.v2"
	//"io/ioutil"
	"time"
	"strings"
	//"net/url"
	"github.com/imroc/req"
	"encoding/json"
	//"reflect"
	//"strconv"
)


//定义conf类型
//类型里的属性，全是配置文件里的属性
type Conf struct {
    Env        string `yaml:"env"`
    SecretId   string `yaml:"secretId"`
    SecretKey  string `yaml:"secretKey"`
}
type Request struct {
    CollectionName  string `json:"collectionName"`
    Limit   int64 `json:"limit"`
	Action  string `json:"action"`
	EnvName string `json:"envName"`
	TimeStamp int64 `json:"timestamp"`

	EventId string `json:"eventId"`
	WxCloudApiToken string `json:"wxCloudApiToken"`
	Tcb_sessionToken string `json:"tcb_sessionToken"`
	Query interface{} `json:"query"`
	Data  interface{} `json:"data,omitempty"`
	Multi bool `json:"multi"`
	Merge bool `json:"merge"`
	Upsert bool `json:"upsert"`
}

type Option struct{
	Path string
	Method string
    Params map[string]interface{}
	Body  map[string]interface{}
	Dst interface{}
}

type TcbAdmin struct{
	yml_conf Conf
	debug bool
}
type ResponseData struct{
	Id string `json:"_id"`
	Offset int `json:"offset"`
	Limit int `json:"limit"`
	Total int `json:"total"`
	Updated int `json:"updated"`
	Deleted int `json:"deleted"`
	Upserted_id string `json:"upserted_id`
	List json.RawMessage `json:"list"`
}
type Response struct{
	RequestId string `json:"requestId`
	Data ResponseData `json:"data"`
}
func New(yml_conf Conf) *TcbAdmin{

	return &TcbAdmin{yml_conf, false}
}
func (self *TcbAdmin) Debug(opt bool){
	self.debug = opt
}
func (self *TcbAdmin) CommonRequest(option Option, action string) Response{
	// 首先翻译接口url
	if strings.HasSuffix(option.Path, "/"){
		// 去掉url最后的 "/"
		option.Path = string([]rune(option.Path)[:len(option.Path)-1])
	}
	if false == strings.HasPrefix(option.Path, "/"){
		// 增加url头部的 "/"
		option.Path = "/" + option.Path
	}
	restPath := strings.Split(option.Path, "/")
	pathField := "" // db, file, or other
	pathName := ""  // collectionName, or other

	pathField = restPath[1]
	pathName = restPath[2]
	
	if nil == option.Params {
		option.Params = map[string]interface{}{}
	}
	if nil == option.Body {
		option.Body = map[string]interface{}{}
	}
	fmt.Println(pathField)

	headers := map[string]string{
		"x-tcb-source": ",not_scf",
		"user-agent": "tcb-admin-sdk/1.15.0"}

	pathname := "/admin"
	method := "post"

	authHeader := req.Header{
		"content-type": "application/json",
		"x-tcb-source": ",not_scf",
		"user-agent": "tcb-admin-sdk/1.15.0"}
	
	// param := req.Param{"envName": Env}

	r := req.New()
	req.Debug = self.debug
	myrequest := Request{}
	
	myrequest.CollectionName = pathName
	myrequest.Limit = 100
	myrequest.Action = action
	myrequest.EnvName = self.yml_conf.Env
	myrequest.WxCloudApiToken = ""
	
	myrequest.Tcb_sessionToken = ""
	myrequest.Multi = false
	myrequest.Merge = true
	// 这里的params实际上就是查询使用的结构体
	// 但是一般上url的键值对要求有等号等等，这里有点不同
	myrequest.Query = option.Params
	myrequest.Data  = option.Body
	
	myrequest.Upsert = false
	// we will fill timestamp and event id later in AuthRequest function
	timeUnixNano := time.Now().UnixNano()
	queryStr,eventId := AuthRequest(timeUnixNano, self.yml_conf.Env,
		self.yml_conf.SecretId,
		self.yml_conf.SecretKey,
		headers,
		pathname,
		method,
		myrequest)
	url := "http://tcb-admin.tencentcloudapi.com/admin?eventId="+eventId+"&seqId="+eventId
	ret, _ := r.Post(url, authHeader, queryStr)

	response := Response{}
	err := ret.ToJSON(&response)
	if err != nil {
		fmt.Println("error:", err)
	}
	return response
}
func (self *TcbAdmin) Post(option Option) string {
	response := self.CommonRequest(option, "database.addDocument")
	return response.Data.Id
}
func (self *TcbAdmin) Delete(option Option)  int {
	response := self.CommonRequest(option, "database.deleteDocument")
	return response.Data.Deleted
}
func (self *TcbAdmin) Patch(option Option)  int {
	response := self.CommonRequest(option, "database.updateDocument")
	return response.Data.Updated
}
func (self *TcbAdmin) Get(option Option)  int {
	response := self.CommonRequest(option, "database.queryDocument")
	err := json.Unmarshal(response.Data.List, option.Dst) 
	if err != nil {
		fmt.Println("error:", err)
	}
	return response.Data.Total
}

//读取Yaml配置文件,
//并转换成conf对象
/*
func getConf(filePath string, yml_conf *Conf){
    //应该是 绝对地址
    yamlFile, err := ioutil.ReadFile(filePath)
    if err != nil {
        fmt.Println(err.Error())
    }

    err = yaml.Unmarshal(yamlFile, yml_conf)

    if err != nil {
        fmt.Println(err.Error())
    }
}
*/
