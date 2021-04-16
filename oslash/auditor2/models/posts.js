const { DataTypes, Model } = require('sequelize');

class Post extends Model { }

async function init(sequelize) {
  return new Promise((resolve, _reject) => {
    Post.init({
      content: DataTypes.STRING,
    }, {
      sequelize, modelName: 'posts'
    });
    resolve()
  })
}

module.exports = {
  init: init,
  model: Post
};
