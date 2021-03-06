package main

var Paints map[string]string = map[string]string{
	"Brushed Black Steel":      "0",
	"Metallic Carbon Black":    "1",
	"Classic Choc Brown":       "2",
	"Metallic Schafter Purple": "3",
	"Metallic Hot Pink":        "4",
	"Metallic Formula Red":     "5",
	"Metallic Blue":            "6",
	"Classic Ultra Blue":       "7",
	"Classic Racing Green":     "8",
	"Matte Lime Green":         "9",
	"Classic Race Yellow":      "10",
	"Classic Orange":           "11",
	"Metallic Gold":            "12",
	"Metallic Silver":          "13",
	"Chrome":                   "14",
	"Classic Ice White":        "15",
	"Black":                    "16",
	"Graphite":                 "17",
	"Anthracite Black":         "18",
	"Black Steel":              "19",
	"Dark Steel":               "20",
	"Bluish Silver":            "21",
	"Rolled Steel":             "22",
	"Shadow Silver":            "23",
	"Stone Silver":             "24",
	"Midnight Silver":          "25",
	"Cast Iron Silver":         "26",
	"Red":                      "27",
	"Torino Red":               "28",
	"Lava Red":                 "29",
	"Blaze Red":                "30",
	"Grace Red":                "31",
	"Garnet Red":               "32",
	"Sunset Red":               "33",
	"Cabernet Red":             "34",
	"Wine Red":                 "35",
	"Candy Red":                "36",
	"Pfister Pink":             "37",
	"Salmon Pink":              "38",
	"Sunrise Orange":           "39",
	"Orange":                   "40",
	"Bright Orange":            "41",
	"Bronze":                   "42",
	"Yellow":                   "43",
	"Race Yellow":              "44",
	"Dew Yellow":               "45",
	"Dark Green":               "46",
	"Racing Green":             "47",
	"Sea Green":                "48",
	"Olive Green":              "49",
	"Bright Green":             "50",
	"Gasoline Green":           "51",
	"Lime Green":               "52",
	"Midnight Blue":            "53",
	"Galaxy Blue":              "54",
	"Dark Blue":                "55",
	"Saxon Blue":               "56",
	"Mariner Blue":             "57",
	"Harbor Blue":              "58",
	"Diamond Blue":             "59",
	"Surf Blue":                "60",
	"Nautical Blue":            "61",
	"Racing Blue":              "62",
	"Ultra Blue":               "63",
	"Light Blue":               "64",
	"Chocolate Brown":          "65",
	"Bison Brown":              "66",
	"Creek Brown":              "67",
	"Feltzer Brown":            "68",
	"Maple Brown":              "69",
	"Beechwood Brown":          "70",
	"Sienna Brown":             "71",
	"Saddle Brown":             "72",
	"Moss Brown":               "73",
	"Woodbeech Brown":          "74",
	"Straw Brown":              "75",
	"Sandy Brown":              "76",
	"Bleached Brown":           "77",
	"Spinnaker Purple":         "78",
	"Midnight Purple":          "79",
	"Bright Purple":            "80",
	"Cream":                    "81",
	"Ice White":                "82",
	"Frost White":              "83",
	"Black2":                   "84",
	"Carbon Black":             "85",
	"Graphite2":                "86",
	"Anthracite Black2":        "87",
	"Black Steel2":             "88",
	"Dark Steel2":              "89",
	"Silver2":                  "90",
	"Bluish Silver2":           "91",
	"Rolled Steel2":            "92",
	"Shadow Silver2":           "93",
	"Stone Silver2":            "94",
	"Midnight Silver2":         "95",
	"Cast Iron Silver2":        "96",
	"Red2":                     "97",
	"Torino Red2":              "98",
	"Formula Red2":             "99",
	"Lava Red2":                "100",
	"Blaze Red2":               "101",
	"Grace Red2":               "102",
	"Garnet Red2":              "103",
	"Sunset Red2":              "104",
	"Cabernet Red2":            "105",
	"Wine Red2":                "106",
	"Candy Red2":               "107",
	"Hot Pink2":                "108",
	"Pfister Pink2":            "109",
	"Salmon Pink2":             "110",
	"Sunrise Orange2":          "111",
	"Bright Orange2":           "112",
	"Gold2":                    "113",
	"Bronze2":                  "114",
	"Yellow2":                  "115",
	"Dew Yellow2":              "116",
	"Dark Green2":              "117",
	"Sea Green2":               "118",
	"Olive Green2":             "119",
	"Bright Green2":            "120",
	"Gasoline Green2":          "121",
	"Lime Green2":              "122",
	"Midnight Blue2":           "123",
	"Galaxy Blue2":             "124",
	"Dark Blue2":               "125",
	"Saxon Blue2":              "126",
	"Blue2":                    "127",
	"Mariner Blue2":            "128",
	"Harbor Blue2":             "129",
	"Diamond Blue2":            "130",
	"Surf Blue2":               "131",
	"Nautical Blue2":           "132",
	"Racing Blue2":             "133",
	"Light Blue2":              "134",
	"Bison Brown2":             "135",
	"Creek Brown2":             "136",
	"Feltzer Brown2":           "137",
	"Maple Brown2":             "138",
	"Beechwood Brown2":         "139",
	"Sienna Brown2":            "140",
	"Saddle Brown2":            "141",
	"Moss Brown2":              "142",
	"Woodbeech Brown2":         "143",
	"Straw Brown2":             "144",
	"Sandy Brown2":             "145",
	"Bleached Brown2":          "146",
	"Schafter Purple2":         "147",
	"Spinnaker Purple2":        "148",
	"Midnight Purple2":         "149",
	"Bright Purple2":           "150",
	"Cream2":                   "151",
	"Frost White2":             "152",
	"Black3":                   "153",
	"Gray3":                    "154",
	"Light Gray2":              "155",
	"Ice White2":               "156",
	"Blue3":                    "157",
	"Dark Blue3":               "158",
	"Midnight Blue3":           "159",
	"Midnight Purple3":         "160",
	"Schafter Purple":          "161",
	"Red3":                     "162",
	"Dark Red":                 "163",
	"Orange3":                  "164",
	"Yellow3":                  "165",
	"Green":                    "166",
	"Forest Green":             "167",
	"Foliage Green":            "168",
	"Olive Drab":               "169",
	"Dark Earth":               "170",
	"Desert Tan":               "171",
	"Brushed Steel":            "172",
	"Brushed Black Steel2":     "0",
	"Brushed Aluminum":         "173",
	"Pure Gold":                "174",
	"Brushed Gold":             "175",
}
