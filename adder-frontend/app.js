const express = require('express');
const axios = require('axios');
const app = express();
const port = 3000;

// Here we are configuring express to use body-parser as middle-ware and set up EJS.
app.use(express.json());
app.use(express.urlencoded({ extended: true }));
app.set('view engine', 'ejs');

app.get('/', (req, res) => {
  res.render('index', { response: '', server: '', method: 'GET', body: '' });
});


app.get('/second', async (req, res) => {
  let server = req.query.server || 'host.docker.internal';

  if (!server.startsWith('http://') && !server.startsWith('https://')) {
      server = 'http://' + server;
  }

  try {
      const response = await axios.get(`${server}:8080`);
      const routesResponse = await axios.get(`${server}:8080/routes`);
      console.log(routesResponse.data);
      res.render('second', { response: JSON.stringify(response.data, null, 2), routes: routesResponse.data });
  } catch (error) {
      res.render('error', { error: error.toString() });
  }
});


app.post('/makeRequest', async (req, res) => {
    let { server, method, body } = req.body;
    let response;
    try {
        // Check if the server URL starts with http:// or https://, and if not, prepend with "http://"
        if (!server.startsWith('http://') && !server.startsWith('https://')) {
            server = 'http://' + server;
        }

        if (method === 'GET') {
            response = await axios.get(server);
        } else if (method === 'POST') {
            // Assuming the body is in JSON format, you might need to handle parsing the JSON safely.
            response = await axios.post(server, JSON.parse(body));
        }
        // Add more methods as needed
        response = response.data;
    } catch (error) {
        response = error.toString();
    }
    res.render('index', { response: JSON.stringify(response, null, 2), server, method, body });
});

app.listen(port, () => {
  console.log(`Example app listening at http://localhost:${port}`)
});