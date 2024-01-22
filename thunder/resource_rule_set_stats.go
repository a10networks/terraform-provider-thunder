package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRuleSetStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_rule_set_stats`: Statistics for the object rule-set\n\n__PLACEHOLDER__",
		ReadContext: resourceRuleSetStatsRead,

		Schema: map[string]*schema.Schema{
			"app": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
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
					},
				},
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Rule set name",
			},
			"rule_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString, Required: true, Description: "Rule name",
						},
						"stats": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"hit_count": {
										Type: schema.TypeInt, Optional: true, Description: "Hit counts",
									},
									"permit_bytes": {
										Type: schema.TypeInt, Optional: true, Description: "Permitted bytes counter",
									},
									"deny_bytes": {
										Type: schema.TypeInt, Optional: true, Description: "Denied bytes counter",
									},
									"reset_bytes": {
										Type: schema.TypeInt, Optional: true, Description: "Reset bytes counter",
									},
									"permit_packets": {
										Type: schema.TypeInt, Optional: true, Description: "Permitted packets counter",
									},
									"deny_packets": {
										Type: schema.TypeInt, Optional: true, Description: "Denied packets counter",
									},
									"reset_packets": {
										Type: schema.TypeInt, Optional: true, Description: "Reset packets counter",
									},
									"active_session_tcp": {
										Type: schema.TypeInt, Optional: true, Description: "Active TCP session counter",
									},
									"active_session_udp": {
										Type: schema.TypeInt, Optional: true, Description: "Active UDP session counter",
									},
									"active_session_icmp": {
										Type: schema.TypeInt, Optional: true, Description: "Active ICMP session counter",
									},
									"active_session_other": {
										Type: schema.TypeInt, Optional: true, Description: "Active other protocol session counter",
									},
									"session_tcp": {
										Type: schema.TypeInt, Optional: true, Description: "TCP session counter",
									},
									"session_udp": {
										Type: schema.TypeInt, Optional: true, Description: "UDP session counter",
									},
									"session_icmp": {
										Type: schema.TypeInt, Optional: true, Description: "ICMP session counter",
									},
									"session_other": {
										Type: schema.TypeInt, Optional: true, Description: "Other protocol session counter",
									},
									"active_session_sctp": {
										Type: schema.TypeInt, Optional: true, Description: "Active SCTP session counter",
									},
									"session_sctp": {
										Type: schema.TypeInt, Optional: true, Description: "SCTP session counter",
									},
									"hitcount_timestamp": {
										Type: schema.TypeInt, Optional: true, Description: "Last hit counts timestamp",
									},
									"rate_limit_drops": {
										Type: schema.TypeInt, Optional: true, Description: "Rate Limit Drops",
									},
								},
							},
						},
					},
				},
			},
			"rules_by_zone": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"stats": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"dummy": {
										Type: schema.TypeInt, Optional: true, Description: "Entry for a10countergen",
									},
								},
							},
						},
					},
				},
			},
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"unmatched_drops": {
							Type: schema.TypeInt, Optional: true, Description: "Unmatched drops counter",
						},
						"permit": {
							Type: schema.TypeInt, Optional: true, Description: "Permitted counter",
						},
						"deny": {
							Type: schema.TypeInt, Optional: true, Description: "Denied counter",
						},
						"reset": {
							Type: schema.TypeInt, Optional: true, Description: "Reset counter",
						},
					},
				},
			},
			"tag": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
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
					},
				},
			},
			"track_app_rule_list": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"stats": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"dummy": {
										Type: schema.TypeInt, Optional: true, Description: "Entry for a10countergen",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func resourceRuleSetStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRuleSetStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRuleSetStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		RuleSetStatsApp := setObjectRuleSetStatsApp(res)
		d.Set("app", RuleSetStatsApp)
		RuleSetStatsRuleList := setSliceRuleSetStatsRuleList(res)
		d.Set("rule_list", RuleSetStatsRuleList)
		RuleSetStatsRulesByZone := setObjectRuleSetStatsRulesByZone(res)
		d.Set("rules_by_zone", RuleSetStatsRulesByZone)
		RuleSetStatsStats := setObjectRuleSetStatsStats(res)
		d.Set("stats", RuleSetStatsStats)
		RuleSetStatsTag := setObjectRuleSetStatsTag(res)
		d.Set("tag", RuleSetStatsTag)
		RuleSetStatsTrackAppRuleList := setObjectRuleSetStatsTrackAppRuleList(res)
		d.Set("track_app_rule_list", RuleSetStatsTrackAppRuleList)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectRuleSetStatsApp(ret edpt.DataRuleSetStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"stats": setObjectRuleSetStatsAppStats(ret.DtRuleSetStats.App.Stats),
		},
	}
}

func setObjectRuleSetStatsAppStats(d edpt.RuleSetStatsAppStats) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["appstat1"] = d.Appstat1

	in["appstat2"] = d.Appstat2

	in["appstat3"] = d.Appstat3

	in["appstat4"] = d.Appstat4

	in["appstat5"] = d.Appstat5

	in["appstat6"] = d.Appstat6

	in["appstat7"] = d.Appstat7

	in["appstat8"] = d.Appstat8

	in["appstat9"] = d.Appstat9

	in["appstat10"] = d.Appstat10

	in["appstat11"] = d.Appstat11

	in["appstat12"] = d.Appstat12

	in["appstat13"] = d.Appstat13

	in["appstat14"] = d.Appstat14

	in["appstat15"] = d.Appstat15

	in["appstat16"] = d.Appstat16

	in["appstat17"] = d.Appstat17

	in["appstat18"] = d.Appstat18

	in["appstat19"] = d.Appstat19

	in["appstat20"] = d.Appstat20

	in["appstat21"] = d.Appstat21

	in["appstat22"] = d.Appstat22

	in["appstat23"] = d.Appstat23

	in["appstat24"] = d.Appstat24

	in["appstat25"] = d.Appstat25

	in["appstat26"] = d.Appstat26

	in["appstat27"] = d.Appstat27

	in["appstat28"] = d.Appstat28

	in["appstat29"] = d.Appstat29

	in["appstat30"] = d.Appstat30

	in["appstat31"] = d.Appstat31

	in["appstat32"] = d.Appstat32

	in["appstat33"] = d.Appstat33

	in["appstat34"] = d.Appstat34

	in["appstat35"] = d.Appstat35

	in["appstat36"] = d.Appstat36

	in["appstat37"] = d.Appstat37

	in["appstat38"] = d.Appstat38

	in["appstat39"] = d.Appstat39

	in["appstat40"] = d.Appstat40

	in["appstat41"] = d.Appstat41

	in["appstat42"] = d.Appstat42

	in["appstat43"] = d.Appstat43

	in["appstat44"] = d.Appstat44

	in["appstat45"] = d.Appstat45

	in["appstat46"] = d.Appstat46

	in["appstat47"] = d.Appstat47

	in["appstat48"] = d.Appstat48

	in["appstat49"] = d.Appstat49

	in["appstat50"] = d.Appstat50

	in["appstat51"] = d.Appstat51

	in["appstat52"] = d.Appstat52

	in["appstat53"] = d.Appstat53

	in["appstat54"] = d.Appstat54

	in["appstat55"] = d.Appstat55

	in["appstat56"] = d.Appstat56

	in["appstat57"] = d.Appstat57

	in["appstat58"] = d.Appstat58

	in["appstat59"] = d.Appstat59

	in["appstat60"] = d.Appstat60

	in["appstat61"] = d.Appstat61

	in["appstat62"] = d.Appstat62

	in["appstat63"] = d.Appstat63

	in["appstat64"] = d.Appstat64

	in["appstat65"] = d.Appstat65

	in["appstat66"] = d.Appstat66

	in["appstat67"] = d.Appstat67

	in["appstat68"] = d.Appstat68

	in["appstat69"] = d.Appstat69

	in["appstat70"] = d.Appstat70

	in["appstat71"] = d.Appstat71

	in["appstat72"] = d.Appstat72

	in["appstat73"] = d.Appstat73

	in["appstat74"] = d.Appstat74

	in["appstat75"] = d.Appstat75

	in["appstat76"] = d.Appstat76

	in["appstat77"] = d.Appstat77

	in["appstat78"] = d.Appstat78

	in["appstat79"] = d.Appstat79

	in["appstat80"] = d.Appstat80

	in["appstat81"] = d.Appstat81

	in["appstat82"] = d.Appstat82

	in["appstat83"] = d.Appstat83

	in["appstat84"] = d.Appstat84

	in["appstat85"] = d.Appstat85

	in["appstat86"] = d.Appstat86

	in["appstat87"] = d.Appstat87

	in["appstat88"] = d.Appstat88

	in["appstat89"] = d.Appstat89

	in["appstat90"] = d.Appstat90

	in["appstat91"] = d.Appstat91

	in["appstat92"] = d.Appstat92

	in["appstat93"] = d.Appstat93

	in["appstat94"] = d.Appstat94

	in["appstat95"] = d.Appstat95

	in["appstat96"] = d.Appstat96

	in["appstat97"] = d.Appstat97

	in["appstat98"] = d.Appstat98

	in["appstat99"] = d.Appstat99

	in["appstat100"] = d.Appstat100

	in["appstat101"] = d.Appstat101

	in["appstat102"] = d.Appstat102

	in["appstat103"] = d.Appstat103

	in["appstat104"] = d.Appstat104

	in["appstat105"] = d.Appstat105

	in["appstat106"] = d.Appstat106

	in["appstat107"] = d.Appstat107

	in["appstat108"] = d.Appstat108

	in["appstat109"] = d.Appstat109

	in["appstat110"] = d.Appstat110

	in["appstat111"] = d.Appstat111

	in["appstat112"] = d.Appstat112

	in["appstat113"] = d.Appstat113

	in["appstat114"] = d.Appstat114

	in["appstat115"] = d.Appstat115

	in["appstat116"] = d.Appstat116

	in["appstat117"] = d.Appstat117

	in["appstat118"] = d.Appstat118

	in["appstat119"] = d.Appstat119

	in["appstat120"] = d.Appstat120

	in["appstat121"] = d.Appstat121

	in["appstat122"] = d.Appstat122

	in["appstat123"] = d.Appstat123

	in["appstat124"] = d.Appstat124

	in["appstat125"] = d.Appstat125

	in["appstat126"] = d.Appstat126

	in["appstat127"] = d.Appstat127

	in["appstat128"] = d.Appstat128

	in["appstat129"] = d.Appstat129

	in["appstat130"] = d.Appstat130

	in["appstat131"] = d.Appstat131

	in["appstat132"] = d.Appstat132

	in["appstat133"] = d.Appstat133

	in["appstat134"] = d.Appstat134

	in["appstat135"] = d.Appstat135

	in["appstat136"] = d.Appstat136

	in["appstat137"] = d.Appstat137

	in["appstat138"] = d.Appstat138

	in["appstat139"] = d.Appstat139

	in["appstat140"] = d.Appstat140

	in["appstat141"] = d.Appstat141

	in["appstat142"] = d.Appstat142

	in["appstat143"] = d.Appstat143

	in["appstat144"] = d.Appstat144

	in["appstat145"] = d.Appstat145

	in["appstat146"] = d.Appstat146

	in["appstat147"] = d.Appstat147

	in["appstat148"] = d.Appstat148

	in["appstat149"] = d.Appstat149

	in["appstat150"] = d.Appstat150

	in["appstat151"] = d.Appstat151

	in["appstat152"] = d.Appstat152

	in["appstat153"] = d.Appstat153

	in["appstat154"] = d.Appstat154

	in["appstat155"] = d.Appstat155

	in["appstat156"] = d.Appstat156

	in["appstat157"] = d.Appstat157

	in["appstat158"] = d.Appstat158

	in["appstat159"] = d.Appstat159

	in["appstat160"] = d.Appstat160

	in["appstat161"] = d.Appstat161

	in["appstat162"] = d.Appstat162

	in["appstat163"] = d.Appstat163

	in["appstat164"] = d.Appstat164

	in["appstat165"] = d.Appstat165

	in["appstat166"] = d.Appstat166

	in["appstat167"] = d.Appstat167

	in["appstat168"] = d.Appstat168

	in["appstat169"] = d.Appstat169

	in["appstat170"] = d.Appstat170

	in["appstat171"] = d.Appstat171

	in["appstat172"] = d.Appstat172

	in["appstat173"] = d.Appstat173

	in["appstat174"] = d.Appstat174

	in["appstat175"] = d.Appstat175

	in["appstat176"] = d.Appstat176

	in["appstat177"] = d.Appstat177

	in["appstat178"] = d.Appstat178

	in["appstat179"] = d.Appstat179

	in["appstat180"] = d.Appstat180

	in["appstat181"] = d.Appstat181

	in["appstat182"] = d.Appstat182

	in["appstat183"] = d.Appstat183

	in["appstat184"] = d.Appstat184

	in["appstat185"] = d.Appstat185

	in["appstat186"] = d.Appstat186

	in["appstat187"] = d.Appstat187

	in["appstat188"] = d.Appstat188

	in["appstat189"] = d.Appstat189

	in["appstat190"] = d.Appstat190

	in["appstat191"] = d.Appstat191

	in["appstat192"] = d.Appstat192

	in["appstat193"] = d.Appstat193

	in["appstat194"] = d.Appstat194

	in["appstat195"] = d.Appstat195

	in["appstat196"] = d.Appstat196

	in["appstat197"] = d.Appstat197

	in["appstat198"] = d.Appstat198

	in["appstat199"] = d.Appstat199

	in["appstat200"] = d.Appstat200

	in["appstat201"] = d.Appstat201

	in["appstat202"] = d.Appstat202

	in["appstat203"] = d.Appstat203

	in["appstat204"] = d.Appstat204

	in["appstat205"] = d.Appstat205

	in["appstat206"] = d.Appstat206

	in["appstat207"] = d.Appstat207

	in["appstat208"] = d.Appstat208

	in["appstat209"] = d.Appstat209

	in["appstat210"] = d.Appstat210

	in["appstat211"] = d.Appstat211

	in["appstat212"] = d.Appstat212

	in["appstat213"] = d.Appstat213

	in["appstat214"] = d.Appstat214

	in["appstat215"] = d.Appstat215

	in["appstat216"] = d.Appstat216

	in["appstat217"] = d.Appstat217

	in["appstat218"] = d.Appstat218

	in["appstat219"] = d.Appstat219

	in["appstat220"] = d.Appstat220

	in["appstat221"] = d.Appstat221

	in["appstat222"] = d.Appstat222

	in["appstat223"] = d.Appstat223

	in["appstat224"] = d.Appstat224

	in["appstat225"] = d.Appstat225

	in["appstat226"] = d.Appstat226

	in["appstat227"] = d.Appstat227

	in["appstat228"] = d.Appstat228

	in["appstat229"] = d.Appstat229

	in["appstat230"] = d.Appstat230

	in["appstat231"] = d.Appstat231

	in["appstat232"] = d.Appstat232

	in["appstat233"] = d.Appstat233

	in["appstat234"] = d.Appstat234

	in["appstat235"] = d.Appstat235

	in["appstat236"] = d.Appstat236

	in["appstat237"] = d.Appstat237

	in["appstat238"] = d.Appstat238

	in["appstat239"] = d.Appstat239

	in["appstat240"] = d.Appstat240

	in["appstat241"] = d.Appstat241

	in["appstat242"] = d.Appstat242

	in["appstat243"] = d.Appstat243

	in["appstat244"] = d.Appstat244

	in["appstat245"] = d.Appstat245

	in["appstat246"] = d.Appstat246

	in["appstat247"] = d.Appstat247

	in["appstat248"] = d.Appstat248

	in["appstat249"] = d.Appstat249

	in["appstat250"] = d.Appstat250

	in["appstat251"] = d.Appstat251

	in["appstat252"] = d.Appstat252

	in["appstat253"] = d.Appstat253

	in["appstat254"] = d.Appstat254

	in["appstat255"] = d.Appstat255

	in["appstat256"] = d.Appstat256

	in["appstat257"] = d.Appstat257

	in["appstat258"] = d.Appstat258

	in["appstat259"] = d.Appstat259

	in["appstat260"] = d.Appstat260

	in["appstat261"] = d.Appstat261

	in["appstat262"] = d.Appstat262

	in["appstat263"] = d.Appstat263

	in["appstat264"] = d.Appstat264

	in["appstat265"] = d.Appstat265

	in["appstat266"] = d.Appstat266

	in["appstat267"] = d.Appstat267

	in["appstat268"] = d.Appstat268

	in["appstat269"] = d.Appstat269

	in["appstat270"] = d.Appstat270

	in["appstat271"] = d.Appstat271

	in["appstat272"] = d.Appstat272

	in["appstat273"] = d.Appstat273

	in["appstat274"] = d.Appstat274

	in["appstat275"] = d.Appstat275

	in["appstat276"] = d.Appstat276

	in["appstat277"] = d.Appstat277

	in["appstat278"] = d.Appstat278

	in["appstat279"] = d.Appstat279

	in["appstat280"] = d.Appstat280

	in["appstat281"] = d.Appstat281

	in["appstat282"] = d.Appstat282

	in["appstat283"] = d.Appstat283

	in["appstat284"] = d.Appstat284

	in["appstat285"] = d.Appstat285

	in["appstat286"] = d.Appstat286

	in["appstat287"] = d.Appstat287

	in["appstat288"] = d.Appstat288

	in["appstat289"] = d.Appstat289

	in["appstat290"] = d.Appstat290

	in["appstat291"] = d.Appstat291

	in["appstat292"] = d.Appstat292

	in["appstat293"] = d.Appstat293

	in["appstat294"] = d.Appstat294

	in["appstat295"] = d.Appstat295

	in["appstat296"] = d.Appstat296

	in["appstat297"] = d.Appstat297

	in["appstat298"] = d.Appstat298

	in["appstat299"] = d.Appstat299

	in["appstat300"] = d.Appstat300

	in["appstat301"] = d.Appstat301

	in["appstat302"] = d.Appstat302

	in["appstat303"] = d.Appstat303

	in["appstat304"] = d.Appstat304

	in["appstat305"] = d.Appstat305

	in["appstat306"] = d.Appstat306

	in["appstat307"] = d.Appstat307

	in["appstat308"] = d.Appstat308

	in["appstat309"] = d.Appstat309

	in["appstat310"] = d.Appstat310

	in["appstat311"] = d.Appstat311

	in["appstat312"] = d.Appstat312

	in["appstat313"] = d.Appstat313

	in["appstat314"] = d.Appstat314

	in["appstat315"] = d.Appstat315

	in["appstat316"] = d.Appstat316

	in["appstat317"] = d.Appstat317

	in["appstat318"] = d.Appstat318

	in["appstat319"] = d.Appstat319

	in["appstat320"] = d.Appstat320

	in["appstat321"] = d.Appstat321

	in["appstat322"] = d.Appstat322

	in["appstat323"] = d.Appstat323

	in["appstat324"] = d.Appstat324

	in["appstat325"] = d.Appstat325

	in["appstat326"] = d.Appstat326

	in["appstat327"] = d.Appstat327

	in["appstat328"] = d.Appstat328

	in["appstat329"] = d.Appstat329

	in["appstat330"] = d.Appstat330

	in["appstat331"] = d.Appstat331

	in["appstat332"] = d.Appstat332

	in["appstat333"] = d.Appstat333

	in["appstat334"] = d.Appstat334

	in["appstat335"] = d.Appstat335

	in["appstat336"] = d.Appstat336

	in["appstat337"] = d.Appstat337

	in["appstat338"] = d.Appstat338

	in["appstat339"] = d.Appstat339

	in["appstat340"] = d.Appstat340

	in["appstat341"] = d.Appstat341

	in["appstat342"] = d.Appstat342

	in["appstat343"] = d.Appstat343

	in["appstat344"] = d.Appstat344

	in["appstat345"] = d.Appstat345

	in["appstat346"] = d.Appstat346

	in["appstat347"] = d.Appstat347

	in["appstat348"] = d.Appstat348

	in["appstat349"] = d.Appstat349

	in["appstat350"] = d.Appstat350

	in["appstat351"] = d.Appstat351

	in["appstat352"] = d.Appstat352

	in["appstat353"] = d.Appstat353

	in["appstat354"] = d.Appstat354

	in["appstat355"] = d.Appstat355

	in["appstat356"] = d.Appstat356

	in["appstat357"] = d.Appstat357

	in["appstat358"] = d.Appstat358

	in["appstat359"] = d.Appstat359

	in["appstat360"] = d.Appstat360

	in["appstat361"] = d.Appstat361

	in["appstat362"] = d.Appstat362

	in["appstat363"] = d.Appstat363

	in["appstat364"] = d.Appstat364

	in["appstat365"] = d.Appstat365

	in["appstat366"] = d.Appstat366

	in["appstat367"] = d.Appstat367

	in["appstat368"] = d.Appstat368

	in["appstat369"] = d.Appstat369

	in["appstat370"] = d.Appstat370

	in["appstat371"] = d.Appstat371

	in["appstat372"] = d.Appstat372

	in["appstat373"] = d.Appstat373

	in["appstat374"] = d.Appstat374

	in["appstat375"] = d.Appstat375

	in["appstat376"] = d.Appstat376

	in["appstat377"] = d.Appstat377

	in["appstat378"] = d.Appstat378

	in["appstat379"] = d.Appstat379

	in["appstat380"] = d.Appstat380

	in["appstat381"] = d.Appstat381

	in["appstat382"] = d.Appstat382

	in["appstat383"] = d.Appstat383

	in["appstat384"] = d.Appstat384

	in["appstat385"] = d.Appstat385

	in["appstat386"] = d.Appstat386

	in["appstat387"] = d.Appstat387

	in["appstat388"] = d.Appstat388

	in["appstat389"] = d.Appstat389

	in["appstat390"] = d.Appstat390

	in["appstat391"] = d.Appstat391

	in["appstat392"] = d.Appstat392

	in["appstat393"] = d.Appstat393

	in["appstat394"] = d.Appstat394

	in["appstat395"] = d.Appstat395

	in["appstat396"] = d.Appstat396

	in["appstat397"] = d.Appstat397

	in["appstat398"] = d.Appstat398

	in["appstat399"] = d.Appstat399

	in["appstat400"] = d.Appstat400

	in["appstat401"] = d.Appstat401

	in["appstat402"] = d.Appstat402

	in["appstat403"] = d.Appstat403

	in["appstat404"] = d.Appstat404

	in["appstat405"] = d.Appstat405

	in["appstat406"] = d.Appstat406

	in["appstat407"] = d.Appstat407

	in["appstat408"] = d.Appstat408

	in["appstat409"] = d.Appstat409

	in["appstat410"] = d.Appstat410

	in["appstat411"] = d.Appstat411

	in["appstat412"] = d.Appstat412

	in["appstat413"] = d.Appstat413

	in["appstat414"] = d.Appstat414

	in["appstat415"] = d.Appstat415

	in["appstat416"] = d.Appstat416

	in["appstat417"] = d.Appstat417

	in["appstat418"] = d.Appstat418

	in["appstat419"] = d.Appstat419

	in["appstat420"] = d.Appstat420

	in["appstat421"] = d.Appstat421

	in["appstat422"] = d.Appstat422

	in["appstat423"] = d.Appstat423

	in["appstat424"] = d.Appstat424

	in["appstat425"] = d.Appstat425

	in["appstat426"] = d.Appstat426

	in["appstat427"] = d.Appstat427

	in["appstat428"] = d.Appstat428

	in["appstat429"] = d.Appstat429

	in["appstat430"] = d.Appstat430

	in["appstat431"] = d.Appstat431

	in["appstat432"] = d.Appstat432

	in["appstat433"] = d.Appstat433

	in["appstat434"] = d.Appstat434

	in["appstat435"] = d.Appstat435

	in["appstat436"] = d.Appstat436

	in["appstat437"] = d.Appstat437

	in["appstat438"] = d.Appstat438

	in["appstat439"] = d.Appstat439

	in["appstat440"] = d.Appstat440

	in["appstat441"] = d.Appstat441

	in["appstat442"] = d.Appstat442

	in["appstat443"] = d.Appstat443

	in["appstat444"] = d.Appstat444

	in["appstat445"] = d.Appstat445

	in["appstat446"] = d.Appstat446

	in["appstat447"] = d.Appstat447

	in["appstat448"] = d.Appstat448

	in["appstat449"] = d.Appstat449

	in["appstat450"] = d.Appstat450

	in["appstat451"] = d.Appstat451

	in["appstat452"] = d.Appstat452

	in["appstat453"] = d.Appstat453

	in["appstat454"] = d.Appstat454

	in["appstat455"] = d.Appstat455

	in["appstat456"] = d.Appstat456

	in["appstat457"] = d.Appstat457

	in["appstat458"] = d.Appstat458

	in["appstat459"] = d.Appstat459

	in["appstat460"] = d.Appstat460

	in["appstat461"] = d.Appstat461

	in["appstat462"] = d.Appstat462

	in["appstat463"] = d.Appstat463

	in["appstat464"] = d.Appstat464

	in["appstat465"] = d.Appstat465

	in["appstat466"] = d.Appstat466

	in["appstat467"] = d.Appstat467

	in["appstat468"] = d.Appstat468

	in["appstat469"] = d.Appstat469

	in["appstat470"] = d.Appstat470

	in["appstat471"] = d.Appstat471

	in["appstat472"] = d.Appstat472

	in["appstat473"] = d.Appstat473

	in["appstat474"] = d.Appstat474

	in["appstat475"] = d.Appstat475

	in["appstat476"] = d.Appstat476

	in["appstat477"] = d.Appstat477

	in["appstat478"] = d.Appstat478

	in["appstat479"] = d.Appstat479

	in["appstat480"] = d.Appstat480

	in["appstat481"] = d.Appstat481

	in["appstat482"] = d.Appstat482

	in["appstat483"] = d.Appstat483

	in["appstat484"] = d.Appstat484

	in["appstat485"] = d.Appstat485

	in["appstat486"] = d.Appstat486

	in["appstat487"] = d.Appstat487

	in["appstat488"] = d.Appstat488

	in["appstat489"] = d.Appstat489

	in["appstat490"] = d.Appstat490

	in["appstat491"] = d.Appstat491

	in["appstat492"] = d.Appstat492

	in["appstat493"] = d.Appstat493

	in["appstat494"] = d.Appstat494

	in["appstat495"] = d.Appstat495

	in["appstat496"] = d.Appstat496

	in["appstat497"] = d.Appstat497

	in["appstat498"] = d.Appstat498

	in["appstat499"] = d.Appstat499

	in["appstat500"] = d.Appstat500

	in["appstat501"] = d.Appstat501

	in["appstat502"] = d.Appstat502

	in["appstat503"] = d.Appstat503

	in["appstat504"] = d.Appstat504

	in["appstat505"] = d.Appstat505

	in["appstat506"] = d.Appstat506

	in["appstat507"] = d.Appstat507

	in["appstat508"] = d.Appstat508

	in["appstat509"] = d.Appstat509

	in["appstat510"] = d.Appstat510

	in["appstat511"] = d.Appstat511
	result = append(result, in)
	return result
}

func setSliceRuleSetStatsRuleList(d edpt.DataRuleSetStats) []map[string]interface{} {
	result := []map[string]interface{}{}

	for _, item := range d.DtRuleSetStats.RuleList {
		in := make(map[string]interface{})
		in["name"] = item.Name
		in["stats"] = setObjectRuleSetStatsRuleListStats(item.Stats)
		result = append(result, in)
	}
	return result
}

func setObjectRuleSetStatsRuleListStats(d edpt.RuleSetStatsRuleListStats) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["hit_count"] = d.HitCount

	in["permit_bytes"] = d.PermitBytes

	in["deny_bytes"] = d.DenyBytes

	in["reset_bytes"] = d.ResetBytes

	in["permit_packets"] = d.PermitPackets

	in["deny_packets"] = d.DenyPackets

	in["reset_packets"] = d.ResetPackets

	in["active_session_tcp"] = d.ActiveSessionTcp

	in["active_session_udp"] = d.ActiveSessionUdp

	in["active_session_icmp"] = d.ActiveSessionIcmp

	in["active_session_other"] = d.ActiveSessionOther

	in["session_tcp"] = d.SessionTcp

	in["session_udp"] = d.SessionUdp

	in["session_icmp"] = d.SessionIcmp

	in["session_other"] = d.SessionOther

	in["active_session_sctp"] = d.ActiveSessionSctp

	in["session_sctp"] = d.SessionSctp

	in["hitcount_timestamp"] = d.HitcountTimestamp

	in["rate_limit_drops"] = d.RateLimitDrops
	result = append(result, in)
	return result
}

func setObjectRuleSetStatsRulesByZone(ret edpt.DataRuleSetStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"stats": setObjectRuleSetStatsRulesByZoneStats(ret.DtRuleSetStats.RulesByZone.Stats),
		},
	}
}

func setObjectRuleSetStatsRulesByZoneStats(d edpt.RuleSetStatsRulesByZoneStats) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["dummy"] = d.Dummy
	result = append(result, in)
	return result
}

func setObjectRuleSetStatsStats(ret edpt.DataRuleSetStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"unmatched_drops": ret.DtRuleSetStats.Stats.UnmatchedDrops,
			"permit":          ret.DtRuleSetStats.Stats.Permit,
			"deny":            ret.DtRuleSetStats.Stats.Deny,
			"reset":           ret.DtRuleSetStats.Stats.Reset,
		},
	}
}

func setObjectRuleSetStatsTag(ret edpt.DataRuleSetStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"stats": setObjectRuleSetStatsTagStats(ret.DtRuleSetStats.Tag.Stats),
		},
	}
}

func setObjectRuleSetStatsTagStats(d edpt.RuleSetStatsTagStats) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["categorystat1"] = d.Categorystat1

	in["categorystat2"] = d.Categorystat2

	in["categorystat3"] = d.Categorystat3

	in["categorystat4"] = d.Categorystat4

	in["categorystat5"] = d.Categorystat5

	in["categorystat6"] = d.Categorystat6

	in["categorystat7"] = d.Categorystat7

	in["categorystat8"] = d.Categorystat8

	in["categorystat9"] = d.Categorystat9

	in["categorystat10"] = d.Categorystat10

	in["categorystat11"] = d.Categorystat11

	in["categorystat12"] = d.Categorystat12

	in["categorystat13"] = d.Categorystat13

	in["categorystat14"] = d.Categorystat14

	in["categorystat15"] = d.Categorystat15

	in["categorystat16"] = d.Categorystat16

	in["categorystat17"] = d.Categorystat17

	in["categorystat18"] = d.Categorystat18

	in["categorystat19"] = d.Categorystat19

	in["categorystat20"] = d.Categorystat20

	in["categorystat21"] = d.Categorystat21

	in["categorystat22"] = d.Categorystat22

	in["categorystat23"] = d.Categorystat23

	in["categorystat24"] = d.Categorystat24

	in["categorystat25"] = d.Categorystat25

	in["categorystat26"] = d.Categorystat26

	in["categorystat27"] = d.Categorystat27

	in["categorystat28"] = d.Categorystat28

	in["categorystat29"] = d.Categorystat29

	in["categorystat30"] = d.Categorystat30

	in["categorystat31"] = d.Categorystat31

	in["categorystat32"] = d.Categorystat32

	in["categorystat33"] = d.Categorystat33

	in["categorystat34"] = d.Categorystat34

	in["categorystat35"] = d.Categorystat35

	in["categorystat36"] = d.Categorystat36

	in["categorystat37"] = d.Categorystat37

	in["categorystat38"] = d.Categorystat38

	in["categorystat39"] = d.Categorystat39

	in["categorystat40"] = d.Categorystat40

	in["categorystat41"] = d.Categorystat41

	in["categorystat42"] = d.Categorystat42

	in["categorystat43"] = d.Categorystat43

	in["categorystat44"] = d.Categorystat44

	in["categorystat45"] = d.Categorystat45

	in["categorystat46"] = d.Categorystat46

	in["categorystat47"] = d.Categorystat47

	in["categorystat48"] = d.Categorystat48

	in["categorystat49"] = d.Categorystat49

	in["categorystat50"] = d.Categorystat50

	in["categorystat51"] = d.Categorystat51

	in["categorystat52"] = d.Categorystat52

	in["categorystat53"] = d.Categorystat53

	in["categorystat54"] = d.Categorystat54

	in["categorystat55"] = d.Categorystat55

	in["categorystat56"] = d.Categorystat56

	in["categorystat57"] = d.Categorystat57

	in["categorystat58"] = d.Categorystat58

	in["categorystat59"] = d.Categorystat59

	in["categorystat60"] = d.Categorystat60

	in["categorystat61"] = d.Categorystat61

	in["categorystat62"] = d.Categorystat62

	in["categorystat63"] = d.Categorystat63

	in["categorystat64"] = d.Categorystat64

	in["categorystat65"] = d.Categorystat65

	in["categorystat66"] = d.Categorystat66

	in["categorystat67"] = d.Categorystat67

	in["categorystat68"] = d.Categorystat68

	in["categorystat69"] = d.Categorystat69

	in["categorystat70"] = d.Categorystat70

	in["categorystat71"] = d.Categorystat71

	in["categorystat72"] = d.Categorystat72

	in["categorystat73"] = d.Categorystat73

	in["categorystat74"] = d.Categorystat74

	in["categorystat75"] = d.Categorystat75

	in["categorystat76"] = d.Categorystat76

	in["categorystat77"] = d.Categorystat77

	in["categorystat78"] = d.Categorystat78

	in["categorystat79"] = d.Categorystat79

	in["categorystat80"] = d.Categorystat80

	in["categorystat81"] = d.Categorystat81

	in["categorystat82"] = d.Categorystat82

	in["categorystat83"] = d.Categorystat83

	in["categorystat84"] = d.Categorystat84

	in["categorystat85"] = d.Categorystat85

	in["categorystat86"] = d.Categorystat86

	in["categorystat87"] = d.Categorystat87

	in["categorystat88"] = d.Categorystat88

	in["categorystat89"] = d.Categorystat89

	in["categorystat90"] = d.Categorystat90

	in["categorystat91"] = d.Categorystat91

	in["categorystat92"] = d.Categorystat92

	in["categorystat93"] = d.Categorystat93

	in["categorystat94"] = d.Categorystat94

	in["categorystat95"] = d.Categorystat95

	in["categorystat96"] = d.Categorystat96

	in["categorystat97"] = d.Categorystat97

	in["categorystat98"] = d.Categorystat98

	in["categorystat99"] = d.Categorystat99

	in["categorystat100"] = d.Categorystat100

	in["categorystat101"] = d.Categorystat101

	in["categorystat102"] = d.Categorystat102

	in["categorystat103"] = d.Categorystat103

	in["categorystat104"] = d.Categorystat104

	in["categorystat105"] = d.Categorystat105

	in["categorystat106"] = d.Categorystat106

	in["categorystat107"] = d.Categorystat107

	in["categorystat108"] = d.Categorystat108

	in["categorystat109"] = d.Categorystat109

	in["categorystat110"] = d.Categorystat110

	in["categorystat111"] = d.Categorystat111

	in["categorystat112"] = d.Categorystat112

	in["categorystat113"] = d.Categorystat113

	in["categorystat114"] = d.Categorystat114

	in["categorystat115"] = d.Categorystat115

	in["categorystat116"] = d.Categorystat116

	in["categorystat117"] = d.Categorystat117

	in["categorystat118"] = d.Categorystat118

	in["categorystat119"] = d.Categorystat119

	in["categorystat120"] = d.Categorystat120

	in["categorystat121"] = d.Categorystat121

	in["categorystat122"] = d.Categorystat122

	in["categorystat123"] = d.Categorystat123

	in["categorystat124"] = d.Categorystat124

	in["categorystat125"] = d.Categorystat125

	in["categorystat126"] = d.Categorystat126

	in["categorystat127"] = d.Categorystat127

	in["categorystat128"] = d.Categorystat128

	in["categorystat129"] = d.Categorystat129

	in["categorystat130"] = d.Categorystat130

	in["categorystat131"] = d.Categorystat131

	in["categorystat132"] = d.Categorystat132

	in["categorystat133"] = d.Categorystat133

	in["categorystat134"] = d.Categorystat134

	in["categorystat135"] = d.Categorystat135

	in["categorystat136"] = d.Categorystat136

	in["categorystat137"] = d.Categorystat137

	in["categorystat138"] = d.Categorystat138

	in["categorystat139"] = d.Categorystat139

	in["categorystat140"] = d.Categorystat140

	in["categorystat141"] = d.Categorystat141

	in["categorystat142"] = d.Categorystat142

	in["categorystat143"] = d.Categorystat143

	in["categorystat144"] = d.Categorystat144

	in["categorystat145"] = d.Categorystat145

	in["categorystat146"] = d.Categorystat146

	in["categorystat147"] = d.Categorystat147

	in["categorystat148"] = d.Categorystat148

	in["categorystat149"] = d.Categorystat149

	in["categorystat150"] = d.Categorystat150

	in["categorystat151"] = d.Categorystat151

	in["categorystat152"] = d.Categorystat152

	in["categorystat153"] = d.Categorystat153

	in["categorystat154"] = d.Categorystat154

	in["categorystat155"] = d.Categorystat155

	in["categorystat156"] = d.Categorystat156

	in["categorystat157"] = d.Categorystat157

	in["categorystat158"] = d.Categorystat158

	in["categorystat159"] = d.Categorystat159

	in["categorystat160"] = d.Categorystat160

	in["categorystat161"] = d.Categorystat161

	in["categorystat162"] = d.Categorystat162

	in["categorystat163"] = d.Categorystat163

	in["categorystat164"] = d.Categorystat164

	in["categorystat165"] = d.Categorystat165

	in["categorystat166"] = d.Categorystat166

	in["categorystat167"] = d.Categorystat167

	in["categorystat168"] = d.Categorystat168

	in["categorystat169"] = d.Categorystat169

	in["categorystat170"] = d.Categorystat170

	in["categorystat171"] = d.Categorystat171

	in["categorystat172"] = d.Categorystat172

	in["categorystat173"] = d.Categorystat173

	in["categorystat174"] = d.Categorystat174

	in["categorystat175"] = d.Categorystat175

	in["categorystat176"] = d.Categorystat176

	in["categorystat177"] = d.Categorystat177

	in["categorystat178"] = d.Categorystat178

	in["categorystat179"] = d.Categorystat179

	in["categorystat180"] = d.Categorystat180

	in["categorystat181"] = d.Categorystat181

	in["categorystat182"] = d.Categorystat182

	in["categorystat183"] = d.Categorystat183

	in["categorystat184"] = d.Categorystat184

	in["categorystat185"] = d.Categorystat185

	in["categorystat186"] = d.Categorystat186

	in["categorystat187"] = d.Categorystat187

	in["categorystat188"] = d.Categorystat188

	in["categorystat189"] = d.Categorystat189

	in["categorystat190"] = d.Categorystat190

	in["categorystat191"] = d.Categorystat191

	in["categorystat192"] = d.Categorystat192

	in["categorystat193"] = d.Categorystat193

	in["categorystat194"] = d.Categorystat194

	in["categorystat195"] = d.Categorystat195

	in["categorystat196"] = d.Categorystat196

	in["categorystat197"] = d.Categorystat197

	in["categorystat198"] = d.Categorystat198

	in["categorystat199"] = d.Categorystat199

	in["categorystat200"] = d.Categorystat200

	in["categorystat201"] = d.Categorystat201

	in["categorystat202"] = d.Categorystat202

	in["categorystat203"] = d.Categorystat203

	in["categorystat204"] = d.Categorystat204

	in["categorystat205"] = d.Categorystat205

	in["categorystat206"] = d.Categorystat206

	in["categorystat207"] = d.Categorystat207

	in["categorystat208"] = d.Categorystat208

	in["categorystat209"] = d.Categorystat209

	in["categorystat210"] = d.Categorystat210

	in["categorystat211"] = d.Categorystat211

	in["categorystat212"] = d.Categorystat212

	in["categorystat213"] = d.Categorystat213

	in["categorystat214"] = d.Categorystat214

	in["categorystat215"] = d.Categorystat215

	in["categorystat216"] = d.Categorystat216

	in["categorystat217"] = d.Categorystat217

	in["categorystat218"] = d.Categorystat218

	in["categorystat219"] = d.Categorystat219

	in["categorystat220"] = d.Categorystat220

	in["categorystat221"] = d.Categorystat221

	in["categorystat222"] = d.Categorystat222

	in["categorystat223"] = d.Categorystat223

	in["categorystat224"] = d.Categorystat224

	in["categorystat225"] = d.Categorystat225

	in["categorystat226"] = d.Categorystat226

	in["categorystat227"] = d.Categorystat227

	in["categorystat228"] = d.Categorystat228

	in["categorystat229"] = d.Categorystat229

	in["categorystat230"] = d.Categorystat230

	in["categorystat231"] = d.Categorystat231

	in["categorystat232"] = d.Categorystat232

	in["categorystat233"] = d.Categorystat233

	in["categorystat234"] = d.Categorystat234

	in["categorystat235"] = d.Categorystat235

	in["categorystat236"] = d.Categorystat236

	in["categorystat237"] = d.Categorystat237

	in["categorystat238"] = d.Categorystat238

	in["categorystat239"] = d.Categorystat239

	in["categorystat240"] = d.Categorystat240

	in["categorystat241"] = d.Categorystat241

	in["categorystat242"] = d.Categorystat242

	in["categorystat243"] = d.Categorystat243

	in["categorystat244"] = d.Categorystat244

	in["categorystat245"] = d.Categorystat245

	in["categorystat246"] = d.Categorystat246

	in["categorystat247"] = d.Categorystat247

	in["categorystat248"] = d.Categorystat248

	in["categorystat249"] = d.Categorystat249

	in["categorystat250"] = d.Categorystat250

	in["categorystat251"] = d.Categorystat251

	in["categorystat252"] = d.Categorystat252

	in["categorystat253"] = d.Categorystat253

	in["categorystat254"] = d.Categorystat254

	in["categorystat255"] = d.Categorystat255

	in["categorystat256"] = d.Categorystat256
	result = append(result, in)
	return result
}

func setObjectRuleSetStatsTrackAppRuleList(ret edpt.DataRuleSetStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"stats": setObjectRuleSetStatsTrackAppRuleListStats(ret.DtRuleSetStats.TrackAppRuleList.Stats),
		},
	}
}

func setObjectRuleSetStatsTrackAppRuleListStats(d edpt.RuleSetStatsTrackAppRuleListStats) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["dummy"] = d.Dummy
	result = append(result, in)
	return result
}

func getObjectRuleSetStatsApp(d []interface{}) edpt.RuleSetStatsApp {

	count1 := len(d)
	var ret edpt.RuleSetStatsApp
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Stats = getObjectRuleSetStatsAppStats(in["stats"].([]interface{}))
	}
	return ret
}

func getObjectRuleSetStatsAppStats(d []interface{}) edpt.RuleSetStatsAppStats {

	count1 := len(d)
	var ret edpt.RuleSetStatsAppStats
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

func getSliceRuleSetStatsRuleList(d []interface{}) []edpt.RuleSetStatsRuleList {

	count1 := len(d)
	ret := make([]edpt.RuleSetStatsRuleList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RuleSetStatsRuleList
		oi.Name = in["name"].(string)
		oi.Stats = getObjectRuleSetStatsRuleListStats(in["stats"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectRuleSetStatsRuleListStats(d []interface{}) edpt.RuleSetStatsRuleListStats {

	count1 := len(d)
	var ret edpt.RuleSetStatsRuleListStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.HitCount = in["hit_count"].(int)
		ret.PermitBytes = in["permit_bytes"].(int)
		ret.DenyBytes = in["deny_bytes"].(int)
		ret.ResetBytes = in["reset_bytes"].(int)
		ret.PermitPackets = in["permit_packets"].(int)
		ret.DenyPackets = in["deny_packets"].(int)
		ret.ResetPackets = in["reset_packets"].(int)
		ret.ActiveSessionTcp = in["active_session_tcp"].(int)
		ret.ActiveSessionUdp = in["active_session_udp"].(int)
		ret.ActiveSessionIcmp = in["active_session_icmp"].(int)
		ret.ActiveSessionOther = in["active_session_other"].(int)
		ret.SessionTcp = in["session_tcp"].(int)
		ret.SessionUdp = in["session_udp"].(int)
		ret.SessionIcmp = in["session_icmp"].(int)
		ret.SessionOther = in["session_other"].(int)
		ret.ActiveSessionSctp = in["active_session_sctp"].(int)
		ret.SessionSctp = in["session_sctp"].(int)
		ret.HitcountTimestamp = in["hitcount_timestamp"].(int)
		ret.RateLimitDrops = in["rate_limit_drops"].(int)
	}
	return ret
}

func getObjectRuleSetStatsRulesByZone(d []interface{}) edpt.RuleSetStatsRulesByZone {

	count1 := len(d)
	var ret edpt.RuleSetStatsRulesByZone
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Stats = getObjectRuleSetStatsRulesByZoneStats(in["stats"].([]interface{}))
	}
	return ret
}

func getObjectRuleSetStatsRulesByZoneStats(d []interface{}) edpt.RuleSetStatsRulesByZoneStats {

	count1 := len(d)
	var ret edpt.RuleSetStatsRulesByZoneStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Dummy = in["dummy"].(int)
	}
	return ret
}

func getObjectRuleSetStatsStats(d []interface{}) edpt.RuleSetStatsStats {

	count1 := len(d)
	var ret edpt.RuleSetStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.UnmatchedDrops = in["unmatched_drops"].(int)
		ret.Permit = in["permit"].(int)
		ret.Deny = in["deny"].(int)
		ret.Reset = in["reset"].(int)
	}
	return ret
}

func getObjectRuleSetStatsTag(d []interface{}) edpt.RuleSetStatsTag {

	count1 := len(d)
	var ret edpt.RuleSetStatsTag
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Stats = getObjectRuleSetStatsTagStats(in["stats"].([]interface{}))
	}
	return ret
}

func getObjectRuleSetStatsTagStats(d []interface{}) edpt.RuleSetStatsTagStats {

	count1 := len(d)
	var ret edpt.RuleSetStatsTagStats
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

func getObjectRuleSetStatsTrackAppRuleList(d []interface{}) edpt.RuleSetStatsTrackAppRuleList {

	count1 := len(d)
	var ret edpt.RuleSetStatsTrackAppRuleList
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Stats = getObjectRuleSetStatsTrackAppRuleListStats(in["stats"].([]interface{}))
	}
	return ret
}

func getObjectRuleSetStatsTrackAppRuleListStats(d []interface{}) edpt.RuleSetStatsTrackAppRuleListStats {

	count1 := len(d)
	var ret edpt.RuleSetStatsTrackAppRuleListStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Dummy = in["dummy"].(int)
	}
	return ret
}

func dataToEndpointRuleSetStats(d *schema.ResourceData) edpt.RuleSetStats {
	var ret edpt.RuleSetStats

	ret.App = getObjectRuleSetStatsApp(d.Get("app").([]interface{}))

	ret.Name = d.Get("name").(string)

	ret.RuleList = getSliceRuleSetStatsRuleList(d.Get("rule_list").([]interface{}))

	ret.RulesByZone = getObjectRuleSetStatsRulesByZone(d.Get("rules_by_zone").([]interface{}))

	ret.Stats = getObjectRuleSetStatsStats(d.Get("stats").([]interface{}))

	ret.Tag = getObjectRuleSetStatsTag(d.Get("tag").([]interface{}))

	ret.TrackAppRuleList = getObjectRuleSetStatsTrackAppRuleList(d.Get("track_app_rule_list").([]interface{}))
	return ret
}
