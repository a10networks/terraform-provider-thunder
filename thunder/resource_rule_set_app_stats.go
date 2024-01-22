package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRuleSetAppStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_rule_set_app_stats`: Statistics for the object app\n\n__PLACEHOLDER__",
		ReadContext: resourceRuleSetAppStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"appstat1": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 1",
						},
						"appstat2": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 2",
						},
						"appstat3": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 3",
						},
						"appstat4": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 4",
						},
						"appstat5": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 5",
						},
						"appstat6": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 6",
						},
						"appstat7": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 7",
						},
						"appstat8": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 8",
						},
						"appstat9": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 9",
						},
						"appstat10": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 10",
						},
						"appstat11": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 11",
						},
						"appstat12": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 12",
						},
						"appstat13": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 13",
						},
						"appstat14": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 14",
						},
						"appstat15": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 15",
						},
						"appstat16": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 16",
						},
						"appstat17": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 17",
						},
						"appstat18": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 18",
						},
						"appstat19": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 19",
						},
						"appstat20": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 20",
						},
						"appstat21": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 21",
						},
						"appstat22": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 22",
						},
						"appstat23": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 23",
						},
						"appstat24": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 24",
						},
						"appstat25": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 25",
						},
						"appstat26": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 26",
						},
						"appstat27": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 27",
						},
						"appstat28": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 28",
						},
						"appstat29": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 29",
						},
						"appstat30": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 30",
						},
						"appstat31": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 31",
						},
						"appstat32": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 32",
						},
						"appstat33": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 33",
						},
						"appstat34": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 34",
						},
						"appstat35": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 35",
						},
						"appstat36": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 36",
						},
						"appstat37": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 37",
						},
						"appstat38": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 38",
						},
						"appstat39": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 39",
						},
						"appstat40": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 40",
						},
						"appstat41": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 41",
						},
						"appstat42": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 42",
						},
						"appstat43": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 43",
						},
						"appstat44": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 44",
						},
						"appstat45": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 45",
						},
						"appstat46": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 46",
						},
						"appstat47": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 47",
						},
						"appstat48": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 48",
						},
						"appstat49": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 49",
						},
						"appstat50": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 50",
						},
						"appstat51": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 51",
						},
						"appstat52": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 52",
						},
						"appstat53": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 53",
						},
						"appstat54": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 54",
						},
						"appstat55": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 55",
						},
						"appstat56": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 56",
						},
						"appstat57": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 57",
						},
						"appstat58": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 58",
						},
						"appstat59": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 59",
						},
						"appstat60": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 60",
						},
						"appstat61": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 61",
						},
						"appstat62": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 62",
						},
						"appstat63": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 63",
						},
						"appstat64": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 64",
						},
						"appstat65": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 65",
						},
						"appstat66": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 66",
						},
						"appstat67": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 67",
						},
						"appstat68": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 68",
						},
						"appstat69": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 69",
						},
						"appstat70": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 70",
						},
						"appstat71": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 71",
						},
						"appstat72": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 72",
						},
						"appstat73": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 73",
						},
						"appstat74": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 74",
						},
						"appstat75": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 75",
						},
						"appstat76": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 76",
						},
						"appstat77": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 77",
						},
						"appstat78": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 78",
						},
						"appstat79": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 79",
						},
						"appstat80": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 80",
						},
						"appstat81": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 81",
						},
						"appstat82": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 82",
						},
						"appstat83": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 83",
						},
						"appstat84": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 84",
						},
						"appstat85": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 85",
						},
						"appstat86": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 86",
						},
						"appstat87": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 87",
						},
						"appstat88": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 88",
						},
						"appstat89": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 89",
						},
						"appstat90": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 90",
						},
						"appstat91": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 91",
						},
						"appstat92": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 92",
						},
						"appstat93": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 93",
						},
						"appstat94": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 94",
						},
						"appstat95": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 95",
						},
						"appstat96": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 96",
						},
						"appstat97": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 97",
						},
						"appstat98": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 98",
						},
						"appstat99": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 99",
						},
						"appstat100": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 100",
						},
						"appstat101": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 101",
						},
						"appstat102": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 102",
						},
						"appstat103": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 103",
						},
						"appstat104": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 104",
						},
						"appstat105": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 105",
						},
						"appstat106": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 106",
						},
						"appstat107": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 107",
						},
						"appstat108": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 108",
						},
						"appstat109": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 109",
						},
						"appstat110": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 110",
						},
						"appstat111": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 111",
						},
						"appstat112": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 112",
						},
						"appstat113": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 113",
						},
						"appstat114": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 114",
						},
						"appstat115": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 115",
						},
						"appstat116": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 116",
						},
						"appstat117": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 117",
						},
						"appstat118": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 118",
						},
						"appstat119": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 119",
						},
						"appstat120": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 120",
						},
						"appstat121": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 121",
						},
						"appstat122": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 122",
						},
						"appstat123": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 123",
						},
						"appstat124": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 124",
						},
						"appstat125": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 125",
						},
						"appstat126": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 126",
						},
						"appstat127": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 127",
						},
						"appstat128": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 128",
						},
						"appstat129": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 129",
						},
						"appstat130": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 130",
						},
						"appstat131": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 131",
						},
						"appstat132": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 132",
						},
						"appstat133": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 133",
						},
						"appstat134": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 134",
						},
						"appstat135": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 135",
						},
						"appstat136": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 136",
						},
						"appstat137": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 137",
						},
						"appstat138": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 138",
						},
						"appstat139": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 139",
						},
						"appstat140": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 140",
						},
						"appstat141": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 141",
						},
						"appstat142": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 142",
						},
						"appstat143": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 143",
						},
						"appstat144": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 144",
						},
						"appstat145": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 145",
						},
						"appstat146": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 146",
						},
						"appstat147": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 147",
						},
						"appstat148": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 148",
						},
						"appstat149": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 149",
						},
						"appstat150": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 150",
						},
						"appstat151": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 151",
						},
						"appstat152": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 152",
						},
						"appstat153": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 153",
						},
						"appstat154": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 154",
						},
						"appstat155": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 155",
						},
						"appstat156": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 156",
						},
						"appstat157": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 157",
						},
						"appstat158": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 158",
						},
						"appstat159": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 159",
						},
						"appstat160": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 160",
						},
						"appstat161": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 161",
						},
						"appstat162": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 162",
						},
						"appstat163": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 163",
						},
						"appstat164": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 164",
						},
						"appstat165": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 165",
						},
						"appstat166": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 166",
						},
						"appstat167": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 167",
						},
						"appstat168": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 168",
						},
						"appstat169": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 169",
						},
						"appstat170": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 170",
						},
						"appstat171": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 171",
						},
						"appstat172": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 172",
						},
						"appstat173": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 173",
						},
						"appstat174": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 174",
						},
						"appstat175": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 175",
						},
						"appstat176": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 176",
						},
						"appstat177": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 177",
						},
						"appstat178": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 178",
						},
						"appstat179": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 179",
						},
						"appstat180": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 180",
						},
						"appstat181": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 181",
						},
						"appstat182": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 182",
						},
						"appstat183": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 183",
						},
						"appstat184": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 184",
						},
						"appstat185": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 185",
						},
						"appstat186": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 186",
						},
						"appstat187": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 187",
						},
						"appstat188": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 188",
						},
						"appstat189": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 189",
						},
						"appstat190": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 190",
						},
						"appstat191": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 191",
						},
						"appstat192": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 192",
						},
						"appstat193": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 193",
						},
						"appstat194": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 194",
						},
						"appstat195": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 195",
						},
						"appstat196": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 196",
						},
						"appstat197": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 197",
						},
						"appstat198": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 198",
						},
						"appstat199": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 199",
						},
						"appstat200": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 200",
						},
						"appstat201": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 201",
						},
						"appstat202": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 202",
						},
						"appstat203": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 203",
						},
						"appstat204": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 204",
						},
						"appstat205": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 205",
						},
						"appstat206": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 206",
						},
						"appstat207": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 207",
						},
						"appstat208": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 208",
						},
						"appstat209": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 209",
						},
						"appstat210": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 210",
						},
						"appstat211": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 211",
						},
						"appstat212": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 212",
						},
						"appstat213": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 213",
						},
						"appstat214": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 214",
						},
						"appstat215": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 215",
						},
						"appstat216": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 216",
						},
						"appstat217": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 217",
						},
						"appstat218": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 218",
						},
						"appstat219": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 219",
						},
						"appstat220": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 220",
						},
						"appstat221": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 221",
						},
						"appstat222": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 222",
						},
						"appstat223": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 223",
						},
						"appstat224": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 224",
						},
						"appstat225": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 225",
						},
						"appstat226": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 226",
						},
						"appstat227": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 227",
						},
						"appstat228": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 228",
						},
						"appstat229": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 229",
						},
						"appstat230": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 230",
						},
						"appstat231": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 231",
						},
						"appstat232": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 232",
						},
						"appstat233": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 233",
						},
						"appstat234": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 234",
						},
						"appstat235": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 235",
						},
						"appstat236": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 236",
						},
						"appstat237": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 237",
						},
						"appstat238": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 238",
						},
						"appstat239": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 239",
						},
						"appstat240": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 240",
						},
						"appstat241": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 241",
						},
						"appstat242": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 242",
						},
						"appstat243": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 243",
						},
						"appstat244": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 244",
						},
						"appstat245": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 245",
						},
						"appstat246": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 246",
						},
						"appstat247": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 247",
						},
						"appstat248": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 248",
						},
						"appstat249": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 249",
						},
						"appstat250": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 250",
						},
						"appstat251": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 251",
						},
						"appstat252": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 252",
						},
						"appstat253": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 253",
						},
						"appstat254": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 254",
						},
						"appstat255": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 255",
						},
						"appstat256": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 256",
						},
						"appstat257": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 257",
						},
						"appstat258": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 258",
						},
						"appstat259": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 259",
						},
						"appstat260": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 260",
						},
						"appstat261": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 261",
						},
						"appstat262": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 262",
						},
						"appstat263": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 263",
						},
						"appstat264": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 264",
						},
						"appstat265": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 265",
						},
						"appstat266": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 266",
						},
						"appstat267": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 267",
						},
						"appstat268": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 268",
						},
						"appstat269": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 269",
						},
						"appstat270": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 270",
						},
						"appstat271": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 271",
						},
						"appstat272": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 272",
						},
						"appstat273": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 273",
						},
						"appstat274": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 274",
						},
						"appstat275": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 275",
						},
						"appstat276": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 276",
						},
						"appstat277": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 277",
						},
						"appstat278": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 278",
						},
						"appstat279": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 279",
						},
						"appstat280": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 280",
						},
						"appstat281": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 281",
						},
						"appstat282": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 282",
						},
						"appstat283": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 283",
						},
						"appstat284": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 284",
						},
						"appstat285": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 285",
						},
						"appstat286": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 286",
						},
						"appstat287": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 287",
						},
						"appstat288": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 288",
						},
						"appstat289": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 289",
						},
						"appstat290": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 290",
						},
						"appstat291": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 291",
						},
						"appstat292": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 292",
						},
						"appstat293": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 293",
						},
						"appstat294": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 294",
						},
						"appstat295": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 295",
						},
						"appstat296": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 296",
						},
						"appstat297": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 297",
						},
						"appstat298": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 298",
						},
						"appstat299": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 299",
						},
						"appstat300": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 300",
						},
						"appstat301": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 301",
						},
						"appstat302": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 302",
						},
						"appstat303": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 303",
						},
						"appstat304": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 304",
						},
						"appstat305": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 305",
						},
						"appstat306": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 306",
						},
						"appstat307": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 307",
						},
						"appstat308": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 308",
						},
						"appstat309": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 309",
						},
						"appstat310": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 310",
						},
						"appstat311": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 311",
						},
						"appstat312": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 312",
						},
						"appstat313": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 313",
						},
						"appstat314": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 314",
						},
						"appstat315": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 315",
						},
						"appstat316": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 316",
						},
						"appstat317": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 317",
						},
						"appstat318": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 318",
						},
						"appstat319": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 319",
						},
						"appstat320": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 320",
						},
						"appstat321": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 321",
						},
						"appstat322": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 322",
						},
						"appstat323": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 323",
						},
						"appstat324": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 324",
						},
						"appstat325": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 325",
						},
						"appstat326": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 326",
						},
						"appstat327": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 327",
						},
						"appstat328": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 328",
						},
						"appstat329": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 329",
						},
						"appstat330": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 330",
						},
						"appstat331": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 331",
						},
						"appstat332": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 332",
						},
						"appstat333": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 333",
						},
						"appstat334": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 334",
						},
						"appstat335": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 335",
						},
						"appstat336": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 336",
						},
						"appstat337": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 337",
						},
						"appstat338": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 338",
						},
						"appstat339": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 339",
						},
						"appstat340": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 340",
						},
						"appstat341": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 341",
						},
						"appstat342": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 342",
						},
						"appstat343": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 343",
						},
						"appstat344": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 344",
						},
						"appstat345": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 345",
						},
						"appstat346": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 346",
						},
						"appstat347": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 347",
						},
						"appstat348": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 348",
						},
						"appstat349": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 349",
						},
						"appstat350": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 350",
						},
						"appstat351": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 351",
						},
						"appstat352": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 352",
						},
						"appstat353": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 353",
						},
						"appstat354": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 354",
						},
						"appstat355": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 355",
						},
						"appstat356": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 356",
						},
						"appstat357": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 357",
						},
						"appstat358": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 358",
						},
						"appstat359": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 359",
						},
						"appstat360": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 360",
						},
						"appstat361": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 361",
						},
						"appstat362": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 362",
						},
						"appstat363": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 363",
						},
						"appstat364": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 364",
						},
						"appstat365": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 365",
						},
						"appstat366": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 366",
						},
						"appstat367": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 367",
						},
						"appstat368": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 368",
						},
						"appstat369": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 369",
						},
						"appstat370": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 370",
						},
						"appstat371": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 371",
						},
						"appstat372": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 372",
						},
						"appstat373": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 373",
						},
						"appstat374": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 374",
						},
						"appstat375": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 375",
						},
						"appstat376": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 376",
						},
						"appstat377": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 377",
						},
						"appstat378": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 378",
						},
						"appstat379": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 379",
						},
						"appstat380": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 380",
						},
						"appstat381": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 381",
						},
						"appstat382": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 382",
						},
						"appstat383": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 383",
						},
						"appstat384": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 384",
						},
						"appstat385": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 385",
						},
						"appstat386": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 386",
						},
						"appstat387": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 387",
						},
						"appstat388": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 388",
						},
						"appstat389": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 389",
						},
						"appstat390": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 390",
						},
						"appstat391": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 391",
						},
						"appstat392": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 392",
						},
						"appstat393": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 393",
						},
						"appstat394": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 394",
						},
						"appstat395": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 395",
						},
						"appstat396": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 396",
						},
						"appstat397": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 397",
						},
						"appstat398": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 398",
						},
						"appstat399": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 399",
						},
						"appstat400": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 400",
						},
						"appstat401": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 401",
						},
						"appstat402": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 402",
						},
						"appstat403": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 403",
						},
						"appstat404": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 404",
						},
						"appstat405": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 405",
						},
						"appstat406": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 406",
						},
						"appstat407": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 407",
						},
						"appstat408": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 408",
						},
						"appstat409": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 409",
						},
						"appstat410": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 410",
						},
						"appstat411": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 411",
						},
						"appstat412": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 412",
						},
						"appstat413": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 413",
						},
						"appstat414": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 414",
						},
						"appstat415": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 415",
						},
						"appstat416": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 416",
						},
						"appstat417": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 417",
						},
						"appstat418": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 418",
						},
						"appstat419": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 419",
						},
						"appstat420": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 420",
						},
						"appstat421": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 421",
						},
						"appstat422": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 422",
						},
						"appstat423": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 423",
						},
						"appstat424": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 424",
						},
						"appstat425": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 425",
						},
						"appstat426": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 426",
						},
						"appstat427": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 427",
						},
						"appstat428": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 428",
						},
						"appstat429": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 429",
						},
						"appstat430": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 430",
						},
						"appstat431": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 431",
						},
						"appstat432": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 432",
						},
						"appstat433": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 433",
						},
						"appstat434": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 434",
						},
						"appstat435": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 435",
						},
						"appstat436": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 436",
						},
						"appstat437": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 437",
						},
						"appstat438": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 438",
						},
						"appstat439": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 439",
						},
						"appstat440": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 440",
						},
						"appstat441": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 441",
						},
						"appstat442": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 442",
						},
						"appstat443": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 443",
						},
						"appstat444": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 444",
						},
						"appstat445": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 445",
						},
						"appstat446": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 446",
						},
						"appstat447": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 447",
						},
						"appstat448": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 448",
						},
						"appstat449": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 449",
						},
						"appstat450": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 450",
						},
						"appstat451": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 451",
						},
						"appstat452": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 452",
						},
						"appstat453": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 453",
						},
						"appstat454": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 454",
						},
						"appstat455": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 455",
						},
						"appstat456": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 456",
						},
						"appstat457": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 457",
						},
						"appstat458": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 458",
						},
						"appstat459": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 459",
						},
						"appstat460": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 460",
						},
						"appstat461": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 461",
						},
						"appstat462": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 462",
						},
						"appstat463": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 463",
						},
						"appstat464": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 464",
						},
						"appstat465": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 465",
						},
						"appstat466": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 466",
						},
						"appstat467": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 467",
						},
						"appstat468": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 468",
						},
						"appstat469": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 469",
						},
						"appstat470": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 470",
						},
						"appstat471": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 471",
						},
						"appstat472": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 472",
						},
						"appstat473": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 473",
						},
						"appstat474": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 474",
						},
						"appstat475": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 475",
						},
						"appstat476": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 476",
						},
						"appstat477": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 477",
						},
						"appstat478": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 478",
						},
						"appstat479": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 479",
						},
						"appstat480": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 480",
						},
						"appstat481": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 481",
						},
						"appstat482": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 482",
						},
						"appstat483": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 483",
						},
						"appstat484": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 484",
						},
						"appstat485": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 485",
						},
						"appstat486": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 486",
						},
						"appstat487": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 487",
						},
						"appstat488": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 488",
						},
						"appstat489": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 489",
						},
						"appstat490": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 490",
						},
						"appstat491": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 491",
						},
						"appstat492": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 492",
						},
						"appstat493": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 493",
						},
						"appstat494": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 494",
						},
						"appstat495": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 495",
						},
						"appstat496": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 496",
						},
						"appstat497": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 497",
						},
						"appstat498": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 498",
						},
						"appstat499": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 499",
						},
						"appstat500": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 500",
						},
						"appstat501": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 501",
						},
						"appstat502": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 502",
						},
						"appstat503": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 503",
						},
						"appstat504": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 504",
						},
						"appstat505": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 505",
						},
						"appstat506": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 506",
						},
						"appstat507": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 507",
						},
						"appstat508": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 508",
						},
						"appstat509": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 509",
						},
						"appstat510": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 510",
						},
						"appstat511": {
							Type: schema.TypeInt, Optional: true, Description: "counter app stat 511",
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

func resourceRuleSetAppStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRuleSetAppStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRuleSetAppStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		RuleSetAppStatsStats := setObjectRuleSetAppStatsStats(res)
		d.Set("stats", RuleSetAppStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectRuleSetAppStatsStats(ret edpt.DataRuleSetAppStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"appstat1":   ret.DtRuleSetAppStats.Stats.Appstat1,
			"appstat2":   ret.DtRuleSetAppStats.Stats.Appstat2,
			"appstat3":   ret.DtRuleSetAppStats.Stats.Appstat3,
			"appstat4":   ret.DtRuleSetAppStats.Stats.Appstat4,
			"appstat5":   ret.DtRuleSetAppStats.Stats.Appstat5,
			"appstat6":   ret.DtRuleSetAppStats.Stats.Appstat6,
			"appstat7":   ret.DtRuleSetAppStats.Stats.Appstat7,
			"appstat8":   ret.DtRuleSetAppStats.Stats.Appstat8,
			"appstat9":   ret.DtRuleSetAppStats.Stats.Appstat9,
			"appstat10":  ret.DtRuleSetAppStats.Stats.Appstat10,
			"appstat11":  ret.DtRuleSetAppStats.Stats.Appstat11,
			"appstat12":  ret.DtRuleSetAppStats.Stats.Appstat12,
			"appstat13":  ret.DtRuleSetAppStats.Stats.Appstat13,
			"appstat14":  ret.DtRuleSetAppStats.Stats.Appstat14,
			"appstat15":  ret.DtRuleSetAppStats.Stats.Appstat15,
			"appstat16":  ret.DtRuleSetAppStats.Stats.Appstat16,
			"appstat17":  ret.DtRuleSetAppStats.Stats.Appstat17,
			"appstat18":  ret.DtRuleSetAppStats.Stats.Appstat18,
			"appstat19":  ret.DtRuleSetAppStats.Stats.Appstat19,
			"appstat20":  ret.DtRuleSetAppStats.Stats.Appstat20,
			"appstat21":  ret.DtRuleSetAppStats.Stats.Appstat21,
			"appstat22":  ret.DtRuleSetAppStats.Stats.Appstat22,
			"appstat23":  ret.DtRuleSetAppStats.Stats.Appstat23,
			"appstat24":  ret.DtRuleSetAppStats.Stats.Appstat24,
			"appstat25":  ret.DtRuleSetAppStats.Stats.Appstat25,
			"appstat26":  ret.DtRuleSetAppStats.Stats.Appstat26,
			"appstat27":  ret.DtRuleSetAppStats.Stats.Appstat27,
			"appstat28":  ret.DtRuleSetAppStats.Stats.Appstat28,
			"appstat29":  ret.DtRuleSetAppStats.Stats.Appstat29,
			"appstat30":  ret.DtRuleSetAppStats.Stats.Appstat30,
			"appstat31":  ret.DtRuleSetAppStats.Stats.Appstat31,
			"appstat32":  ret.DtRuleSetAppStats.Stats.Appstat32,
			"appstat33":  ret.DtRuleSetAppStats.Stats.Appstat33,
			"appstat34":  ret.DtRuleSetAppStats.Stats.Appstat34,
			"appstat35":  ret.DtRuleSetAppStats.Stats.Appstat35,
			"appstat36":  ret.DtRuleSetAppStats.Stats.Appstat36,
			"appstat37":  ret.DtRuleSetAppStats.Stats.Appstat37,
			"appstat38":  ret.DtRuleSetAppStats.Stats.Appstat38,
			"appstat39":  ret.DtRuleSetAppStats.Stats.Appstat39,
			"appstat40":  ret.DtRuleSetAppStats.Stats.Appstat40,
			"appstat41":  ret.DtRuleSetAppStats.Stats.Appstat41,
			"appstat42":  ret.DtRuleSetAppStats.Stats.Appstat42,
			"appstat43":  ret.DtRuleSetAppStats.Stats.Appstat43,
			"appstat44":  ret.DtRuleSetAppStats.Stats.Appstat44,
			"appstat45":  ret.DtRuleSetAppStats.Stats.Appstat45,
			"appstat46":  ret.DtRuleSetAppStats.Stats.Appstat46,
			"appstat47":  ret.DtRuleSetAppStats.Stats.Appstat47,
			"appstat48":  ret.DtRuleSetAppStats.Stats.Appstat48,
			"appstat49":  ret.DtRuleSetAppStats.Stats.Appstat49,
			"appstat50":  ret.DtRuleSetAppStats.Stats.Appstat50,
			"appstat51":  ret.DtRuleSetAppStats.Stats.Appstat51,
			"appstat52":  ret.DtRuleSetAppStats.Stats.Appstat52,
			"appstat53":  ret.DtRuleSetAppStats.Stats.Appstat53,
			"appstat54":  ret.DtRuleSetAppStats.Stats.Appstat54,
			"appstat55":  ret.DtRuleSetAppStats.Stats.Appstat55,
			"appstat56":  ret.DtRuleSetAppStats.Stats.Appstat56,
			"appstat57":  ret.DtRuleSetAppStats.Stats.Appstat57,
			"appstat58":  ret.DtRuleSetAppStats.Stats.Appstat58,
			"appstat59":  ret.DtRuleSetAppStats.Stats.Appstat59,
			"appstat60":  ret.DtRuleSetAppStats.Stats.Appstat60,
			"appstat61":  ret.DtRuleSetAppStats.Stats.Appstat61,
			"appstat62":  ret.DtRuleSetAppStats.Stats.Appstat62,
			"appstat63":  ret.DtRuleSetAppStats.Stats.Appstat63,
			"appstat64":  ret.DtRuleSetAppStats.Stats.Appstat64,
			"appstat65":  ret.DtRuleSetAppStats.Stats.Appstat65,
			"appstat66":  ret.DtRuleSetAppStats.Stats.Appstat66,
			"appstat67":  ret.DtRuleSetAppStats.Stats.Appstat67,
			"appstat68":  ret.DtRuleSetAppStats.Stats.Appstat68,
			"appstat69":  ret.DtRuleSetAppStats.Stats.Appstat69,
			"appstat70":  ret.DtRuleSetAppStats.Stats.Appstat70,
			"appstat71":  ret.DtRuleSetAppStats.Stats.Appstat71,
			"appstat72":  ret.DtRuleSetAppStats.Stats.Appstat72,
			"appstat73":  ret.DtRuleSetAppStats.Stats.Appstat73,
			"appstat74":  ret.DtRuleSetAppStats.Stats.Appstat74,
			"appstat75":  ret.DtRuleSetAppStats.Stats.Appstat75,
			"appstat76":  ret.DtRuleSetAppStats.Stats.Appstat76,
			"appstat77":  ret.DtRuleSetAppStats.Stats.Appstat77,
			"appstat78":  ret.DtRuleSetAppStats.Stats.Appstat78,
			"appstat79":  ret.DtRuleSetAppStats.Stats.Appstat79,
			"appstat80":  ret.DtRuleSetAppStats.Stats.Appstat80,
			"appstat81":  ret.DtRuleSetAppStats.Stats.Appstat81,
			"appstat82":  ret.DtRuleSetAppStats.Stats.Appstat82,
			"appstat83":  ret.DtRuleSetAppStats.Stats.Appstat83,
			"appstat84":  ret.DtRuleSetAppStats.Stats.Appstat84,
			"appstat85":  ret.DtRuleSetAppStats.Stats.Appstat85,
			"appstat86":  ret.DtRuleSetAppStats.Stats.Appstat86,
			"appstat87":  ret.DtRuleSetAppStats.Stats.Appstat87,
			"appstat88":  ret.DtRuleSetAppStats.Stats.Appstat88,
			"appstat89":  ret.DtRuleSetAppStats.Stats.Appstat89,
			"appstat90":  ret.DtRuleSetAppStats.Stats.Appstat90,
			"appstat91":  ret.DtRuleSetAppStats.Stats.Appstat91,
			"appstat92":  ret.DtRuleSetAppStats.Stats.Appstat92,
			"appstat93":  ret.DtRuleSetAppStats.Stats.Appstat93,
			"appstat94":  ret.DtRuleSetAppStats.Stats.Appstat94,
			"appstat95":  ret.DtRuleSetAppStats.Stats.Appstat95,
			"appstat96":  ret.DtRuleSetAppStats.Stats.Appstat96,
			"appstat97":  ret.DtRuleSetAppStats.Stats.Appstat97,
			"appstat98":  ret.DtRuleSetAppStats.Stats.Appstat98,
			"appstat99":  ret.DtRuleSetAppStats.Stats.Appstat99,
			"appstat100": ret.DtRuleSetAppStats.Stats.Appstat100,
			"appstat101": ret.DtRuleSetAppStats.Stats.Appstat101,
			"appstat102": ret.DtRuleSetAppStats.Stats.Appstat102,
			"appstat103": ret.DtRuleSetAppStats.Stats.Appstat103,
			"appstat104": ret.DtRuleSetAppStats.Stats.Appstat104,
			"appstat105": ret.DtRuleSetAppStats.Stats.Appstat105,
			"appstat106": ret.DtRuleSetAppStats.Stats.Appstat106,
			"appstat107": ret.DtRuleSetAppStats.Stats.Appstat107,
			"appstat108": ret.DtRuleSetAppStats.Stats.Appstat108,
			"appstat109": ret.DtRuleSetAppStats.Stats.Appstat109,
			"appstat110": ret.DtRuleSetAppStats.Stats.Appstat110,
			"appstat111": ret.DtRuleSetAppStats.Stats.Appstat111,
			"appstat112": ret.DtRuleSetAppStats.Stats.Appstat112,
			"appstat113": ret.DtRuleSetAppStats.Stats.Appstat113,
			"appstat114": ret.DtRuleSetAppStats.Stats.Appstat114,
			"appstat115": ret.DtRuleSetAppStats.Stats.Appstat115,
			"appstat116": ret.DtRuleSetAppStats.Stats.Appstat116,
			"appstat117": ret.DtRuleSetAppStats.Stats.Appstat117,
			"appstat118": ret.DtRuleSetAppStats.Stats.Appstat118,
			"appstat119": ret.DtRuleSetAppStats.Stats.Appstat119,
			"appstat120": ret.DtRuleSetAppStats.Stats.Appstat120,
			"appstat121": ret.DtRuleSetAppStats.Stats.Appstat121,
			"appstat122": ret.DtRuleSetAppStats.Stats.Appstat122,
			"appstat123": ret.DtRuleSetAppStats.Stats.Appstat123,
			"appstat124": ret.DtRuleSetAppStats.Stats.Appstat124,
			"appstat125": ret.DtRuleSetAppStats.Stats.Appstat125,
			"appstat126": ret.DtRuleSetAppStats.Stats.Appstat126,
			"appstat127": ret.DtRuleSetAppStats.Stats.Appstat127,
			"appstat128": ret.DtRuleSetAppStats.Stats.Appstat128,
			"appstat129": ret.DtRuleSetAppStats.Stats.Appstat129,
			"appstat130": ret.DtRuleSetAppStats.Stats.Appstat130,
			"appstat131": ret.DtRuleSetAppStats.Stats.Appstat131,
			"appstat132": ret.DtRuleSetAppStats.Stats.Appstat132,
			"appstat133": ret.DtRuleSetAppStats.Stats.Appstat133,
			"appstat134": ret.DtRuleSetAppStats.Stats.Appstat134,
			"appstat135": ret.DtRuleSetAppStats.Stats.Appstat135,
			"appstat136": ret.DtRuleSetAppStats.Stats.Appstat136,
			"appstat137": ret.DtRuleSetAppStats.Stats.Appstat137,
			"appstat138": ret.DtRuleSetAppStats.Stats.Appstat138,
			"appstat139": ret.DtRuleSetAppStats.Stats.Appstat139,
			"appstat140": ret.DtRuleSetAppStats.Stats.Appstat140,
			"appstat141": ret.DtRuleSetAppStats.Stats.Appstat141,
			"appstat142": ret.DtRuleSetAppStats.Stats.Appstat142,
			"appstat143": ret.DtRuleSetAppStats.Stats.Appstat143,
			"appstat144": ret.DtRuleSetAppStats.Stats.Appstat144,
			"appstat145": ret.DtRuleSetAppStats.Stats.Appstat145,
			"appstat146": ret.DtRuleSetAppStats.Stats.Appstat146,
			"appstat147": ret.DtRuleSetAppStats.Stats.Appstat147,
			"appstat148": ret.DtRuleSetAppStats.Stats.Appstat148,
			"appstat149": ret.DtRuleSetAppStats.Stats.Appstat149,
			"appstat150": ret.DtRuleSetAppStats.Stats.Appstat150,
			"appstat151": ret.DtRuleSetAppStats.Stats.Appstat151,
			"appstat152": ret.DtRuleSetAppStats.Stats.Appstat152,
			"appstat153": ret.DtRuleSetAppStats.Stats.Appstat153,
			"appstat154": ret.DtRuleSetAppStats.Stats.Appstat154,
			"appstat155": ret.DtRuleSetAppStats.Stats.Appstat155,
			"appstat156": ret.DtRuleSetAppStats.Stats.Appstat156,
			"appstat157": ret.DtRuleSetAppStats.Stats.Appstat157,
			"appstat158": ret.DtRuleSetAppStats.Stats.Appstat158,
			"appstat159": ret.DtRuleSetAppStats.Stats.Appstat159,
			"appstat160": ret.DtRuleSetAppStats.Stats.Appstat160,
			"appstat161": ret.DtRuleSetAppStats.Stats.Appstat161,
			"appstat162": ret.DtRuleSetAppStats.Stats.Appstat162,
			"appstat163": ret.DtRuleSetAppStats.Stats.Appstat163,
			"appstat164": ret.DtRuleSetAppStats.Stats.Appstat164,
			"appstat165": ret.DtRuleSetAppStats.Stats.Appstat165,
			"appstat166": ret.DtRuleSetAppStats.Stats.Appstat166,
			"appstat167": ret.DtRuleSetAppStats.Stats.Appstat167,
			"appstat168": ret.DtRuleSetAppStats.Stats.Appstat168,
			"appstat169": ret.DtRuleSetAppStats.Stats.Appstat169,
			"appstat170": ret.DtRuleSetAppStats.Stats.Appstat170,
			"appstat171": ret.DtRuleSetAppStats.Stats.Appstat171,
			"appstat172": ret.DtRuleSetAppStats.Stats.Appstat172,
			"appstat173": ret.DtRuleSetAppStats.Stats.Appstat173,
			"appstat174": ret.DtRuleSetAppStats.Stats.Appstat174,
			"appstat175": ret.DtRuleSetAppStats.Stats.Appstat175,
			"appstat176": ret.DtRuleSetAppStats.Stats.Appstat176,
			"appstat177": ret.DtRuleSetAppStats.Stats.Appstat177,
			"appstat178": ret.DtRuleSetAppStats.Stats.Appstat178,
			"appstat179": ret.DtRuleSetAppStats.Stats.Appstat179,
			"appstat180": ret.DtRuleSetAppStats.Stats.Appstat180,
			"appstat181": ret.DtRuleSetAppStats.Stats.Appstat181,
			"appstat182": ret.DtRuleSetAppStats.Stats.Appstat182,
			"appstat183": ret.DtRuleSetAppStats.Stats.Appstat183,
			"appstat184": ret.DtRuleSetAppStats.Stats.Appstat184,
			"appstat185": ret.DtRuleSetAppStats.Stats.Appstat185,
			"appstat186": ret.DtRuleSetAppStats.Stats.Appstat186,
			"appstat187": ret.DtRuleSetAppStats.Stats.Appstat187,
			"appstat188": ret.DtRuleSetAppStats.Stats.Appstat188,
			"appstat189": ret.DtRuleSetAppStats.Stats.Appstat189,
			"appstat190": ret.DtRuleSetAppStats.Stats.Appstat190,
			"appstat191": ret.DtRuleSetAppStats.Stats.Appstat191,
			"appstat192": ret.DtRuleSetAppStats.Stats.Appstat192,
			"appstat193": ret.DtRuleSetAppStats.Stats.Appstat193,
			"appstat194": ret.DtRuleSetAppStats.Stats.Appstat194,
			"appstat195": ret.DtRuleSetAppStats.Stats.Appstat195,
			"appstat196": ret.DtRuleSetAppStats.Stats.Appstat196,
			"appstat197": ret.DtRuleSetAppStats.Stats.Appstat197,
			"appstat198": ret.DtRuleSetAppStats.Stats.Appstat198,
			"appstat199": ret.DtRuleSetAppStats.Stats.Appstat199,
			"appstat200": ret.DtRuleSetAppStats.Stats.Appstat200,
			"appstat201": ret.DtRuleSetAppStats.Stats.Appstat201,
			"appstat202": ret.DtRuleSetAppStats.Stats.Appstat202,
			"appstat203": ret.DtRuleSetAppStats.Stats.Appstat203,
			"appstat204": ret.DtRuleSetAppStats.Stats.Appstat204,
			"appstat205": ret.DtRuleSetAppStats.Stats.Appstat205,
			"appstat206": ret.DtRuleSetAppStats.Stats.Appstat206,
			"appstat207": ret.DtRuleSetAppStats.Stats.Appstat207,
			"appstat208": ret.DtRuleSetAppStats.Stats.Appstat208,
			"appstat209": ret.DtRuleSetAppStats.Stats.Appstat209,
			"appstat210": ret.DtRuleSetAppStats.Stats.Appstat210,
			"appstat211": ret.DtRuleSetAppStats.Stats.Appstat211,
			"appstat212": ret.DtRuleSetAppStats.Stats.Appstat212,
			"appstat213": ret.DtRuleSetAppStats.Stats.Appstat213,
			"appstat214": ret.DtRuleSetAppStats.Stats.Appstat214,
			"appstat215": ret.DtRuleSetAppStats.Stats.Appstat215,
			"appstat216": ret.DtRuleSetAppStats.Stats.Appstat216,
			"appstat217": ret.DtRuleSetAppStats.Stats.Appstat217,
			"appstat218": ret.DtRuleSetAppStats.Stats.Appstat218,
			"appstat219": ret.DtRuleSetAppStats.Stats.Appstat219,
			"appstat220": ret.DtRuleSetAppStats.Stats.Appstat220,
			"appstat221": ret.DtRuleSetAppStats.Stats.Appstat221,
			"appstat222": ret.DtRuleSetAppStats.Stats.Appstat222,
			"appstat223": ret.DtRuleSetAppStats.Stats.Appstat223,
			"appstat224": ret.DtRuleSetAppStats.Stats.Appstat224,
			"appstat225": ret.DtRuleSetAppStats.Stats.Appstat225,
			"appstat226": ret.DtRuleSetAppStats.Stats.Appstat226,
			"appstat227": ret.DtRuleSetAppStats.Stats.Appstat227,
			"appstat228": ret.DtRuleSetAppStats.Stats.Appstat228,
			"appstat229": ret.DtRuleSetAppStats.Stats.Appstat229,
			"appstat230": ret.DtRuleSetAppStats.Stats.Appstat230,
			"appstat231": ret.DtRuleSetAppStats.Stats.Appstat231,
			"appstat232": ret.DtRuleSetAppStats.Stats.Appstat232,
			"appstat233": ret.DtRuleSetAppStats.Stats.Appstat233,
			"appstat234": ret.DtRuleSetAppStats.Stats.Appstat234,
			"appstat235": ret.DtRuleSetAppStats.Stats.Appstat235,
			"appstat236": ret.DtRuleSetAppStats.Stats.Appstat236,
			"appstat237": ret.DtRuleSetAppStats.Stats.Appstat237,
			"appstat238": ret.DtRuleSetAppStats.Stats.Appstat238,
			"appstat239": ret.DtRuleSetAppStats.Stats.Appstat239,
			"appstat240": ret.DtRuleSetAppStats.Stats.Appstat240,
			"appstat241": ret.DtRuleSetAppStats.Stats.Appstat241,
			"appstat242": ret.DtRuleSetAppStats.Stats.Appstat242,
			"appstat243": ret.DtRuleSetAppStats.Stats.Appstat243,
			"appstat244": ret.DtRuleSetAppStats.Stats.Appstat244,
			"appstat245": ret.DtRuleSetAppStats.Stats.Appstat245,
			"appstat246": ret.DtRuleSetAppStats.Stats.Appstat246,
			"appstat247": ret.DtRuleSetAppStats.Stats.Appstat247,
			"appstat248": ret.DtRuleSetAppStats.Stats.Appstat248,
			"appstat249": ret.DtRuleSetAppStats.Stats.Appstat249,
			"appstat250": ret.DtRuleSetAppStats.Stats.Appstat250,
			"appstat251": ret.DtRuleSetAppStats.Stats.Appstat251,
			"appstat252": ret.DtRuleSetAppStats.Stats.Appstat252,
			"appstat253": ret.DtRuleSetAppStats.Stats.Appstat253,
			"appstat254": ret.DtRuleSetAppStats.Stats.Appstat254,
			"appstat255": ret.DtRuleSetAppStats.Stats.Appstat255,
			"appstat256": ret.DtRuleSetAppStats.Stats.Appstat256,
			"appstat257": ret.DtRuleSetAppStats.Stats.Appstat257,
			"appstat258": ret.DtRuleSetAppStats.Stats.Appstat258,
			"appstat259": ret.DtRuleSetAppStats.Stats.Appstat259,
			"appstat260": ret.DtRuleSetAppStats.Stats.Appstat260,
			"appstat261": ret.DtRuleSetAppStats.Stats.Appstat261,
			"appstat262": ret.DtRuleSetAppStats.Stats.Appstat262,
			"appstat263": ret.DtRuleSetAppStats.Stats.Appstat263,
			"appstat264": ret.DtRuleSetAppStats.Stats.Appstat264,
			"appstat265": ret.DtRuleSetAppStats.Stats.Appstat265,
			"appstat266": ret.DtRuleSetAppStats.Stats.Appstat266,
			"appstat267": ret.DtRuleSetAppStats.Stats.Appstat267,
			"appstat268": ret.DtRuleSetAppStats.Stats.Appstat268,
			"appstat269": ret.DtRuleSetAppStats.Stats.Appstat269,
			"appstat270": ret.DtRuleSetAppStats.Stats.Appstat270,
			"appstat271": ret.DtRuleSetAppStats.Stats.Appstat271,
			"appstat272": ret.DtRuleSetAppStats.Stats.Appstat272,
			"appstat273": ret.DtRuleSetAppStats.Stats.Appstat273,
			"appstat274": ret.DtRuleSetAppStats.Stats.Appstat274,
			"appstat275": ret.DtRuleSetAppStats.Stats.Appstat275,
			"appstat276": ret.DtRuleSetAppStats.Stats.Appstat276,
			"appstat277": ret.DtRuleSetAppStats.Stats.Appstat277,
			"appstat278": ret.DtRuleSetAppStats.Stats.Appstat278,
			"appstat279": ret.DtRuleSetAppStats.Stats.Appstat279,
			"appstat280": ret.DtRuleSetAppStats.Stats.Appstat280,
			"appstat281": ret.DtRuleSetAppStats.Stats.Appstat281,
			"appstat282": ret.DtRuleSetAppStats.Stats.Appstat282,
			"appstat283": ret.DtRuleSetAppStats.Stats.Appstat283,
			"appstat284": ret.DtRuleSetAppStats.Stats.Appstat284,
			"appstat285": ret.DtRuleSetAppStats.Stats.Appstat285,
			"appstat286": ret.DtRuleSetAppStats.Stats.Appstat286,
			"appstat287": ret.DtRuleSetAppStats.Stats.Appstat287,
			"appstat288": ret.DtRuleSetAppStats.Stats.Appstat288,
			"appstat289": ret.DtRuleSetAppStats.Stats.Appstat289,
			"appstat290": ret.DtRuleSetAppStats.Stats.Appstat290,
			"appstat291": ret.DtRuleSetAppStats.Stats.Appstat291,
			"appstat292": ret.DtRuleSetAppStats.Stats.Appstat292,
			"appstat293": ret.DtRuleSetAppStats.Stats.Appstat293,
			"appstat294": ret.DtRuleSetAppStats.Stats.Appstat294,
			"appstat295": ret.DtRuleSetAppStats.Stats.Appstat295,
			"appstat296": ret.DtRuleSetAppStats.Stats.Appstat296,
			"appstat297": ret.DtRuleSetAppStats.Stats.Appstat297,
			"appstat298": ret.DtRuleSetAppStats.Stats.Appstat298,
			"appstat299": ret.DtRuleSetAppStats.Stats.Appstat299,
			"appstat300": ret.DtRuleSetAppStats.Stats.Appstat300,
			"appstat301": ret.DtRuleSetAppStats.Stats.Appstat301,
			"appstat302": ret.DtRuleSetAppStats.Stats.Appstat302,
			"appstat303": ret.DtRuleSetAppStats.Stats.Appstat303,
			"appstat304": ret.DtRuleSetAppStats.Stats.Appstat304,
			"appstat305": ret.DtRuleSetAppStats.Stats.Appstat305,
			"appstat306": ret.DtRuleSetAppStats.Stats.Appstat306,
			"appstat307": ret.DtRuleSetAppStats.Stats.Appstat307,
			"appstat308": ret.DtRuleSetAppStats.Stats.Appstat308,
			"appstat309": ret.DtRuleSetAppStats.Stats.Appstat309,
			"appstat310": ret.DtRuleSetAppStats.Stats.Appstat310,
			"appstat311": ret.DtRuleSetAppStats.Stats.Appstat311,
			"appstat312": ret.DtRuleSetAppStats.Stats.Appstat312,
			"appstat313": ret.DtRuleSetAppStats.Stats.Appstat313,
			"appstat314": ret.DtRuleSetAppStats.Stats.Appstat314,
			"appstat315": ret.DtRuleSetAppStats.Stats.Appstat315,
			"appstat316": ret.DtRuleSetAppStats.Stats.Appstat316,
			"appstat317": ret.DtRuleSetAppStats.Stats.Appstat317,
			"appstat318": ret.DtRuleSetAppStats.Stats.Appstat318,
			"appstat319": ret.DtRuleSetAppStats.Stats.Appstat319,
			"appstat320": ret.DtRuleSetAppStats.Stats.Appstat320,
			"appstat321": ret.DtRuleSetAppStats.Stats.Appstat321,
			"appstat322": ret.DtRuleSetAppStats.Stats.Appstat322,
			"appstat323": ret.DtRuleSetAppStats.Stats.Appstat323,
			"appstat324": ret.DtRuleSetAppStats.Stats.Appstat324,
			"appstat325": ret.DtRuleSetAppStats.Stats.Appstat325,
			"appstat326": ret.DtRuleSetAppStats.Stats.Appstat326,
			"appstat327": ret.DtRuleSetAppStats.Stats.Appstat327,
			"appstat328": ret.DtRuleSetAppStats.Stats.Appstat328,
			"appstat329": ret.DtRuleSetAppStats.Stats.Appstat329,
			"appstat330": ret.DtRuleSetAppStats.Stats.Appstat330,
			"appstat331": ret.DtRuleSetAppStats.Stats.Appstat331,
			"appstat332": ret.DtRuleSetAppStats.Stats.Appstat332,
			"appstat333": ret.DtRuleSetAppStats.Stats.Appstat333,
			"appstat334": ret.DtRuleSetAppStats.Stats.Appstat334,
			"appstat335": ret.DtRuleSetAppStats.Stats.Appstat335,
			"appstat336": ret.DtRuleSetAppStats.Stats.Appstat336,
			"appstat337": ret.DtRuleSetAppStats.Stats.Appstat337,
			"appstat338": ret.DtRuleSetAppStats.Stats.Appstat338,
			"appstat339": ret.DtRuleSetAppStats.Stats.Appstat339,
			"appstat340": ret.DtRuleSetAppStats.Stats.Appstat340,
			"appstat341": ret.DtRuleSetAppStats.Stats.Appstat341,
			"appstat342": ret.DtRuleSetAppStats.Stats.Appstat342,
			"appstat343": ret.DtRuleSetAppStats.Stats.Appstat343,
			"appstat344": ret.DtRuleSetAppStats.Stats.Appstat344,
			"appstat345": ret.DtRuleSetAppStats.Stats.Appstat345,
			"appstat346": ret.DtRuleSetAppStats.Stats.Appstat346,
			"appstat347": ret.DtRuleSetAppStats.Stats.Appstat347,
			"appstat348": ret.DtRuleSetAppStats.Stats.Appstat348,
			"appstat349": ret.DtRuleSetAppStats.Stats.Appstat349,
			"appstat350": ret.DtRuleSetAppStats.Stats.Appstat350,
			"appstat351": ret.DtRuleSetAppStats.Stats.Appstat351,
			"appstat352": ret.DtRuleSetAppStats.Stats.Appstat352,
			"appstat353": ret.DtRuleSetAppStats.Stats.Appstat353,
			"appstat354": ret.DtRuleSetAppStats.Stats.Appstat354,
			"appstat355": ret.DtRuleSetAppStats.Stats.Appstat355,
			"appstat356": ret.DtRuleSetAppStats.Stats.Appstat356,
			"appstat357": ret.DtRuleSetAppStats.Stats.Appstat357,
			"appstat358": ret.DtRuleSetAppStats.Stats.Appstat358,
			"appstat359": ret.DtRuleSetAppStats.Stats.Appstat359,
			"appstat360": ret.DtRuleSetAppStats.Stats.Appstat360,
			"appstat361": ret.DtRuleSetAppStats.Stats.Appstat361,
			"appstat362": ret.DtRuleSetAppStats.Stats.Appstat362,
			"appstat363": ret.DtRuleSetAppStats.Stats.Appstat363,
			"appstat364": ret.DtRuleSetAppStats.Stats.Appstat364,
			"appstat365": ret.DtRuleSetAppStats.Stats.Appstat365,
			"appstat366": ret.DtRuleSetAppStats.Stats.Appstat366,
			"appstat367": ret.DtRuleSetAppStats.Stats.Appstat367,
			"appstat368": ret.DtRuleSetAppStats.Stats.Appstat368,
			"appstat369": ret.DtRuleSetAppStats.Stats.Appstat369,
			"appstat370": ret.DtRuleSetAppStats.Stats.Appstat370,
			"appstat371": ret.DtRuleSetAppStats.Stats.Appstat371,
			"appstat372": ret.DtRuleSetAppStats.Stats.Appstat372,
			"appstat373": ret.DtRuleSetAppStats.Stats.Appstat373,
			"appstat374": ret.DtRuleSetAppStats.Stats.Appstat374,
			"appstat375": ret.DtRuleSetAppStats.Stats.Appstat375,
			"appstat376": ret.DtRuleSetAppStats.Stats.Appstat376,
			"appstat377": ret.DtRuleSetAppStats.Stats.Appstat377,
			"appstat378": ret.DtRuleSetAppStats.Stats.Appstat378,
			"appstat379": ret.DtRuleSetAppStats.Stats.Appstat379,
			"appstat380": ret.DtRuleSetAppStats.Stats.Appstat380,
			"appstat381": ret.DtRuleSetAppStats.Stats.Appstat381,
			"appstat382": ret.DtRuleSetAppStats.Stats.Appstat382,
			"appstat383": ret.DtRuleSetAppStats.Stats.Appstat383,
			"appstat384": ret.DtRuleSetAppStats.Stats.Appstat384,
			"appstat385": ret.DtRuleSetAppStats.Stats.Appstat385,
			"appstat386": ret.DtRuleSetAppStats.Stats.Appstat386,
			"appstat387": ret.DtRuleSetAppStats.Stats.Appstat387,
			"appstat388": ret.DtRuleSetAppStats.Stats.Appstat388,
			"appstat389": ret.DtRuleSetAppStats.Stats.Appstat389,
			"appstat390": ret.DtRuleSetAppStats.Stats.Appstat390,
			"appstat391": ret.DtRuleSetAppStats.Stats.Appstat391,
			"appstat392": ret.DtRuleSetAppStats.Stats.Appstat392,
			"appstat393": ret.DtRuleSetAppStats.Stats.Appstat393,
			"appstat394": ret.DtRuleSetAppStats.Stats.Appstat394,
			"appstat395": ret.DtRuleSetAppStats.Stats.Appstat395,
			"appstat396": ret.DtRuleSetAppStats.Stats.Appstat396,
			"appstat397": ret.DtRuleSetAppStats.Stats.Appstat397,
			"appstat398": ret.DtRuleSetAppStats.Stats.Appstat398,
			"appstat399": ret.DtRuleSetAppStats.Stats.Appstat399,
			"appstat400": ret.DtRuleSetAppStats.Stats.Appstat400,
			"appstat401": ret.DtRuleSetAppStats.Stats.Appstat401,
			"appstat402": ret.DtRuleSetAppStats.Stats.Appstat402,
			"appstat403": ret.DtRuleSetAppStats.Stats.Appstat403,
			"appstat404": ret.DtRuleSetAppStats.Stats.Appstat404,
			"appstat405": ret.DtRuleSetAppStats.Stats.Appstat405,
			"appstat406": ret.DtRuleSetAppStats.Stats.Appstat406,
			"appstat407": ret.DtRuleSetAppStats.Stats.Appstat407,
			"appstat408": ret.DtRuleSetAppStats.Stats.Appstat408,
			"appstat409": ret.DtRuleSetAppStats.Stats.Appstat409,
			"appstat410": ret.DtRuleSetAppStats.Stats.Appstat410,
			"appstat411": ret.DtRuleSetAppStats.Stats.Appstat411,
			"appstat412": ret.DtRuleSetAppStats.Stats.Appstat412,
			"appstat413": ret.DtRuleSetAppStats.Stats.Appstat413,
			"appstat414": ret.DtRuleSetAppStats.Stats.Appstat414,
			"appstat415": ret.DtRuleSetAppStats.Stats.Appstat415,
			"appstat416": ret.DtRuleSetAppStats.Stats.Appstat416,
			"appstat417": ret.DtRuleSetAppStats.Stats.Appstat417,
			"appstat418": ret.DtRuleSetAppStats.Stats.Appstat418,
			"appstat419": ret.DtRuleSetAppStats.Stats.Appstat419,
			"appstat420": ret.DtRuleSetAppStats.Stats.Appstat420,
			"appstat421": ret.DtRuleSetAppStats.Stats.Appstat421,
			"appstat422": ret.DtRuleSetAppStats.Stats.Appstat422,
			"appstat423": ret.DtRuleSetAppStats.Stats.Appstat423,
			"appstat424": ret.DtRuleSetAppStats.Stats.Appstat424,
			"appstat425": ret.DtRuleSetAppStats.Stats.Appstat425,
			"appstat426": ret.DtRuleSetAppStats.Stats.Appstat426,
			"appstat427": ret.DtRuleSetAppStats.Stats.Appstat427,
			"appstat428": ret.DtRuleSetAppStats.Stats.Appstat428,
			"appstat429": ret.DtRuleSetAppStats.Stats.Appstat429,
			"appstat430": ret.DtRuleSetAppStats.Stats.Appstat430,
			"appstat431": ret.DtRuleSetAppStats.Stats.Appstat431,
			"appstat432": ret.DtRuleSetAppStats.Stats.Appstat432,
			"appstat433": ret.DtRuleSetAppStats.Stats.Appstat433,
			"appstat434": ret.DtRuleSetAppStats.Stats.Appstat434,
			"appstat435": ret.DtRuleSetAppStats.Stats.Appstat435,
			"appstat436": ret.DtRuleSetAppStats.Stats.Appstat436,
			"appstat437": ret.DtRuleSetAppStats.Stats.Appstat437,
			"appstat438": ret.DtRuleSetAppStats.Stats.Appstat438,
			"appstat439": ret.DtRuleSetAppStats.Stats.Appstat439,
			"appstat440": ret.DtRuleSetAppStats.Stats.Appstat440,
			"appstat441": ret.DtRuleSetAppStats.Stats.Appstat441,
			"appstat442": ret.DtRuleSetAppStats.Stats.Appstat442,
			"appstat443": ret.DtRuleSetAppStats.Stats.Appstat443,
			"appstat444": ret.DtRuleSetAppStats.Stats.Appstat444,
			"appstat445": ret.DtRuleSetAppStats.Stats.Appstat445,
			"appstat446": ret.DtRuleSetAppStats.Stats.Appstat446,
			"appstat447": ret.DtRuleSetAppStats.Stats.Appstat447,
			"appstat448": ret.DtRuleSetAppStats.Stats.Appstat448,
			"appstat449": ret.DtRuleSetAppStats.Stats.Appstat449,
			"appstat450": ret.DtRuleSetAppStats.Stats.Appstat450,
			"appstat451": ret.DtRuleSetAppStats.Stats.Appstat451,
			"appstat452": ret.DtRuleSetAppStats.Stats.Appstat452,
			"appstat453": ret.DtRuleSetAppStats.Stats.Appstat453,
			"appstat454": ret.DtRuleSetAppStats.Stats.Appstat454,
			"appstat455": ret.DtRuleSetAppStats.Stats.Appstat455,
			"appstat456": ret.DtRuleSetAppStats.Stats.Appstat456,
			"appstat457": ret.DtRuleSetAppStats.Stats.Appstat457,
			"appstat458": ret.DtRuleSetAppStats.Stats.Appstat458,
			"appstat459": ret.DtRuleSetAppStats.Stats.Appstat459,
			"appstat460": ret.DtRuleSetAppStats.Stats.Appstat460,
			"appstat461": ret.DtRuleSetAppStats.Stats.Appstat461,
			"appstat462": ret.DtRuleSetAppStats.Stats.Appstat462,
			"appstat463": ret.DtRuleSetAppStats.Stats.Appstat463,
			"appstat464": ret.DtRuleSetAppStats.Stats.Appstat464,
			"appstat465": ret.DtRuleSetAppStats.Stats.Appstat465,
			"appstat466": ret.DtRuleSetAppStats.Stats.Appstat466,
			"appstat467": ret.DtRuleSetAppStats.Stats.Appstat467,
			"appstat468": ret.DtRuleSetAppStats.Stats.Appstat468,
			"appstat469": ret.DtRuleSetAppStats.Stats.Appstat469,
			"appstat470": ret.DtRuleSetAppStats.Stats.Appstat470,
			"appstat471": ret.DtRuleSetAppStats.Stats.Appstat471,
			"appstat472": ret.DtRuleSetAppStats.Stats.Appstat472,
			"appstat473": ret.DtRuleSetAppStats.Stats.Appstat473,
			"appstat474": ret.DtRuleSetAppStats.Stats.Appstat474,
			"appstat475": ret.DtRuleSetAppStats.Stats.Appstat475,
			"appstat476": ret.DtRuleSetAppStats.Stats.Appstat476,
			"appstat477": ret.DtRuleSetAppStats.Stats.Appstat477,
			"appstat478": ret.DtRuleSetAppStats.Stats.Appstat478,
			"appstat479": ret.DtRuleSetAppStats.Stats.Appstat479,
			"appstat480": ret.DtRuleSetAppStats.Stats.Appstat480,
			"appstat481": ret.DtRuleSetAppStats.Stats.Appstat481,
			"appstat482": ret.DtRuleSetAppStats.Stats.Appstat482,
			"appstat483": ret.DtRuleSetAppStats.Stats.Appstat483,
			"appstat484": ret.DtRuleSetAppStats.Stats.Appstat484,
			"appstat485": ret.DtRuleSetAppStats.Stats.Appstat485,
			"appstat486": ret.DtRuleSetAppStats.Stats.Appstat486,
			"appstat487": ret.DtRuleSetAppStats.Stats.Appstat487,
			"appstat488": ret.DtRuleSetAppStats.Stats.Appstat488,
			"appstat489": ret.DtRuleSetAppStats.Stats.Appstat489,
			"appstat490": ret.DtRuleSetAppStats.Stats.Appstat490,
			"appstat491": ret.DtRuleSetAppStats.Stats.Appstat491,
			"appstat492": ret.DtRuleSetAppStats.Stats.Appstat492,
			"appstat493": ret.DtRuleSetAppStats.Stats.Appstat493,
			"appstat494": ret.DtRuleSetAppStats.Stats.Appstat494,
			"appstat495": ret.DtRuleSetAppStats.Stats.Appstat495,
			"appstat496": ret.DtRuleSetAppStats.Stats.Appstat496,
			"appstat497": ret.DtRuleSetAppStats.Stats.Appstat497,
			"appstat498": ret.DtRuleSetAppStats.Stats.Appstat498,
			"appstat499": ret.DtRuleSetAppStats.Stats.Appstat499,
			"appstat500": ret.DtRuleSetAppStats.Stats.Appstat500,
			"appstat501": ret.DtRuleSetAppStats.Stats.Appstat501,
			"appstat502": ret.DtRuleSetAppStats.Stats.Appstat502,
			"appstat503": ret.DtRuleSetAppStats.Stats.Appstat503,
			"appstat504": ret.DtRuleSetAppStats.Stats.Appstat504,
			"appstat505": ret.DtRuleSetAppStats.Stats.Appstat505,
			"appstat506": ret.DtRuleSetAppStats.Stats.Appstat506,
			"appstat507": ret.DtRuleSetAppStats.Stats.Appstat507,
			"appstat508": ret.DtRuleSetAppStats.Stats.Appstat508,
			"appstat509": ret.DtRuleSetAppStats.Stats.Appstat509,
			"appstat510": ret.DtRuleSetAppStats.Stats.Appstat510,
			"appstat511": ret.DtRuleSetAppStats.Stats.Appstat511,
		},
	}
}

func getObjectRuleSetAppStatsStats(d []interface{}) edpt.RuleSetAppStatsStats {

	count1 := len(d)
	var ret edpt.RuleSetAppStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Appstat1 = in["appstat1"].(int)
		ret.Appstat2 = in["appstat2"].(int)
		ret.Appstat3 = in["appstat3"].(int)
		ret.Appstat4 = in["appstat4"].(int)
		ret.Appstat5 = in["appstat5"].(int)
		ret.Appstat6 = in["appstat6"].(int)
		ret.Appstat7 = in["appstat7"].(int)
		ret.Appstat8 = in["appstat8"].(int)
		ret.Appstat9 = in["appstat9"].(int)
		ret.Appstat10 = in["appstat10"].(int)
		ret.Appstat11 = in["appstat11"].(int)
		ret.Appstat12 = in["appstat12"].(int)
		ret.Appstat13 = in["appstat13"].(int)
		ret.Appstat14 = in["appstat14"].(int)
		ret.Appstat15 = in["appstat15"].(int)
		ret.Appstat16 = in["appstat16"].(int)
		ret.Appstat17 = in["appstat17"].(int)
		ret.Appstat18 = in["appstat18"].(int)
		ret.Appstat19 = in["appstat19"].(int)
		ret.Appstat20 = in["appstat20"].(int)
		ret.Appstat21 = in["appstat21"].(int)
		ret.Appstat22 = in["appstat22"].(int)
		ret.Appstat23 = in["appstat23"].(int)
		ret.Appstat24 = in["appstat24"].(int)
		ret.Appstat25 = in["appstat25"].(int)
		ret.Appstat26 = in["appstat26"].(int)
		ret.Appstat27 = in["appstat27"].(int)
		ret.Appstat28 = in["appstat28"].(int)
		ret.Appstat29 = in["appstat29"].(int)
		ret.Appstat30 = in["appstat30"].(int)
		ret.Appstat31 = in["appstat31"].(int)
		ret.Appstat32 = in["appstat32"].(int)
		ret.Appstat33 = in["appstat33"].(int)
		ret.Appstat34 = in["appstat34"].(int)
		ret.Appstat35 = in["appstat35"].(int)
		ret.Appstat36 = in["appstat36"].(int)
		ret.Appstat37 = in["appstat37"].(int)
		ret.Appstat38 = in["appstat38"].(int)
		ret.Appstat39 = in["appstat39"].(int)
		ret.Appstat40 = in["appstat40"].(int)
		ret.Appstat41 = in["appstat41"].(int)
		ret.Appstat42 = in["appstat42"].(int)
		ret.Appstat43 = in["appstat43"].(int)
		ret.Appstat44 = in["appstat44"].(int)
		ret.Appstat45 = in["appstat45"].(int)
		ret.Appstat46 = in["appstat46"].(int)
		ret.Appstat47 = in["appstat47"].(int)
		ret.Appstat48 = in["appstat48"].(int)
		ret.Appstat49 = in["appstat49"].(int)
		ret.Appstat50 = in["appstat50"].(int)
		ret.Appstat51 = in["appstat51"].(int)
		ret.Appstat52 = in["appstat52"].(int)
		ret.Appstat53 = in["appstat53"].(int)
		ret.Appstat54 = in["appstat54"].(int)
		ret.Appstat55 = in["appstat55"].(int)
		ret.Appstat56 = in["appstat56"].(int)
		ret.Appstat57 = in["appstat57"].(int)
		ret.Appstat58 = in["appstat58"].(int)
		ret.Appstat59 = in["appstat59"].(int)
		ret.Appstat60 = in["appstat60"].(int)
		ret.Appstat61 = in["appstat61"].(int)
		ret.Appstat62 = in["appstat62"].(int)
		ret.Appstat63 = in["appstat63"].(int)
		ret.Appstat64 = in["appstat64"].(int)
		ret.Appstat65 = in["appstat65"].(int)
		ret.Appstat66 = in["appstat66"].(int)
		ret.Appstat67 = in["appstat67"].(int)
		ret.Appstat68 = in["appstat68"].(int)
		ret.Appstat69 = in["appstat69"].(int)
		ret.Appstat70 = in["appstat70"].(int)
		ret.Appstat71 = in["appstat71"].(int)
		ret.Appstat72 = in["appstat72"].(int)
		ret.Appstat73 = in["appstat73"].(int)
		ret.Appstat74 = in["appstat74"].(int)
		ret.Appstat75 = in["appstat75"].(int)
		ret.Appstat76 = in["appstat76"].(int)
		ret.Appstat77 = in["appstat77"].(int)
		ret.Appstat78 = in["appstat78"].(int)
		ret.Appstat79 = in["appstat79"].(int)
		ret.Appstat80 = in["appstat80"].(int)
		ret.Appstat81 = in["appstat81"].(int)
		ret.Appstat82 = in["appstat82"].(int)
		ret.Appstat83 = in["appstat83"].(int)
		ret.Appstat84 = in["appstat84"].(int)
		ret.Appstat85 = in["appstat85"].(int)
		ret.Appstat86 = in["appstat86"].(int)
		ret.Appstat87 = in["appstat87"].(int)
		ret.Appstat88 = in["appstat88"].(int)
		ret.Appstat89 = in["appstat89"].(int)
		ret.Appstat90 = in["appstat90"].(int)
		ret.Appstat91 = in["appstat91"].(int)
		ret.Appstat92 = in["appstat92"].(int)
		ret.Appstat93 = in["appstat93"].(int)
		ret.Appstat94 = in["appstat94"].(int)
		ret.Appstat95 = in["appstat95"].(int)
		ret.Appstat96 = in["appstat96"].(int)
		ret.Appstat97 = in["appstat97"].(int)
		ret.Appstat98 = in["appstat98"].(int)
		ret.Appstat99 = in["appstat99"].(int)
		ret.Appstat100 = in["appstat100"].(int)
		ret.Appstat101 = in["appstat101"].(int)
		ret.Appstat102 = in["appstat102"].(int)
		ret.Appstat103 = in["appstat103"].(int)
		ret.Appstat104 = in["appstat104"].(int)
		ret.Appstat105 = in["appstat105"].(int)
		ret.Appstat106 = in["appstat106"].(int)
		ret.Appstat107 = in["appstat107"].(int)
		ret.Appstat108 = in["appstat108"].(int)
		ret.Appstat109 = in["appstat109"].(int)
		ret.Appstat110 = in["appstat110"].(int)
		ret.Appstat111 = in["appstat111"].(int)
		ret.Appstat112 = in["appstat112"].(int)
		ret.Appstat113 = in["appstat113"].(int)
		ret.Appstat114 = in["appstat114"].(int)
		ret.Appstat115 = in["appstat115"].(int)
		ret.Appstat116 = in["appstat116"].(int)
		ret.Appstat117 = in["appstat117"].(int)
		ret.Appstat118 = in["appstat118"].(int)
		ret.Appstat119 = in["appstat119"].(int)
		ret.Appstat120 = in["appstat120"].(int)
		ret.Appstat121 = in["appstat121"].(int)
		ret.Appstat122 = in["appstat122"].(int)
		ret.Appstat123 = in["appstat123"].(int)
		ret.Appstat124 = in["appstat124"].(int)
		ret.Appstat125 = in["appstat125"].(int)
		ret.Appstat126 = in["appstat126"].(int)
		ret.Appstat127 = in["appstat127"].(int)
		ret.Appstat128 = in["appstat128"].(int)
		ret.Appstat129 = in["appstat129"].(int)
		ret.Appstat130 = in["appstat130"].(int)
		ret.Appstat131 = in["appstat131"].(int)
		ret.Appstat132 = in["appstat132"].(int)
		ret.Appstat133 = in["appstat133"].(int)
		ret.Appstat134 = in["appstat134"].(int)
		ret.Appstat135 = in["appstat135"].(int)
		ret.Appstat136 = in["appstat136"].(int)
		ret.Appstat137 = in["appstat137"].(int)
		ret.Appstat138 = in["appstat138"].(int)
		ret.Appstat139 = in["appstat139"].(int)
		ret.Appstat140 = in["appstat140"].(int)
		ret.Appstat141 = in["appstat141"].(int)
		ret.Appstat142 = in["appstat142"].(int)
		ret.Appstat143 = in["appstat143"].(int)
		ret.Appstat144 = in["appstat144"].(int)
		ret.Appstat145 = in["appstat145"].(int)
		ret.Appstat146 = in["appstat146"].(int)
		ret.Appstat147 = in["appstat147"].(int)
		ret.Appstat148 = in["appstat148"].(int)
		ret.Appstat149 = in["appstat149"].(int)
		ret.Appstat150 = in["appstat150"].(int)
		ret.Appstat151 = in["appstat151"].(int)
		ret.Appstat152 = in["appstat152"].(int)
		ret.Appstat153 = in["appstat153"].(int)
		ret.Appstat154 = in["appstat154"].(int)
		ret.Appstat155 = in["appstat155"].(int)
		ret.Appstat156 = in["appstat156"].(int)
		ret.Appstat157 = in["appstat157"].(int)
		ret.Appstat158 = in["appstat158"].(int)
		ret.Appstat159 = in["appstat159"].(int)
		ret.Appstat160 = in["appstat160"].(int)
		ret.Appstat161 = in["appstat161"].(int)
		ret.Appstat162 = in["appstat162"].(int)
		ret.Appstat163 = in["appstat163"].(int)
		ret.Appstat164 = in["appstat164"].(int)
		ret.Appstat165 = in["appstat165"].(int)
		ret.Appstat166 = in["appstat166"].(int)
		ret.Appstat167 = in["appstat167"].(int)
		ret.Appstat168 = in["appstat168"].(int)
		ret.Appstat169 = in["appstat169"].(int)
		ret.Appstat170 = in["appstat170"].(int)
		ret.Appstat171 = in["appstat171"].(int)
		ret.Appstat172 = in["appstat172"].(int)
		ret.Appstat173 = in["appstat173"].(int)
		ret.Appstat174 = in["appstat174"].(int)
		ret.Appstat175 = in["appstat175"].(int)
		ret.Appstat176 = in["appstat176"].(int)
		ret.Appstat177 = in["appstat177"].(int)
		ret.Appstat178 = in["appstat178"].(int)
		ret.Appstat179 = in["appstat179"].(int)
		ret.Appstat180 = in["appstat180"].(int)
		ret.Appstat181 = in["appstat181"].(int)
		ret.Appstat182 = in["appstat182"].(int)
		ret.Appstat183 = in["appstat183"].(int)
		ret.Appstat184 = in["appstat184"].(int)
		ret.Appstat185 = in["appstat185"].(int)
		ret.Appstat186 = in["appstat186"].(int)
		ret.Appstat187 = in["appstat187"].(int)
		ret.Appstat188 = in["appstat188"].(int)
		ret.Appstat189 = in["appstat189"].(int)
		ret.Appstat190 = in["appstat190"].(int)
		ret.Appstat191 = in["appstat191"].(int)
		ret.Appstat192 = in["appstat192"].(int)
		ret.Appstat193 = in["appstat193"].(int)
		ret.Appstat194 = in["appstat194"].(int)
		ret.Appstat195 = in["appstat195"].(int)
		ret.Appstat196 = in["appstat196"].(int)
		ret.Appstat197 = in["appstat197"].(int)
		ret.Appstat198 = in["appstat198"].(int)
		ret.Appstat199 = in["appstat199"].(int)
		ret.Appstat200 = in["appstat200"].(int)
		ret.Appstat201 = in["appstat201"].(int)
		ret.Appstat202 = in["appstat202"].(int)
		ret.Appstat203 = in["appstat203"].(int)
		ret.Appstat204 = in["appstat204"].(int)
		ret.Appstat205 = in["appstat205"].(int)
		ret.Appstat206 = in["appstat206"].(int)
		ret.Appstat207 = in["appstat207"].(int)
		ret.Appstat208 = in["appstat208"].(int)
		ret.Appstat209 = in["appstat209"].(int)
		ret.Appstat210 = in["appstat210"].(int)
		ret.Appstat211 = in["appstat211"].(int)
		ret.Appstat212 = in["appstat212"].(int)
		ret.Appstat213 = in["appstat213"].(int)
		ret.Appstat214 = in["appstat214"].(int)
		ret.Appstat215 = in["appstat215"].(int)
		ret.Appstat216 = in["appstat216"].(int)
		ret.Appstat217 = in["appstat217"].(int)
		ret.Appstat218 = in["appstat218"].(int)
		ret.Appstat219 = in["appstat219"].(int)
		ret.Appstat220 = in["appstat220"].(int)
		ret.Appstat221 = in["appstat221"].(int)
		ret.Appstat222 = in["appstat222"].(int)
		ret.Appstat223 = in["appstat223"].(int)
		ret.Appstat224 = in["appstat224"].(int)
		ret.Appstat225 = in["appstat225"].(int)
		ret.Appstat226 = in["appstat226"].(int)
		ret.Appstat227 = in["appstat227"].(int)
		ret.Appstat228 = in["appstat228"].(int)
		ret.Appstat229 = in["appstat229"].(int)
		ret.Appstat230 = in["appstat230"].(int)
		ret.Appstat231 = in["appstat231"].(int)
		ret.Appstat232 = in["appstat232"].(int)
		ret.Appstat233 = in["appstat233"].(int)
		ret.Appstat234 = in["appstat234"].(int)
		ret.Appstat235 = in["appstat235"].(int)
		ret.Appstat236 = in["appstat236"].(int)
		ret.Appstat237 = in["appstat237"].(int)
		ret.Appstat238 = in["appstat238"].(int)
		ret.Appstat239 = in["appstat239"].(int)
		ret.Appstat240 = in["appstat240"].(int)
		ret.Appstat241 = in["appstat241"].(int)
		ret.Appstat242 = in["appstat242"].(int)
		ret.Appstat243 = in["appstat243"].(int)
		ret.Appstat244 = in["appstat244"].(int)
		ret.Appstat245 = in["appstat245"].(int)
		ret.Appstat246 = in["appstat246"].(int)
		ret.Appstat247 = in["appstat247"].(int)
		ret.Appstat248 = in["appstat248"].(int)
		ret.Appstat249 = in["appstat249"].(int)
		ret.Appstat250 = in["appstat250"].(int)
		ret.Appstat251 = in["appstat251"].(int)
		ret.Appstat252 = in["appstat252"].(int)
		ret.Appstat253 = in["appstat253"].(int)
		ret.Appstat254 = in["appstat254"].(int)
		ret.Appstat255 = in["appstat255"].(int)
		ret.Appstat256 = in["appstat256"].(int)
		ret.Appstat257 = in["appstat257"].(int)
		ret.Appstat258 = in["appstat258"].(int)
		ret.Appstat259 = in["appstat259"].(int)
		ret.Appstat260 = in["appstat260"].(int)
		ret.Appstat261 = in["appstat261"].(int)
		ret.Appstat262 = in["appstat262"].(int)
		ret.Appstat263 = in["appstat263"].(int)
		ret.Appstat264 = in["appstat264"].(int)
		ret.Appstat265 = in["appstat265"].(int)
		ret.Appstat266 = in["appstat266"].(int)
		ret.Appstat267 = in["appstat267"].(int)
		ret.Appstat268 = in["appstat268"].(int)
		ret.Appstat269 = in["appstat269"].(int)
		ret.Appstat270 = in["appstat270"].(int)
		ret.Appstat271 = in["appstat271"].(int)
		ret.Appstat272 = in["appstat272"].(int)
		ret.Appstat273 = in["appstat273"].(int)
		ret.Appstat274 = in["appstat274"].(int)
		ret.Appstat275 = in["appstat275"].(int)
		ret.Appstat276 = in["appstat276"].(int)
		ret.Appstat277 = in["appstat277"].(int)
		ret.Appstat278 = in["appstat278"].(int)
		ret.Appstat279 = in["appstat279"].(int)
		ret.Appstat280 = in["appstat280"].(int)
		ret.Appstat281 = in["appstat281"].(int)
		ret.Appstat282 = in["appstat282"].(int)
		ret.Appstat283 = in["appstat283"].(int)
		ret.Appstat284 = in["appstat284"].(int)
		ret.Appstat285 = in["appstat285"].(int)
		ret.Appstat286 = in["appstat286"].(int)
		ret.Appstat287 = in["appstat287"].(int)
		ret.Appstat288 = in["appstat288"].(int)
		ret.Appstat289 = in["appstat289"].(int)
		ret.Appstat290 = in["appstat290"].(int)
		ret.Appstat291 = in["appstat291"].(int)
		ret.Appstat292 = in["appstat292"].(int)
		ret.Appstat293 = in["appstat293"].(int)
		ret.Appstat294 = in["appstat294"].(int)
		ret.Appstat295 = in["appstat295"].(int)
		ret.Appstat296 = in["appstat296"].(int)
		ret.Appstat297 = in["appstat297"].(int)
		ret.Appstat298 = in["appstat298"].(int)
		ret.Appstat299 = in["appstat299"].(int)
		ret.Appstat300 = in["appstat300"].(int)
		ret.Appstat301 = in["appstat301"].(int)
		ret.Appstat302 = in["appstat302"].(int)
		ret.Appstat303 = in["appstat303"].(int)
		ret.Appstat304 = in["appstat304"].(int)
		ret.Appstat305 = in["appstat305"].(int)
		ret.Appstat306 = in["appstat306"].(int)
		ret.Appstat307 = in["appstat307"].(int)
		ret.Appstat308 = in["appstat308"].(int)
		ret.Appstat309 = in["appstat309"].(int)
		ret.Appstat310 = in["appstat310"].(int)
		ret.Appstat311 = in["appstat311"].(int)
		ret.Appstat312 = in["appstat312"].(int)
		ret.Appstat313 = in["appstat313"].(int)
		ret.Appstat314 = in["appstat314"].(int)
		ret.Appstat315 = in["appstat315"].(int)
		ret.Appstat316 = in["appstat316"].(int)
		ret.Appstat317 = in["appstat317"].(int)
		ret.Appstat318 = in["appstat318"].(int)
		ret.Appstat319 = in["appstat319"].(int)
		ret.Appstat320 = in["appstat320"].(int)
		ret.Appstat321 = in["appstat321"].(int)
		ret.Appstat322 = in["appstat322"].(int)
		ret.Appstat323 = in["appstat323"].(int)
		ret.Appstat324 = in["appstat324"].(int)
		ret.Appstat325 = in["appstat325"].(int)
		ret.Appstat326 = in["appstat326"].(int)
		ret.Appstat327 = in["appstat327"].(int)
		ret.Appstat328 = in["appstat328"].(int)
		ret.Appstat329 = in["appstat329"].(int)
		ret.Appstat330 = in["appstat330"].(int)
		ret.Appstat331 = in["appstat331"].(int)
		ret.Appstat332 = in["appstat332"].(int)
		ret.Appstat333 = in["appstat333"].(int)
		ret.Appstat334 = in["appstat334"].(int)
		ret.Appstat335 = in["appstat335"].(int)
		ret.Appstat336 = in["appstat336"].(int)
		ret.Appstat337 = in["appstat337"].(int)
		ret.Appstat338 = in["appstat338"].(int)
		ret.Appstat339 = in["appstat339"].(int)
		ret.Appstat340 = in["appstat340"].(int)
		ret.Appstat341 = in["appstat341"].(int)
		ret.Appstat342 = in["appstat342"].(int)
		ret.Appstat343 = in["appstat343"].(int)
		ret.Appstat344 = in["appstat344"].(int)
		ret.Appstat345 = in["appstat345"].(int)
		ret.Appstat346 = in["appstat346"].(int)
		ret.Appstat347 = in["appstat347"].(int)
		ret.Appstat348 = in["appstat348"].(int)
		ret.Appstat349 = in["appstat349"].(int)
		ret.Appstat350 = in["appstat350"].(int)
		ret.Appstat351 = in["appstat351"].(int)
		ret.Appstat352 = in["appstat352"].(int)
		ret.Appstat353 = in["appstat353"].(int)
		ret.Appstat354 = in["appstat354"].(int)
		ret.Appstat355 = in["appstat355"].(int)
		ret.Appstat356 = in["appstat356"].(int)
		ret.Appstat357 = in["appstat357"].(int)
		ret.Appstat358 = in["appstat358"].(int)
		ret.Appstat359 = in["appstat359"].(int)
		ret.Appstat360 = in["appstat360"].(int)
		ret.Appstat361 = in["appstat361"].(int)
		ret.Appstat362 = in["appstat362"].(int)
		ret.Appstat363 = in["appstat363"].(int)
		ret.Appstat364 = in["appstat364"].(int)
		ret.Appstat365 = in["appstat365"].(int)
		ret.Appstat366 = in["appstat366"].(int)
		ret.Appstat367 = in["appstat367"].(int)
		ret.Appstat368 = in["appstat368"].(int)
		ret.Appstat369 = in["appstat369"].(int)
		ret.Appstat370 = in["appstat370"].(int)
		ret.Appstat371 = in["appstat371"].(int)
		ret.Appstat372 = in["appstat372"].(int)
		ret.Appstat373 = in["appstat373"].(int)
		ret.Appstat374 = in["appstat374"].(int)
		ret.Appstat375 = in["appstat375"].(int)
		ret.Appstat376 = in["appstat376"].(int)
		ret.Appstat377 = in["appstat377"].(int)
		ret.Appstat378 = in["appstat378"].(int)
		ret.Appstat379 = in["appstat379"].(int)
		ret.Appstat380 = in["appstat380"].(int)
		ret.Appstat381 = in["appstat381"].(int)
		ret.Appstat382 = in["appstat382"].(int)
		ret.Appstat383 = in["appstat383"].(int)
		ret.Appstat384 = in["appstat384"].(int)
		ret.Appstat385 = in["appstat385"].(int)
		ret.Appstat386 = in["appstat386"].(int)
		ret.Appstat387 = in["appstat387"].(int)
		ret.Appstat388 = in["appstat388"].(int)
		ret.Appstat389 = in["appstat389"].(int)
		ret.Appstat390 = in["appstat390"].(int)
		ret.Appstat391 = in["appstat391"].(int)
		ret.Appstat392 = in["appstat392"].(int)
		ret.Appstat393 = in["appstat393"].(int)
		ret.Appstat394 = in["appstat394"].(int)
		ret.Appstat395 = in["appstat395"].(int)
		ret.Appstat396 = in["appstat396"].(int)
		ret.Appstat397 = in["appstat397"].(int)
		ret.Appstat398 = in["appstat398"].(int)
		ret.Appstat399 = in["appstat399"].(int)
		ret.Appstat400 = in["appstat400"].(int)
		ret.Appstat401 = in["appstat401"].(int)
		ret.Appstat402 = in["appstat402"].(int)
		ret.Appstat403 = in["appstat403"].(int)
		ret.Appstat404 = in["appstat404"].(int)
		ret.Appstat405 = in["appstat405"].(int)
		ret.Appstat406 = in["appstat406"].(int)
		ret.Appstat407 = in["appstat407"].(int)
		ret.Appstat408 = in["appstat408"].(int)
		ret.Appstat409 = in["appstat409"].(int)
		ret.Appstat410 = in["appstat410"].(int)
		ret.Appstat411 = in["appstat411"].(int)
		ret.Appstat412 = in["appstat412"].(int)
		ret.Appstat413 = in["appstat413"].(int)
		ret.Appstat414 = in["appstat414"].(int)
		ret.Appstat415 = in["appstat415"].(int)
		ret.Appstat416 = in["appstat416"].(int)
		ret.Appstat417 = in["appstat417"].(int)
		ret.Appstat418 = in["appstat418"].(int)
		ret.Appstat419 = in["appstat419"].(int)
		ret.Appstat420 = in["appstat420"].(int)
		ret.Appstat421 = in["appstat421"].(int)
		ret.Appstat422 = in["appstat422"].(int)
		ret.Appstat423 = in["appstat423"].(int)
		ret.Appstat424 = in["appstat424"].(int)
		ret.Appstat425 = in["appstat425"].(int)
		ret.Appstat426 = in["appstat426"].(int)
		ret.Appstat427 = in["appstat427"].(int)
		ret.Appstat428 = in["appstat428"].(int)
		ret.Appstat429 = in["appstat429"].(int)
		ret.Appstat430 = in["appstat430"].(int)
		ret.Appstat431 = in["appstat431"].(int)
		ret.Appstat432 = in["appstat432"].(int)
		ret.Appstat433 = in["appstat433"].(int)
		ret.Appstat434 = in["appstat434"].(int)
		ret.Appstat435 = in["appstat435"].(int)
		ret.Appstat436 = in["appstat436"].(int)
		ret.Appstat437 = in["appstat437"].(int)
		ret.Appstat438 = in["appstat438"].(int)
		ret.Appstat439 = in["appstat439"].(int)
		ret.Appstat440 = in["appstat440"].(int)
		ret.Appstat441 = in["appstat441"].(int)
		ret.Appstat442 = in["appstat442"].(int)
		ret.Appstat443 = in["appstat443"].(int)
		ret.Appstat444 = in["appstat444"].(int)
		ret.Appstat445 = in["appstat445"].(int)
		ret.Appstat446 = in["appstat446"].(int)
		ret.Appstat447 = in["appstat447"].(int)
		ret.Appstat448 = in["appstat448"].(int)
		ret.Appstat449 = in["appstat449"].(int)
		ret.Appstat450 = in["appstat450"].(int)
		ret.Appstat451 = in["appstat451"].(int)
		ret.Appstat452 = in["appstat452"].(int)
		ret.Appstat453 = in["appstat453"].(int)
		ret.Appstat454 = in["appstat454"].(int)
		ret.Appstat455 = in["appstat455"].(int)
		ret.Appstat456 = in["appstat456"].(int)
		ret.Appstat457 = in["appstat457"].(int)
		ret.Appstat458 = in["appstat458"].(int)
		ret.Appstat459 = in["appstat459"].(int)
		ret.Appstat460 = in["appstat460"].(int)
		ret.Appstat461 = in["appstat461"].(int)
		ret.Appstat462 = in["appstat462"].(int)
		ret.Appstat463 = in["appstat463"].(int)
		ret.Appstat464 = in["appstat464"].(int)
		ret.Appstat465 = in["appstat465"].(int)
		ret.Appstat466 = in["appstat466"].(int)
		ret.Appstat467 = in["appstat467"].(int)
		ret.Appstat468 = in["appstat468"].(int)
		ret.Appstat469 = in["appstat469"].(int)
		ret.Appstat470 = in["appstat470"].(int)
		ret.Appstat471 = in["appstat471"].(int)
		ret.Appstat472 = in["appstat472"].(int)
		ret.Appstat473 = in["appstat473"].(int)
		ret.Appstat474 = in["appstat474"].(int)
		ret.Appstat475 = in["appstat475"].(int)
		ret.Appstat476 = in["appstat476"].(int)
		ret.Appstat477 = in["appstat477"].(int)
		ret.Appstat478 = in["appstat478"].(int)
		ret.Appstat479 = in["appstat479"].(int)
		ret.Appstat480 = in["appstat480"].(int)
		ret.Appstat481 = in["appstat481"].(int)
		ret.Appstat482 = in["appstat482"].(int)
		ret.Appstat483 = in["appstat483"].(int)
		ret.Appstat484 = in["appstat484"].(int)
		ret.Appstat485 = in["appstat485"].(int)
		ret.Appstat486 = in["appstat486"].(int)
		ret.Appstat487 = in["appstat487"].(int)
		ret.Appstat488 = in["appstat488"].(int)
		ret.Appstat489 = in["appstat489"].(int)
		ret.Appstat490 = in["appstat490"].(int)
		ret.Appstat491 = in["appstat491"].(int)
		ret.Appstat492 = in["appstat492"].(int)
		ret.Appstat493 = in["appstat493"].(int)
		ret.Appstat494 = in["appstat494"].(int)
		ret.Appstat495 = in["appstat495"].(int)
		ret.Appstat496 = in["appstat496"].(int)
		ret.Appstat497 = in["appstat497"].(int)
		ret.Appstat498 = in["appstat498"].(int)
		ret.Appstat499 = in["appstat499"].(int)
		ret.Appstat500 = in["appstat500"].(int)
		ret.Appstat501 = in["appstat501"].(int)
		ret.Appstat502 = in["appstat502"].(int)
		ret.Appstat503 = in["appstat503"].(int)
		ret.Appstat504 = in["appstat504"].(int)
		ret.Appstat505 = in["appstat505"].(int)
		ret.Appstat506 = in["appstat506"].(int)
		ret.Appstat507 = in["appstat507"].(int)
		ret.Appstat508 = in["appstat508"].(int)
		ret.Appstat509 = in["appstat509"].(int)
		ret.Appstat510 = in["appstat510"].(int)
		ret.Appstat511 = in["appstat511"].(int)
	}
	return ret
}

func dataToEndpointRuleSetAppStats(d *schema.ResourceData) edpt.RuleSetAppStats {
	var ret edpt.RuleSetAppStats

	ret.Stats = getObjectRuleSetAppStatsStats(d.Get("stats").([]interface{}))

	ret.Name = d.Get("name").(string)
	return ret
}
