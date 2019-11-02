package tcbadmingo
import (
	"testing"
	// "time"
	"github.com/stretchr/testify/assert"
	"strconv"
)

func TestAuthRequest(t *testing.T) {
	t.Log("Test")
	var yml_conf = Conf{}
    //读取yaml配置文件
    getConf("config.yaml", &yml_conf)

	headers := map[string]string{
		"x-tcb-source": ",not_scf",
		"user-agent": "tcb-admin-sdk/1.15.0"}

	pathname := "/admin"
	method := "post"

	query := Query{}
	query.CollectionName = "test"
	query.Limit = 100
	query.Action = "database.queryDocument"
	query.EnvName = yml_conf.Env
	query.WxCloudApiToken = ""
	query.Tcb_sessionToken = ""
	
	timeUnixNano,_ := strconv.ParseInt("1572333388387000000", 10, 64)
	queryStr,_ := AuthRequest(timeUnixNano, yml_conf.Env,
		yml_conf.SecretId,
		yml_conf.SecretKey,
		headers,
		pathname,
		method,
		query)
	length := len(queryStr)
	assert.Equal(t, "338a9c8464daa2a9ebca1b7006c29f02153aa305", queryStr[length-42:length-2])
}