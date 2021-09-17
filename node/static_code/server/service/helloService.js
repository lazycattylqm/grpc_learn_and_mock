const { messages } = require('proto')

const sayHello = (call, callback) => {
    let reply = new messages.HelloReply();
    const name = call.request.getName();
    reply.setMessage(`hello ${name}, this is node static code gen server`)
    callback(null, reply)
}

module.exports = {
    sayHello
}