const { DataTypes, Model } = require('sequelize');

class User extends Model { }

async function init(sequelize) {
  return new Promise((resolve, _reject) => {
    User.init({
      username: {
        type: DataTypes.STRING,
        allowNull: false,
        unique: true
      },
      role: {
        type: DataTypes.ENUM('user', 'admin', 'super_admin'),
        allowNull: false,
        defaultValue: "user"
      }
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
