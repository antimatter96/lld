const { requestContext } = require('fastify-request-context')

module.exports = async function (fastify, opts) {
  fastify.addHook('preHandler', async (request, reply) => {
    const user = requestContext.get('user');

    if (user.role != 'super_admin') {
      reply.code(401).send({
        status: "ERROR",
        error: "You need to be a super admin",
      });
      return
    }
  })

  fastify.post('/actions/list', async function (request, reply) {
    try {
      let djs = await fastify.models.DelayedJob.findAll()

      reply.code(400).send({
        status: "ERROR",
        data: djs
      });
    } catch (error) {

    }
  });

  fastify.post('/actions/approve', async function (request, reply) {
    let id = parseInt(request.body.jobId);

    if (isNaN(id)) {
      reply.code(400).send({
        status: "ERROR",
        error: "Need a id",
      });
      return
    }

    try {
      await fastify.models.AuditLog.create({
        'action': "APPROVE",
        'content': id,
        'creatorId': request.body.userId,
      });

      let dj = await fastify.models.DelayedJob.findByPk(id);

      if (dj == null) {
        reply.code(404).send({
          status: "ERROR",
          error: "Not found",
        });
        return
      }

      if (dj.state != "pending") {
        reply.code(400).send({
          status: "ERROR",
          error: "State is " + dj.state,
        });
        return
      }

      await dj.update({
        state: "approved",
        approvedAt: new Date(),
        approvedBy: request.body.userId,
      });

      console.log(dj.dataValues);

      switch (dj.dataValues.action) {
        case "CREATE": {
          console.log(dj.action);

          let post = await fastify.models.Post.create({
            'userId': dj.dataValues.onBehalOfId,
            'content': dj.dataValues.content
          });

          await dj.update({
            state: "done",
            approvedAt: new Date()
          });

          reply.code(200).send({
            status: "SUCCESS",
            data: post.dataValues
          });

          break;
        }
        case "EDIT": {
          let post = await fastify.models.Post.findByPk(dj.postId)

          let content = JSON.parse(dj.content);

          if (JSON.stringify(post.content) != JSON.stringify(content.from)) {
            await dj.update({
              state: "approved",
              approvedAt: new Date()
            });

            reply.code(400).send({
              status: "FAILURE",
              error: `Changed is ${post.content}, needed it to be ${JSON.stringify(content.from)}`
            });

            return
          }

          await post.update({ content: JSON.stringify(content.to) })

          await dj.update({
            state: "done",
            approvedAt: new Date()
          });

          reply.code(200).send({
            status: "SUCCESS",
            data: post.dataValues
          });

          break;
        }
        default:
          break;
      }

      await fastify.models.AuditLog.create({
        'action': "APPROVED",
        'content': id,
        'creatorId': request.body.userId,
      });

    } catch (error) {
      return {
        status: "ERROR",
        error: error.toString()
      }
    }
  });
}
