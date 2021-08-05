package lib

type StatType int

//stat types
//goland:noinspection ALL
const (
	DEFP StatType = iota
	DEF
	HP
	HPP
	ATK
	ATKP
	ER
	EM
	CR
	CD
	Heal
	PyroP
	HydroP
	CryoP
	ElectroP
	AnemoP
	GeoP
	PhyP
	EndStatType //delim
)

var SubWeights = [...]float64{
	100, //def%
	150, //def
	150, //hp
	100, //hp%
	150, //atk
	100, //atk%
	100, //er
	100, //em
	75,  //cr
	75,  //cd
	0,   //heal
	0,   //pyro
	0,   //hydro
	0,   //cryo
	0,   //electro
	0,   //anemo
	0,   //geo
	0,   //phys
}

var SubTier = [][]float64{
	{0.051, 0.0583, 0.0656, 0.0729},  //def%
	{16.2, 18.52, 20.83, 23.15},      //def
	{209.13, 239, 269.88, 299.75},    //hp
	{0.0408, 0.0466, 0.0525, 0.0583}, //hp%
	{13.62, 15.56, 17.51, 19.45},     //atk
	{0.0408, 0.0466, 0.0525, 0.0583}, //atk%
	{0.0453, 0.0518, 0.0583, 0.0648}, //er
	{16.32, 18.65, 20.98, 23.31},     //em
	{0.0272, 0.0311, 0.035, 0.0389},  //cr
	{0.0544, 0.0622, 0.0699, 0.0777}, //cd
}

var StatTypeString = [...]string{
	"def%",
	"def",
	"hp",
	"hp%",
	"atk",
	"atk%",
	"er",
	"em",
	"cr",
	"cd",
	"heal",
	"pyro%",
	"hydro%",
	"cryo%",
	"electro%",
	"anemo%",
	"geo%",
	"phys%",
	"n/a",
}

//goland:noinspection GoUnusedExportedFunction
func StrToStatType(s string) StatType {
	for i, v := range StatTypeString {
		if v == s {
			return StatType(i)
		}
	}
	return -1
}

type SlotType int

//goland:noinspection ALL
const (
	Flower SlotType = iota
	Feather
	Sand
	Goblet
	Circlet
	//delim
	EndSlotType
)

var SlotTypeString = [...]string{
	"flower",
	"feather",
	"sand",
	"goblet",
	"circlet",
}

func StrToSlotType(s string) SlotType {
	for i, v := range SlotTypeString {
		if v == s {
			return SlotType(i)
		}
	}
	return -1
}

var MainStatVal = [][]float64{
	//flower
	{
		0,    //def%
		0,    //def
		4780, //hp
		0,    //hp%
		0,    //atk
		0,    //atk%
		0,    //er
		0,    //em
		0,    //cr
		0,    //cd
		0,    //heal
		0,    //pyro
		0,    //hydro
		0,    //cryo
		0,    //electro
		0,    //anemo
		0,    //geo
		0,    //phys
	},
	//feather,
	{
		0,   //def%
		0,   //def
		0,   //hp
		0,   //hp%
		311, //atk
		0,   //atk%
		0,   //er
		0,   //em
		0,   //cr
		0,   //cd
		0,   //heal
		0,   //pyro
		0,   //hydro
		0,   //cryo
		0,   //electro
		0,   //anemo
		0,   //geo
		0,   //phys
	},
	//sand
	{
		.583, //def%
		0,    //def
		0,    //hp
		.466, //hp%
		0,    //atk
		.466, //atk%
		.518, //er
		187,  //em
		0,    //cr
		0,    //cd
		0,    //heal
		0,    //pyro
		0,    //hydro
		0,    //cryo
		0,    //electro
		0,    //anemo
		0,    //geo
		0,    //phys
	},
	//goblet
	{
		.583, //def%
		0,    //def
		0,    //hp
		.466, //hp%
		0,    //atk
		.466, //atk%
		0,    //er
		187,  //em
		0,    //cr
		0,    //cd
		0,    //heal
		.466, //pyro
		.466, //hydro
		.466, //cryo
		.466, //electro
		.466, //anemo
		.466, //geo
		.466, //phys
	},
	//circlet
	{
		.583, //def%
		0,    //def
		0,    //hp
		.466, //hp%
		0,    //atk
		.466, //atk%
		0,    //er
		187,  //em
		.311, //cr
		.622, //cd
		.359, //heal
		0,    //pyro
		0,    //hydro
		0,    //cryo
		0,    //electro
		0,    //anemo
		0,    //geo
		0,    //phys
	},
}

var MainStatWeight = [][]float64{
	//flower
	{
		0, //def%
		0, //def
		1, //hp
		0, //hp%
		0, //atk
		0, //atk%
		0, //er
		0, //em
		0, //cr
		0, //cd
		0, //heal
		0, //pyro
		0, //hydro
		0, //cryo
		0, //electro
		0, //anemo
		0, //geo
		0, //phys
	},
	//feather,
	{
		0, //def%
		0, //def
		0, //hp
		0, //hp%
		1, //atk
		0, //atk%
		0, //er
		0, //em
		0, //cr
		0, //cd
		0, //heal
		0, //pyro
		0, //hydro
		0, //cryo
		0, //electro
		0, //anemo
		0, //geo
		0, //phys
	},
	//sand
	{
		1333, //def%
		0,    //def
		0,    //hp
		1334, //hp%
		0,    //atk
		1333, //atk%
		500,  //er
		500,  //em
		0,    //cr
		0,    //cd
		0,    //heal
		0,    //pyro
		0,    //hydro
		0,    //cryo
		0,    //electro
		0,    //anemo
		0,    //geo
		0,    //phys
	},
	//goblet
	{
		800, //def%
		0,   //def
		0,   //hp
		850, //hp%
		0,   //atk
		850, //atk%
		0,   //er
		100, //em
		0,   //cr
		0,   //cd
		0,   //heal
		200, //pyro
		200, //hydro
		200, //cryo
		200, //electro
		200, //anemo
		200, //geo
		200, //phys
	},
	//circlet
	{
		1100, //def%
		0,    //def
		0,    //hp
		1100, //hp%
		0,    //atk
		1100, //atk%
		0,    //er
		200,  //em
		500,  //cr
		500,  //cd
		500,  //heal
		0,    //pyro
		0,    //hydro
		0,    //cryo
		0,    //electro
		0,    //anemo
		0,    //geo
		0,    //phys
	},
}
