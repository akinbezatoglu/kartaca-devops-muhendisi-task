from flask import Flask, jsonify, make_response
from elasticsearch import Elasticsearch
import random

class FlaskAppWrapper(object):

    def __init__(self, app):
        self.app = app
        self.client = Elasticsearch(
            hosts=[
                "http://elasticsearch1:9200",
                "http://elasticsearch2:9200",
                "http://elasticsearch3:9200",
            ]
        )
    
    def get_one_random_province(self):
        return self.client.get(index="iller", id=str(random.randint(1, 10)))

    def add_endpoint(self, endpoint=None, endpoint_name=None, handler=None, methods=['GET'], *args, **kwargs):
        self.app.add_url_rule(endpoint, endpoint_name, handler, methods=methods, *args, **kwargs)

    def run(self, **kwargs):
        self.app.run(**kwargs)


flask_app = Flask(__name__)
app = FlaskAppWrapper(flask_app)

def hello():
    return "Merhaba, Python!"

def staj():
    resp = app.get_one_random_province()
    if resp["found"]:
        return jsonify(resp['_source'])
    else:
        return make_response({"error": "Internal Server Error"}, 500)

app.add_endpoint('/', 'hello', hello, methods=['GET'])
app.add_endpoint('/staj', 'staj', staj, methods=['GET'])

if __name__ == "__main__":
    app.run(host='0.0.0.0', debug=True, port=4444)