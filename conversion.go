package dynjson

func convToObject(data interface{}) (JsonObject, bool) {
	if data != nil {
		if dataObject, ok := data.(map[string]interface{}); ok {
			return dataObject, ok
		}
		if dataObject, ok := data.(JsonObject); ok {
			return dataObject, ok
		}
	}

	return nil, false
}

func convToList(data interface{}) (JsonList, bool) {
	if data != nil {
		if dataList, ok := data.([]interface{}); ok {
			return NewJsonList(dataList), ok
		}
		if dataList, ok := data.(JsonList); ok {
			return dataList, ok
		}
	}

	return nil, false
}

func convToString(data interface{}) (string, bool) {
	if data != nil {
		if dataString, ok := data.(string); ok {
			return dataString, true
		}
	}

	return "", false
}

func convToFloat64(data interface{}) (float64, bool) {
	if data != nil {
		if dataFloat64, ok := data.(float64); ok {
			return dataFloat64, true
		}
	}

	return 0, false
}

func convToBool(data interface{}) (bool, bool) {
	if data != nil {
		if dataBool, ok := data.(bool); ok {
			return dataBool, true
		}
	}

	return false, false
}
