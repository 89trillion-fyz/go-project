# 配置文件说明

[toc]



# config.army.model.json  士兵基础配置

### 下发配置文件接口
|信息|说明|
|:-------|:-----------------------|
| 接口方式 | Http GET|
| 路由 | /sys/config?type = 1|

### 下发的配置文件
|配置名|说明|
|:-------|:--------------|
|armyUnlock|竞技场杯数解锁士兵的配置|
|armyUnlockLevel|pve关卡数解锁士兵的配置|

### 配置文件详解
```json5
{
	"id": "10101",  //士兵id
	"note": "步兵lv1",  //士兵名字
	"Rarity": "1", //士兵稀有度
	"ArmyCount": "3", //士兵数量
	"KillGold": "5",//击杀小兵获取金币数量
	"KillWood": "8",//击杀小兵获取木头数量
	"UnlockLevel": "5", //pve多少关数解锁士兵
	"UnlockArena": "0",//竞技场多少杯数解锁士兵
	"CVC": "1470",//对应的客户端版本
	"Source": 255 //资源产出位置 采用2进制位表示  
}
```

# config.achievement.json  成就配置

### 下发配置文件接口
|信息|说明|
|:-------|:-----------------------|
| 接口方式 | Http GET|
| 路由 | /ach/config?version=1|

### 下发的配置文件
|配置名|说明|
|:-------|:--------------|
|acheConfig|成就配置|

### 配置文件详解
```json5
{
	"achievementID": 8,  // 勋章ID
	"subType": 3,   // 勋章大类
	"displayType": 2, // 展示类型 1需要数字类型展示 2不需要数字类型展示
	"level": 1,  // 等级
	"targetValue": 3,  // 目标值
	"rewardType": 201,  // 奖励类型
	"rewardNum": 20  // 奖励数量
}
```

# config.arena.potion.json  pvp战斗药水配置

### 下发配置文件接口
|信息|说明|
|:-------|:-----------------------|
| 接口方式 | Http GET|
| 路由 | /sys/config?type = 1|

### 下发的配置文件
|配置名|说明|
|:-------|:--------------|
|arenaPotion|pvp战斗药水配置|

### 配置文件详解
```json5
{
	"arenaPotion": [
		{
			"type": 1,  //药水类型
			"cnt": 0,   //药水数量
			"ratio": 20, //每次恢复生命的百分比
			"max": 40000,  //可恢复的最大生命
			"cooldown": 5000,  //冷却时间 单位: ms
			"costList": [
				{
					"cnt": 1,  //可点击次数
					"cost": {
						"rewardId": 201,  //消耗品: 钻石
						"count": 10  //消耗的数量
					}
				}
			]
		}
	]
}
```

# config.arena.road.json  新版奖杯之路配置
# 新版奖杯之路配置 （bp赛季 >= 18 使用）

### 下发配置文件接口
|信息|说明|
|:-------|:-----------------------|
| 接口方式 | Http GET|
| 路由 | /sys/config?type = 1|

### 下发的配置文件
|配置名|说明|
|:-------|:--------------|
|arenaTrophy|新版奖杯之路配置|

### 配置文件详解
```json5
{
	"trophy": "25",   //赛季的id
	"rewardId": "1020102",   //奖励的id
	"count": "2"   //数量
}
```

# config.arena.spoils.json

### 配置文件详解
```json5
{
	"Ease": {    
		"0": {
			"OpponentLevel": "Ease",   
			"Arena": "0",
			"WinTrophies": "0",
			"LoseTrophies": "0",
			"MinWoods": "0",          //最小木头数
			"MaxWoods": "0"           //最大木头数
		}
	}
}
```

# config.army.dynamic.source.json 活动定期产出士兵配置

### 配置文件详解
```json5
{
	"start": 1626480000,    //开始时间
	"end": -1,     //结束时间
	"armySourceMap": {
		"190":{   
			"id":190,  //士兵id
			"source":6    //士兵产出来源  
		},
	}
}
```
### 士兵产出来源详解
|士兵产出来源|source|说明|
|:-------|:-------|:-------|
|ShopCoinDrawSource|2|商店金币抽卡|
|ShopGemDrawSource|3|商店钻石抽卡|
|ShopChestSource|4|商店宝箱|
|FuseSource|5|熔炉|
|BattleCoinDrawSource|6|战场金币抽卡|
|LuckDrawSource|7|幸运抽卡|
|TrophyChestSource|8|奖杯之路宝箱/battlepass宝箱|
|LeagueChestSource|9|巅峰赛宝箱|
|WeekCardSource|10|周卡月卡|
|SpecialChestSource|11|特殊宝箱/节日宝箱|
|CompensateSource|12|补偿|
|RecruitSource|13|招募|
|RecruitLegendSource|14|传奇招募|

# config.army.fuse.weight.json 士兵熔炼权重配置

### 配置文件详解
```json5
{
	"160": 75,   //"士兵id": 士兵的熔炉权重
	"163": 100,
	"169": 125,
	"184": 100,
	"191": 100,
	"150": 125,
	"188": 100
}
```

# config.army.sell.json 出售士兵配置

### 配置文件详解
```json5
{
	"101": 200,  //"rarity*10+level": 士兵出售的价格
	"102": 360,
	"103": 640,
}
```

# config.army.upgrade.json 士兵升级所需材料配置

### 下发配置文件接口
|信息|说明|
|:-------|:-----------------------|
| 接口方式 | Http GET|
| 路由 | /sys/config?type = 13|

### 下发的配置文件
|配置名|说明|
|:-------|:--------------|
|armyConditions|士兵升级所需材料配置|

### 配置文件详解
```json5
{
	"rarityLevel": 101,    //士兵的稀有度和等级
	"costCoin": 0,         //升级消耗金币数
	"costMedal": 0,        //升级消耗奖牌数
	"costGem": 0           //审计消耗钻石数
}
```

## config.barracks.capacity.json 军营容量配置

下发配置文件接口

| 信息     | 说明               |
| -------- | ------------------ |
| 接口方式 | http GET           |
| 路由     | /sys/config?type=1 |

下发的配置文件

| 配置名           | 说明         |
| ---------------- | ------------ |
| barracksCapacity | 军营容量配置 |

配置文件详解

```
{
	"capacities": [				// 配置key
		{
			"lv": 1,	        // 军营等级值
			"capacity": 50		// 军营容量
		},
		{
			"lv": 2,
			"capacity": 75
		},
		...
	}
```

## config.battle.spoils.json 战斗战利品配置（项目中无引用关系）

下发配置文件接口

无

下发的配置文件

无

配置文件详解

```
{
		"Ease":{
					"0": {
						"OpponentLevel": "Ease",
						"Arena": "0",
						"MinWoods": "0",
						"MaxWoods": "0",
						"WinTrophies": "0",
						"LoseTrophies": "0"
					},
					...
		},
		"Medium":{
			"0": {
				"OpponentLevel": "Ease",
				"Arena": "0",
				"MinWoods": "0",
				"MaxWoods": "0",
				"WinTrophies": "0",
				"LoseTrophies": "0"
			},
			...
		},
		"Medium":{
				"0": {
						"OpponentLevel": "Ease",
						"Arena": "0",
						"MinWoods": "0",
						"MaxWoods": "0",
						"WinTrophies": "0",
						"LoseTrophies": "0"
				},
				...
		}
}
```

## config.box.json 宝箱配置

下发配置文件接口

| 信息     | 说明               |
| -------- | ------------------ |
| 接口方式 | http GET           |
| 路由     | /sys/config?type=1 |

下发的配置文件

| 配置名                | 说明        |
| --------------------- | ----------- |
| chests                | Pvp宝箱配置 |
| chestDetail.boxConfig | 宝箱详情    |

配置文件详解

```
{
	"21": {				// 宝箱ID
		"id": "21",             // 宝箱ID
		"rewards": "1108:2",	//宝箱奖励
		"quality": "01"		//品质
	},
	...
```

## config.box.old.version.json 旧宝箱配置，版本小于 NewBoxVersion = 2420时使用

下发配置文件接口

| 信息     | 说明               |
| -------- | ------------------ |
| 接口方式 | http GET           |
| 路由     | /sys/config?type=1 |

下发的配置文件

| 配置名                | 说明     |
| --------------------- | -------- |
| chestDetail.boxConfig | 宝箱详情 |

配置文件详解

```
{
	"21": {				// 宝箱ID
		"id": "21", // 宝箱ID
		"rewards": "1108:2",	//宝箱奖励
		"quality": "01"		//品质
	},
	...
```

## config.boxrwd.json 宝箱奖励配置

下发配置文件接口

| 信息     | 说明               |
| -------- | ------------------ |
| 接口方式 | http GET           |
| 路由     | /sys/config?type=1 |

下发的配置文件

| 配置名                  | 说明     |
| ----------------------- | -------- |
| chestDetail.cardsConfig | 抽卡概率 |

配置文件详解

```
{
	"11": {						//宝箱ID
		"Rwd": {				// 奖励
			"0": {
				"Type": 2,		// 类型
				"Rarity": 1,	        // 稀有度
				"Level": 1,		//等级
				"Count": 1		//数量
			}
		},
		"Weight": {				// 权重
			"0": 100
		}
	},
	...
```

## config.build.new.num.json 新建筑数量和等级配置

使用配置文件接口

| 信息     | 说明          |
| -------- | ------------- |
| 接口方式 | Ws            |
| 路由     | lay.bd.new.v2 |

使用的配置文件

| 配置                                                         | 说明             |
| ------------------------------------------------------------ | ---------------- |
| func  GetBuildingNewNum(cityHallLv, buildingType)(num uint32, ok bool) | 建筑等级数量校验 |

配置文件详解

```
{
	{
		"id": 2,		//建筑ID
		"cityHallLv": 1,	//市政厅等级
		"num": 1		// 数量
	},
	...
```

## config.build.up.config.json 建筑升级配置

下发配置文件接口

| 信息     | 说明                |
| -------- | ------------------- |
| 接口方式 | http GET            |
| 路由     | /sys/config?type=13 |

下发的配置文件

| 配置           | 说明         |
| -------------- | ------------ |
| bdUpConditions | 建筑升级限制 |

配置文件详解

```
{
	{
		"id": 201,// bdType*100 + bdLv 要升级的建筑类型 和 要升级到的等级
		"targetType": 1,		// 目标类型 需要达到条件。目标建筑的类型
		"isMultiple": 0,		// 多建筑标识
		"targetLv": 1,			// 目标等级 需要达到条件。目标建筑的类型
		"targetCount": 1		// 目标数量 需要达到条件。目标建筑的类型
	},
	...
```

```
byType 1 大本 2金库 6兵营 11墙 10 CrossTower 18 Catapult 9 Cannon大炮 17BattleMine
```

## config.build.up.cost.json 建筑升级消耗配置

下发配置文件接口

| 信息     | 说明         |
| -------- | ------------ |
| 接口方式 | http GET     |
| 路由     | /lay/up/conf |

下发的配置文件

| 配置     | 说明           |
| -------- | -------------- |
| bdUpCost | 建筑升级消配置 |

建筑升级校验

| 信息                                                         | 说明         |
| ------------------------------------------------------------ | ------------ |
| func GetBdCostV2(buildingType uint32, buildingLv uint32) (uint32, uint32, *myerr.MyErr) | 建筑升级校验 |

配置文件详解

```
{
	{
		"id": 101,	                //ID
		"buildTime": 0,	                //建造时间
		"ItemCost": [		        //物品消耗
			{
				"itemId": 301,	//物品ID
				"itemCnt": 0	//物品数量
			}
		]
	},
	...
```

## 


# config.champ.road.json 奖杯之路领取奖励老配置
###【服务器计算奖励用 bp赛季 < 18 使用】

### 配置文件详解
```json5
{
	"trophy": "6300",  //bp赛季id
	"rewardId": "101", //奖励id和奖励类型
	"count": "100000"  //奖励数量
}
```


# config.champ.road.v2.json 奖杯之路领取奖励新配置
###【服务器计算奖励用 bp赛季 >= 18 使用】

### 配置文件详解
```json5
{
	"trophy": "4100",   //bp赛季id
	"rewardId": "101",  //奖励id和奖励类型
	"count": "100000"   //奖励数量
}
```

# config.coin.balance.json  大本等级对应金币配置

### 配置文件详解
```json5
{
	"CoinProtect": 500,    //大本等级对应的金币
	"RobRatio": 0.95,      //计算基础奖励时乘的百分比
	"RobMax": 475          //计算基础奖励最大金币
}
```

## config.draw.json 抽卡配置

下发配置文件接口

| 信息     | 说明               |
| -------- | ------------------ |
| 接口方式 | http GET           |
| 路由     | /sys/config?type=1 |

下发的配置文件

| 配置 | 说明     |
| ---- | -------- |
| draw | 抽卡配置 |

配置文件详解

```
{
	"bd": {								//战斗抽卡配置
		"items": [
			{
				"star": 1,				//星星
				"price": 20000,		//价格
				"lucky_base": 2,	//基础幸运值
				"prob": [					//概率
					{
						"rarity": 1,	//稀有度
						"percent": 55	//百分比
					},
					{
						"rarity": 2,
						"percent": 45
					}
				]
			},
	...
	"sc": {								//商店金币抽卡配置
		"unlock": 0,
		"unlockLevel": 0,		//解锁等级
		"price": 50000,			//价格
		"multiprice": 150000,
		"prob": [						//概率
			{
				"rarity": 1,		//稀有度
				"percent": 40		//百分比
			},
			{
				"rarity": 2,
				"percent": 60
			}
		]
	},
	...
	"sg": {								//商店钻石抽卡配置
		"unlock": 0,
		"unlockLevel": 0,		//解锁等级
		"price": 200,				//价格
		"multiprice": 1800,
		"prob": [						//概率
			{
				"rarity": 1,		//稀有度
				"percent": 10		//百分比
			},
			...
		]
	},
	"sh": {
		"unlock": 0,
		"unlockLevel": 0,		//解锁等级
		"price": 1000,			//价格
		"multiprice": 9000
	}
```

## config.guild.league.battle.building.json 公会联盟战斗建筑配置

使用配置

| 配置                                  | 说明                                |
| ------------------------------------- | ----------------------------------- |
| pkg/config/leagueBattle.GuildBuildCfg | 公会联盟赛\|pve\|积分赛建筑奖励配置 |

配置文件详解

```
{
	"10": {											//塔类型ID
		"1": {										//等级值
			"Type": "10",						//塔类型ID
			"Lv": "1",						  //等级值
			"Name": "Archer Tower",	//塔类型名称，箭塔
			"Wood": "7.13",					//木材
			"Gold": "75",						//金币
			"Score": "6"						//分数
		},
	...

```

## config.guild.league.battle.troops.json 公会联盟战斗军队配置

使用配置

| 配置                                   | 说明                                |
| -------------------------------------- | ----------------------------------- |
| pkg/config/leagueBattle.GuildTroopsCfg | 公会联盟赛\|pve\|积分赛军队奖励配置 |

配置文件详解

```
{
	"10101": {						//士兵ID
		"Id": "10101",			//士兵ID
		"Name": "Swordsman",//士兵名称
		"note": "步兵lv1",	//士兵名称+等级
		"Wood": "1.11",			//木材
		"Gold": "10",				//金币
		"Score": "1"				//分数
	},
	...

```

## config.head.hunt.json 	赏金关卡｜pvp配置

使用配置

| 配置                            | 说明                  |
| ------------------------------- | --------------------- |
| pkg/config/headhunt.HeadHuntCfg | 赏金关卡｜pvp解析配置 |

配置文件详解

```
{
	"Simple_1": {													//简单模式1
		"MaxRound": 2,											//最大回合
		"Rounds": [
			{	
				"MaxLevel": 10,									//最大等级
				"Levels": [
					{
						"HeadHuntType": "Simple",		//简单困难类型，简单
						"Arena": 1,									//竞技场
						"Round": 1,									//回合
						"Turn": 1,
						"RewardCoins": 475,					//金币奖励
						"BasicCoins": 20,						//基础金币奖励
						"RewardWoods": 125,					// 木材奖励
						"RewardGems": 20,						//钻石奖励
						"LevelId": 1,								//等级ID
						"CombatPoints": 640					//战斗力
					},
	...

```


# config.idle.reward.json 金矿对应的详情及离线奖励配置

### 配置文件详解
```json5
{
	"Lv": "0",             //金矿等级
	"Name": "GoldMine",    //金矿名字
	"Hp": "0",             //血量
	"BuildTime": "0",      //建造所花费时间
	"BuildCostWoods": "0", //建造所花费木头
	"Productivity": "0",   //金币生产力
	"GemProductivity": "0", //钻石生产力
	"Capacity": "0",       //容量
	"GemCapacity": "0",    //钻石容量
	"ProtectRate": "0",    //保护率
	"PlunderRate": "0"     //掠夺率
}
```

# config.infinite.war.base.json 无尽试炼基础配置+赛季配置

### 下发配置文件接口

| 信息     | 说明         |
| -------- | ------------ |
| 接口方式 | http GET     |
| 路由     | /infwar/config |

### 下发的配置文件

| 配置     | 说明           |
| -------- | -------------- |
| InfWarBaseConf | 无尽试炼基础配置+赛季配置 |

### 配置文件详解
```json5
{
	"unlockLevel": 200,   //pve200关解锁无尽试炼
	"unlock": 1000,       //1000奖杯解锁无尽试炼
	"truceTime": 600,     //停止时间
	"potions": [          //无尽试炼治疗药水配置
		{
			"type": 1,    //药水类型
			"cnt": 0,     //药水数量
			"ratio": 15,  //恢复生命百分比
			"max": 30000, //恢复最大生命值
			"cooldown": 1000,  //冷却时间
			"cost": [
				{
					"rewardId": 201,  //钻石
					"count": 15       //花费数量
				}
			]
		}
	],
	"challenge": {     //无尽试炼挑战
		"free": 3,     //免费挑战次数
		"limit": -1,   //每日可挑战次数
		"cost": [
			{
				"rewardId": 201,   //花费钻石
				"count": 30        //花费数量
			}
		]
	},
	"difficulty": [      //没有使用
		{
			"score": 0,
			"troopLv": 1,
			"buildingLv": 1
		}
	],
	"diff_addition": {      //没有使用
		"damage": 225,
		"health": 0,
		"score": 500
	}
}
```

# config.layout.unavailable.holders.json 建筑士兵禁止坐标配置

### 配置文件详解
```json5
{
	"3": {              //岛id
		"65281": true,  // "坐标":是否禁用
		"65279": true
	}
}
```

# config.layout.unavailable.walls.json 围墙禁止坐标配置

### 配置文件详解
```json5
{
	"2": {              //岛id
		"64515": true,  // "坐标":是否禁用
		"65027": true,
		"3": true,
		"64258": true,
		"65282": true,
		"258": true,
		"1": true,
		"64256": true,
		"65280": true,
		"64510": true,
		"64765": true,
		"65534": true,
		"65277": true
	}
}
```

# config.layout.weather.json  天气配置

### 下发配置文件接口

| 信息     | 说明         |
| -------- | ------------ |
| 接口方式 | http GET     |
| 路由     | /lay/battle/conf |

### 下发的配置文件

| 配置     | 说明           |
| -------- | -------------- |
| weather | 天气配置 |

### 配置文件详解
```json5
{
	"sun": [    //天气
		{
			"id": 405,   //天气id
			"atkRatio": 0.75,    //命中率
			"splAtkRatio": 0.75,  
			"msRatio": 0.8       //miss率
		}
	]
}
```

# config.lucky.card.json  幸运抽卡配置

### 配置文件详解
```json5
{
	"low": {         //抽卡模块
		"3": 30,     
		"4": 70
	},
	"high": {
		"2": 20,
		"3": 60,
		"4": 20
	}
}
```


## config.morale.army.camp.json 军营士气配置

下发配置文件接口

| 信息     | 说明               |
| -------- | ------------------ |
| 接口方式 | http GET           |
| 路由     | /sys/config?type=1 |

下发的配置文件

| 配置名称     | 说明         |
| ------------ | ------------ |
| productivity | 军营士气配置 |

配置文件详解

```
{
	{
		"lv": 1,							//军营等级
		"speed": 1400,				//速度
		"capacity": 3000,			//容量
		"cost": 20						
	},
	...
```

## config.morale.decrease.json 士气减少配置

使用配置

| 配置名称                                   | 说明             |
| ------------------------------------------ | ---------------- |
| pkg/config/moraleContent.MoraleDecreaseCfg | 解析后的士气配置 |

配置文件详解

```
{
	{
		"armyLv": 1,		//士兵等级
		"rarity": 1,		//稀有度
		"count": 56.19	//减少值
	},
	...
```

## config.newtrophy.road.json  奖杯之路配置,（bp赛季 < 18 使用）

下发配置文件接口

| 信息     | 说明               |
| -------- | ------------------ |
| 接口方式 | http GET           |
| 路由     | /sys/config?type=1 |

下发的配置文件

| 配置名称    | 说明         |
| ----------- | ------------ |
| arenaTrophy | 奖杯之路配置 |

配置文件详解

```
{
	"0": {
		"trophy": "0",				//奖杯数量
		"rewardId": "",				//奖励ID
		"count": ""						//数量
	},
	"25": {
		"trophy": "25",
		"rewardId": "1020102",
		"count": "2"
	},
	...
```


# config.pointrace.json  积分赛配置

### 配置文件详解
```json5
{
	"elite": {     //精英积分赛
		"StartTime": "1605657600",  //开始时间
		"Continue": "4",            //持续几个赛季
		"Interval": "0",            //时间间隔
		"UnitSecond": "86400",      //时间单位
		"Status": "1"
	},
	"master": {    //大师积分赛
		"StartTime": "1605657600",  
		"Continue": "4",           
		"Interval": "0",
		"UnitSecond": "86400",
		"Status": "1"
	}
}
```

# config.pointrace.room.json   积分赛房间配置

### 配置文件详解
```json5
{
	"elite": [      //精英积分赛
		{
			"combats": 0,  //战斗力
			"stage": 1     //阶段
		}
	],
	"master": [     //大师积分赛
		{
			"combats": 0,
			"stage": 1
		}
	],
}
```

# config.pve.basis.json  pve/pvp/公会战/积分赛击杀奖励配置

### 下发配置文件接口

| 信息     | 说明         |
| -------- | ------------ |
| 接口方式 | http GET     |
| 路由     | /lay/battle/conf |

### 下发的配置文件

| 配置     | 说明           |
| -------- | -------------- |
| battleConfig | pve击杀奖励配置 |

### 配置文件详解
```json5
{
	"pveBasicCfg": {       //pve击杀奖励配置
		"trooper": [       //小兵
			{
				"type": 1,   //等级
				"KillReward": [   //击杀奖励
					{
						"KillCoin": 28,   //击杀奖励金币
						"KillWood": 6.65  //击杀奖励木头
					}
				]
			}
		],
		"building": [       //建筑
			{
				"type": 10,   //建筑id
				"KillReward": [   //击杀奖励
					{
						"KillCoin": 36,    //击杀奖励金币
						"KillWood": 16.94  //击杀奖励木头
					}
				]
			}
		]
	},
	"pvpBasicCfg": {      //pvp｜积分赛｜公会战 击杀奖励
		"trooper": [
			{
				"type": 1,
				"KillReward": [
					{
						"KillCoin": 10,
						"KillWood": 6
					}
				]
			}
		],
		"building": [
			{
				"type": 10,
				"KillReward": [
					{
						"KillCoin": 36,
						"KillWood": 16.94
					}
				]
			}
		]
	}
}
```

# config.pve.chapter.json  pve章节的规则配置

### 配置文件详解
```json55
[
	{
		//第1章
		"chapter": 1,
		//多少星星解锁
		"unlockStar": 20,
		//章节中每关阵容配置md5
		"levelFile": "bb6874dec3b015d080cfdad45b2f8a7f",
		//胜利规则
		"starsRule": {
			//获取2颗星需要剩余士兵率为15%
			"twoStar": 15,
			//获取3颗星需要剩余士兵率为25%
			"threeStar": 25
		},
		"rewardList": [
		]
	}
]
```

# config.pve.levels.json  pve关卡的基础奖励配置

### 配置文件详解
```json5
[
	{
		"levelId": "1",   //第1关
		"rewards": [      //基础奖励
			{
				"rewardId": "101",   //金币
				"count": "412"       //奖励数量
			},
			{
				"rewardId": "301",   //木头
				"count": "57"        //奖励数量
			}
		]
	},
]
```

# config.pve.potion.json  pve治疗药水配置

### 下发配置文件接口

| 信息     | 说明         |
| -------- | ------------ |
| 接口方式 | http GET     |
| 路由     | /sys/config?type = 1 |

### 下发的配置文件

| 配置     | 说明           |
| -------- | -------------- |
| pvePotion | pve治疗药水配置 |

### 配置文件详解
```json5
{
	"pvePotion": [     //pve治疗药水
		{
			"type": 1,         //治疗药水类型
			"cnt": 0,          //数量
			"ratio": 15,       //恢复生命百分比
			"max": 30000,      //可恢复的最大生命值
			"cooldown": 1000,  //冷却时间
			"cost": [          //消耗品
				{
					"rewardId": 201,   //钻石
					"count": 5         //数量
				}
			]
		}
	],
	"unlockLevel": 50          //pve50关解锁治疗药水
}
```

# config.pvp.battlepass.json  没有使用

# config.pvp.coin.json     pvp金币配置
### 配置文件详解
```json5
{
	"1": [   
		{
			"lv": 1,         
			"percent": 9872
		}
	]
}
```

# config.pvp.plunder.json  pvp掠夺配置文件
### 配置文件详解
```json5
{
	"0": {
		"Productivity": "0",      //生产效率
		"GemProductivity": "0",   //钻石生产效率
		"Capacity": "0",          //掠夺容量
		"GemCapacity": "0",       //钻石容量
		"ProtectRate": "0",       //保护率
		"PlunderRate": "0"        //掠夺率
	}
}
```

## config.shop.dealchest.json	商店宝箱配置

下发配置文件接口

| 信息     | 说明               |
| -------- | ------------------ |
| 接口方式 | http GET           |
| 路由     | /sys/config?type=2 |

下发的配置文件

| 配置名称   | 说明             |
| ---------- | ---------------- |
| dealChests | 商店宝箱奖励配置 |

配置文件详解

```
{
	"unlockLevels": 0,		//解锁等级
	"chests": {						//宝箱
		"61": {							//宝箱ID
			"id": 61,					//宝箱ID
			"rewards": [			//奖励内容
				{
					"type": 2,		//奖励类型
					"rarity": 2,	//稀有度
					"cnt": 1,			//数量
					"level": 3		//等级
				},
				...
			}
			...
```

## config.shop.draw.lv.json	商店抽卡等级概率配置

使用配置

| 配置名称                          | 说明                     |
| --------------------------------- | ------------------------ |
| pkg.config.drawCard.drawCardLvCfg | 解析后的商店抽卡概率配置 |

配置文件详解

```
{
"sc": [								//商店金币
		{
			"lv": 1,				//等级
			"percent": 90		//百分比
		},
		...
	],
	"sg": {							//商店钻石
		"one": [
			{
				"rarity": 1,	//稀有度
				"prob": [			//概率
					{
						"lv": 1,
						"percent": 90 //百分比
					},
					{
						"lv": 2,
						"percent": 10
					}
				]
			},
			...
```

## config.stage.bonus.json  内部加奖励用户、锦标赛专属奖励配置

下发配置文件接口

| 信息     | 说明               |
| -------- | ------------------ |
| 接口方式 | http GET           |
| 路由     | /sys/config?type=1 |

下发的配置文件

| 配置名称   | 说明                               |
| ---------- | ---------------------------------- |
| champBonus | 内部加奖励用户、锦标赛专属奖励配置 |

配置文件详解

```
{
	"11": {
		"stage": "11",			//阶段
		"trophy": "6000",		//奖杯数
		"rewards": [
			{
				"rewardId": 301,//奖励ID
				"count": 1200		//奖励数量
			},
			{
				"rewardId": 101,
				"count": 4000
			}
		]
	},
	...
```

## config.stage.bonus.v2.json  内部加奖励用户、锦标赛专属奖励配置,（bp赛季 >= 18 使用）

下发配置文件接口

| 信息     | 说明               |
| -------- | ------------------ |
| 接口方式 | http GET           |
| 路由     | /sys/config?type=1 |

下发的配置文件

| 配置名称   | 说明                               |
| ---------- | ---------------------------------- |
| champBonus | 内部加奖励用户、锦标赛专属奖励配置 |

配置文件详解

```
{
	"11": {
		"stage": "11",			//阶段
		"trophy": "4000",		//奖杯数
		"rewards": [
			{
				"rewardId": 301,//奖励ID
				"count": 1200		//奖励数量
			},
			{
				"rewardId": 101,
				"count": 4000
			}
		]
	},
	...
```

## config.stars.random.json 出星概率配置

下发配置文件接口

| 信息     | 说明               |
| -------- | ------------------ |
| 接口方式 | http GET           |
| 路由     | /sys/config?type=1 |

下发的配置文件

| 配置名称                      | 说明                     |
| ----------------------------- | ------------------------ |
| chestDetail.starsRandomConfig | 宝箱详情中的出星概率配置 |

配置文件详解

```
{
	"1101": {											//宝箱ID
		"Id": 1101,									//宝箱ID
		"StarsInterval": "1,1",			//星星范围
		"StarsProbability": "1.0"		//星星概率
	},
	...
```


# config.task.dailytask.json 每日任务配置文件

### 下发配置文件接口

| 信息     | 说明         |
| -------- | ------------ |
| 接口方式 | http GET     |
| 路由     | /sys/dailyTaskConfig |


### 配置文件详解
```json5
{
	"task": [                 //任务
		{
			"id": 10101,      //任务id
			"active": 20,     //活跃值
			"target": 200,    //进度条最大值
			"jumpType": 3,    //跳转的地方
			"rewards": [      //奖励
				{
					"rewardId": 101,   //金币
					"count": 10000     //数量
				},
				{
					"rewardId": 201,   //钻石
					"count": 50        //数量
				}
			]
		}
	]
}
```

# config.team.champ.reward.json  公会联赛奖励

### 配置文件详解
```json5
{
	"1": {
		"Rank": 1,        //排名
		"Coin": 450000,   //金币奖励
		"Wood": 100000,   //木头奖励
		"Gem": 4800       //钻石奖励
	}
}
```

# config.team.task.box.json  公会升级宝箱配置

### 配置文件详解
```json5
[
	{
		"id": 1,                 //等级
		"upgrade": 0,            //没有使用到该字段
		"totalValue": 0,         //经验值
		"rewardInfo": [          //奖励内容
			{
				"rewardId": 301, //木头
				"count": 0       //数量
			},
			{
				"rewardId": 101, //金币
				"count": 0       //数量
			}
		]
	}
]
```

# config.team.task.json  公会任务配置

### 配置文件详解
```json5
{
	"targetTask": [          //公会任务
		{
			"id": 10101,     //任务id
			"active": 600,   //获得的活跃值
			"target": 3000,  //任务达成的目标
			"jumpType": 3    //完成按钮跳转到哪
		}
	]
}
```

# config.team.war.reward.json 公会战奖励配置

### 配置文件详解
```json5
{
	"1": {
		"Rank": 1,         //公会战排行榜的排名
		"Coin": 600000,    //金币奖励
		"Wood": 100000,    //木头奖励
		"Gem": 2400,       //钻石奖励
		"TeamExp": 2100    //公会经验值奖励
	}
}
```

# config.trophy.road.json  奖杯之路奖励配置

### 配置文件详解
```json5
{
	"0": {
		"Phase": 0,           //奖励阶段
		"ArenaStage": 1,      //竞技场阶段
		"Trophy": 0,          //奖杯数量
		"RewardCategory": 0,  //奖励种类
		"RewardType": 0,      //奖励id
		"Count": 0            //数量
	}
}
```

# config.team.war.v1.reward.json    没有使用

# config.teamwar.match.weight.json  没有使用

# config.unlock.boat.json   解锁船只配置
```json5
{
	"1": 1,    //"大本营等级": 船只数量
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 6,
	"8": 6,
	"9": 6,
	"10": 6,
	"11": 6,
	"12": 6,
	"13": 6,
	"14": 6,
	"15": 6,
	"16": 6
}
```

## config.vip.json VIP 配置

下发配置文件接口

| 信息     | 说明               |
| -------- | ------------------ |
| 接口方式 | http GET           |
| 路由     | /sys/config?type=2 |

下发的配置文件

| 配置名称 | 说明     |
| -------- | -------- |
| vip      | VIP 配置 |

配置文件详解

```
{
	"autoBattle": 50,									//自动战斗次数
	"drawCardDiscount": 10,						//抽卡折扣
	"buildSpeedup": 10,								//建造时间增幅百分比
  "placeBoatIntervalTime": 1000,		//放船的间隔时间
	"endPageIntervalTime": 3,					//页面间隔时间
	"airSoliderStartTime": 1000,			//放空投兵的间隔时间
	"dailyReward": [									//VIP每日奖励邮件
		{
			"rewardId": 101,							//奖励ID
			"count": 20000								//奖励数量
		},
		{
			"rewardId": 201,
			"count": 200
		}
	]
}
```

## config.wood.balance.json 大本等级木材配置

使用配置

| 配置名称                                   | 说明                     |
| ------------------------------------------ | ------------------------ |
| pkg.config.battleContent.WoodBalanceRobCfg | 解析后的大本等级木材配置 |

配置文件详解

```
{
	{
		"WoodProtect": 200,	//大本等级对应的木材
		"RobRatio": 0.450,	//计算基础奖励时乘的百分比
		"RobMax": 1000			//计算基础奖励最大木材
	},
	...
}
```

## config.pvp.stage.json pvp赛季段位对应的奖励

``` json5
{
  {
		"stage":6,//段位
  	"score":1500,//杯数
		"rewards":[//奖励
			{
				"rewardId":101,
				"count":20000
			}
		]
	},
...
}
```

## config.pvp.top.json pvp赛季前50奖励

``` json5
{
  {
		"top":1,//top排名
		"rewards":[//奖励
			{
				"rewardId":1920102,
				"count":12
			}
		]
	},
...
}
```



## config.clan.war.reward.json  公会ClanWar奖励内容配置

配置文件详解

```
	{
		"rankStart":1,	// 排名区间开始
		"rankEnd":1,	// 排名区间的结束
		"teamExp":2100,  // 经验值
		"rewards":[	// 奖励列表
			{
				"itemCnt":600000,
				"itemId":101
			},
			{
				"itemCnt":100000,
				"itemId":301
			},
			{
				"itemCnt":2400,
				"itemId":201
			}
		]
	},
```
# config.module.unlock.json 各个模块功能的解锁条件
配置文件详解
```json
{
		"module": 26,	// 功能模块id
		"level": 41,
		"score": 0
}

```
现有各模块id列表

| module id | 功能模块             |
| --------- | -------------------- |
| 1         | 每日任务             |
| 2         | 空投士兵             |
| 3         | 赏金模式             |
| 4         | 商店转盘             |
| 5         | 商店翻牌             |
| 6         | 战斗加速             |
| 7         | pvp模式              |
| 8         | pve药水              |
| 9         | 训练模式             |
| 10        | 天气                 |
| 11        | 公会                 |
| 12        | 商店冠军礼包         |
| 13        | 无尽试炼模式         |
| 14        | 积分精英赛           |
| 15        | 积分大师赛           |
| 16        | 商店熔炉             |
| 17        | 商店士兵宝箱         |
| 18        | 商店抽卡（金币钻石） |
| 19        | 商店周卡月卡         |
| 20        | 商店成长基金         |
| 21        | pvp排行榜            |
| 23        | 首页海岛入口         |
| 24        | 首页军营入口         |
| 25        | 高星招募             |
| 26        | pve章节难度选择      |
|           |                      |

## config.robot.json 机器人战力区间配置

```json
[{
	"start": 176,
	"end": 2000,
	"data": [
		"123","12312","123","123","3123"
	]
},{
	"start": 2000,
	"end": 5000,
	"data": [
		"123","12312","123","123","3123"
	]
}
]
```