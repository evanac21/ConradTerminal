const exp = require('express')
const app = exp()
const port = 3000
var path = require('path');

app.get('/', function(req, res) {
    res.sendFile(path.join(__dirname + '/index.html'));
});

  app.listen(3000)