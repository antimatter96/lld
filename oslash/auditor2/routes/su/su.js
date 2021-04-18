module.exports = async function (fastify, opts) {
  fastify.get('/actions/list', async function (request, reply) {
    console.log(fastify.models.Posts.create)
    return 'this is an example'
  });

  fastify.post('/actions/approve', async function (request, reply) {
    //console.log(request.body);
    try {
      await fastify.models.AuditLog.create({
        'action': "CREATE",
        'content': JSON.stringify(request.body.content),
        'userId': request.body.userId,
      })

      let params = {
        userId : request.body.userId,
        content : JSON.stringify(request.body.content),
      }

      let post = await fastify.models.Posts.create(params)

      await fastify.models.AuditLog.create({
        'action': "CREATED",
        'postId': post.id,
        'userId': request.body.userId,
      })

      return {
        status: "SUCCESS",
        data : post.dataValues
      }
    } catch (error) {
      return {
        status: "ERROR",
        error : error.toString()
      }
    }
  });
}
