package logging

func logParamsToZapParams(keys map[ExtraKey]interface{}) []interface{} {
	params := make([]interface{}, 0, len(keys))

	for k, v := range keys {
		params = append(params, string(k))
		params = append(params, v)
	}

	return params
}

func logParamsToZeroParams(keys map[ExtraKey]interface{}) map[string]interface{} {
	params := map[string]interface{}{}

	for k, v := range keys {
		params[string(k)] = v
	}

	return params
}
