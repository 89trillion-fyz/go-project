# go-project
go语言项目

# 接口文档

## 1）创建礼包码
POST http://localhost:8000/cdkey/createCdkey

Accept: application/json

请求体
```json
{"cdkeyType":3,"cdkeyUser":"admin","createTime":"2021-05-26 15:00:00","creator":"admin","desc":"兑换吗","contents":[{"item":"金币","count":"10"},{"item":"钻石","count":"20"}],"expireTime":"2021-05-26 19:00:00","totalExchangeNum":3}
```
响应
```json
{"code":0,"data":"QEFI1VG8","msg":"ok"}
```

## 2）查询礼包码详情

GET http://localhost:8000/cdkey/getCdkeyDetails?cdkey=QEFI1VG8

Accept: application/json

响应
```json
{"code":0,"data":{"cdkeyType":3,"cdkeyUser":"admin","cdkey":"QEFI1VG8","createTime":"2021-05-26 15:00:00","creator":"admin","desc":"兑换吗","contents":[{"item":"金币","count":"10"},{"item":"钻石","count":"20"}],"expireTime":"2021-05-26 19:00:00","totalExchangeNum":3,"alreadyExchangeNum":0,"exchangeList":null},"msg":"ok"}
```


## 3）验证礼包码

GET http://localhost:8000/cdkey/verifyCdkey?cdkey=QEFI1VG8&user=admin

Accept: application/json

响应
```json
{"code":0,"data":{"cdkeyType":3,"cdkeyUser":"admin","cdkey":"QEFI1VG8","createTime":"2021-05-26 15:00:00","creator":"admin","desc":"兑换吗","contents":[{"item":"金币","count":"10"},{"item":"钻石","count":"20"}],"expireTime":"2021-05-26 19:00:00","totalExchangeNum":3,"alreadyExchangeNum":1,"exchangeList":[{"user":"admin","exchangeTime":"2021-05-26 20:59:36"}]},"msg":"ok"}
```
