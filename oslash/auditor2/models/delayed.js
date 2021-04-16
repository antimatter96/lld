const { DataTypes, Model } = require('sequelize');

class DelayedJob extends Model { }

async function init(sequelize) {
  return new Promise((resolve, _reject) => {
    DelayedJob.init({
      approved: DataTypes.BOOLEAN,
      approvedAt: DataTypes.TIME,
      content: DataTypes.STRING,
      action: DataTypes.STRING,
    }, {
      sequelize, modelName: 'delayed_job'
    });
    resolve()
  })
}
module.exports = {
  init: init,
  model: DelayedJob
};
