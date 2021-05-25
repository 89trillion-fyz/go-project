from locust import HttpUser, task, between

class QuickstartUser(HttpUser):

    @task
    #输入稀有度，当前解锁阶段和cvc，获取该稀有度cvc合法且已解锁的所有士兵
    def findByRarityAndLock(self):
        self.client.get("/army/findByRarityAndLock?rarity=1&lock=0")
    #输入士兵id获取稀有度
    @task
    def findRarityById(self):
        self.client.get("/army/findRarityById?id=10101")
    #输入士兵id获取战力
    @task
    def findQualityById(self):
        self.client.get("/army/findQualityById?id=10101")
    #获取每个阶段解锁相应士兵的json数据
    @task
    def findByLock(self):
        self.client.get("/army/findByLock?lock=1")

