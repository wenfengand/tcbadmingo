{ request:
   { debugId: 1,
     uri:
      'https://tcb-admin.tencentcloudapi.com/admin?&eventId=1572334937596_1_51613&seqId=1572334937596_1_51613',
     method: 'post',
     headers:
      { 'content-type': 'application/json',
        'user-agent': 'tcb-admin-sdk/1.15.0',
        'x-tcb-source': ',not_scf',
        host: 'tcb-admin.tencentcloudapi.com',
        accept: 'application/json',
        'content-length': 597 },
     body:
      '{"collectionName":"test","query":{"age":{"$gt":8}},"limit":100,"action":"database.queryDocument","envName":"tets-91347b","timestamp":1572334937596,"eventId":"1572334937596_1_51613","wxCloudApiToken":"","tcb_sessionToken":"","authorization":"q-sign-algorithm=sha1&q-ak=AK*****************************************Zd&q-sign-time=1572334936;1572335836&q-key-time=1572334936;1572335836&q-header-list=user-agent;x-tcb-source&q-url-param-list=action;collectionname;envname;eventid;limit;query;tcb_sessiontoken;timestamp;wxcloudapitoken&q-signature=fc836a2e51365991fa5c0aee1ad0a5591cee065d","sdk_version":"1.15.0"}' } }
{ response:
   { debugId: 1,
     headers:
      { server: 'nginx',
        date: 'Tue, 29 Oct 2019 07:42:19 GMT',
        'content-type': 'application/json; charset=utf-8',
        'content-length': '303',
        connection: 'close' },
     statusCode: 200,
     body: { requestId: '1572334937596_1_51613', data: [Object] } } }
{ data:
   [ { _id: '9888d322-fe11-4aaa-af7f-bc91ea870f4c',
       age: 100,
       name: 'wenfeng' },
     { _id: '57b20d40-e50d-4aec-b9e1-92df0550294b',
       age: 33,
       name: 'hello' },
     { _id: '3c8b53f8-a7f5-429f-ae9f-5d7fb2d8d9e2',
       age: 45,
       name: 'test' } ],
  requestId: '1572334937596_1_51613' }


比较运算 加上美元符即可
eq	字段 ==
neq	字段 !=
gt	字段 >
gte	字段 >=
lt	字段 <
lte	字段 <=
in	字段值在数组里
nin	字段值不在数组里
逻辑运算	and	表示需同时满足指定的所有条件
or	表示需同时满足指定条件中的至少一个
