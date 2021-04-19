const { requestContext } = require('fastify-request-context')


module.exports = async function (fastify, opts) {
  fastify.addHook('preHandler', async (request, reply) => {
    const user = requestContext.get('user');

    if (user.role != 'admin' && user.role != 'super_admin') {
      reply.code(401).send({
        status: "ERROR",
        error: "You need to be an admin",
      })
      return
    }
  })

  fastify.get('/post/:id', async function (request, reply) {
    console.log(fastify.models.Posts.create)
    return 'this is an example'
  });

  fastify.post('/post/create', async function (request, reply) {
    try {
      await fastify.models.AuditLog.create({
        'action': "CREATE",
        'content': JSON.stringify(request.body.content),
        'onBehalfOfId': request.body.actingAs,
        'creatorId': request.body.userId,
      })

      let delayedJob = await fastify.models.DelayedJob.create({
        "action": "CREATE",
        'content': JSON.stringify(request.body.content),
        'creatorId': request.body.userId,
        'onBehalOfId': request.body.actingAs,
      })

      await fastify.models.AuditLog.create({
        'action': "CREATED",
        'content': JSON.stringify(request.body.content),
        'onBehalfOfId': request.body.actingAs,
        'creatorId': request.body.userId,
      })

      return {
        status: "SUCCESS",
        data: delayedJob.dataValues
      }
    } catch (error) {
      return {
        status: "ERROR",
        error: error.toString()
      }
    }
  });

  fastify.post('/post/edit/:id', async function (request, reply) {
    let id = parseInt(request.params.id);

    if (isNaN(id)) {
      reply.code(400).send({
        status: "ERROR",
        error: "Need a valid id",
      })
      return
    }

    try {
      let post = await fastify.models.Post.findByPk(id);

      if (post == null) {
        reply.code(404).send({
          status: "ERROR",
          error: "Not found",
        })
        return
      }

      await fastify.models.AuditLog.create({
        'action': "EDIT",
        'postId': id,
        'content': JSON.stringify({
          from: post.content,
          to: request.body.content
        }),
        'onBehalfOfId': request.body.actingAs,
        'creatorId': request.body.userId,
      })

      let delayedJob = await fastify.models.DelayedJob.create({
        "action": "EDIT",
        'postId': id,
        'content': JSON.stringify({
          from: post.content,
          to: request.body.content
        }),
        'creatorId': request.body.userId,
        'onBehalOfId': request.body.actingAs,
      })

      await fastify.models.AuditLog.create({
        'action': "EDITED",
        'postId': id,
        'content': JSON.stringify({
          from: post.content,
          to: request.body.content
        }),
        'onBehalfOfId': request.body.actingAs,
        'creatorId': request.body.userId,
      })

      return {
        status: "SUCCESS",
        data : delayedJob.dataValues
      }
    } catch (error) {
      return {
        status: "ERROR",
        error : error.toString()
      }
    }
  });
}
