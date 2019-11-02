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
	option.Method = "delete"
    option.Params = map[string]interface{}{"_id":"4a741dc95dbd51d5006ea699315a850c"}
	//option.Body  = map[string]interface{}{"name":"20191102", "age":100}
    ret := TcbAdmin.Delete(option)
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