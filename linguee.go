package main

import (
	"encoding/json"
	"fmt"
	"github.com/buger/jsonparser"
	"io/ioutil"
	"net/http"
	"os"
)

var translation_entity translation_model

func get_exact_matches(matches []interface{}) {
	total_size := len(matches)
	for translation_entity.number_of_translation < total_size {
		translation_entity.translation = append(translation_entity.translation, matches[translation_entity.number_of_translation].(map[string]interface{})["text"].(string))
		translation_entity.number_of_translation++
	}
}

func request_translation() string {
	resp, _ := http.Get(fmt.Sprintf("http://localhost:8000/api?q=%v&src=%v&dst=%v", translation_entity.query, translation_entity.src_lang, translation_entity.dest_lang))
	defer resp.Body.Close()

	if resp.StatusCode == 200 {

		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		err := json.Unmarshal(bodyBytes, &translation_entity.response_body)

		if err == nil {
			exact_match := translation_entity.response_body.(map[string]interface{})["exact_matches"]
			if exact_match != nil && len(exact_match.([]interface{})) > 0 {
				element := translation_entity.response_body.(map[string]interface{})["exact_matches"].([]interface{})[0].(map[string]interface{})["translations"].([]interface{})
				get_exact_matches(element)
				fmt.Println(translation_entity.number_of_translation, "Exact match: ", translation_entity.translation)

			} else {
				return "No exact match."
			}
		} else {
			return "Error happened during the conversion of resp.Body to JSON format."
		}
	}
	return ""
}

func main() {

	var json_brut string

	if len(os.Args) == 4 {
		translation_entity.src_lang = os.Args[1]
		translation_entity.dest_lang = os.Args[2]
		translation_entity.query = os.Args[3]
		json_brut = request_translation()
	} else {
		/* Do translate by default EN-FR */
		translation_entity.src_lang = "en"
		translation_entity.dest_lang = "fr"
		translation_entity.query = os.Args[1]
		json_brut = request_translation()
	}

	fmt.Println(json_brut)

	//fmt.Println("Linguee CLI tool provide you an easy way to translate text directly in your terminal.")
}
