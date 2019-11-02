package main
import tcb "github.com/wenfengand/tcbadmingo"
import "gopkg.in/yaml.v2"
import "io/ioutil"
import "fmt"

var filePath = "config.yaml"

type User struct{
    Id string `json:"_id,omitempty"`
    Name string `json:"name,omitempty"`
    Age int `json:"age,omitempty"`
}
func main(){
	yml_conf := tcb.Conf{}
    getConf(filePath, &yml_conf)
    TcbAdmin := tcb.New(yml_conf)
    TcbAdmin.Debug(true)

	option := tcb.Option{}
	option.Path = "/db/test"
	option.Method = "post"
    //option.Params = map[string]interface{}{"_id":"9888d322-fe11-4aaa-af7f-bc91ea870f4c"}
    user := User{}
    user.Name = "testetet"
    user.Age  = 1010
    option.Body  = tcb.Struct2Map(user)
    ret := TcbAdmin.Post(option)
    fmt.Println(ret)
}

func getConf(filePath string, yml_conf *tcb.Conf){
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