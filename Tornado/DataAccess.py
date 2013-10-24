import tornado.ioloop
import tornado.web
import tornadoredis
import re

class DataAccess(object):
    
    @tornado.gen.engine
    def cardsForSet(self, setName, callback):
        c = tornadoredis.Client()
        c.connect()
        name = self.escape(setName)
        guids = yield tornado.gen.Task(c.smembers, "set-cards-" + name)
        keys = set("card-" + guid for guid in guids)
        
        rawResults = yield tornado.gen.Task(c.mget, keys)
        callback("[" + ','.join(rawResults) + "]")
    
    def escape(self, string):
        return re.sub(r'-+', "-", re.sub(r'\W', "-", string).lower())