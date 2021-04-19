'use strict'

const fp = require('fastify-plugin')

async function addUser(fastify, options, next) {

  await fastify.addHook('preHandler', async (req, reply) => {
    let user = await fastify.models.User.findByPk(req.body.userId);

    if (user == null) {
      reply.code(401).send({
        status: "ERROR",
        error: "You need to give userId, apiKey",
      })
      return
    }

    // Do API KEY VALIDATION

    req.requestContext.set('user', user.dataValues);
  });

  next()
}

module.exports = fp(addUser, {
  fastify: '>=1.0.0',
  name: 'addUser'
})
