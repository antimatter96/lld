'use strict'

const fp = require('fastify-plugin')
const models = require('../models/index')

function fastifyMysql (fastify, options, next) {
  models._init(options.dbConnection).then((done) => {
    delete models._init;

    for (let model in models ) {
      delete models[model].init
      models[model] = models[model].model
    }

    fastify.addHook('onClose', (fastify, done) => {
      options.dbConnection.close()
        .then(() => { 
          console.log("Closed db connections");
          done;
        }).catch(done)
    });

    fastify.decorate('models', models)

    next()
  }).catch((err) => {
    console.log(err, "ERROR")
  });
}

module.exports = fp(fastifyMysql, {
  fastify: '>=1.0.0',
  name: 'models'
})
