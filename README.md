# go-project
go语言项目
#接口文档
####1）输入稀有度，当前解锁阶段和cvc，获取该稀有度cvc合法且已解锁的所有士兵
GET http://localhost:8000/army/findByRarityAndLock?rarity=1&lock=0
Accept: application/json
####2）输入士兵id获取稀有度
GET http://localhost:8000/army/findRarityById?id=10101
Accept: application/json
####3）输入士兵id获取战力
GET http://localhost:8000/army/findQualityById?id=10101
Accept: application/json

####4）获取每个阶段解锁相应士兵的json数据
GET http://localhost:8000/army/findByLock?lock=1
Accept: application/json
