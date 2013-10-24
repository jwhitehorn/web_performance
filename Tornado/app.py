import tornado.ioloop
import tornado.web
import tornadoredis
import re
import DataAccess

class MainHandler(tornado.web.RequestHandler):
    
    @tornado.web.asynchronous
    def get(self, setName):
        def completion(results):
            self.write(results)
            self.finish()
            
        self.set_header('Content-Type', 'application/json')
        dao = DataAccess.DataAccess()
        dao.cardsForSet(setName, completion)

application = tornado.web.Application([
    (r'/api/sets/(.*)/cards', MainHandler),
])

if __name__ == "__main__":
    application.listen(9292)
    tornado.ioloop.IOLoop.instance().start()