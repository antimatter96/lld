var Users = require("./user");
var AuditLog = require("./log");
var DelayedJob = require("./delayed");
var Posts = require("./posts");

var Associations = require("./associations");

module.exports = {
  _init: async function (sequelize) {
    return new Promise(async (resolve, _reject) => {
      await Users.init(sequelize);
      await AuditLog.init(sequelize);
      await DelayedJob.init(sequelize);
      await Posts.init(sequelize)

      await Associations.init(Users.model, AuditLog.model, DelayedJob.model, Posts.model);

      await sequelize.sync()
      resolve();
    })
  },
  Users: Users,
  AuditLog: AuditLog,
  DelayedJob: DelayedJob,
  Posts: Posts
};
