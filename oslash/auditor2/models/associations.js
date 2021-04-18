async function init(User, AuditLog, DelayedJob, Post) {
  return new Promise((resolve, _reject) => {
    AuditLog.belongsTo(User, { as: 'creator' })
    AuditLog.belongsTo(User, { as: 'onBehalfOf' })

    DelayedJob.belongsTo(User, { as: 'approver' })
    DelayedJob.belongsTo(User, { as: 'creator' })
    DelayedJob.belongsTo(User, { as: 'onBehalOf' })

    Post.belongsTo(User)

    resolve()
  })
}

module.exports = {
  init: init
}
