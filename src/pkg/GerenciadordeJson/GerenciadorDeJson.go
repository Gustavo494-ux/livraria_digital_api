package GerenciadordeJson

import (
	"encoding/json"
	"reflect"
	"strings"
)

// InterfaceParaJsonString: Converte uma interface genérica em um json em formato string.
func InterfaceParaJsonString(i any) (string, error) {

	jsonBytes, err := json.Marshal(i)
	if err != nil {
		return "", err
	}
	jsonStr := string(jsonBytes)
	return jsonStr, nil
}

// JsonStringParaInterface: Converte um json em formato string para uma interface genérica.
func JsonStringParaInterface(jsonStr string) (any, error) {
	var i interface{}
	err := json.Unmarshal([]byte(jsonStr), &i)
	if err != nil {
		return nil, err
	}
	return i, nil
}

// IgnorarCamposPelaTag: ignora os campos na serialização caso a tag e conteúdo esperado sejam encontrados.
func IgnorarCamposPelaTag(st any, tagAlvo string, conteudoEsperado string) (j []byte, err error) {
	m := make(map[string]any)

	val := reflect.ValueOf(st)
	typeOfVal := val.Type()

	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := typeOfVal.Field(i)

		valorTag := typeField.Tag.Get(tagAlvo)
		if strings.EqualFold(conteudoEsperado, valorTag) {
			continue
		}

		m[typeField.Name] = valueField.Interface()
	}

	j, err = json.Marshal(m)
	return
}
