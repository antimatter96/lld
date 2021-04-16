const { Sequelize } = require('sequelize');
const sequelize = new Sequelize('os', 'root', '661996_aj', {
  host: 'localhost',
  dialect: 'mysql'
});

module.exports = sequelize
