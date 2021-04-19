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

  fastify.get('/actions/list', async function (request, reply) {
    console.log(fastify.models.Post.create)
    return 'this is an example'
  });

  fastify.post('/actions/approve', async function (request, reply) {
    console.log(request.body);
    let id = parseInt(request.body.jobId);

    if (isNaN(id)) {
      reply.code(400).send({
        status: "ERROR",
        error: "Need a id",
      })
      return
    }

    try {
      await fastify.models.AuditLog.create({
        'action': "APPROVE",
        'content': id,
        'creatorId': request.body.userId,
      })

      let dj = await fastify.models.DelayedJob.findByPk(id);

      if (dj == null) {
        reply.code(404).send({
          status: "ERROR",
          error: "Not found",
        })
        return
      }

      if (dj.state != "pending") {
        reply.code(400).send({
          status: "ERROR",
          error: "State is " + dj.state,
        })
        return
      }

      await dj.update({
        state: "approved",
        approvedAt: new Date()
      })

      console.log(dj.dataValues);

      switch (dj.dataValues.action) {
        case "CREATE": {
          console.log(dj.action);

          let post = await fastify.models.Post.create({
            'userId': dj.dataValues.onBehalOfId,
            'content': dj.dataValues.content
          })

          await dj.update({
            state: "done",
            approvedAt: new Date()
          })

          reply.code(200).send({
            status: "SUCCESS",
            data: post.dataValues
          });

          break;
        }
        default:
          break;
      }

      reply.code(200).send({
        status: "SUCCESS",
        data: "log",
      });

    } catch (error) {
      return {
        status: "ERROR",
        error: error.toString()
      }
    }
  });
}
