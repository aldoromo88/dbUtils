package dbUtils

import(
	"database/sql/driver"
	"io"
	"encoding/json"
)



func  ParseAsMap(r driver.Rows) ([]map[string]interface{},error) {
	columns := r.Columns()
	columnsLen := len(columns)

	rawData := make([]map[string]interface{}, 0)
	values := make([]driver.Value, columnsLen)


	err := r.Next(values)
	for err == nil  {

		record := map[string]interface{}{}
		for i, c := range columns {

			record[c] = values[i]
		}
		rawData = append(rawData, record)

		err = r.Next(values)
	}

	if err != io.EOF{
		return nil, err
	}


	return rawData, nil
}

func ParseAsJson(r driver.Rows) ([]byte,error){
	data, err := ParseAsMap(r)
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)

}