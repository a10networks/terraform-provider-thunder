package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRuleSetTagStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_rule_set_tag_stats`: Statistics for the object tag\n\n__PLACEHOLDER__",
		ReadContext: resourceRuleSetTagStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"categorystat1": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 1",
						},
						"categorystat2": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 2",
						},
						"categorystat3": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 3",
						},
						"categorystat4": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 4",
						},
						"categorystat5": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 5",
						},
						"categorystat6": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 6",
						},
						"categorystat7": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 7",
						},
						"categorystat8": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 8",
						},
						"categorystat9": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 9",
						},
						"categorystat10": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 10",
						},
						"categorystat11": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 11",
						},
						"categorystat12": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 12",
						},
						"categorystat13": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 13",
						},
						"categorystat14": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 14",
						},
						"categorystat15": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 15",
						},
						"categorystat16": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 16",
						},
						"categorystat17": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 17",
						},
						"categorystat18": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 18",
						},
						"categorystat19": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 19",
						},
						"categorystat20": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 20",
						},
						"categorystat21": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 21",
						},
						"categorystat22": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 22",
						},
						"categorystat23": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 23",
						},
						"categorystat24": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 24",
						},
						"categorystat25": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 25",
						},
						"categorystat26": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 26",
						},
						"categorystat27": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 27",
						},
						"categorystat28": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 28",
						},
						"categorystat29": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 29",
						},
						"categorystat30": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 30",
						},
						"categorystat31": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 31",
						},
						"categorystat32": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 32",
						},
						"categorystat33": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 33",
						},
						"categorystat34": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 34",
						},
						"categorystat35": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 35",
						},
						"categorystat36": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 36",
						},
						"categorystat37": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 37",
						},
						"categorystat38": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 38",
						},
						"categorystat39": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 39",
						},
						"categorystat40": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 40",
						},
						"categorystat41": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 41",
						},
						"categorystat42": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 42",
						},
						"categorystat43": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 43",
						},
						"categorystat44": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 44",
						},
						"categorystat45": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 45",
						},
						"categorystat46": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 46",
						},
						"categorystat47": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 47",
						},
						"categorystat48": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 48",
						},
						"categorystat49": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 49",
						},
						"categorystat50": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 50",
						},
						"categorystat51": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 51",
						},
						"categorystat52": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 52",
						},
						"categorystat53": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 53",
						},
						"categorystat54": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 54",
						},
						"categorystat55": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 55",
						},
						"categorystat56": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 56",
						},
						"categorystat57": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 57",
						},
						"categorystat58": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 58",
						},
						"categorystat59": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 59",
						},
						"categorystat60": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 60",
						},
						"categorystat61": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 61",
						},
						"categorystat62": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 62",
						},
						"categorystat63": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 63",
						},
						"categorystat64": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 64",
						},
						"categorystat65": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 65",
						},
						"categorystat66": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 66",
						},
						"categorystat67": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 67",
						},
						"categorystat68": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 68",
						},
						"categorystat69": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 69",
						},
						"categorystat70": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 70",
						},
						"categorystat71": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 71",
						},
						"categorystat72": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 72",
						},
						"categorystat73": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 73",
						},
						"categorystat74": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 74",
						},
						"categorystat75": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 75",
						},
						"categorystat76": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 76",
						},
						"categorystat77": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 77",
						},
						"categorystat78": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 78",
						},
						"categorystat79": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 79",
						},
						"categorystat80": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 80",
						},
						"categorystat81": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 81",
						},
						"categorystat82": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 82",
						},
						"categorystat83": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 83",
						},
						"categorystat84": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 84",
						},
						"categorystat85": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 85",
						},
						"categorystat86": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 86",
						},
						"categorystat87": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 87",
						},
						"categorystat88": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 88",
						},
						"categorystat89": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 89",
						},
						"categorystat90": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 90",
						},
						"categorystat91": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 91",
						},
						"categorystat92": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 92",
						},
						"categorystat93": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 93",
						},
						"categorystat94": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 94",
						},
						"categorystat95": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 95",
						},
						"categorystat96": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 96",
						},
						"categorystat97": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 97",
						},
						"categorystat98": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 98",
						},
						"categorystat99": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 99",
						},
						"categorystat100": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 100",
						},
						"categorystat101": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 101",
						},
						"categorystat102": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 102",
						},
						"categorystat103": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 103",
						},
						"categorystat104": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 104",
						},
						"categorystat105": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 105",
						},
						"categorystat106": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 106",
						},
						"categorystat107": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 107",
						},
						"categorystat108": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 108",
						},
						"categorystat109": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 109",
						},
						"categorystat110": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 110",
						},
						"categorystat111": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 111",
						},
						"categorystat112": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 112",
						},
						"categorystat113": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 113",
						},
						"categorystat114": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 114",
						},
						"categorystat115": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 115",
						},
						"categorystat116": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 116",
						},
						"categorystat117": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 117",
						},
						"categorystat118": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 118",
						},
						"categorystat119": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 119",
						},
						"categorystat120": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 120",
						},
						"categorystat121": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 121",
						},
						"categorystat122": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 122",
						},
						"categorystat123": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 123",
						},
						"categorystat124": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 124",
						},
						"categorystat125": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 125",
						},
						"categorystat126": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 126",
						},
						"categorystat127": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 127",
						},
						"categorystat128": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 128",
						},
						"categorystat129": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 129",
						},
						"categorystat130": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 130",
						},
						"categorystat131": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 131",
						},
						"categorystat132": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 132",
						},
						"categorystat133": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 133",
						},
						"categorystat134": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 134",
						},
						"categorystat135": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 135",
						},
						"categorystat136": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 136",
						},
						"categorystat137": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 137",
						},
						"categorystat138": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 138",
						},
						"categorystat139": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 139",
						},
						"categorystat140": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 140",
						},
						"categorystat141": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 141",
						},
						"categorystat142": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 142",
						},
						"categorystat143": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 143",
						},
						"categorystat144": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 144",
						},
						"categorystat145": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 145",
						},
						"categorystat146": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 146",
						},
						"categorystat147": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 147",
						},
						"categorystat148": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 148",
						},
						"categorystat149": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 149",
						},
						"categorystat150": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 150",
						},
						"categorystat151": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 151",
						},
						"categorystat152": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 152",
						},
						"categorystat153": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 153",
						},
						"categorystat154": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 154",
						},
						"categorystat155": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 155",
						},
						"categorystat156": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 156",
						},
						"categorystat157": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 157",
						},
						"categorystat158": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 158",
						},
						"categorystat159": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 159",
						},
						"categorystat160": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 160",
						},
						"categorystat161": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 161",
						},
						"categorystat162": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 162",
						},
						"categorystat163": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 163",
						},
						"categorystat164": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 164",
						},
						"categorystat165": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 165",
						},
						"categorystat166": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 166",
						},
						"categorystat167": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 167",
						},
						"categorystat168": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 168",
						},
						"categorystat169": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 169",
						},
						"categorystat170": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 170",
						},
						"categorystat171": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 171",
						},
						"categorystat172": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 172",
						},
						"categorystat173": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 173",
						},
						"categorystat174": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 174",
						},
						"categorystat175": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 175",
						},
						"categorystat176": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 176",
						},
						"categorystat177": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 177",
						},
						"categorystat178": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 178",
						},
						"categorystat179": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 179",
						},
						"categorystat180": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 180",
						},
						"categorystat181": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 181",
						},
						"categorystat182": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 182",
						},
						"categorystat183": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 183",
						},
						"categorystat184": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 184",
						},
						"categorystat185": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 185",
						},
						"categorystat186": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 186",
						},
						"categorystat187": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 187",
						},
						"categorystat188": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 188",
						},
						"categorystat189": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 189",
						},
						"categorystat190": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 190",
						},
						"categorystat191": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 191",
						},
						"categorystat192": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 192",
						},
						"categorystat193": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 193",
						},
						"categorystat194": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 194",
						},
						"categorystat195": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 195",
						},
						"categorystat196": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 196",
						},
						"categorystat197": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 197",
						},
						"categorystat198": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 198",
						},
						"categorystat199": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 199",
						},
						"categorystat200": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 200",
						},
						"categorystat201": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 201",
						},
						"categorystat202": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 202",
						},
						"categorystat203": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 203",
						},
						"categorystat204": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 204",
						},
						"categorystat205": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 205",
						},
						"categorystat206": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 206",
						},
						"categorystat207": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 207",
						},
						"categorystat208": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 208",
						},
						"categorystat209": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 209",
						},
						"categorystat210": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 210",
						},
						"categorystat211": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 211",
						},
						"categorystat212": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 212",
						},
						"categorystat213": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 213",
						},
						"categorystat214": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 214",
						},
						"categorystat215": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 215",
						},
						"categorystat216": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 216",
						},
						"categorystat217": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 217",
						},
						"categorystat218": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 218",
						},
						"categorystat219": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 219",
						},
						"categorystat220": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 220",
						},
						"categorystat221": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 221",
						},
						"categorystat222": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 222",
						},
						"categorystat223": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 223",
						},
						"categorystat224": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 224",
						},
						"categorystat225": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 225",
						},
						"categorystat226": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 226",
						},
						"categorystat227": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 227",
						},
						"categorystat228": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 228",
						},
						"categorystat229": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 229",
						},
						"categorystat230": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 230",
						},
						"categorystat231": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 231",
						},
						"categorystat232": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 232",
						},
						"categorystat233": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 233",
						},
						"categorystat234": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 234",
						},
						"categorystat235": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 235",
						},
						"categorystat236": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 236",
						},
						"categorystat237": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 237",
						},
						"categorystat238": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 238",
						},
						"categorystat239": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 239",
						},
						"categorystat240": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 240",
						},
						"categorystat241": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 241",
						},
						"categorystat242": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 242",
						},
						"categorystat243": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 243",
						},
						"categorystat244": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 244",
						},
						"categorystat245": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 245",
						},
						"categorystat246": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 246",
						},
						"categorystat247": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 247",
						},
						"categorystat248": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 248",
						},
						"categorystat249": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 249",
						},
						"categorystat250": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 250",
						},
						"categorystat251": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 251",
						},
						"categorystat252": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 252",
						},
						"categorystat253": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 253",
						},
						"categorystat254": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 254",
						},
						"categorystat255": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 255",
						},
						"categorystat256": {
							Type: schema.TypeInt, Optional: true, Description: "counter app category stat 255",
						},
					},
				},
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}

func resourceRuleSetTagStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRuleSetTagStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRuleSetTagStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		RuleSetTagStatsStats := setObjectRuleSetTagStatsStats(res)
		d.Set("stats", RuleSetTagStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectRuleSetTagStatsStats(ret edpt.DataRuleSetTagStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"categorystat1":   ret.DtRuleSetTagStats.Stats.Categorystat1,
			"categorystat2":   ret.DtRuleSetTagStats.Stats.Categorystat2,
			"categorystat3":   ret.DtRuleSetTagStats.Stats.Categorystat3,
			"categorystat4":   ret.DtRuleSetTagStats.Stats.Categorystat4,
			"categorystat5":   ret.DtRuleSetTagStats.Stats.Categorystat5,
			"categorystat6":   ret.DtRuleSetTagStats.Stats.Categorystat6,
			"categorystat7":   ret.DtRuleSetTagStats.Stats.Categorystat7,
			"categorystat8":   ret.DtRuleSetTagStats.Stats.Categorystat8,
			"categorystat9":   ret.DtRuleSetTagStats.Stats.Categorystat9,
			"categorystat10":  ret.DtRuleSetTagStats.Stats.Categorystat10,
			"categorystat11":  ret.DtRuleSetTagStats.Stats.Categorystat11,
			"categorystat12":  ret.DtRuleSetTagStats.Stats.Categorystat12,
			"categorystat13":  ret.DtRuleSetTagStats.Stats.Categorystat13,
			"categorystat14":  ret.DtRuleSetTagStats.Stats.Categorystat14,
			"categorystat15":  ret.DtRuleSetTagStats.Stats.Categorystat15,
			"categorystat16":  ret.DtRuleSetTagStats.Stats.Categorystat16,
			"categorystat17":  ret.DtRuleSetTagStats.Stats.Categorystat17,
			"categorystat18":  ret.DtRuleSetTagStats.Stats.Categorystat18,
			"categorystat19":  ret.DtRuleSetTagStats.Stats.Categorystat19,
			"categorystat20":  ret.DtRuleSetTagStats.Stats.Categorystat20,
			"categorystat21":  ret.DtRuleSetTagStats.Stats.Categorystat21,
			"categorystat22":  ret.DtRuleSetTagStats.Stats.Categorystat22,
			"categorystat23":  ret.DtRuleSetTagStats.Stats.Categorystat23,
			"categorystat24":  ret.DtRuleSetTagStats.Stats.Categorystat24,
			"categorystat25":  ret.DtRuleSetTagStats.Stats.Categorystat25,
			"categorystat26":  ret.DtRuleSetTagStats.Stats.Categorystat26,
			"categorystat27":  ret.DtRuleSetTagStats.Stats.Categorystat27,
			"categorystat28":  ret.DtRuleSetTagStats.Stats.Categorystat28,
			"categorystat29":  ret.DtRuleSetTagStats.Stats.Categorystat29,
			"categorystat30":  ret.DtRuleSetTagStats.Stats.Categorystat30,
			"categorystat31":  ret.DtRuleSetTagStats.Stats.Categorystat31,
			"categorystat32":  ret.DtRuleSetTagStats.Stats.Categorystat32,
			"categorystat33":  ret.DtRuleSetTagStats.Stats.Categorystat33,
			"categorystat34":  ret.DtRuleSetTagStats.Stats.Categorystat34,
			"categorystat35":  ret.DtRuleSetTagStats.Stats.Categorystat35,
			"categorystat36":  ret.DtRuleSetTagStats.Stats.Categorystat36,
			"categorystat37":  ret.DtRuleSetTagStats.Stats.Categorystat37,
			"categorystat38":  ret.DtRuleSetTagStats.Stats.Categorystat38,
			"categorystat39":  ret.DtRuleSetTagStats.Stats.Categorystat39,
			"categorystat40":  ret.DtRuleSetTagStats.Stats.Categorystat40,
			"categorystat41":  ret.DtRuleSetTagStats.Stats.Categorystat41,
			"categorystat42":  ret.DtRuleSetTagStats.Stats.Categorystat42,
			"categorystat43":  ret.DtRuleSetTagStats.Stats.Categorystat43,
			"categorystat44":  ret.DtRuleSetTagStats.Stats.Categorystat44,
			"categorystat45":  ret.DtRuleSetTagStats.Stats.Categorystat45,
			"categorystat46":  ret.DtRuleSetTagStats.Stats.Categorystat46,
			"categorystat47":  ret.DtRuleSetTagStats.Stats.Categorystat47,
			"categorystat48":  ret.DtRuleSetTagStats.Stats.Categorystat48,
			"categorystat49":  ret.DtRuleSetTagStats.Stats.Categorystat49,
			"categorystat50":  ret.DtRuleSetTagStats.Stats.Categorystat50,
			"categorystat51":  ret.DtRuleSetTagStats.Stats.Categorystat51,
			"categorystat52":  ret.DtRuleSetTagStats.Stats.Categorystat52,
			"categorystat53":  ret.DtRuleSetTagStats.Stats.Categorystat53,
			"categorystat54":  ret.DtRuleSetTagStats.Stats.Categorystat54,
			"categorystat55":  ret.DtRuleSetTagStats.Stats.Categorystat55,
			"categorystat56":  ret.DtRuleSetTagStats.Stats.Categorystat56,
			"categorystat57":  ret.DtRuleSetTagStats.Stats.Categorystat57,
			"categorystat58":  ret.DtRuleSetTagStats.Stats.Categorystat58,
			"categorystat59":  ret.DtRuleSetTagStats.Stats.Categorystat59,
			"categorystat60":  ret.DtRuleSetTagStats.Stats.Categorystat60,
			"categorystat61":  ret.DtRuleSetTagStats.Stats.Categorystat61,
			"categorystat62":  ret.DtRuleSetTagStats.Stats.Categorystat62,
			"categorystat63":  ret.DtRuleSetTagStats.Stats.Categorystat63,
			"categorystat64":  ret.DtRuleSetTagStats.Stats.Categorystat64,
			"categorystat65":  ret.DtRuleSetTagStats.Stats.Categorystat65,
			"categorystat66":  ret.DtRuleSetTagStats.Stats.Categorystat66,
			"categorystat67":  ret.DtRuleSetTagStats.Stats.Categorystat67,
			"categorystat68":  ret.DtRuleSetTagStats.Stats.Categorystat68,
			"categorystat69":  ret.DtRuleSetTagStats.Stats.Categorystat69,
			"categorystat70":  ret.DtRuleSetTagStats.Stats.Categorystat70,
			"categorystat71":  ret.DtRuleSetTagStats.Stats.Categorystat71,
			"categorystat72":  ret.DtRuleSetTagStats.Stats.Categorystat72,
			"categorystat73":  ret.DtRuleSetTagStats.Stats.Categorystat73,
			"categorystat74":  ret.DtRuleSetTagStats.Stats.Categorystat74,
			"categorystat75":  ret.DtRuleSetTagStats.Stats.Categorystat75,
			"categorystat76":  ret.DtRuleSetTagStats.Stats.Categorystat76,
			"categorystat77":  ret.DtRuleSetTagStats.Stats.Categorystat77,
			"categorystat78":  ret.DtRuleSetTagStats.Stats.Categorystat78,
			"categorystat79":  ret.DtRuleSetTagStats.Stats.Categorystat79,
			"categorystat80":  ret.DtRuleSetTagStats.Stats.Categorystat80,
			"categorystat81":  ret.DtRuleSetTagStats.Stats.Categorystat81,
			"categorystat82":  ret.DtRuleSetTagStats.Stats.Categorystat82,
			"categorystat83":  ret.DtRuleSetTagStats.Stats.Categorystat83,
			"categorystat84":  ret.DtRuleSetTagStats.Stats.Categorystat84,
			"categorystat85":  ret.DtRuleSetTagStats.Stats.Categorystat85,
			"categorystat86":  ret.DtRuleSetTagStats.Stats.Categorystat86,
			"categorystat87":  ret.DtRuleSetTagStats.Stats.Categorystat87,
			"categorystat88":  ret.DtRuleSetTagStats.Stats.Categorystat88,
			"categorystat89":  ret.DtRuleSetTagStats.Stats.Categorystat89,
			"categorystat90":  ret.DtRuleSetTagStats.Stats.Categorystat90,
			"categorystat91":  ret.DtRuleSetTagStats.Stats.Categorystat91,
			"categorystat92":  ret.DtRuleSetTagStats.Stats.Categorystat92,
			"categorystat93":  ret.DtRuleSetTagStats.Stats.Categorystat93,
			"categorystat94":  ret.DtRuleSetTagStats.Stats.Categorystat94,
			"categorystat95":  ret.DtRuleSetTagStats.Stats.Categorystat95,
			"categorystat96":  ret.DtRuleSetTagStats.Stats.Categorystat96,
			"categorystat97":  ret.DtRuleSetTagStats.Stats.Categorystat97,
			"categorystat98":  ret.DtRuleSetTagStats.Stats.Categorystat98,
			"categorystat99":  ret.DtRuleSetTagStats.Stats.Categorystat99,
			"categorystat100": ret.DtRuleSetTagStats.Stats.Categorystat100,
			"categorystat101": ret.DtRuleSetTagStats.Stats.Categorystat101,
			"categorystat102": ret.DtRuleSetTagStats.Stats.Categorystat102,
			"categorystat103": ret.DtRuleSetTagStats.Stats.Categorystat103,
			"categorystat104": ret.DtRuleSetTagStats.Stats.Categorystat104,
			"categorystat105": ret.DtRuleSetTagStats.Stats.Categorystat105,
			"categorystat106": ret.DtRuleSetTagStats.Stats.Categorystat106,
			"categorystat107": ret.DtRuleSetTagStats.Stats.Categorystat107,
			"categorystat108": ret.DtRuleSetTagStats.Stats.Categorystat108,
			"categorystat109": ret.DtRuleSetTagStats.Stats.Categorystat109,
			"categorystat110": ret.DtRuleSetTagStats.Stats.Categorystat110,
			"categorystat111": ret.DtRuleSetTagStats.Stats.Categorystat111,
			"categorystat112": ret.DtRuleSetTagStats.Stats.Categorystat112,
			"categorystat113": ret.DtRuleSetTagStats.Stats.Categorystat113,
			"categorystat114": ret.DtRuleSetTagStats.Stats.Categorystat114,
			"categorystat115": ret.DtRuleSetTagStats.Stats.Categorystat115,
			"categorystat116": ret.DtRuleSetTagStats.Stats.Categorystat116,
			"categorystat117": ret.DtRuleSetTagStats.Stats.Categorystat117,
			"categorystat118": ret.DtRuleSetTagStats.Stats.Categorystat118,
			"categorystat119": ret.DtRuleSetTagStats.Stats.Categorystat119,
			"categorystat120": ret.DtRuleSetTagStats.Stats.Categorystat120,
			"categorystat121": ret.DtRuleSetTagStats.Stats.Categorystat121,
			"categorystat122": ret.DtRuleSetTagStats.Stats.Categorystat122,
			"categorystat123": ret.DtRuleSetTagStats.Stats.Categorystat123,
			"categorystat124": ret.DtRuleSetTagStats.Stats.Categorystat124,
			"categorystat125": ret.DtRuleSetTagStats.Stats.Categorystat125,
			"categorystat126": ret.DtRuleSetTagStats.Stats.Categorystat126,
			"categorystat127": ret.DtRuleSetTagStats.Stats.Categorystat127,
			"categorystat128": ret.DtRuleSetTagStats.Stats.Categorystat128,
			"categorystat129": ret.DtRuleSetTagStats.Stats.Categorystat129,
			"categorystat130": ret.DtRuleSetTagStats.Stats.Categorystat130,
			"categorystat131": ret.DtRuleSetTagStats.Stats.Categorystat131,
			"categorystat132": ret.DtRuleSetTagStats.Stats.Categorystat132,
			"categorystat133": ret.DtRuleSetTagStats.Stats.Categorystat133,
			"categorystat134": ret.DtRuleSetTagStats.Stats.Categorystat134,
			"categorystat135": ret.DtRuleSetTagStats.Stats.Categorystat135,
			"categorystat136": ret.DtRuleSetTagStats.Stats.Categorystat136,
			"categorystat137": ret.DtRuleSetTagStats.Stats.Categorystat137,
			"categorystat138": ret.DtRuleSetTagStats.Stats.Categorystat138,
			"categorystat139": ret.DtRuleSetTagStats.Stats.Categorystat139,
			"categorystat140": ret.DtRuleSetTagStats.Stats.Categorystat140,
			"categorystat141": ret.DtRuleSetTagStats.Stats.Categorystat141,
			"categorystat142": ret.DtRuleSetTagStats.Stats.Categorystat142,
			"categorystat143": ret.DtRuleSetTagStats.Stats.Categorystat143,
			"categorystat144": ret.DtRuleSetTagStats.Stats.Categorystat144,
			"categorystat145": ret.DtRuleSetTagStats.Stats.Categorystat145,
			"categorystat146": ret.DtRuleSetTagStats.Stats.Categorystat146,
			"categorystat147": ret.DtRuleSetTagStats.Stats.Categorystat147,
			"categorystat148": ret.DtRuleSetTagStats.Stats.Categorystat148,
			"categorystat149": ret.DtRuleSetTagStats.Stats.Categorystat149,
			"categorystat150": ret.DtRuleSetTagStats.Stats.Categorystat150,
			"categorystat151": ret.DtRuleSetTagStats.Stats.Categorystat151,
			"categorystat152": ret.DtRuleSetTagStats.Stats.Categorystat152,
			"categorystat153": ret.DtRuleSetTagStats.Stats.Categorystat153,
			"categorystat154": ret.DtRuleSetTagStats.Stats.Categorystat154,
			"categorystat155": ret.DtRuleSetTagStats.Stats.Categorystat155,
			"categorystat156": ret.DtRuleSetTagStats.Stats.Categorystat156,
			"categorystat157": ret.DtRuleSetTagStats.Stats.Categorystat157,
			"categorystat158": ret.DtRuleSetTagStats.Stats.Categorystat158,
			"categorystat159": ret.DtRuleSetTagStats.Stats.Categorystat159,
			"categorystat160": ret.DtRuleSetTagStats.Stats.Categorystat160,
			"categorystat161": ret.DtRuleSetTagStats.Stats.Categorystat161,
			"categorystat162": ret.DtRuleSetTagStats.Stats.Categorystat162,
			"categorystat163": ret.DtRuleSetTagStats.Stats.Categorystat163,
			"categorystat164": ret.DtRuleSetTagStats.Stats.Categorystat164,
			"categorystat165": ret.DtRuleSetTagStats.Stats.Categorystat165,
			"categorystat166": ret.DtRuleSetTagStats.Stats.Categorystat166,
			"categorystat167": ret.DtRuleSetTagStats.Stats.Categorystat167,
			"categorystat168": ret.DtRuleSetTagStats.Stats.Categorystat168,
			"categorystat169": ret.DtRuleSetTagStats.Stats.Categorystat169,
			"categorystat170": ret.DtRuleSetTagStats.Stats.Categorystat170,
			"categorystat171": ret.DtRuleSetTagStats.Stats.Categorystat171,
			"categorystat172": ret.DtRuleSetTagStats.Stats.Categorystat172,
			"categorystat173": ret.DtRuleSetTagStats.Stats.Categorystat173,
			"categorystat174": ret.DtRuleSetTagStats.Stats.Categorystat174,
			"categorystat175": ret.DtRuleSetTagStats.Stats.Categorystat175,
			"categorystat176": ret.DtRuleSetTagStats.Stats.Categorystat176,
			"categorystat177": ret.DtRuleSetTagStats.Stats.Categorystat177,
			"categorystat178": ret.DtRuleSetTagStats.Stats.Categorystat178,
			"categorystat179": ret.DtRuleSetTagStats.Stats.Categorystat179,
			"categorystat180": ret.DtRuleSetTagStats.Stats.Categorystat180,
			"categorystat181": ret.DtRuleSetTagStats.Stats.Categorystat181,
			"categorystat182": ret.DtRuleSetTagStats.Stats.Categorystat182,
			"categorystat183": ret.DtRuleSetTagStats.Stats.Categorystat183,
			"categorystat184": ret.DtRuleSetTagStats.Stats.Categorystat184,
			"categorystat185": ret.DtRuleSetTagStats.Stats.Categorystat185,
			"categorystat186": ret.DtRuleSetTagStats.Stats.Categorystat186,
			"categorystat187": ret.DtRuleSetTagStats.Stats.Categorystat187,
			"categorystat188": ret.DtRuleSetTagStats.Stats.Categorystat188,
			"categorystat189": ret.DtRuleSetTagStats.Stats.Categorystat189,
			"categorystat190": ret.DtRuleSetTagStats.Stats.Categorystat190,
			"categorystat191": ret.DtRuleSetTagStats.Stats.Categorystat191,
			"categorystat192": ret.DtRuleSetTagStats.Stats.Categorystat192,
			"categorystat193": ret.DtRuleSetTagStats.Stats.Categorystat193,
			"categorystat194": ret.DtRuleSetTagStats.Stats.Categorystat194,
			"categorystat195": ret.DtRuleSetTagStats.Stats.Categorystat195,
			"categorystat196": ret.DtRuleSetTagStats.Stats.Categorystat196,
			"categorystat197": ret.DtRuleSetTagStats.Stats.Categorystat197,
			"categorystat198": ret.DtRuleSetTagStats.Stats.Categorystat198,
			"categorystat199": ret.DtRuleSetTagStats.Stats.Categorystat199,
			"categorystat200": ret.DtRuleSetTagStats.Stats.Categorystat200,
			"categorystat201": ret.DtRuleSetTagStats.Stats.Categorystat201,
			"categorystat202": ret.DtRuleSetTagStats.Stats.Categorystat202,
			"categorystat203": ret.DtRuleSetTagStats.Stats.Categorystat203,
			"categorystat204": ret.DtRuleSetTagStats.Stats.Categorystat204,
			"categorystat205": ret.DtRuleSetTagStats.Stats.Categorystat205,
			"categorystat206": ret.DtRuleSetTagStats.Stats.Categorystat206,
			"categorystat207": ret.DtRuleSetTagStats.Stats.Categorystat207,
			"categorystat208": ret.DtRuleSetTagStats.Stats.Categorystat208,
			"categorystat209": ret.DtRuleSetTagStats.Stats.Categorystat209,
			"categorystat210": ret.DtRuleSetTagStats.Stats.Categorystat210,
			"categorystat211": ret.DtRuleSetTagStats.Stats.Categorystat211,
			"categorystat212": ret.DtRuleSetTagStats.Stats.Categorystat212,
			"categorystat213": ret.DtRuleSetTagStats.Stats.Categorystat213,
			"categorystat214": ret.DtRuleSetTagStats.Stats.Categorystat214,
			"categorystat215": ret.DtRuleSetTagStats.Stats.Categorystat215,
			"categorystat216": ret.DtRuleSetTagStats.Stats.Categorystat216,
			"categorystat217": ret.DtRuleSetTagStats.Stats.Categorystat217,
			"categorystat218": ret.DtRuleSetTagStats.Stats.Categorystat218,
			"categorystat219": ret.DtRuleSetTagStats.Stats.Categorystat219,
			"categorystat220": ret.DtRuleSetTagStats.Stats.Categorystat220,
			"categorystat221": ret.DtRuleSetTagStats.Stats.Categorystat221,
			"categorystat222": ret.DtRuleSetTagStats.Stats.Categorystat222,
			"categorystat223": ret.DtRuleSetTagStats.Stats.Categorystat223,
			"categorystat224": ret.DtRuleSetTagStats.Stats.Categorystat224,
			"categorystat225": ret.DtRuleSetTagStats.Stats.Categorystat225,
			"categorystat226": ret.DtRuleSetTagStats.Stats.Categorystat226,
			"categorystat227": ret.DtRuleSetTagStats.Stats.Categorystat227,
			"categorystat228": ret.DtRuleSetTagStats.Stats.Categorystat228,
			"categorystat229": ret.DtRuleSetTagStats.Stats.Categorystat229,
			"categorystat230": ret.DtRuleSetTagStats.Stats.Categorystat230,
			"categorystat231": ret.DtRuleSetTagStats.Stats.Categorystat231,
			"categorystat232": ret.DtRuleSetTagStats.Stats.Categorystat232,
			"categorystat233": ret.DtRuleSetTagStats.Stats.Categorystat233,
			"categorystat234": ret.DtRuleSetTagStats.Stats.Categorystat234,
			"categorystat235": ret.DtRuleSetTagStats.Stats.Categorystat235,
			"categorystat236": ret.DtRuleSetTagStats.Stats.Categorystat236,
			"categorystat237": ret.DtRuleSetTagStats.Stats.Categorystat237,
			"categorystat238": ret.DtRuleSetTagStats.Stats.Categorystat238,
			"categorystat239": ret.DtRuleSetTagStats.Stats.Categorystat239,
			"categorystat240": ret.DtRuleSetTagStats.Stats.Categorystat240,
			"categorystat241": ret.DtRuleSetTagStats.Stats.Categorystat241,
			"categorystat242": ret.DtRuleSetTagStats.Stats.Categorystat242,
			"categorystat243": ret.DtRuleSetTagStats.Stats.Categorystat243,
			"categorystat244": ret.DtRuleSetTagStats.Stats.Categorystat244,
			"categorystat245": ret.DtRuleSetTagStats.Stats.Categorystat245,
			"categorystat246": ret.DtRuleSetTagStats.Stats.Categorystat246,
			"categorystat247": ret.DtRuleSetTagStats.Stats.Categorystat247,
			"categorystat248": ret.DtRuleSetTagStats.Stats.Categorystat248,
			"categorystat249": ret.DtRuleSetTagStats.Stats.Categorystat249,
			"categorystat250": ret.DtRuleSetTagStats.Stats.Categorystat250,
			"categorystat251": ret.DtRuleSetTagStats.Stats.Categorystat251,
			"categorystat252": ret.DtRuleSetTagStats.Stats.Categorystat252,
			"categorystat253": ret.DtRuleSetTagStats.Stats.Categorystat253,
			"categorystat254": ret.DtRuleSetTagStats.Stats.Categorystat254,
			"categorystat255": ret.DtRuleSetTagStats.Stats.Categorystat255,
			"categorystat256": ret.DtRuleSetTagStats.Stats.Categorystat256,
		},
	}
}

func getObjectRuleSetTagStatsStats(d []interface{}) edpt.RuleSetTagStatsStats {

	count1 := len(d)
	var ret edpt.RuleSetTagStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Categorystat1 = in["categorystat1"].(int)
		ret.Categorystat2 = in["categorystat2"].(int)
		ret.Categorystat3 = in["categorystat3"].(int)
		ret.Categorystat4 = in["categorystat4"].(int)
		ret.Categorystat5 = in["categorystat5"].(int)
		ret.Categorystat6 = in["categorystat6"].(int)
		ret.Categorystat7 = in["categorystat7"].(int)
		ret.Categorystat8 = in["categorystat8"].(int)
		ret.Categorystat9 = in["categorystat9"].(int)
		ret.Categorystat10 = in["categorystat10"].(int)
		ret.Categorystat11 = in["categorystat11"].(int)
		ret.Categorystat12 = in["categorystat12"].(int)
		ret.Categorystat13 = in["categorystat13"].(int)
		ret.Categorystat14 = in["categorystat14"].(int)
		ret.Categorystat15 = in["categorystat15"].(int)
		ret.Categorystat16 = in["categorystat16"].(int)
		ret.Categorystat17 = in["categorystat17"].(int)
		ret.Categorystat18 = in["categorystat18"].(int)
		ret.Categorystat19 = in["categorystat19"].(int)
		ret.Categorystat20 = in["categorystat20"].(int)
		ret.Categorystat21 = in["categorystat21"].(int)
		ret.Categorystat22 = in["categorystat22"].(int)
		ret.Categorystat23 = in["categorystat23"].(int)
		ret.Categorystat24 = in["categorystat24"].(int)
		ret.Categorystat25 = in["categorystat25"].(int)
		ret.Categorystat26 = in["categorystat26"].(int)
		ret.Categorystat27 = in["categorystat27"].(int)
		ret.Categorystat28 = in["categorystat28"].(int)
		ret.Categorystat29 = in["categorystat29"].(int)
		ret.Categorystat30 = in["categorystat30"].(int)
		ret.Categorystat31 = in["categorystat31"].(int)
		ret.Categorystat32 = in["categorystat32"].(int)
		ret.Categorystat33 = in["categorystat33"].(int)
		ret.Categorystat34 = in["categorystat34"].(int)
		ret.Categorystat35 = in["categorystat35"].(int)
		ret.Categorystat36 = in["categorystat36"].(int)
		ret.Categorystat37 = in["categorystat37"].(int)
		ret.Categorystat38 = in["categorystat38"].(int)
		ret.Categorystat39 = in["categorystat39"].(int)
		ret.Categorystat40 = in["categorystat40"].(int)
		ret.Categorystat41 = in["categorystat41"].(int)
		ret.Categorystat42 = in["categorystat42"].(int)
		ret.Categorystat43 = in["categorystat43"].(int)
		ret.Categorystat44 = in["categorystat44"].(int)
		ret.Categorystat45 = in["categorystat45"].(int)
		ret.Categorystat46 = in["categorystat46"].(int)
		ret.Categorystat47 = in["categorystat47"].(int)
		ret.Categorystat48 = in["categorystat48"].(int)
		ret.Categorystat49 = in["categorystat49"].(int)
		ret.Categorystat50 = in["categorystat50"].(int)
		ret.Categorystat51 = in["categorystat51"].(int)
		ret.Categorystat52 = in["categorystat52"].(int)
		ret.Categorystat53 = in["categorystat53"].(int)
		ret.Categorystat54 = in["categorystat54"].(int)
		ret.Categorystat55 = in["categorystat55"].(int)
		ret.Categorystat56 = in["categorystat56"].(int)
		ret.Categorystat57 = in["categorystat57"].(int)
		ret.Categorystat58 = in["categorystat58"].(int)
		ret.Categorystat59 = in["categorystat59"].(int)
		ret.Categorystat60 = in["categorystat60"].(int)
		ret.Categorystat61 = in["categorystat61"].(int)
		ret.Categorystat62 = in["categorystat62"].(int)
		ret.Categorystat63 = in["categorystat63"].(int)
		ret.Categorystat64 = in["categorystat64"].(int)
		ret.Categorystat65 = in["categorystat65"].(int)
		ret.Categorystat66 = in["categorystat66"].(int)
		ret.Categorystat67 = in["categorystat67"].(int)
		ret.Categorystat68 = in["categorystat68"].(int)
		ret.Categorystat69 = in["categorystat69"].(int)
		ret.Categorystat70 = in["categorystat70"].(int)
		ret.Categorystat71 = in["categorystat71"].(int)
		ret.Categorystat72 = in["categorystat72"].(int)
		ret.Categorystat73 = in["categorystat73"].(int)
		ret.Categorystat74 = in["categorystat74"].(int)
		ret.Categorystat75 = in["categorystat75"].(int)
		ret.Categorystat76 = in["categorystat76"].(int)
		ret.Categorystat77 = in["categorystat77"].(int)
		ret.Categorystat78 = in["categorystat78"].(int)
		ret.Categorystat79 = in["categorystat79"].(int)
		ret.Categorystat80 = in["categorystat80"].(int)
		ret.Categorystat81 = in["categorystat81"].(int)
		ret.Categorystat82 = in["categorystat82"].(int)
		ret.Categorystat83 = in["categorystat83"].(int)
		ret.Categorystat84 = in["categorystat84"].(int)
		ret.Categorystat85 = in["categorystat85"].(int)
		ret.Categorystat86 = in["categorystat86"].(int)
		ret.Categorystat87 = in["categorystat87"].(int)
		ret.Categorystat88 = in["categorystat88"].(int)
		ret.Categorystat89 = in["categorystat89"].(int)
		ret.Categorystat90 = in["categorystat90"].(int)
		ret.Categorystat91 = in["categorystat91"].(int)
		ret.Categorystat92 = in["categorystat92"].(int)
		ret.Categorystat93 = in["categorystat93"].(int)
		ret.Categorystat94 = in["categorystat94"].(int)
		ret.Categorystat95 = in["categorystat95"].(int)
		ret.Categorystat96 = in["categorystat96"].(int)
		ret.Categorystat97 = in["categorystat97"].(int)
		ret.Categorystat98 = in["categorystat98"].(int)
		ret.Categorystat99 = in["categorystat99"].(int)
		ret.Categorystat100 = in["categorystat100"].(int)
		ret.Categorystat101 = in["categorystat101"].(int)
		ret.Categorystat102 = in["categorystat102"].(int)
		ret.Categorystat103 = in["categorystat103"].(int)
		ret.Categorystat104 = in["categorystat104"].(int)
		ret.Categorystat105 = in["categorystat105"].(int)
		ret.Categorystat106 = in["categorystat106"].(int)
		ret.Categorystat107 = in["categorystat107"].(int)
		ret.Categorystat108 = in["categorystat108"].(int)
		ret.Categorystat109 = in["categorystat109"].(int)
		ret.Categorystat110 = in["categorystat110"].(int)
		ret.Categorystat111 = in["categorystat111"].(int)
		ret.Categorystat112 = in["categorystat112"].(int)
		ret.Categorystat113 = in["categorystat113"].(int)
		ret.Categorystat114 = in["categorystat114"].(int)
		ret.Categorystat115 = in["categorystat115"].(int)
		ret.Categorystat116 = in["categorystat116"].(int)
		ret.Categorystat117 = in["categorystat117"].(int)
		ret.Categorystat118 = in["categorystat118"].(int)
		ret.Categorystat119 = in["categorystat119"].(int)
		ret.Categorystat120 = in["categorystat120"].(int)
		ret.Categorystat121 = in["categorystat121"].(int)
		ret.Categorystat122 = in["categorystat122"].(int)
		ret.Categorystat123 = in["categorystat123"].(int)
		ret.Categorystat124 = in["categorystat124"].(int)
		ret.Categorystat125 = in["categorystat125"].(int)
		ret.Categorystat126 = in["categorystat126"].(int)
		ret.Categorystat127 = in["categorystat127"].(int)
		ret.Categorystat128 = in["categorystat128"].(int)
		ret.Categorystat129 = in["categorystat129"].(int)
		ret.Categorystat130 = in["categorystat130"].(int)
		ret.Categorystat131 = in["categorystat131"].(int)
		ret.Categorystat132 = in["categorystat132"].(int)
		ret.Categorystat133 = in["categorystat133"].(int)
		ret.Categorystat134 = in["categorystat134"].(int)
		ret.Categorystat135 = in["categorystat135"].(int)
		ret.Categorystat136 = in["categorystat136"].(int)
		ret.Categorystat137 = in["categorystat137"].(int)
		ret.Categorystat138 = in["categorystat138"].(int)
		ret.Categorystat139 = in["categorystat139"].(int)
		ret.Categorystat140 = in["categorystat140"].(int)
		ret.Categorystat141 = in["categorystat141"].(int)
		ret.Categorystat142 = in["categorystat142"].(int)
		ret.Categorystat143 = in["categorystat143"].(int)
		ret.Categorystat144 = in["categorystat144"].(int)
		ret.Categorystat145 = in["categorystat145"].(int)
		ret.Categorystat146 = in["categorystat146"].(int)
		ret.Categorystat147 = in["categorystat147"].(int)
		ret.Categorystat148 = in["categorystat148"].(int)
		ret.Categorystat149 = in["categorystat149"].(int)
		ret.Categorystat150 = in["categorystat150"].(int)
		ret.Categorystat151 = in["categorystat151"].(int)
		ret.Categorystat152 = in["categorystat152"].(int)
		ret.Categorystat153 = in["categorystat153"].(int)
		ret.Categorystat154 = in["categorystat154"].(int)
		ret.Categorystat155 = in["categorystat155"].(int)
		ret.Categorystat156 = in["categorystat156"].(int)
		ret.Categorystat157 = in["categorystat157"].(int)
		ret.Categorystat158 = in["categorystat158"].(int)
		ret.Categorystat159 = in["categorystat159"].(int)
		ret.Categorystat160 = in["categorystat160"].(int)
		ret.Categorystat161 = in["categorystat161"].(int)
		ret.Categorystat162 = in["categorystat162"].(int)
		ret.Categorystat163 = in["categorystat163"].(int)
		ret.Categorystat164 = in["categorystat164"].(int)
		ret.Categorystat165 = in["categorystat165"].(int)
		ret.Categorystat166 = in["categorystat166"].(int)
		ret.Categorystat167 = in["categorystat167"].(int)
		ret.Categorystat168 = in["categorystat168"].(int)
		ret.Categorystat169 = in["categorystat169"].(int)
		ret.Categorystat170 = in["categorystat170"].(int)
		ret.Categorystat171 = in["categorystat171"].(int)
		ret.Categorystat172 = in["categorystat172"].(int)
		ret.Categorystat173 = in["categorystat173"].(int)
		ret.Categorystat174 = in["categorystat174"].(int)
		ret.Categorystat175 = in["categorystat175"].(int)
		ret.Categorystat176 = in["categorystat176"].(int)
		ret.Categorystat177 = in["categorystat177"].(int)
		ret.Categorystat178 = in["categorystat178"].(int)
		ret.Categorystat179 = in["categorystat179"].(int)
		ret.Categorystat180 = in["categorystat180"].(int)
		ret.Categorystat181 = in["categorystat181"].(int)
		ret.Categorystat182 = in["categorystat182"].(int)
		ret.Categorystat183 = in["categorystat183"].(int)
		ret.Categorystat184 = in["categorystat184"].(int)
		ret.Categorystat185 = in["categorystat185"].(int)
		ret.Categorystat186 = in["categorystat186"].(int)
		ret.Categorystat187 = in["categorystat187"].(int)
		ret.Categorystat188 = in["categorystat188"].(int)
		ret.Categorystat189 = in["categorystat189"].(int)
		ret.Categorystat190 = in["categorystat190"].(int)
		ret.Categorystat191 = in["categorystat191"].(int)
		ret.Categorystat192 = in["categorystat192"].(int)
		ret.Categorystat193 = in["categorystat193"].(int)
		ret.Categorystat194 = in["categorystat194"].(int)
		ret.Categorystat195 = in["categorystat195"].(int)
		ret.Categorystat196 = in["categorystat196"].(int)
		ret.Categorystat197 = in["categorystat197"].(int)
		ret.Categorystat198 = in["categorystat198"].(int)
		ret.Categorystat199 = in["categorystat199"].(int)
		ret.Categorystat200 = in["categorystat200"].(int)
		ret.Categorystat201 = in["categorystat201"].(int)
		ret.Categorystat202 = in["categorystat202"].(int)
		ret.Categorystat203 = in["categorystat203"].(int)
		ret.Categorystat204 = in["categorystat204"].(int)
		ret.Categorystat205 = in["categorystat205"].(int)
		ret.Categorystat206 = in["categorystat206"].(int)
		ret.Categorystat207 = in["categorystat207"].(int)
		ret.Categorystat208 = in["categorystat208"].(int)
		ret.Categorystat209 = in["categorystat209"].(int)
		ret.Categorystat210 = in["categorystat210"].(int)
		ret.Categorystat211 = in["categorystat211"].(int)
		ret.Categorystat212 = in["categorystat212"].(int)
		ret.Categorystat213 = in["categorystat213"].(int)
		ret.Categorystat214 = in["categorystat214"].(int)
		ret.Categorystat215 = in["categorystat215"].(int)
		ret.Categorystat216 = in["categorystat216"].(int)
		ret.Categorystat217 = in["categorystat217"].(int)
		ret.Categorystat218 = in["categorystat218"].(int)
		ret.Categorystat219 = in["categorystat219"].(int)
		ret.Categorystat220 = in["categorystat220"].(int)
		ret.Categorystat221 = in["categorystat221"].(int)
		ret.Categorystat222 = in["categorystat222"].(int)
		ret.Categorystat223 = in["categorystat223"].(int)
		ret.Categorystat224 = in["categorystat224"].(int)
		ret.Categorystat225 = in["categorystat225"].(int)
		ret.Categorystat226 = in["categorystat226"].(int)
		ret.Categorystat227 = in["categorystat227"].(int)
		ret.Categorystat228 = in["categorystat228"].(int)
		ret.Categorystat229 = in["categorystat229"].(int)
		ret.Categorystat230 = in["categorystat230"].(int)
		ret.Categorystat231 = in["categorystat231"].(int)
		ret.Categorystat232 = in["categorystat232"].(int)
		ret.Categorystat233 = in["categorystat233"].(int)
		ret.Categorystat234 = in["categorystat234"].(int)
		ret.Categorystat235 = in["categorystat235"].(int)
		ret.Categorystat236 = in["categorystat236"].(int)
		ret.Categorystat237 = in["categorystat237"].(int)
		ret.Categorystat238 = in["categorystat238"].(int)
		ret.Categorystat239 = in["categorystat239"].(int)
		ret.Categorystat240 = in["categorystat240"].(int)
		ret.Categorystat241 = in["categorystat241"].(int)
		ret.Categorystat242 = in["categorystat242"].(int)
		ret.Categorystat243 = in["categorystat243"].(int)
		ret.Categorystat244 = in["categorystat244"].(int)
		ret.Categorystat245 = in["categorystat245"].(int)
		ret.Categorystat246 = in["categorystat246"].(int)
		ret.Categorystat247 = in["categorystat247"].(int)
		ret.Categorystat248 = in["categorystat248"].(int)
		ret.Categorystat249 = in["categorystat249"].(int)
		ret.Categorystat250 = in["categorystat250"].(int)
		ret.Categorystat251 = in["categorystat251"].(int)
		ret.Categorystat252 = in["categorystat252"].(int)
		ret.Categorystat253 = in["categorystat253"].(int)
		ret.Categorystat254 = in["categorystat254"].(int)
		ret.Categorystat255 = in["categorystat255"].(int)
		ret.Categorystat256 = in["categorystat256"].(int)
	}
	return ret
}

func dataToEndpointRuleSetTagStats(d *schema.ResourceData) edpt.RuleSetTagStats {
	var ret edpt.RuleSetTagStats

	ret.Stats = getObjectRuleSetTagStatsStats(d.Get("stats").([]interface{}))

	ret.Name = d.Get("name").(string)
	return ret
}
