from flask import Flask, jsonify, make_response
from elasticsearch import Elasticsearch
import random

class ES:
    def __init__(self):
        self.client = Elasticsearch(
            hosts=[
                "http://elasticsearch1:9200",
                "http://elasticsearch2:9200",
                "http://elasticsearch3:9200",
            ]
        )

    def get_one_random_province(self):
        idx = "iller"
        doc_id = str(random.randint(1, 10))
        resp = self.client.get(index=idx, id=doc_id)
        if resp["found"]:
            return resp['_source']
        else:
            return False

app = Flask(__name__)

@app.route("/")
def greet():
    return "Merhaba Python!"

@app.route("/staj")
def province():
    elastic = ES()
    c = elastic.get_one_random_province()
    if c:
        return jsonify(c)
    else:
        make_response({"error": "Internal Server Error"}, 500)

if __name__ == "__main__":
    app.run(host='0.0.0.0', debug=True, port=4444)