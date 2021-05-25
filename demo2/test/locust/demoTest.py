from locust import HttpUser,task,between
class Test(HttpUser):
    @task
    def getResult(self):
        response =  self.client.post("/getResult",json={"str": "3+2*2+13"})
        print("response",response.status_code)
