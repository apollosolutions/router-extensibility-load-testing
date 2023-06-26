from flask import Flask, request
import jwt
from os import environ
from uuid import uuid4 as uuid

app = Flask(__name__)

JWT_SECRET = environ.get('JWT_SECRET') if environ.get('JWT_SECRET') is not None else 'apollo'

def _get_payload():
    payload = request.get_json(force=True)
    if (payload.get('headers') is None):
        payload['headers'] = {}

    return payload

def _send_unauthenticated(payload):
    payload['control'] = { 'break': 401 }
    return payload

@app.post('/client-awareness')
def handle_client_awareness():
    payload = _get_payload()
    if (payload['stage'] != 'RouterRequest'):
        return payload

    if (payload['headers'].get('authentication') == None):
        return _send_unauthenticated(payload)

    token = payload['headers']['authentication'][0].split('Bearer ')[1]
    if (token == None):
        return _send_unauthenticated(payload)

    try:
        jwtPayload = jwt.decode(token, key=JWT_SECRET, algorithms='HS256')
        clientName = [jwtPayload.get('client_name')] if jwtPayload.get('client_name') is not None else ['coprocessor']
        clientVersion = [jwtPayload.get('client_version')] if jwtPayload.get('client_version') is not None else ['loadtest']
        payload['headers']['apollographql-client-name'] = clientName
        payload['headers']['apollographql-client-version'] = clientVersion
    except jwt.exceptions.InvalidSignatureError:
        return _send_unauthenticated(payload)

    return payload

@app.post('/guid-response')
def handle_guid_response():
    payload = _get_payload()
    if (payload['stage'] != 'RouterResponse'):
        return payload

    payload['headers']['GUID'] = []
    for _ in range(10):
        payload['headers']['GUID'].append(uuid())

    return payload

@app.post('/static-subgraph')
def handle_static_subgraph():
    payload = _get_payload()
    if (payload['stage'] != 'SubgraphRequest'):
        return payload

    payload['headers']['source'] = 'coprocessor'

    return payload

if __name__ == '__main__':
    from waitress import serve

    port = int(environ.get('PORT')) if environ.get('PORT') is not None else 3000
    print('Starting server on port', port)
    serve(app, port=port)
