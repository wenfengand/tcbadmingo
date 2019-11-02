```sh
{ request:
   { debugId: 1,
     uri:
      'https://tcb-admin.tencentcloudapi.com/admin?&eventId=1572348591893_1_32275&seqId=1572348591893_1_32275',
     method: 'post',
     headers:
      { 'content-type': 'application/json',
        'user-agent': 'tcb-admin-sdk/1.15.0',
        'x-tcb-source': ',not_scf',
        host: 'tcb-admin.tencentcloudapi.com',
        accept: 'application/json',
        'content-length': 548 },
     body:
      '{"collectionName":"test2","action":"database.addCollection","envName":"tets-91347b","timestamp":1572348591893,"eventId":"1572348591893_1_32275","wxCloudApiToken":"","tcb_sessionToken":"","authorization":"q-sign-algorithm=sha1&q-ak=AK*****************************************Zd&q-sign-time=1572348590;1572349490&q-key-time=1572348590;1572349490&q-header-list=user-agent;x-tcb-source&q-url-param-list=action;collectionname;envname;eventid;tcb_sessiontoken;timestamp;wxcloudapitoken&q-signature=7055ab7e25530f2aa9e481dcca46c8e801717ae4","sdk_version":"1.15.0"}' } }
{ response:
   { debugId: 1,
     headers:
      { server: 'openresty',
        date: 'Tue, 29 Oct 2019 11:29:53 GMT',
        'content-type': 'application/json; charset=utf-8',
        'content-length': '57',
        connection: 'close' },
     statusCode: 200,
     body: { requestId: '1572348591893_1_32275', message: 'success' } } }
{ requestId: '1572348591893_1_32275', message: 'success' }
```