package request

import jsoniter "github.com/json-iterator/go"

var (
	_json = jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		UseNumber:              true,
	}.Froze()
)

func OptionJson(option jsoniter.Config) {
	_json = option.Froze()
}
