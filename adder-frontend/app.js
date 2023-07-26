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

app.listen(port, () => {
  console.log(`App listening on port ${port}!`);
});

module.exports = app;