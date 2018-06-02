var express = require('express')
var crypto = require('crypto')
var app = express()

var hmac = function (key, content) {
  var method = crypto.createHmac('sha1', key)
  method.setEncoding('base64')
  method.write(content)
  method.end()
  return method.read()
}

var f = function (req, resp) {
  var query = req.query
  var key = '4080218913'
  var time_to_live = 600
  var timestamp = Math.floor(Date.now() / 1000) + time_to_live
  var turn_username = timestamp + ':ninefingers'
  var password = hmac(key, turn_username)
  
  resp.header("Access-Control-Allow-Origin", "*")

  return resp.send({
    iceServers: [
      {
        urls: [
          'stun:stun1.l.google.com:19302',
          'turn:ICE_SERVER_ADDR:3478'
        ]
      }
    ]
  })
}

app.get('/iceconfig', f)
app.post('/iceconfig', f)

app.listen('3033', function () {
  console.log('server started')
})
