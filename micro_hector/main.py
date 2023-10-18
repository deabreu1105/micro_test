from flask import Flask
from flask import request

app = Flask(__name__)
@app.route('/saludo', methods=['GET'])
def saludo():
   if request.method == 'GET':
      return "hola munndo"
if __name__ == "__main__":
    app.run(host="0.0.0.0", port=8080, debug=True)
