const { DataTypes, Model } = require('sequelize');

class AuditLog extends Model { }

async function init(sequelize) {
  return new Promise((resolve, _reject) => {
    AuditLog.init({
      action: DataTypes.STRING,
      postId: DataTypes.INTEGER,
      content: DataTypes.STRING
    }, {
      sequelize,
      modelName: 'audit_log',
      paranoid: true,
    });
    resolve()
  });
}

module.exports = {
  init: init,
  model: AuditLog
};
