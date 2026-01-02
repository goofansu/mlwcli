package format

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/itchyny/gojq"
)

func Output(data any, fields string, jqExpr string) error {
	var outputData any = data

	if fields != "" {
		filtered, err := filterFields(data, fields)
		if err != nil {
			return err
		}
		outputData = filtered
	}

	if jqExpr != "" {
		var dataForJQ any = outputData
		if items, ok := outputData.(map[string]any); ok {
			if v, exists := items["items"]; exists {
				dataForJQ = v
			}
		}

		jsonData, err := json.Marshal(dataForJQ)
		if err != nil {
			return err
		}
		var cleanData any
		if err := json.Unmarshal(jsonData, &cleanData); err != nil {
			return err
		}

		results, err := applyJQ(cleanData, jqExpr)
		if err != nil {
			return err
		}

		for _, result := range results {
			if result == nil {
				fmt.Println()
				continue
			}
			switch v := result.(type) {
			case string:
				fmt.Println(v)
			case float64, int, int64, bool:
				fmt.Println(v)
			default:
				output, err := json.MarshalIndent(result, "", "  ")
				if err != nil {
					return err
				}
				fmt.Println(string(output))
			}
		}
		return nil
	}

	output, err := json.MarshalIndent(outputData, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(output))
	return nil
}

func filterFields(data any, fields string) (map[string]any, error) {
	if fields == "" {
		return nil, fmt.Errorf("no fields specified")
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	var raw any
	if err := json.Unmarshal(jsonData, &raw); err != nil {
		return nil, err
	}

	result := make(map[string]any)
	fieldList := strings.Split(fields, ",")

	switch v := raw.(type) {
	case []any:
		if len(v) == 0 {
			return map[string]any{"items": []any{}}, nil
		}
		items := make([]map[string]any, len(v))
		for i, item := range v {
			if m, ok := item.(map[string]any); ok {
				items[i] = filterMapFields(m, fieldList)
			}
		}
		return map[string]any{"items": items}, nil
	case map[string]any:
		return filterMapFields(v, fieldList), nil
	default:
		return result, nil
	}
}

func filterMapFields(m map[string]any, fields []string) map[string]any {
	result := make(map[string]any)
	for _, field := range fields {
		if val, exists := m[field]; exists {
			result[field] = val
		}
	}
	return result
}

func applyJQ(data any, jqExpr string) ([]any, error) {
	query, err := gojq.Parse(jqExpr)
	if err != nil {
		return nil, fmt.Errorf("jq parse error: %w", err)
	}

	iter := query.Run(data)
	var results []any
	for {
		v, ok := iter.Next()
		if !ok {
			break
		}
		if err, ok := v.(error); ok {
			return nil, fmt.Errorf("jq error: %s", err.Error())
		}
		results = append(results, v)
	}

	return results, nil
}
