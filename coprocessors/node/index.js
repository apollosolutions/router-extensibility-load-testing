const { randomUUID: uuid } = require('crypto')
const express = require('express')
const jwt = require('jsonwebtoken')

const JWT_SECRET = process.env.JWT_SECRET || 'apollo'

function handleCreateJwt(req, res) {
  const payload = {
    client_name: req.body?.clientName || 'node coprocessor',
    client_version: req.body?.clientVersion || 1,
  }
  res.json({ token: jwt.sign(payload, JWT_SECRET) })
}

function getRequestPayload(req) {
  const payload = req.body
  if (!payload.headers) {
    payload.headers = {}
  }

  return payload
}

function sendUnauthenticated(res, body) {
  res.json({
    ...body,
    control: { break: 401 }
  })

  return
}

function handleClientAwareness(req, res) {
  if (req.body.stage !== 'RouterRequest') {
    res.json(req.body)
    return
  }

  const payload = getRequestPayload(req)
  if (!payload.headers['authentication']) {
    return sendUnauthenticated(res, payload)
  }

  const token = payload.headers['authentication'][0].split('Bearer ')[1]
  if (!token) {
    return sendUnauthenticated(res, payload)
  }

  try {
    const jwtPayload = jwt.verify(token, JWT_SECRET)
    payload.headers['apollographql-client-name'] = [jwtPayload.client_name || 'coprocessor']
    payload.headers['apollographql-client-version'] = [jwtPayload.client_version || 'loadtest']

    res.json(payload)
  } catch {
    sendUnauthenticated(res, payload)
  }
}

function handleGuidResponse(req, res) {
  if (req.body.stage !== 'RouterResponse') {
    res.json(req.body)
    return
  }

  const payload = getRequestPayload(req)

  payload.headers['GUID'] = []
  for (let i = 0; i < 10; i++) {
    payload.headers['GUID'].push(uuid())
  }

  res.json(payload)
}

function handleStaticSubgraph(req, res) {
  if (req.body.stage !== 'SubgraphRequest') {
    res.json(req.body)
    return
  }

  const payload = getRequestPayload(req)

  payload.headers['source'] = ['coprocessor']

  res.json(payload)
}

const port = process.env.PORT || 3000
const app = express()
app.use(express.json())
app.post('/create-jwt', handleCreateJwt)
app.post('/client-awareness', handleClientAwareness)
app.post('/guid-response', handleGuidResponse)
app.post('/static-subgraph', handleStaticSubgraph)
app.listen(port, () => {
  console.log(`🚀 Coprocessor running on port ${port}`)
})
