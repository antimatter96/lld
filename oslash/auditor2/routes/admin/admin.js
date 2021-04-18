module.exports = async function (fastify, opts) {
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
}
