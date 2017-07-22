var ws = require('ws');
var notify = require('./notify');
var driver = new ws('ws://localhost:8080/restconf/streams','', {origin:'localhost:8080'});
var n = new notify.handler(driver);
n.on('', 'update', 'car', (car, err) => {
  console.log(car);
});
