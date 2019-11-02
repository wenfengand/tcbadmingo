//将yaml文件，转换成对象，再转换成json格式输出
package tcbadmingo

import (
    "fmt"
	"crypto/hmac"
    "crypto/sha1"
	"strings"
	"net/url"
	"strconv"
	"encoding/json"
	"reflect"
	"sort"
)

func AuthRequest(timeUnixNano int64,
			 env string,
			 secretId string,
			 secretKey string, 
			 headers map[string]string,
			 pathname string,
			 method string,
			 query Request) (string,string){

	timeStamp := timeUnixNano / 1000000
	authStart := timeStamp / 1000
	authEnd := authStart + 900

	query.TimeStamp = timeStamp
	query.EventId = strconv.FormatInt(timeStamp, 10) + "_1_46219"

	queryBytes,_ := json.Marshal(query)
	queryStr := string(queryBytes)

	var qSignAlgorithm = "sha1"
	var qAk = secretId
	var qSignTime = strconv.FormatInt(authStart, 10) + ";" + strconv.FormatInt(authEnd, 10)
	var qKeyTime = strconv.FormatInt(authStart, 10) + ";" + strconv.FormatInt(authEnd, 10)
	
	var qHeaderList = strings.Join(getKeys(headers), ";")
	qHeaderList = strings.ToLower(qHeaderList)

	var qUrlParamList = strings.Join(interfaceKeys(Struct2Map(query)), ";")
	qUrlParamList = strings.ToLower(qUrlParamList)

	

	var hashFunc = sha1.New
	h := hmac.New(hashFunc, []byte(secretKey))
	h.Write([]byte(string(qKeyTime)))
	signKey := fmt.Sprintf("%x", h.Sum(nil))
	//fmt.Println("sign key")
	//fmt.Println(signKey)

	var formatString = make([]string, 5)
	formatString[0] = method
	formatString[1] = pathname
	formatString[2] = interface2str(Struct2Map(query))
	formatString[3] = obj2str(headers)
	formatString[4] = ""

	
	httpString := strings.Join(formatString, "\n")
	fmt.Println("httpString")
	fmt.Println(httpString)

	h = sha1.New()
	h.Write([]byte(httpString))
	sha1HttpString := h.Sum(nil)

	arryToSign := []string{
		"sha1",
		qSignTime,
		fmt.Sprintf("%x", sha1HttpString),
		""}
	stringToSign := strings.Join(arryToSign, "\n")
	//fmt.Println("string to sign")
	//fmt.Println(stringToSign)


	h = hmac.New(sha1.New, []byte(signKey))
	h.Write([]byte(stringToSign))
	tempSig := h.Sum(nil)
	qSignature := fmt.Sprintf("%x", tempSig)

	var authorization = make([]string, 7)
	authorization[0] = "q-sign-algorithm=" + qSignAlgorithm
	authorization[1] =  "q-ak=" + qAk
	authorization[2] =  "q-sign-time=" + qSignTime
	authorization[3] =  "q-key-time=" + qKeyTime

	authorization[4] =  "q-header-list=" + qHeaderList
	authorization[5] =  "q-url-param-list=" + qUrlParamList
	authorization[6] =  "q-signature=" + qSignature

	authString := strings.Join(authorization, "&")


	//fmt.Println(url)
	//fmt.Println(authString)
	// fmt.Printf(string(param))
	queryStr = queryStr[:len(queryStr)-1] + "," + "\"authorization\":\""+ authString + "\"}"
	return queryStr,query.EventId
}

func getKeys(mymap map[string]string) []string {
	count := len(mymap)
	keys := make([]string, count)
	i := 0
	for k,_:= range mymap{
		keys[i] = k
    	i++
	}
	sort.Strings(keys)
	return keys
}

func obj2str(obj map[string]string) string {
	var key,val string
    var list = make([]string, len(obj))
    var keyList = getKeys(obj)
    for i := 0; i < len(keyList); i++ {
      key = strings.ToLower(keyList[i])
      val = obj[key] 

      key = camSafeUrlEncode(key)
      val = camSafeUrlEncode(val)
      list[i] = key + "=" + val
    }
    return strings.Join(list, "&")
}
func interface2str(obj map[string]interface{}) string {
	var list = make([]string, 0)
	var keys = interfaceKeys(obj)
	for i := 0; i < len(keys); i++ {
		k := keys[i]
		v := obj[k]
		fmt.Printf(k)
		fmt.Println(fmt.Sprintf("%T", v))
		if str, ok := v.(string); ok {
			key := camSafeUrlEncode(strings.ToLower(k))
			val := camSafeUrlEncode(str)
			list = append(list, key + "=" + val)
		} else if num, ok := v.(int64); ok {
			key := camSafeUrlEncode(strings.ToLower(k))
			val := camSafeUrlEncode(strconv.FormatInt(num, 10))
			list = append(list, key + "=" + val)
		}else if boolval, ok := v.(bool); ok {
			key := camSafeUrlEncode(strings.ToLower(k))
			val := camSafeUrlEncode(strconv.FormatBool(boolval))
			list = append(list, key + "=" + val)
		}else if mapData, ok := v.(map[string]interface{}); ok {
			if len(mapData) == 0 {
				key := camSafeUrlEncode(strings.ToLower(k))
				val := camSafeUrlEncode("{}")
				list = append(list, key + "=" + val)
			}else{
				mapBytes,_ := json.Marshal(mapData)
				mapStr := string(mapBytes)

				key := camSafeUrlEncode(strings.ToLower(k))
				val := camSafeUrlEncode(mapStr)
				list = append(list, key + "=" + val)
			}
			
			/*
			if len(mapData)==0 {
				key := camSafeUrlEncode(strings.ToLower(k))
				val := camSafeUrlEncode("{}")
				list = append(list, key + "=" + val)
			}else{
				fmt.Println("Not implemented!")
			}*/
		} else{
			fmt.Println("Error")/* not string */
		}
	}
    return strings.Join(list, "&")
}
func interfaceKeys(obj map[string]interface{}) []string {
    var list = make([]string, 0)
	for k,_ := range obj{
		list = append(list, k)
	}
	sort.Strings(list)
    return list
}

func Struct2Map(obj interface{}) map[string]interface{} {
    t := reflect.TypeOf(obj)
    v := reflect.ValueOf(obj)

    var data = make(map[string]interface{})
    for i := 0; i < t.NumField(); i++ {
        data[t.Field(i).Name] = v.Field(i).Interface()
    }
    return data
}
func mapString2MapInterface(m map[string]string) map[string]interface{} {
	mi := map[string]interface{}{}
	for k, v := range m {
		mi[k] = v
	}
	return mi
}
func mapInterface2MapString(m map[string]interface{}) map[string]string {
	mi := map[string]string{}
	for k, v := range m {
		if str, ok := v.(string); ok {
			mi[k] = str/* act on str */
		} else {
			fmt.Println("Error")/* not string */
		}
	}
	return mi
}
func camSafeUrlEncode(s string) string {
	return url.QueryEscape(s)
}