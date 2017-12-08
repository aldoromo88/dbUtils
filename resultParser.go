package dbUtils

import(
	"encoding/json"
	"database/sql"
)

func  ParseAsMap(r sql.Rows) ([]map[string]interface{},error) {

	columns, err := r.Columns()
	if err!= nil{
		return nil, err
	}

	columnsLen := len(columns)

	rawData := make([]map[string]interface{}, 0)
	values := make([]interface{}, columnsLen)



	for r.Next()  {
		r.Scan(values...)
		record := map[string]interface{}{}
		for i, c := range columns {

			record[c] = values[i]
		}
		rawData = append(rawData, record)
	}


	return rawData, nil
}

func ParseAsJson(r sql.Rows) ([]byte,error){
	data, err := ParseAsMap(r)
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}