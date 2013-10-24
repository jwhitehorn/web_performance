var redisLib = require("redis")

DataAccess = {
  init: function(){
    var self = Object.create(DataAccess);
    var redis = redisLib.createClient();
    var escape = function(string){
      return string.replace(/\W/g, "-").toLowerCase().replace(/-+/g, '-');
    };
    var objectsForKeys = function(keys, block){
      if(keys.length <= 0){
        block("[]");
      }else{
        redis.mget(keys, function(e, objs){
          block("[" + objs.join(",") +  "]");
        });
      }
    };
    
    self.allSets = function(block){
      redis.smembers("sets", function(e, set_names){
        var sets = [];
        set_names.forEach(function(name){
          sets.push({id: escape(name), name: name});
        });
        block(sets);
      });
    };
    
    self.cardsInSet = function(name, block){
      setName = escape(name);
      redis.smembers("set-cards-" + setName, function(e, guids){
        var keys = guids.map(function(guid){ return "card-" + guid; });
        objectsForKeys(keys, block);
      });
    };
    
    self.close = function(){
      redis.quit();
    };
    
    return self;
  },
  
  begin: function(block){
    block(DataAccess.init());
  }
}