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
	option.Method = "get"
    option.Params = map[string]interface{}{"age": map[string]interface{}{"$gt":8}}
    // option.Schema = &User{}
    // option.Params = user
    //option.Body  =  map[string]interface{}{}
    var users []User
    option.Dst = &users
    TcbAdmin.Get(option)
    fmt.Println("I got this!")
    fmt.Println(users)
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