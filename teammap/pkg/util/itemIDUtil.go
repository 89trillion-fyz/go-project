package util

func GetTypeByID(id uint32) int {
	switch id % 100 {
	case 2:
		return TypeArmy
	case 3:
		switch id / 100000 {
		case 11:
			return TypeWall
		default:
			return TypeBuild
		}
	case 5:
		return TypeHero
	}
	return 0
}

func IsMainBd(bdId uint32) bool {
	id := int32(bdId)
	if id == CityHallId || id == GoldMineId || id == ArmyCampId {
		return true
	}
	return false
}

func IsBuildIDValid(bdId uint32) bool {
	if bdId%100 != 3 {
		return false
	}
	name := bdId / 100000
	if name < 1 || name > MaxBDID {
		return false
	}
	if name == 1 || name == 2 || name == 6 {
		if (bdId/100)%1000 != 1 {
			return false
		}
	}
	return true
}

func IsArmyValid(armyId uint32) bool {
	if armyId%100 != 2 {
		return false
	}
	armyId /= 100
	lv := armyId % 100
	name := armyId / 100
	if lv <= 0 || lv > MaxArmylevel {
		return false
	}
	if name < 100 || name > 999 {
		return false
	}
	return true
}

func IsWallValid(bdId uint32) bool {
	if bdId%100 == 3 && bdId/100000 == 11 {
		return true
	}
	return false
}

func IsHeroValid(heroId uint32) bool {
	if heroId%100 != 5 {
		return false
	}
	heroId /= 100
	lv := heroId % (uint32(HeroNameSpliter) / 100)
	name := heroId / (uint32(HeroNameSpliter) / 100)
	if lv <= 0 || lv > HeroMaxLv {
		return false
	}
	if name < 100 || name > 999 {
		return false
	}
	return true
}

func IsIdValid(id uint32) bool {
	switch id % 100 {
	case 2:
		id /= 100
		lv := id % 100
		name := id / 100
		if lv > 0 && lv <= MaxArmylevel && name >= 100 && name <= 999 {
			return true
		}
	case 3:
		name := id / 100000
		if name >= 1 && name <= MaxBDID {
			return true
		}
	case 5:
		id /= 100
		lv := id % (uint32(HeroNameSpliter) / 100)
		name := id / (uint32(HeroNameSpliter) / 100)
		if lv > 0 && lv <= HeroMaxLv && name >= 100 && name <= 999 {
			return true
		}
	}
	return false
}

func GetNameFromHeroId(heroId uint32) uint32 {
	if heroId%100 != 5 {
		return 0
	}
	return heroId / uint32(HeroNameSpliter)
}

func GetLevelFromHeroId(heroId uint32) uint32 {
	if heroId%100 != 5 {
		return 0
	}
	return (heroId % uint32(HeroNameSpliter)) / 100
}

func StickHeroId(name interface{}, lv interface{}) uint32 {
	return uint32(AnyToInt(name)*HeroNameSpliter + AnyToInt(lv)*100 + 5)
}

func IsThroneId(id uint32) bool {
	if id%100 == 3 {
		if id/100000 == 19 {
			return true
		}
	}
	return false
}

func GetScrapIdByHeroName(name uint32) uint32 {
	return name*100 + 7
}
