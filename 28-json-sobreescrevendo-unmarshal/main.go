package main

import (
	"encoding/json"
	"fmt"
)

type Valor struct {
	Numero int `json:"numero"`
}

func (v Valor) MarshalJSON() ([]byte, error) {
	return json.Marshal(fmt.Sprintf("valor: %d", v.Numero))
}

func (v *Valor) UnmarshalJSON(data []byte) error {
	var raw string
	err := json.Unmarshal(data, &raw)
	if err != nil {
		return err
	}

	_, err = fmt.Sscanf(raw, "valor: %d", &v.Numero)
	return err
}

func main() {

	v := Valor{Numero: 10}
	jsonData, _ := json.Marshal(v)
	fmt.Println("JSON: ", string(jsonData))

	inputData := `"valor: 20"`
	var vDecoded Valor
	json.Unmarshal([]byte(inputData), &vDecoded)
	fmt.Println("Valor decodificado: ", vDecoded.Numero)

}
