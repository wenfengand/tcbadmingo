{ request:
   { debugId: 1,
     uri:
      'https://tcb-admin.tencentcloudapi.com/admin?&eventId=1572334666876_1_08121&seqId=1572334666876_1_08121',
     method: 'post',
     headers:
      { 'content-type': 'application/json',
        'user-agent': 'tcb-admin-sdk/1.15.0',
        'x-tcb-source': ',not_scf',
        host: 'tcb-admin.tencentcloudapi.com',
        accept: 'application/json',
        'content-length': 598 },
     body:
      '{"collectionName":"test","query":{"name":"wenfeng"},"limit":100,"action":"database.queryDocument","envName":"tets-91347b","timestamp":1572334666876,"eventId":"1572334666876_1_08121","wxCloudApiToken":"","tcb_sessionToken":"","authorization":"q-sign-algorithm=sha1&q-ak=AK*****************************************Zd&q-sign-time=1572334665;1572335565&q-key-time=1572334665;1572335565&q-header-list=user-agent;x-tcb-source&q-url-param-list=action;collectionname;envname;eventid;limit;query;tcb_sessiontoken;timestamp;wxcloudapitoken&q-signature=52910170b0190aa540446535516a69b7264167b4","sdk_version":"1.15.0"}' } }
{ response:
   { debugId: 1,
     headers:
      { server: 'openresty',
        date: 'Tue, 29 Oct 2019 07:37:47 GMT',
        'content-type': 'application/json; charset=utf-8',
        'content-length': '162',
        connection: 'close' },
     statusCode: 200,
     body: { requestId: '1572334666876_1_08121', data: [Object] } } }
{ data:
   [ { _id: '9888d322-fe11-4aaa-af7f-bc91ea870f4c',
       age: 100,
       name: 'wenfeng' } ],
  requestId: '1572334666876_1_08121' }
