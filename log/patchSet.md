```sh
{ request:
   { debugId: 1,
     uri:
      'https://tcb-admin.tencentcloudapi.com/admin?&eventId=1572335409009_1_60142&seqId=1572335409009_1_60142',
     method: 'post',
     headers:
      { 'content-type': 'application/json',
        'user-agent': 'tcb-admin-sdk/1.15.0',
        'x-tcb-source': ',not_scf',
        host: 'tcb-admin.tencentcloudapi.com',
        accept: 'application/json',
        'content-length': 707 },
     body:
      '{"collectionName":"test","data":{"$set":{"name":"hello400"}},"query":{"_id":"c914be335db7ee63021ebd1b143b491f"},"multi":false,"merge":true,"upsert":false,"action":"database.updateDocument","envName":"tets-91347b","timestamp":1572335409010,"eventId":"1572335409009_1_60142","wxCloudApiToken":"","tcb_sessionToken":"","authorization":"q-sign-algorithm=sha1&q-ak=AK*****************************************Zd&q-sign-time=1572335408;1572336308&q-key-time=1572335408;1572336308&q-header-list=user-agent;x-tcb-source&q-url-param-list=action;collectionname;data;envname;eventid;merge;multi;query;tcb_sessiontoken;timestamp;upsert;wxcloudapitoken&q-signature=7a9f5dd16d0a6ce56994ef2aadb6d761cc85543c","sdk_version":"1.15.0"}' } }
{ response:
   { debugId: 1,
     headers:
      { server: 'openresty',
        date: 'Tue, 29 Oct 2019 07:50:09 GMT',
        'content-type': 'application/json; charset=utf-8',
        'content-length': '77',
        connection: 'close' },
     statusCode: 200,
     body: { requestId: '1572335409009_1_60142', data: [Object] } } }
{ updated: 1,
  upsertedId: null,
  requestId: '1572335409009_1_60142' }
```
set	设置字段值
remove	删除字段
inc	加一个数值，原子自增
mul	乘一个数值，原子自乘
push	数组类型字段追加尾元素，支持数组
pop	数组类型字段删除尾元素，支持数组
shift	数组类型字段删除头元素，支持数组
unshift	数组类型字段追加头元素，支持数组