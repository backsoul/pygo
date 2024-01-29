import os
import mysql.connector
from flask import Flask, request, jsonify

app = Flask(__name__)

# CREATE


@app.route('/create', methods=['POST'])
def create():
    name = request.json['Name']
    conn = mysql.connector.connect(
        host=os.environ.get("MYSQL_HOST", "localhost"),
        user=os.environ.get("MYSQL_USER", "your_mysql_username"),
        password=os.environ.get("MYSQL_PASSWORD", "your_mysql_password"),
        database=os.environ.get("MYSQL_DATABASE", "your_mysql_database")
    )
    cur = conn.cursor()
    sql = "INSERT INTO data (name) VALUES (%s)"
    cur.execute(sql, (name,))
    conn.commit()
    conn.close()
    response = {'Message': 'Success'}
    return jsonify(response)

# READ


@app.route('/read', methods=['GET'])
def read():
    conn = mysql.connector.connect(
        host=os.environ.get("MYSQL_HOST", "localhost"),
        user=os.environ.get("MYSQL_USER", "your_mysql_username"),
        password=os.environ.get("MYSQL_PASSWORD", "your_mysql_password"),
        database=os.environ.get("MYSQL_DATABASE", "your_mysql_database")
    )
    cur = conn.cursor()
    sql = "SELECT * FROM data"
    cur.execute(sql)
    data = cur.fetchone()
    conn.close()
    response = {'Name': data[0]} if data else {'Message': 'No data available'}
    return jsonify(response)

# UPDATE


@app.route('/update', methods=['PUT'])
def update():
    name = request.json['Name']
    conn = mysql.connector.connect(
        host=os.environ.get("MYSQL_HOST", "localhost"),
        user=os.environ.get("MYSQL_USER", "your_mysql_username"),
        password=os.environ.get("MYSQL_PASSWORD", "your_mysql_password"),
        database=os.environ.get("MYSQL_DATABASE", "your_mysql_database")
    )
    cur = conn.cursor()
    sql = "UPDATE data SET name = %s"
    cur.execute(sql, (name,))
    conn.commit()
    conn.close()
    response = {'Message': 'Success'}
    return jsonify(response)

# DELETE


@app.route('/delete', methods=['DELETE'])
def delete():
    name = request.json['Name']
    conn = mysql.connector.connect(
        host=os.environ.get("MYSQL_HOST", "localhost"),
        user=os.environ.get("MYSQL_USER", "your_mysql_username"),
        password=os.environ.get("MYSQL_PASSWORD", "your_mysql_password"),
        database=os.environ.get("MYSQL_DATABASE", "your_mysql_database")
    )
    cur = conn.cursor()
    sql = "DELETE FROM data WHERE name = %s"
    cur.execute(sql, (name,))
    conn.commit()
    conn.close()
    response = {'Message': 'Success'}
    return jsonify(response)


if __name__ == '__main__':
    app.run(threaded=True, host="0.0.0.0", port=8000)
