import discovery_query
import pytest
import json


def read_json_file(file_path):
    """Reads and parse a json file.

    Parameters
    ----------
    file_path : {str} the path to the json file.

    Returns
    -------
    dict : a dictionary containing the json structure read from the file.
    """
    with open(file_path) as json_file:
        json_content = json_file.read()
        json_data = json.loads(json_content)

    return(json_data)


def test_display_results():
    json_data = read_json_file('pytest_data/query_ex1.json')
    output = discovery_query.display_results(json_data)
    assert type(output) is dict, "display_results() should return a dict"
    assert "query" in output, "display_results() should have a key 'query'"
    assert "count" in output, "display_results() should have a key 'count'"
    assert type(output['query']) is str, "key 'query' of output dict should contain a string"

    json_data = read_json_file('pytest_data/query_ex2.json')
    output = discovery_query.display_results(json_data)
    assert type(output) is dict, "display_results() should return a dict"
    assert "count" in output, "display_results() should have a key 'count'"
    assert type(output['query']) is str, "key 'query' of output dict should contain a string"
