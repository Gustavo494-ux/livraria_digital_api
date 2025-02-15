package Generics

import (
	"fmt"
	"reflect"
)

// RetornarCampo: retorna o campo desejado de um generics/interface caso ele não seja privado
func RetornarCampo(valor interface{}, campo string, valorPadraoCampo any) any {
	val := reflect.ValueOf(valor)

	if val.Kind() == reflect.Struct {
		field := val.FieldByName(campo)
		if field.IsValid() && field.CanInterface() {
			return field.Interface()
		}
	}
	return valorPadraoCampo
}

// SetarCampo: seta o valor para o campo desejado de um struct de tipo T
func SetarCampo(valor interface{}, campo string, valorCampo interface{}) error {
	// Obtém o valor refletido do struct
	v := reflect.ValueOf(valor)

	// Verifica se o valor é um ponteiro para um struct
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("valor deve ser um ponteiro para um struct")
	}

	// Obtém o valor subjacente do ponteiro (o struct em si)
	v = v.Elem()

	// Obtém o campo do struct pelo nome
	field := v.FieldByName(campo)

	// Verifica se o campo existe
	if !field.IsValid() {
		return fmt.Errorf("campo '%s' não existe no struct", campo)
	}

	// Verifica se o campo é exportado (ou seja, se começa com letra maiúscula)
	if !field.CanSet() {
		return fmt.Errorf("campo '%s' não é exportado (privado)", campo)
	}

	// Obtém o valor refletido do valorCampo
	valorReflect := reflect.ValueOf(valorCampo)

	// Verifica se o tipo do valorCampo é compatível com o tipo do campo
	if !valorReflect.Type().AssignableTo(field.Type()) {
		return fmt.Errorf("tipo incompatível para o campo '%s': esperado %s, obtido %s", campo, field.Type(), valorReflect.Type())
	}

	// Define o valor do campo
	field.Set(valorReflect)

	return nil
}
