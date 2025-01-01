from flask import Flask, request, jsonify
from langchain.chains import RetrievalQA
from langchain.chat_models import ChatOpenAI
from langchain.vectorstores import Weaviate

app = Flask(__name__)

llm = ChatOpenAI(model="gpt-4", openai_api_key="YOUR_API_KEY")
retriever = Weaviate(client_config={"host": "localhost", "port": 8081})
qa = RetrievalQA(llm=llm, retriever=retriever)

@app.route('/chat', methods=['POST'])
def chat():
    data = request.json
    query = data.get('query', '')
    response = qa.run(query)
    return jsonify({"response": response})

if __name__ == "__main__":
    app.run(host="0.0.0.0", port=5000)
