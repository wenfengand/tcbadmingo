package main
import tcb "github.com/wenfengand/tcbadmingo"
import "gopkg.in/yaml.v2"
import "io/ioutil"
import "fmt"

var filePath = "config.yaml"

func main(){
	yml_conf := tcb.Conf{}
    getConf(filePath, &yml_conf)
    TcbAdmin := tcb.New(yml_conf)
    TcbAdmin.Debug(true)

	option := tcb.Option{}
	option.Path = "/db/test"
	option.Method = "patch"
    option.Params = map[string]interface{}{"_id":"9888d322-fe11-4aaa-af7f-bc91ea870f4c"}
	option.Body  = map[string]interface{}{"$set":map[string]interface{}{"name":"gotest2"}}
    ret := TcbAdmin.Patch(option)
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