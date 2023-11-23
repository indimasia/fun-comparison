import mysql.connector
import time
import random
import string

def random_string(length):
    return ''.join(random.choice(string.ascii_letters + string.digits) for _ in range(length))

cnx = mysql.connector.connect(user='root', password='',
                              port=33061,
                              host='localhost',
                              database='comparison')
cursor = cnx.cursor()

startTime = time.time()

for _ in range(100000):
    name = random_string(10)
    salary = random.randint(10000, 100000)
    greeting = random.choice(['Mr', 'Ms'])
    cursor.execute("INSERT INTO employees (name, salary, greeting) VALUES (%s, %s, %s)", (name, salary, greeting))

for _ in range(100000):
    name = random_string(10)
    salary = random.randint(10000, 100000)
    greeting = random.choice(['Mr', 'Ms'])
    cursor.execute("INSERT INTO employees2 (name, salary, greeting) VALUES (%s, %s, %s)", (name, salary, greeting))

cnx.commit()
endTime = time.time()
print(f"Python script execution time: {endTime - startTime} seconds")

cursor.close()
cnx.close()
