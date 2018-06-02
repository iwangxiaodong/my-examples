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
    "lifetimeDuration": "86400s",
    "iceServers": [
      {
        "urls": [
          "stun:108.177.98.127:19302",
          "stun:[2607:F8B0:400E:C06::7F]:19302"
        ]
      },
      {
        "urls": [
          "turn:173.194.202.127:19305?transport=udp",
          "turn:[2607:F8B0:400E:C06::7F]:19305?transport=udp",
          "turn:173.194.202.127:19305?transport=tcp",
          "turn:[2607:F8B0:400E:C06::7F]:19305?transport=tcp"
        ],
        "username": "CO+Uz9gFEgZIZx7cKRMYzc/s6OMTIICjBQ",
        "credential": "j/qS8db5h+/6WrjZ0x/bGaWgBxA="
      }
    ],
    "blockStatus": "NOT_BLOCKED",
    "iceTransportPolicy": "all"
  })
}

app.get('/iceconfig', f)
app.post('/iceconfig', f)

app.listen('3033', function () {
  console.log('server started')
})
