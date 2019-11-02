```sh
{ request:
   { debugId: 1,
     uri:
      'https://tcb-admin.tencentcloudapi.com/admin?&eventId=1572335648748_1_46795&seqId=1572335648748_1_46795',
     method: 'post',
     headers:
      { 'content-type': 'application/json',
        'user-agent': 'tcb-admin-sdk/1.15.0',
        'x-tcb-source': ',not_scf',
        host: 'tcb-admin.tencentcloudapi.com',
        accept: 'application/json',
        'content-length': 697 },
     body:
      '{"collectionName":"test","data":{"$inc":{"age":1}},"query":{"_id":"c914be335db7ee63021ebd1b143b491f"},"multi":false,"merge":true,"upsert":false,"action":"database.updateDocument","envName":"tets-91347b","timestamp":1572335648748,"eventId":"1572335648748_1_46795","wxCloudApiToken":"","tcb_sessionToken":"","authorization":"q-sign-algorithm=sha1&q-ak=AK*****************************************Zd&q-sign-time=1572335647;1572336547&q-key-time=1572335647;1572336547&q-header-list=user-agent;x-tcb-source&q-url-param-list=action;collectionname;data;envname;eventid;merge;multi;query;tcb_sessiontoken;timestamp;upsert;wxcloudapitoken&q-signature=096168ab5ae1f5eebc5b57208d14eb2fd9ed99cf","sdk_version":"1.15.0"}' } }
{ response:
   { debugId: 1,
     headers:
      { server: 'openresty',
        date: 'Tue, 29 Oct 2019 07:54:08 GMT',
        'content-type': 'application/json; charset=utf-8',
        'content-length': '77',
        connection: 'close' },
     statusCode: 200,
     body: { requestId: '1572335648748_1_46795', data: [Object] } } }
{ updated: 1,
  upsertedId: null,
  requestId: '1572335648748_1_46795' }
```