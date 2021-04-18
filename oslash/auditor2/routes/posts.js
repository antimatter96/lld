'use strict'

module.exports = async function (fastify, opts) {
  fastify.get('/post/:id', async function (request, reply) {
    let id = parseInt(request.params.id);

    if (isNaN(id)) {
      reply.code(400).send({
        status: "ERROR",
        error: "Need a id",
      })
      return
    }

    try {
      let log = await fastify.models.Post.findByPk(id);

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

  fastify.post('/post/create', async function (request, reply) {
    //console.log(request.body);
    try {
      console.log(request.body.content)
      await fastify.models.AuditLog.create({
        'action': "CREATE",
        'content': JSON.stringify(request.body.content),
        'userId': request.body.userId,
      })

      let params = {
        userId : request.body.userId,
        content : JSON.stringify(request.body.content),
      }

      let post = await fastify.models.Post.create(params)

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
        'content': JSON.stringify({
          from: post.content,
          to: request.body.content
        }),
        'userId': request.body.userId,
      })

      await post.update({content: JSON.stringify(request.body.content)})

      await fastify.models.AuditLog.create({
        'action': "EDITED",
        'postId': post.id,
        'content': JSON.stringify({
          from: post.content,
          to: request.body.content
        }),
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
