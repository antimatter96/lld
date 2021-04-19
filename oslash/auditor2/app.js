'use strict'

const { fastifyRequestContextPlugin } = require('fastify-request-context')
const path = require('path')
const AutoLoad = require('fastify-autoload')

var db = require('./db/db');

module.exports = async function (fastify, opts) {
  opts.dbConnection = db;

  fastify.register(require('middie'))

  fastify.register(fastifyRequestContextPlugin);

  // Place here your custom code!

  // Do not touch the following lines

  // This loads all plugins defined in plugins
  // those should be support plugins that are reused
  // through your application
  fastify.register(AutoLoad, {
    dir: path.join(__dirname, 'plugins'),
    options: Object.assign({}, opts)
  })

  // This loads all plugins defined in routes
  // define your routes in one of these
  fastify.register(AutoLoad, {
    dir: path.join(__dirname, 'routes'),
    options: Object.assign({}, opts)
  })
}
