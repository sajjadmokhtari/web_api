package common

import "encoding/json"

func TypeConverter[T any](data any) (*T, error) {// تابع تبدیل داده به نوع جنریک  تی 
	var result T
	dataJson, err := json.Marshal(&data)//داده ورودی تبدیل میشه یه بایت ارایه از نوع جیسون
	if err != nil  {
		return nil,err
	}
	err = json.Unmarshal(dataJson,&result)
	if err != nil {
		return nil , err 
	}
	return &result , nil
}