package main

import (
	"encoding/json"
	"fmt"
	"github.com/mr-tron/base58"
	"io/ioutil"
	"net/http"
	"strconv"
)

const (
	databaseName = "HJM9XONGDY"
	baseUrl      = "https://controlpanel.dbchain.cloud/relay/dbchain/"
)

func main() {
	resp := QueryTableData("student")
	fmt.Println("=============Query Table Data Resp====================")
	fmt.Println(resp)
	fmt.Println("==============Query Table Data Resp End======================")
	fmt.Println("=================Boundary===================")
	fmt.Println("=================ExtractData From Resp===================")
	DataMap :=ExtractData(resp)
	for _,v :=range DataMap{
		for k2,v2:=range v{
			fmt.Println("Data key :",k2,",value :",v2)
		}
		fmt.Println("=================Boundary===================")
	}
	fmt.Println("=================ExtractData  End===================")
	fmt.Println("=================Boundary===================")
	fmt.Println("=================Query Table DataType Resp===================")
	resp2 := QueryTableStruct("student")
	fmt.Println(resp2)
	fmt.Println("=================Query Table DataType Resp===================")
	fmt.Println("=================Boundary===================")
	fmt.Println("=================ExtractData From  Resp===================")
	fieldsMap :=ExtractFields(resp2)
	for k,v :=range fieldsMap{
		fmt.Println("fieldsMap key :",k,",value :",v)
	}
	fmt.Println("=================ExtractData End===================")

}

func ExtractFields(resp string)map[string]interface{}{
	myMap :=JSONToMap(resp)
	//for k,v :=range myMap{
	//	fmt.Println("key :",k,",value :",v)
	//}
	//fmt.Println("submap result:",myMap["result"])
	resultinterface :=myMap["result"]
	//fmt.Println("result type",reflect.TypeOf(resultinterface))
	resultMap,_ :=resultinterface.(map[string]interface{})
	//for k,v :=range resultMap{
	//	fmt.Println("resultmap key :",k,",value :",v)
	//}
	//fmt.Println("fieilds type",reflect.TypeOf(resultMap["fields"]))
	//fmt.Println("fieilds content",resultMap["fields"])
	fieldsInterfaces := resultMap["fields"].([]interface{})
	var fieldsMap =make(map[string] interface{})
	for _,v :=range fieldsInterfaces{
		//fmt.Println("resultmap key :",k,",value :",v)
		fieldsMap[v.(string)]=""
	}
	//for k,v :=range fieldsMap{
	//	fmt.Println("strs key :",k,",value :",v)
	//}
	return fieldsMap
}

func ExtractData(resp string) map[string]map[string]interface{}{
	myMap :=JSONToMap(resp)
	resultinterfaces :=myMap["result"]
	//fmt.Println("submap result type:",reflect.TypeOf(resultinterfaces))
	 resultMap,_ :=resultinterfaces.([]interface{})
	var DataMap =make(map[string]map[string] interface{})
	//fmt.Println("submap resultmap type:",reflect.TypeOf(resultMap))
	for k,v :=range resultMap{
		//fmt.Println("resultmap key :",k,",value :",v)
		DataMap["index"+strconv.Itoa(k)]=v.(map[string] interface{})
	}
	//fmt.Println("values type: ",reflect.TypeOf(resultMap[0]))
	//for k:=range DataMap{
	//	fmt.Println("DataMap key :",k)
	//}
	//for k,v :=range DataMap{
	//	fmt.Println("DataMap key :",k,",value :",v["name"])
	//}
	return DataMap
}

func QueryTableData(tableName string) string {
	queryCondition := "[{\"method\":\"table\",\"table\":\"" + tableName + "\"},{\"method\":\"offset\",\"value\":\"0\"},{\"method\":\"limit\",\"value\":\"50\"}]"
	qCEncode := base58.Encode([]byte(queryCondition))
	token := MakeAccessCode()
	QueryDataUrl := baseUrl + "/querier/" + token + "/" + databaseName + "/" + qCEncode
	//fmt.Println("url : ",QueryDataUrl)
	return Get(QueryDataUrl)
}

func QueryTableStruct(tableName string) string {
	token := MakeAccessCode()
	QueryDataUrl := baseUrl + "/tables/" + token + "/" + databaseName + "/" + tableName
	return Get(QueryDataUrl)
}

func Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		return err.Error()
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err.Error()
	}
	return string(body)
}

func JSONToMap(str string) map[string]interface{} {
	var tempMap map[string]interface{}
	err := json.Unmarshal([]byte(str), &tempMap)
	if err != nil {
		panic(err)
	}
	return tempMap
}
