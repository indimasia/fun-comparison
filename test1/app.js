const mysql = require('mysql2');

const connection = mysql.createConnection({
  host     : 'localhost',
  port     : '33061',
  user     : 'node',
  password : 'password',
  database : 'comparison'
});

function randomString(length) {
  let result           = '';
  let characters       = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789';
  let charactersLength = characters.length;
  for (let i = 0; i < length; i++) {
    result += characters.charAt(Math.floor(Math.random() * charactersLength));
  }
  return result;
}

connection.connect();

let startTime = process.hrtime();

for (let i = 0; i < 100000; i++) {
  let name = randomString(10);
  let salary = Math.floor(Math.random() * (100000 - 10000 + 1)) + 10000;
  let greeting = (Math.random() < 0.5) ? 'Mr' : 'Ms';

  connection.query('INSERT INTO employees SET ?', {name: name, salary: salary, greeting: greeting}, (error) => {
    if (error) throw error;

    // Check if it's the last iteration
    if (i === 99999) {
      let endTime = process.hrtime(startTime);
      console.log(`Node.js script execution time: ${endTime[0] + endTime[1] / 1e9} seconds`);
      connection.end();
    }
  });
}
