import toneanalyzer
import pytest
import json

def test_parse_toneanalyzer_response():
    json_content = toneanalyzer.read_json_file("pytest_data/tones_1.json")
    expected_dict = {
                        "Anger":0.102932,
                        "Disgust":0.175853,
                        "Fear":0.091308,
                        "Joy":0.134345,
                        "Sadness":0.446459
                    }
    assert toneanalyzer.parse_toneanalyzer_response(json_content) == expected_dict, \
           "when run on file 'data/tones_1.json' should return {}".format(expected_dict)


    json_content = toneanalyzer.read_json_file("pytest_data/tones_2.json")
    expected_dict = {
                        "Anger":0.057571,
                        "Disgust":0.018737,
                        "Fear":0.062812,
                        "Joy":0.670701,
                        "Sadness":0.06032
                    }
    assert toneanalyzer.parse_toneanalyzer_response(json_content) == expected_dict, \
           "when run on file 'data/tones_1.json' should return {}".format(expected_dict)


    json_content = toneanalyzer.read_json_file("pytest_data/tones_3.json")
    expected_dict = {
                        "Anger":0.040993,
                        "Disgust":0.013815,
                        "Fear":0.232289,
                        "Joy":1.2e-05,
                        "Sadness":0.757131
                    }
    assert toneanalyzer.parse_toneanalyzer_response(json_content) == expected_dict, \
           "when run on file 'data/tones_1.json' should return {}".format(expected_dict)

    json_content = toneanalyzer.read_json_file("pytest_data/tones_4_fail.json")
    expected_dict = {
                        "Disgust":0.013815,
                        "Fear":0.232289,
                        "Joy":1.2e-05,
                        "Sadness":0.757131
                    }
    assert toneanalyzer.parse_toneanalyzer_response(json_content) == expected_dict, \
           "when run on file 'data/tones_1.json' should return {}".format(expected_dict)

    json_content = toneanalyzer.read_json_file("pytest_data/tones_5_fail.json")
    expected_dict = {}
    assert toneanalyzer.parse_toneanalyzer_response(json_content) == expected_dict, \
           "when run on file 'data/tones_1.json' should return {}".format(expected_dict)
