package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/astaxie/beego"
)

func Metodo_get(host, endpoint string) ([]byte, error) {
	url := beego.AppConfig.String(host) + endpoint
	fmt.Println("Ruta Get", url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return body, nil
}

func ProcesarJsonArreglos(datos []byte) ([]map[string]interface{}, error) {
	var result []map[string]interface{}

	err := json.Unmarshal(datos, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func Metodo_post(nombre_servicio string, data []byte) ([]byte, error) {
	url := beego.AppConfig.String(nombre_servicio)
	response, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	return body, nil
}

func ProcesarJson(datos []byte) (map[string]interface{}, error) {
	var result map[string]interface{}
	err5 := json.Unmarshal(datos, &result)
	if err5 != nil {
		log.Fatal(err5)
		return nil, err5
	}
	return result, nil
}

func Metodo_put(nombre_servicio string, id string, data []byte) ([]byte, error) {
	//obtener la URL base desde la configuracion de Beego
	baseURL := beego.AppConfig.String(nombre_servicio)

	//construir una url final con el ID
	url := fmt.Sprintf("%s/%s", baseURL, id)

	//crear la solicitud put
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	// Establecer el encabezado  COntent-Type
	req.Header.Set("Content-Type", "application/json")

	// Enviar solicitud
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// Leer la respuesta
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	return body, nil
}

// Metodo Delete
func Metodo_delete(nombre_servicio, endpoint, id string) ([]byte, error) {
	baseURL := beego.AppConfig.String(nombre_servicio)
	if baseURL == "" {
		return nil, fmt.Errorf("no se encontró la configuración para %s", nombre_servicio)
	}

	url := fmt.Sprintf("%s%s/%s", baseURL, endpoint, id)
	fmt.Println("URL construida:", url)

	// Crear la solicitud DELETE
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return nil, fmt.Errorf("error al crear la solicitud DELETE: %v", err)
	}

	// Enviar la solicitud
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error en DELETE a %s: %v", url, err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK && response.StatusCode != http.StatusNoContent {
		body, _ := ioutil.ReadAll(response.Body)
		return nil, fmt.Errorf("error en API DELETE: %d - %s", response.StatusCode, string(body))
	}

	// Leer la respuesta
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error al leer la respuesta: %v", err)
	}

	fmt.Println("Respuesta de la API:", string(body))
	return body, nil
}
func ConvertToSliceMap(input interface{}) ([]map[string]interface{}, error) {
	// Intentamos convertir a []interface{}
	rawSlice, ok := input.([]interface{})
	if !ok {
		return nil, errors.New("el input no es un slice de interface{}")
	}

	// Recorremos el slice y convertimos cada elemento a map[string]interface{}
	var result []map[string]interface{}
	for i, item := range rawSlice {
		m, ok := item.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("el elemento en el índice %d no es map[string]interface{}", i)
		}
		result = append(result, m)
	}

	return result, nil
}
