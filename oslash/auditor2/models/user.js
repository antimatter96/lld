const { DataTypes, Model } = require('sequelize');

class User extends Model { }

async function init(sequelize) {
  return new Promise((resolve, _reject) => {
    User.init({
      username: DataTypes.STRING,
      role: DataTypes.ENUM('user', 'admin', 'super_admin')
    }, {
      sequelize, modelName: 'user'
    });
    resolve()
  });
}


module.exports = {
  init: init,
  model: User
};
