const { randomUUID: uuid } = require('crypto')
const express = require('express')
const jwt = require('jsonwebtoken')

const JWT_SECRET = process.env.JWT_SECRET || 'apollo'

function handleCreateJwt(_req, res) {
  const token = jwt.sign({ client_name: 'node coprocessor' }, JWT_SECRET)
  res.json(token)
}

function getRequestPayload(req) {
  const payload = req.body
  if (!payload.headers) {
    payload.headers = {}
  }

  return payload
}

function sendUnauthenticated(res) {
  res.json({
    control: { break: 401 }
  })

  return
}

function handleClientAwareness(req, res) {
  if (req.body.stage !== 'RouterRequest') {
    return
  }

  const payload = getRequestPayload(req)
  if (!payload.headers['authentication']) {
    return sendUnauthenticated(res)
  }

  const token = payload.headers['authentication'][0].split('Bearer ')[1]
  if (!token) {
    return sendUnauthenticated(res)
  }

  try {
    const jwtPayload = jwt.verify(token, JWT_SECRET)
    payload.headers['apollographql-client-name'] = [jwtPayload.client_name || 'coprocessor']
    payload.headers['apollographql-client-version'] = [jwtPayload.client_version || 'loadtest']

    res.json(payload)
  } catch {
    sendUnauthenticated(res)
  }
}

function handleGuidResponse(req, res) {
  if (req.body.stage !== 'RouterResponse') {
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
    return
  }

  const payload = getRequestPayload(req)

  payload.headers['source'] = ['coprocessor']

  res.json(payload)
}

const port = process.env.PORT || 8000
const app = express()
app.use(express.json())
app.post('/create-jwt', handleCreateJwt)
app.post('/client-awareness', handleClientAwareness)
app.post('/guid-response', handleGuidResponse)
app.post('/static-subgraph', handleStaticSubgraph)
app.listen(port, () => {
  console.log(`🚀 Coprocessor running on port ${port}`)
})
