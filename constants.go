package main

// techLevels - lists the tech levels supported by the app
var (
	techLevels = []string{"A", "B", "C", "D", "E", "F", "G"}

	noneString = "0"

	// tons - lists the ship tonnage amounts supported by the app
	tons = []string{
		"100", "200", "300", "400", "500", "600", "700", "800", "900", "1000",
		"1100", "1200", "1300", "1400", "1500", "1600", "1700", "1800", "1900", "2000",
		"3000", "4000", "5000",
	}
)

type hullSelection struct {
	code  string
	tons  int
	price int
	weeks int
}

var hullSelectionCode = []string{
	"1", "2", "3", "4", "5", "6", "7", "8", "9",
	"A", "B", "C", "D", "E", "F", "G", "H", // Skipping I
	"J", "K", "L", "M", "N", // Skipping O
	"P",
}

func getIndexFromHull(dString string) int {
	for resultInt, dMatch := range hullSelectionCode {
		if dMatch == dString {
			return resultInt
		}
	}

	return -1
}

var hullSelections = []hullSelection{
	{hullSelectionCode[0], 100, 2, 36},      // 1
	{hullSelectionCode[1], 200, 8, 44},      // 2
	{hullSelectionCode[2], 300, 12, 52},     // 3
	{hullSelectionCode[3], 400, 16, 60},     // 4
	{hullSelectionCode[4], 500, 32, 68},     // 5
	{hullSelectionCode[5], 600, 48, 76},     // 6
	{hullSelectionCode[6], 700, 64, 84},     // 7
	{hullSelectionCode[7], 800, 80, 92},     // 8
	{hullSelectionCode[8], 900, 90, 100},    // 9
	{hullSelectionCode[9], 1000, 100, 108}, // A
	{hullSelectionCode[10], 1100, 110, 116}, // B
	{hullSelectionCode[11], 1200, 120, 124}, // C
	{hullSelectionCode[12], 1300, 130, 132}, // D
	{hullSelectionCode[13], 1400, 140, 140}, // E
	{hullSelectionCode[14], 1500, 150, 148}, // F
	{hullSelectionCode[15], 1600, 160, 156}, // G
	{hullSelectionCode[16], 1700, 170, 164}, // H
	{hullSelectionCode[17], 1800, 180, 172}, // J
	{hullSelectionCode[18], 1900, 190, 180}, // K
	{hullSelectionCode[19], 2000, 200, 188}, // L
	{hullSelectionCode[20], 3000, 300, 268}, // M
	{hullSelectionCode[21], 4000, 400, 348}, // N
	{hullSelectionCode[22], 5000, 500, 428}, // P
}

var hullSelectionCfgs = []string{"streamlined", "distributed", "standard"}

type hullSelectionConfiguration struct {
	cfg    string
	cost   float32
	scoops bool
}

var hullSelectionConfigurations = []hullSelectionConfiguration{
	{cfg: hullSelectionCfgs[0], cost: 1.1, scoops: true},
	{cfg: hullSelectionCfgs[1], cost: 0.9, scoops: false},
	{cfg: hullSelectionCfgs[2], cost: 1.0, scoops: true},
}

type offering struct {
	offer      string
	tlMin      int
	costPerTon float32
}

var hullSelectionOfferings = []offering{
	{offer: "Reflec", tlMin: 10, costPerTon: 0.1},
	{offer: "Self-Sealing", tlMin: 9, costPerTon: 0.01},
	{offer: "Stealth", tlMin: 11, costPerTon: 0.1},
}

type armor struct {
	armorType  string
	tlMin      int
	protection int
	cost       int
}

var (
	armors = []armor{
		{armorType: "Titanium Steel", tlMin: 7, protection: 2, cost: 5},
		{armorType: "CrystalIron", tlMin: 10, protection: 4, cost: 20},
		{armorType: "Bonded Superdense", tlMin: 14, protection: 6, cost: 50},
	}

	weaponLevel = []string{
		"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15",
		"16", "17", "18", "19", "20", "21", "22", "23", "24", "25", "26", "27", "28", "29", "30", "31",
		"32", "33", "34", "35", "36", "37", "38", "39", "40",
	}

	vehicleCount = []string{
		"0", "1", "2", "3", "4",
	}

	computer = [16]int{
		1, 2, 3, 4, 5, 7, 9, 11, 13, 15, 17, 20, 23, 26, 30, 35,
	}
)

var TrvIndex = []string{
	"A", "B", "C", "D", "E", "F", "G", "H", "J", "K", "L", "M", "N",
	"P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
}

const (
	defaultDriveCode  = "B"
	defaultHullCode   = "2"
	defaultTons       = 200
	defaultMaxhardpts = 2
	defaultHullPrice  = 8
	roundUp           = 0.999999

	// Vehicles
	vehicleTypes      = "Vehicles, Boats & Fighters"
	wATVType          = "Wheeled ATV"
	tATVType          = "Tracked ATV"
	airRaftType       = "Air Raft"
	speederType       = "Speeder"
	gCarrierType      = "GCarrier"
	shipsLaunchType   = "Ship's Launch"
	shipsBoatType     = "Ship's Boat"
	pinnaceType       = "Pinnace"
	cutterType        = "Modular Cutter"
	slowBoatType      = "Slow Boat"
	slowPinnaceType   = "Slow Pinnace"
	shuttleType       = "Shuttle"
	lightFighterType  = "Light Fighter"
	mediumFighterType = "Medium Fighter"
	heavyFighterType  = "Heavy Fighter"
)
