require('./data_access.js');
var express = require('express');
var app = express();

app.get('/api/sets/:name/cards', function(req, res){
  var setName = req.params.name;
  DataAccess.begin(function(dao){
    dao.cardsInSet(setName, function(cards){
      res.set('Content-Type', 'application/json');
      res.send(cards);
      dao.close();
    });
  });
});

app.get('/api/sets', function(req, res){
  DataAccess.begin(function(dao){
    dao.allSets(function(sets){
      res.set('Content-Type', 'application/json');
      res.send(sets);
      dao.close();
    });
  });
});

app.listen(9292);
console.log('Listening on port 9292');