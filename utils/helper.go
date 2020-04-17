package utils

import (
	"encoding/json"
	"github.com/bitly/go-simplejson"
	"reflect"
	"strings"
)

//region Remark: 将查询出来的数据结构只取部分字段展示给前端 Author:Qing

//region Remark: 获取指定字段 Author; chijian
// fields "id, order_number, user_info:{id as user_id; mobile; user_name}, bank_card_info:{bank;name;bank_card}")
// 主表字段之间以","隔开, 副表格式 user_info:{...;...;... as ...}, 其中字段之间以";"隔开, 用别名时, 使用 " as "
func Struct2Map(val interface{}, fields string) []interface{} {
	dataByte, _ := json.Marshal(val)
	js, _ := simplejson.NewJson(dataByte)
	ids := strings.Split(fields, ",")
	nodes, err := js.Array()
	if err != nil { //单条数据
		return []interface{}{dealSingleData(val, ids)}
	}
	data := make([]interface{}, len(nodes))

	for key, value := range nodes {
		data[key] = dealSingleData(value, ids)
		//fmt.Println("item:", item)
	}
	//fmt.Println(data)
	return data
}

func dealSingleData(value interface{}, ids []string) map[string]interface{} {
	jsonMap := make(map[string]interface{})
	jsonMap = Json2Map(value)
	item := make(map[string]interface{})
	for _, v := range ids {
		aim, sour := string2key(v)
		for k, _ := range sour {
			s := strings.Replace(sour[k], " ", "", -1)
			a := strings.Replace(aim[k], " ", "", -1)
			item[a] = jsonMap[s]
		}
	}
	return item
}

//region Remark: 对传入的需要取值的字段解析, field是已经用","分隔开的字符串 Author; chijian

//endregion
func string2key(field string) ([]string, []string) {
	//该 field 中是否是要从副表中取字段(格式 user_info:{name; mobile; id as user_id})
	aim := make([]string, 0)  //最后需要展示的key键切片
	sour := make([]string, 0) //源数据中的key键切片
	ids := strings.Split(field, "{")
	if len(ids) == 1 { //主表字段
		return append(aim, field), append(sour, field)
	}
	//需要从副表取字段
	fields := ids[1][:len(ids[1])-1] //具体字段 name; mobile; id as user_id} 需要去除最后的"}"字段
	prefix := ids[0][:len(ids[0])-1] //前缀 ids[0] user_info: 需要去除最后的":"元素

	f_ids := strings.Split(fields, ";") //副表字段以 ";" 隔开

	for _, val := range f_ids { //循环 name; mobile; id as user_id 副表字段
		f_as := strings.Split(val, " as ")
		if len(f_as) == 1 { //副表字段中没有 " as "标签, 副表val字段, 展示的键就是val
			sour = append(sour, prefix+"_"+val)
			aim = append(aim, val)
		} else {
			sour = append(sour, prefix+"_"+f_as[0])
			aim = append(aim, f_as[1]) //副表字段中有 " as "标签, 副表val字段, 展示的键是f_as[1]
		}
	}
	return aim, sour
}

//endregion

//region Remark: 单条数据 json格式转为 map[string]interface{} 格式(其中副表以 table_field 为键存储) Author; chijian
func Json2Map(json_inter interface{}) (jsonMap map[string]interface{}) {
	json_bt, _ := json.Marshal(json_inter)
	//fmt.Println("string:", string(json_bt))
	json.Unmarshal(json_bt, &jsonMap)
	for key, value := range jsonMap {
		if m := reflect.ValueOf(value); m.Kind() == reflect.Map {
			item := make(map[string]interface{})
			v_bt, _ := json.Marshal(value)
			json.Unmarshal(v_bt, &item)
			for k, val := range item {
				jsonMap[key+"_"+k] = val
			}
		}
	}
	return jsonMap
}

//endregion

//endregion
