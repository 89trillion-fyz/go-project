package util

const (
	Strnull = ""
	RdsNil  = "redis: nil"
	RdsPInf = "+inf"
	RdsNInf = "-inf"
	ANDROID = "android"
	IOS     = "ios"
	CoinId  = uint32(101)
	GemId   = uint32(201)
	WoodId  = uint32(301)
	MedalId = uint32(401)

	ArenaResetTrophySeason = 16

	// layout
	CityHallId        = int32(100103)
	GoldMineId        = int32(200103)
	ArmyCampId        = int32(600103)
	Timelate          = uint32(3)
	MaxArmylevel      = uint32(9)
	NeedCostArmyLevel = uint32(8)
	MaxBdlevel        = int(15)
	MaxBDID           = uint32(99)
	MaxFuseWslistNum  = int(8)
	BdUpTimeMin       = int(1600000000)

	// 强制升级
	EnforceUpgradeVersion = 2660       // 强制升级版本
	EnforceUpgradeTime    = 1621987200 // 强制升级开启时间

	ATKLAY  = uint32(1)
	DEFLAY  = uint32(2)
	WALLLAY = uint32(3)
	HEROLAY = uint32(4)

	TypeArmy  = 2
	TypeBuild = 3
	TypeWall  = 311
	TypeHero  = 5

	BuildTypeCityHall = 1
	BuildTypeGoldMine = 2
	BuildTypeArmyCamp = 6
	BuildTypeWall     = 11

	WallMove         = 1
	SoldierMove      = 2
	SoldierMerge     = 3
	OneClickMerge    = 4
	OneClickRetrieve = 5
	HeroAction       = 6

	ChangeIslandGemCost = 50

	//hero
	HeroNameSpliter = int(10000)
	HeroMinLv       = uint32(1)
	HeroMaxLv       = uint32(50)

	HeroMinStar = uint32(1)
	HeroMaxStar = uint32(5)

	MaxSkillNum    = uint32(4)
	MaxHeroSkillLv = uint32(5)

	// morale version
	MoraleVersion      = 2160
	MoraleRestoreProp  = 0.05
	WorkerCost         = 2000
	WorkerExpire       = int64(7 * DayTime)
	WorkerExpireV2     = int64(14 * DayTime)
	WorkerMax          = 2
	WorkerMaxV2        = 3
	WorkerNewVer       = 2310
	NewWorkerStartTime = 1624924800 //3个工人开启时间

	// time
	MapTime  = 120   // 一局战斗时长（秒）
	DayTime  = 86400 // 一天时长（秒）
	HourTime = 3600  // 一小时时长（秒）

	//Rwd claim 模块
	RwdClaimIdle             = 1  // idle reward
	RwdClaimBattlePass       = 2  // battlepass
	RwdClaimPointRace        = 3  // 积分赛 赛季奖励
	RwdClaimTrophyRoad       = 4  // 奖杯之路
	RwdClaimSpinExtra        = 5  // 转盘额外宝箱
	RwdClaimGrowthFound      = 6  // 成长基金
	RwdClaimWeekMonth        = 7  // 月卡周卡
	RwdClaimTeamLeagueIdle   = 8  // 公会联赛离线收益
	RwdClaimTeamLeagueSeason = 9  // 公会联赛赛季奖励
	RwdClaimTeamClimaxSeason = 10 // 公会巅峰赛赛季奖励
	RwdClaimRedeemCode       = 11 // 兑换码
	RwdClaimArenaRoad        = 12 // 新版奖杯之路
	RwdClaimChampRoad        = 13 // 奖杯之路-锦标赛
	RwdClaimInfWarDaily      = 14 // 无尽试炼-每日奖励
	RwdClaimInfWarStage      = 15 // 无尽试炼-进度奖励
	RwdClaimInfWarSeason     = 16 // 无尽试炼-赛季奖励
	RwdClaimAchievementMedal = 17 // 成就奖章奖励
	RwdClaimBPProcess        = 18 // 新版battlePass奖励
	RwdClaimBattlePassExtra  = 19 //领取 bp 额外宝箱
	RwdClaimDailyTask        = 20 // 每日任务奖励
	RwdClaimTeamTaskBox      = 21 //领取公会奖励
	// currencyService
	CoinField       = "coin"
	GemField        = "gem"
	WoodsField      = "wood"
	TroopMedalField = "trpm"

	// battleScore 战斗评分
	SLevel            = 3
	ALevel            = 2
	BLevel            = 1
	SLevelScore       = 100
	ALevelScore       = 80
	BLevelScore       = 50
	SLevelScoreChamp  = 100
	ALevelScoreChamp  = 85
	BLevelScoreChamp  = 60
	SLevelRwdRatio    = 1.2
	ALevelRwdRatio    = 1.1
	BLevelRwdRatio    = 1.05
	SLevelTrophyRatio = 1.2
	ALevelTrophyRatio = 1.1
	BLevelTrophyRatio = 1.05

	// bd up&new
	NormalFlag        = 0 // 建筑升级or新建正常流程Flag
	CostGemToWoodFlag = 1 // 花费钻石补足木头Flag
	CostGemUpBdFlag   = 2 // 花费钻石完成升级Flag
	BuildImmediateV1  = 1 // v1 立即升级
	BuildImmediateV2  = 2 // v2 立即升级

	// weather
	WeatherStart           = 1619625600   // 天气时序起点
	WeatherRefreshInterval = 4 * HourTime // 天气刷新间隔（秒）
	WeatherInitTotalNum    = 7 * DayTime / WeatherRefreshInterval
	WeatherForecastNum     = 5

	// ELO
	PointRaceScoreRatio     = 500 // 积分赛
	TeamChampScoreRatio     = 300 //公会巅峰赛
	ArenaScoreLowRatio      = 1000
	ArenaScoreMediumRatio   = 3000
	ArenaScoreHighRatio     = 5000 // pvp
	PREloOpenSeason         = 42   // 积分赛ELO开启赛季
	TeamLeagueEloOpenSeason = 5    // 公会联赛ELO开启赛季
	TeamChampEloOpenSeason  = 4    // 公会巅峰赛ELO开启赛季
	EloMaxScore             = 32   // 最大杯数
	EloPVPMaxScore          = 64   // 最大杯数
	EloTCMaxScore           = 200  //公会巅峰赛最大杯数

	// battleLog
	TeamLeagueBattleLogLength = 50 // 公会联赛日志长度

	// pointrace
	PointRaceInitScore  = 500 // 积分赛赛季初始积分
	PointRaceInitSeason = 44  //积分赛初始化积分开始赛季

	// 无尽试炼
	SeasonMaxPage = 4 // 50 一页
	DailyMaxPage  = 2

	// achievement
	AchVersion = 2

	// dailyTask 个人任务
	TaskVersion = 2

	//公户任务成员领取公会任务权限
	TeamTaskRwdAuthority = 3

	//钻石抽奖 卡券flag
	GemTicketDrawFlag = 2

	//pve 控制关卡解锁的用户id
	PveControlUnlockUser = 6411213
	//士兵来源显示版本
	ArmySourceCvc = 2700
)

//新兵配置下发分段版本控制
var ArmySourceCvcList = []int{2740, ArmySourceCvc}

//黑名单类型
const (
	BlockPermafrost = iota + 2 //永久拉黑
	BlockUnknown3
	BlockLimitedTime //限时拉黑
)

//战斗类型
const (
	AttackTypeArena = iota + 1
	AttackTypeTeamWar
	AttackTypePointRace
	AttackTypeHeadHunt
	AttackTypeTeamWarV1
	AttackTypeTeamLeague
	AttackTypeTeamChamp
	AttackTypeInfiniteWar
)

//房间类型
const (
	AtkRoomTypeArenaMatch = iota + 1
	AtkRoomTypeArenaRank
	AtkRoomTypePointElite
	AtkRoomTypePointMaster
	AtkRoomTypeInfiniteWar
)

//赛季奖励类型
const (
	SeasonRwdTypeElitePointRace = iota + 1
	SeasonRwdTypeMasterPointRace
	SeasonRwdTypeTeamLeague
	SeasonRwdTypeTeamChamp
	SeasonRwdTypeInfiniteWar
	DailyRwdTypeInfiniteWar
	SeasonRwdTypeClanWar
	SeasonRwdTypeArenaRace
)

//被打弹窗推送优化版本
const NewBeAttackPushVersion = 2350

//客户端版本号
const (
	NewPvpVersion2500            = 2500
	NewPvpRemoveScoreVersion2600 = 2580 //去除战斗评级版本号
	NewPveVersion                = 2700 //新版pve版本号
	PvpRobotVersion              = 2720 //地形下发版本号
)
