# go-project
go语言项目

# 接口文档

## 1）输入稀有度，当前解锁阶段和cvc，获取该稀有度cvc合法且已解锁的所有士兵
GET http://localhost:8000/army/findByRarityAndLock?rarity=1&lock=0
Accept: application/json

```json
{"code":0,"data":[{"id":"10101","Name":"Swordsman","note":"步兵lv1","UnlockArena":"0","ArmyType":"1","Race":"1","Rarity":"1","Quality":"1","AttackType":"","Prefab":"Infantry","PrefabFX":"","PrefabEnemyFX":"","AttackPattern":"Melee","Radius":"2.5","AvoidancePriority":"2","MoveType":"","AtkTargets":"0","MaxHp":"550","Atk":"100","SplashAtk":"","AtkRange":"10","ViewRange":"80","Width":"","Length":"","Def":"30","Heal":"","ShootSpeed":"1","MoveSpeed":"35","FirstAtkCount":"","AttackBuildingScaleArray":"","ArmyCount":"3","Skill":"","BeatBackDistance":"","RecycleMoney":"200","Desc":"army_desc_10101","CombatPoints":"167","KillGold":"5"},{"id":"10103","Name":"Swordsman","note":"步兵lv3","UnlockArena":"0","ArmyType":"1","Race":"1","Rarity":"1","Quality":"3","AttackType":"","Prefab":"Infantry","PrefabFX":"ArmyCommonGradelv2_3","PrefabEnemyFX":"ArmyEnemyCommonGradelv2_3","AttackPattern":"Melee","Radius":"2.5","AvoidancePriority":"2","MoveType":"","AtkTargets":"0","MaxHp":"800","Atk":"215","SplashAtk":"","AtkRange":"10","ViewRange":"80","Width":"","Length":"","Def":"40","Heal":"","ShootSpeed":"1","MoveSpeed":"35","FirstAtkCount":"","AttackBuildingScaleArray":"","ArmyCount":"7","Skill":"","BeatBackDistance":"","RecycleMoney":"800","Desc":"army_desc_10101","CombatPoints":"691","KillGold":"5"}],"msg":"ok"}
```



## 2）输入士兵id获取稀有度

GET http://localhost:8000/army/findRarityById?id=10101
Accept: application/json

```json
{"code":0,"data":"1","msg":"ok"}
```



## 3）输入士兵id获取战力

GET http://localhost:8000/army/findQualityById?id=10101
Accept: application/json

```json
{"code":0,"data":"1","msg":"ok"}
```

## 4）获取每个阶段解锁相应士兵的json数据

GET http://localhost:8000/army/findByLock?lock=1
Accept: application/json

```json
{"code":0,"data":[{"id":"10206","Name":"Archer","note":"弓箭手lv6","UnlockArena":"1","ArmyType":"2","Race":"1","Rarity":"1","Quality":"6","AttackType":"","Prefab":"Archer","PrefabFX":"ArmyCommonGradelv6","PrefabEnemyFX":"ArmyEnemyCommonGradelv6","AttackPattern":"Ranged","Radius":"2.5","AvoidancePriority":"14","MoveType":"","AtkTargets":"0,1","MaxHp":"1500","Atk":"300","SplashAtk":"","AtkRange":"85","ViewRange":"104","Width":"","Length":"","Def":"45","Heal":"","ShootSpeed":"1.5","MoveSpeed":"30","FirstAtkCount":"","AttackBuildingScaleArray":"","ArmyCount":"9","Skill":"","BeatBackDistance":"","RecycleMoney":"6400","Desc":"army_desc_10201","CombatPoints":"1826","KillGold":"10"},{"id":"10609","Name":"Arctic Wolves","note":"北极狼lv9","UnlockArena":"1","ArmyType":"6","Race":"1","Rarity":"2","Quality":"9","AttackType":"","Prefab":"Ninja","PrefabFX":"ArmyCommonGradelv7","PrefabEnemyFX":"ArmyEnemyCommonGradelv7","AttackPattern":"Melee","Radius":"0.1","AvoidancePriority":"2","MoveType":"","AtkTargets":"0","MaxHp":"5063","Atk":"1250","SplashAtk":"","AtkRange":"10","ViewRange":"80","Width":"","Length":"","Def":"","Heal":"","ShootSpeed":"1.5","MoveSpeed":"80","FirstAtkCount":"6","AttackBuildingScaleArray":"","ArmyCount":"9","Skill":"","BeatBackDistance":"","RecycleMoney":"51200","Desc":"army_desc_10601","CombatPoints":"7032","KillGold":"23"},{"id":"10305","Name":"Spearman","note":"长矛兵lv5","UnlockArena":"1","ArmyType":"3","Race":"1","Rarity":"1","Quality":"5","AttackType":"","Prefab":"Ironguard","PrefabFX":"ArmyCommonGradelv4_5","PrefabEnemyFX":"ArmyEnemyCommonGradelv4_5","AttackPattern":"Melee","Radius":"2.5","AvoidancePriority":"3","MoveType":"","AtkTargets":"0","MaxHp":"1600","Atk":"830","SplashAtk":"","AtkRange":"15","ViewRange":"80","Width":"","Length":"","Def":"60","Heal":"","ShootSpeed":"0.5","MoveSpeed":"20","FirstAtkCount":"","AttackBuildingScaleArray":"","ArmyCount":"9","Skill":"","BeatBackDistance":"10","RecycleMoney":"3200","Desc":"army_desc_10301","CombatPoints":"1802","KillGold":"7"},{"id":"10403","Name":"Bomber","note":"自爆人lv3","UnlockArena":"1","ArmyType":"4","Race":"1","Rarity":"1","Quality":"3","AttackType":"","Prefab":"Bombman","PrefabFX":"Myself_BombManGradelv2_3","PrefabEnemyFX":"Enemy_BombManGradelv2_3","AttackPattern":"Ranged","Radius":"2.5","AvoidancePriority":"2","MoveType":"","AtkTargets":"0,1","MaxHp":"3000","Atk":"2550","SplashAtk":"850","AtkRange":"10","ViewRange":"104","Width":"10","Length":"","Def":"20","Heal":"","ShootSpeed":"0.1","MoveSpeed":"45","FirstAtkCount":"","AttackBuildingScaleArray":"","ArmyCount":"1","Skill":"","BeatBackDistance":"","RecycleMoney":"800","Desc":"army_desc_10401","CombatPoints":"770","KillGold":"35"},{"id":"10601","Name":"Arctic Wolves","note":"北极狼lv1","UnlockArena":"1","ArmyType":"6","Race":"1","Rarity":"2","Quality":"1","AttackType":"","Prefab":"Ninja","PrefabFX":"","PrefabEnemyFX":"","AttackPattern":"Melee","Radius":"0.1","AvoidancePriority":"2","MoveType":"","AtkTargets":"0","MaxHp":"800","Atk":"150","SplashAtk":"","AtkRange":"10","ViewRange":"80","Width":"","Length":"","Def":"20","Heal":"","ShootSpeed":"1.5","MoveSpeed":"80","FirstAtkCount":"3","AttackBuildingScaleArray":"","ArmyCount":"3","Skill":"","BeatBackDistance":"","RecycleMoney":"200","Desc":"army_desc_10601","CombatPoints":"287","KillGold":"6"}],"msg":"ok"}
```

## locust 压测数据

```sh
locust --host=http://127.0.0.1:8000 -f locust.py --logfile=locustfile.log
```

```python
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
```

访问网址 http://127.0.0.1:8089



