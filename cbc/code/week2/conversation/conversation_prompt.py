# -*- coding: utf-8 -*-
import json
from watson_developer_cloud import ConversationV1


# BEGIN of python-dotenv section
from os.path import join, dirname
from dotenv import load_dotenv
import os
# END of python-dotenv section

# Replace with the context obtained from the initial request

def parse_conversation_response(json_data):
    out_dict = {}
    for output in json_data:
        output_type = type(json_data[output])
        if output_type is dict:
            for key in json_data[output]:
                if output == 'context':
                        out_dict[output] = json_data[output]
                if output == 'output':
                    for value in json_data[output][key]:
                        if key == 'text':
                            if key in out_dict:
                                out_dict[key] = out_dict[key] + "\n" + value
                            else:
                                out_dict[key] = value
    return(out_dict)


def handle_conversation_data(parsed_data):
    print("Conversation>",parsed_data['text'])
    return (parsed_data['context'])

def main_loop(conversation, workspace_id):

    # just for debugging
    print('*** Connecting to your workspace...')

    # connects to conversation to start dialog, retrieves the json respose.
    response = conversation.message(workspace_id=workspace_id)

    # sends the response to our functions
    data = parse_conversation_response(response)
    context = handle_conversation_data(data)

    # prompt at the beginning of the conversation
    print('(say anything... type quit to quit)')
    # let's do an infinite loop
    while True:
        # gets some input from the prompt
        input_content = input('You> ')

        # if you type one of those words, it will exit the while loop
        if (input_content.lower() in {'exit', 'quit', 'q', 'n'}):
            break

        # sends the input to conversation and retrieves the json response.
        response = conversation.message(workspace_id=workspace_id,
                                        message_input={'text': input_content},
                                        context=context)

        # sends the response to our functions
        data = parse_conversation_response(response)
        context = handle_conversation_data(data)




if __name__ == '__main__':
    # BEGIN of python-dotenv section
    dotenv_path = join(dirname(__file__), '.env')
    load_dotenv(dotenv_path)
    # END of python-dotenv section

    # create an instance of the API ConversationV1 object
    conversation_apiobj = ConversationV1(
        username=os.environ.get("CONVERSATION_USERNAME"),
        password=os.environ.get("CONVERSATION_PASSWORD"),
        version='2016-09-20')

    # obtain the workspace id from dotenv
    workspace_id_env = os.environ.get("CONVERSATION_WORKSPACE_ID")

    main_loop(conversation_apiobj, workspace_id_env)
