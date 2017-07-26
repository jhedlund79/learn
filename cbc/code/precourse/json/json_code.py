import json


# Lesson using pytest
def simple_tdd_function(a,b):
    return(a + b)


# Exercise 1
def read_json_file(file_path):
    with open(file_path) as json_data:
        data = json.load(json_data)
    return (data)


# Exercise 2
def course_weeks_count(json_data):
    if 'weeks' in json_data:
        return len(json_data['weeks'])
    else:
        return 0



# Exercise 3
def course_content_count(json_data):
    content_count = {}
    if 'weeks' in json_data:
        for key in json_data['content_units']:
            if key in json_data['weeks']:
                content_count[key] = len(json_data['content_units'][key])
    return content_count


# Exercise 4
def tones_parse_anger(json_data):
    score = 0.0
    for key in json_data["document_tone"]["tone_categories"]:
        for value in key["tones"]:
            if value['tone_id'] == 'anger':
                score = value["score"]
    return score
