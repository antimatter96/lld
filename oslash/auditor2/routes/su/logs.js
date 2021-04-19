const { requestContext } = require('fastify-request-context')

module.exports = async function (fastify, opts) {

  fastify.addHook('preHandler', async (request, reply) => {
    const user = requestContext.get('user');

    if (user.role != 'super_admin') {
      reply.code(401).send({
        status: "ERROR",
        error: "You need to be a super admin",
      })
      return
    }
  })

  fastify.get('/logs/:id', async function (request, reply) {
    let id = parseInt(request.params.id);

    if (isNaN(id)) {
      reply.code(400).send({
        status: "ERROR",
        error: "Need a id",
      })
      return
    }

    try {
      let log = await fastify.models.AuditLog.findByPk(id);

      if (log == null) {
        reply.code(404).send({
          status: "ERROR",
          error: "Not found",
        })
        return
      }

      reply.code(200).send({
        status: "SUCCESS",
        data: log,
      });

    } catch (error) {
      return {
        status: "ERROR",
        error: error.toString()
      }
    }
  });

  fastify.get('/logs', async function (request, reply) {
    try {
      return {
        status: "SUCCESS",
        data: "a"
      }
    } catch (error) {
      return {
        status: "ERROR",
        error: error.toString()
      }
    }
  });
}
