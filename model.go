package main

type translation_model struct {
	src_lang                string
	dest_lang               string
	query                   string
	correct_query           string
	translation             []string
	less_common_translation []string
	response_body           interface{}
	number_of_translation   int
}
