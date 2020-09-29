package seralize

import "encoding/json"

func Seralize(a interface{}) (string,error) {
	str,err := json.Marshal(a)
	if err !=nil{
		return "",err
	}
	return string(str),nil
}

func UnSeralize(str string) (interface{},error) {
	var ret interface{}
	err := json.Unmarshal([]byte(str),&ret)
	if err !=nil{
		return nil,err
	}
	return ret,nil
}
