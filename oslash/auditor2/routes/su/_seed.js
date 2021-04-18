module.exports = async function (fastify, opts) {
  fastify.post('/_seed', async function (request, reply) {
    try {
      let u1 = await fastify.models.User.create({
        username: "some basic"
      });
      let u2 = await fastify.models.User.create({
        username: "baby admin",
        role: "admin"
      })
      let u3 = await fastify.models.User.create({
        username: "super admin",
        role: "super_admin"
      })

      console.log(u1, u2, u3);

      reply.code(200).send({
        status: "SUCCESS",
        data: "OK",
      });

    } catch (error) {
      reply.code(500).send({
        status: "ERROR",
        data: error.toString(),
      });
    }
  });
}
