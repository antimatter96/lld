const { DataTypes, Model } = require('sequelize');

class DelayedJob extends Model { }

async function init(sequelize) {
  return new Promise((resolve, _reject) => {
    DelayedJob.init({
      approved: {
        type: DataTypes.ENUM("pending", "rejected",
          "approved", "done"),
        defaultValue: "pending"
      },
      approvedAt: DataTypes.TIME,
      content: DataTypes.STRING,
      action: {type:DataTypes.STRING, allowNull:false},
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
