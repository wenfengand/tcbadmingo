```sh
{ request:
   { debugId: 1,
     uri:
      'https://tcb-admin.tencentcloudapi.com/admin?&eventId=1572335203626_1_47898&seqId=1572335203626_1_47898',
     method: 'post',
     headers:
      { 'content-type': 'application/json',
        'user-agent': 'tcb-admin-sdk/1.15.0',
        'x-tcb-source': ',not_scf',
        host: 'tcb-admin.tencentcloudapi.com',
        accept: 'application/json',
        'content-length': 583 },
     body:
      '{"collectionName":"test","data":{"name":"hello4","age":5},"action":"database.addDocument","envName":"tets-91347b","timestamp":1572335203626,"eventId":"1572335203626_1_47898","wxCloudApiToken":"","tcb_sessionToken":"","authorization":"q-sign-algorithm=sha1&q-ak=AK*****************************************Zd&q-sign-time=1572335202;1572336102&q-key-time=1572335202;1572336102&q-header-list=user-agent;x-tcb-source&q-url-param-list=action;collectionname;data;envname;eventid;tcb_sessiontoken;timestamp;wxcloudapitoken&q-signature=240b7d0c4348a2bf4c539f18a3212a5d4d2c4e30","sdk_version":"1.15.0"}' } }
{ response:
   { debugId: 1,
     headers:
      { server: 'nginx',
        date: 'Tue, 29 Oct 2019 07:46:43 GMT',
        'content-type': 'application/json; charset=utf-8',
        'content-length': '87',
        connection: 'close' },
     statusCode: 200,
     body: { requestId: '1572335203626_1_47898', data: [Object] } } }
{ id: 'c914be335db7ee63021ebd1b143b491f',
  requestId: '1572335203626_1_47898' }
```
