const express = require('express');
const axios = require('axios');
const bodyParser = require('body-parser');
const { addHttp } = require('./util');
const app = express();
const port = 3000;

let defaultServer = process.env.HOST_ADDRESS || "localhost"; // renamed to defaultServer

app.use(bodyParser.urlencoded({ extended: true }));
app.set('view engine', 'ejs');

app.get('/', (req, res) => {
  res.render('index', { server: '', method: '', path: '', body: '', response: '' });
});

app.post('/', async (req, res) => {
  let { server, method, path, body } = req.body;
  server = addHttp(server || defaultServer);

  try {
    let response = await axios({ method, url: server + path, data: body });
    console.log(response.data);  // Log the response data
    res.render('index', { server, method, path, body, response: JSON.stringify(response.data, null, 2) });
  } catch (error) {
    res.render('index', { server, method, path, body, response: JSON.stringify({ message: error.message, stack: error.stack }, null, 2) });
  }
});

app.get('/second', async (req, res) => {
  let server = process.env.HOST_ADDRESS || "localhost";
  try {
    let testResponse = await axios.get(addHttp(server + ':8080'));
    console.log("Test response: ", testResponse);

    try {
      let response = await axios.get(addHttp(server + ':8080/routes'));
      let routes = response.data.routes;
      res.render('second', { response: testResponse.data, routes });
    } catch (error) {
      console.log("Error getting routes: ", error);
      res.render('error', { error: JSON.stringify(error, null, 2) });
    }
  } catch (error) {
    console.log("Error making test request: ", error);
    res.render('error', { error: JSON.stringify(error, null, 2) });
  }
});

app.get('/third', (req, res) => {
  res.render('third');
});

app.get('/fourth', (req, res) => {
  res.render('fourth');
});


// Route to get the list of databases
app.get('/databases', async (req, res) => {
  let server = process.env.HOST_ADDRESS || "localhost";
  try {
      let response = await axios.get(`http://${server}:8080/dblist`);
      res.json(response.data);  // Adjust this line
  } catch (error) {
      console.log("Error fetching database list: ", error);
      res.status(500).json({ message: "Failed to fetch database list." });
  }
});



// Route to get the fields of a specific database
app.get('/fields/:dbName', async (req, res) => {
  let server = process.env.HOST_ADDRESS || "localhost";
  try {
      let response = await axios.get(`http://${server}:8080/fields/${req.params.dbName}`);
      res.json(response.data);
  } catch (error) {
      console.log("Error fetching fields: ", error);
      res.status(500).json({ message: "Failed to fetch fields." });
  }
});



app.listen(port, () => {
  console.log(`App listening on port ${port}!`);
});

module.exports = app;