signKey
df58cf6a5fb191735fcd124c5020101fec521e4e
End
formatString
post
/admin
action=database.queryDocument&collectionname=test&envname=tets-91347b&eventid=1572282739127_1_46219&limit=100&tcb_sessiontoken=&timestamp=1572282739127&wxcloudapitoken=
user-agent=tcb-admin-sdk%2F1.15.0&x-tcb-source=%2Cnot_scf

End
String to sign
sha1
1572282738;1572283638
9950853c2b6b003555dc430b17a8e94590ce6983

End
Signature
82380b0f9757d82d91b12767ab0f6cdadc914aa1
End
{ request:
   { debugId: 1,
     uri:
      'https://tcb-admin.tencentcloudapi.com/admin?&eventId=1572282739127_1_46219&seqId=1572282739127_1_46219',
     method: 'post',
     headers:
      { 'content-type': 'application/json',
        'user-agent': 'tcb-admin-sdk/1.15.0',
        'x-tcb-source': ',not_scf',
        host: 'tcb-admin.tencentcloudapi.com',
        accept: 'application/json',
        'content-length': 565 },
     body:
      '{"collectionName":"test","limit":100,"action":"database.queryDocument","envName":"tets-91347b","timestamp":1572282739127,"eventId":"1572282739127_1_46219","wxCloudApiToken":"","tcb_sessionToken":"","authorization":"q-sign-algorithm=sha1&q-ak=AK*****************************************Zd&q-sign-time=1572282738;1572283638&q-key-time=1572282738;1572283638&q-header-list=user-agent;x-tcb-source&q-url-param-list=action;collectionname;envname;eventid;limit;tcb_sessiontoken;timestamp;wxcloudapitoken&q-signature=82380b0f9757d82d91b12767ab0f6cdadc914aa1","sdk_version":"1.15.0"}' } }
{ response:
   { debugId: 1,
     headers:
      { server: 'openresty',
        date: 'Mon, 28 Oct 2019 17:12:19 GMT',
        'content-type': 'application/json; charset=utf-8',
        'content-length': '162',
        connection: 'close' },
     statusCode: 200,
     body: { requestId: '1572282739127_1_46219', data: [Object] } } }
{ data:
   [ { _id: '9888d322-fe11-4aaa-af7f-bc91ea870f4c',
       age: 100,
       name: 'wenfeng' } ],
  requestId: '1572282739127_1_46219' }
