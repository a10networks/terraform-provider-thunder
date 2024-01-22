package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemDataCpuStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_data_cpu_stats`: Statistics for the object data-cpu\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemDataCpuStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"data_cpu_number": {
							Type: schema.TypeInt, Optional: true, Description: "Number of data cpus",
						},
						"cpu_1": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-1",
						},
						"cpu_2": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-2",
						},
						"cpu_3": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-3",
						},
						"cpu_4": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-4",
						},
						"cpu_5": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-5",
						},
						"cpu_6": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-6",
						},
						"cpu_7": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-7",
						},
						"cpu_8": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-8",
						},
						"cpu_9": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-9",
						},
						"cpu_10": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-10",
						},
						"cpu_11": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-11",
						},
						"cpu_12": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-12",
						},
						"cpu_13": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-13",
						},
						"cpu_14": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-14",
						},
						"cpu_15": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-15",
						},
						"cpu_16": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-16",
						},
						"cpu_17": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-17",
						},
						"cpu_18": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-18",
						},
						"cpu_19": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-19",
						},
						"cpu_20": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-20",
						},
						"cpu_21": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-21",
						},
						"cpu_22": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-22",
						},
						"cpu_23": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-23",
						},
						"cpu_24": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-24",
						},
						"cpu_25": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-25",
						},
						"cpu_26": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-26",
						},
						"cpu_27": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-27",
						},
						"cpu_28": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-28",
						},
						"cpu_29": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-29",
						},
						"cpu_30": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-30",
						},
						"cpu_31": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-31",
						},
						"cpu_32": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-32",
						},
						"cpu_33": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-33",
						},
						"cpu_34": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-34",
						},
						"cpu_35": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-35",
						},
						"cpu_36": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-36",
						},
						"cpu_37": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-37",
						},
						"cpu_38": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-38",
						},
						"cpu_39": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-39",
						},
						"cpu_40": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-40",
						},
						"cpu_41": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-41",
						},
						"cpu_42": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-42",
						},
						"cpu_43": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-43",
						},
						"cpu_44": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-44",
						},
						"cpu_45": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-45",
						},
						"cpu_46": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-46",
						},
						"cpu_47": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-47",
						},
						"cpu_48": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-48",
						},
						"cpu_49": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-49",
						},
						"cpu_50": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-50",
						},
						"cpu_51": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-51",
						},
						"cpu_52": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-52",
						},
						"cpu_53": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-53",
						},
						"cpu_54": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-54",
						},
						"cpu_55": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-55",
						},
						"cpu_56": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-56",
						},
						"cpu_57": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-57",
						},
						"cpu_58": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-58",
						},
						"cpu_59": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-59",
						},
						"cpu_60": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-60",
						},
						"cpu_61": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-61",
						},
						"cpu_62": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-62",
						},
						"cpu_63": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-63",
						},
						"cpu_64": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-64",
						},
						"cpu_65": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-65",
						},
						"cpu_66": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-66",
						},
						"cpu_67": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-67",
						},
						"cpu_68": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-68",
						},
						"cpu_69": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-69",
						},
						"cpu_70": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-70",
						},
						"cpu_71": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-71",
						},
						"cpu_72": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-72",
						},
						"cpu_73": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-73",
						},
						"cpu_74": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-74",
						},
						"cpu_75": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-75",
						},
						"cpu_76": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-76",
						},
						"cpu_77": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-77",
						},
						"cpu_78": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-78",
						},
						"cpu_79": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-79",
						},
						"cpu_80": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-80",
						},
						"cpu_81": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-81",
						},
						"cpu_82": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-82",
						},
						"cpu_83": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-83",
						},
						"cpu_84": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-84",
						},
						"cpu_85": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-85",
						},
						"cpu_86": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-86",
						},
						"cpu_87": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-87",
						},
						"cpu_88": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-88",
						},
						"cpu_89": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-89",
						},
						"cpu_90": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-90",
						},
						"cpu_91": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-91",
						},
						"cpu_92": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-92",
						},
						"cpu_93": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-93",
						},
						"cpu_94": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-94",
						},
						"cpu_95": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-95",
						},
						"cpu_96": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-96",
						},
						"cpu_97": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-97",
						},
						"cpu_98": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-98",
						},
						"cpu_99": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-99",
						},
						"cpu_100": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU-100",
						},
					},
				},
			},
		},
	}
}

func resourceSystemDataCpuStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemDataCpuStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemDataCpuStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemDataCpuStatsStats := setObjectSystemDataCpuStatsStats(res)
		d.Set("stats", SystemDataCpuStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemDataCpuStatsStats(ret edpt.DataSystemDataCpuStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"data_cpu_number": ret.DtSystemDataCpuStats.Stats.DataCpuNumber,
			"cpu_1":           ret.DtSystemDataCpuStats.Stats.Cpu1,
			"cpu_2":           ret.DtSystemDataCpuStats.Stats.Cpu2,
			"cpu_3":           ret.DtSystemDataCpuStats.Stats.Cpu3,
			"cpu_4":           ret.DtSystemDataCpuStats.Stats.Cpu4,
			"cpu_5":           ret.DtSystemDataCpuStats.Stats.Cpu5,
			"cpu_6":           ret.DtSystemDataCpuStats.Stats.Cpu6,
			"cpu_7":           ret.DtSystemDataCpuStats.Stats.Cpu7,
			"cpu_8":           ret.DtSystemDataCpuStats.Stats.Cpu8,
			"cpu_9":           ret.DtSystemDataCpuStats.Stats.Cpu9,
			"cpu_10":          ret.DtSystemDataCpuStats.Stats.Cpu10,
			"cpu_11":          ret.DtSystemDataCpuStats.Stats.Cpu11,
			"cpu_12":          ret.DtSystemDataCpuStats.Stats.Cpu12,
			"cpu_13":          ret.DtSystemDataCpuStats.Stats.Cpu13,
			"cpu_14":          ret.DtSystemDataCpuStats.Stats.Cpu14,
			"cpu_15":          ret.DtSystemDataCpuStats.Stats.Cpu15,
			"cpu_16":          ret.DtSystemDataCpuStats.Stats.Cpu16,
			"cpu_17":          ret.DtSystemDataCpuStats.Stats.Cpu17,
			"cpu_18":          ret.DtSystemDataCpuStats.Stats.Cpu18,
			"cpu_19":          ret.DtSystemDataCpuStats.Stats.Cpu19,
			"cpu_20":          ret.DtSystemDataCpuStats.Stats.Cpu20,
			"cpu_21":          ret.DtSystemDataCpuStats.Stats.Cpu21,
			"cpu_22":          ret.DtSystemDataCpuStats.Stats.Cpu22,
			"cpu_23":          ret.DtSystemDataCpuStats.Stats.Cpu23,
			"cpu_24":          ret.DtSystemDataCpuStats.Stats.Cpu24,
			"cpu_25":          ret.DtSystemDataCpuStats.Stats.Cpu25,
			"cpu_26":          ret.DtSystemDataCpuStats.Stats.Cpu26,
			"cpu_27":          ret.DtSystemDataCpuStats.Stats.Cpu27,
			"cpu_28":          ret.DtSystemDataCpuStats.Stats.Cpu28,
			"cpu_29":          ret.DtSystemDataCpuStats.Stats.Cpu29,
			"cpu_30":          ret.DtSystemDataCpuStats.Stats.Cpu30,
			"cpu_31":          ret.DtSystemDataCpuStats.Stats.Cpu31,
			"cpu_32":          ret.DtSystemDataCpuStats.Stats.Cpu32,
			"cpu_33":          ret.DtSystemDataCpuStats.Stats.Cpu33,
			"cpu_34":          ret.DtSystemDataCpuStats.Stats.Cpu34,
			"cpu_35":          ret.DtSystemDataCpuStats.Stats.Cpu35,
			"cpu_36":          ret.DtSystemDataCpuStats.Stats.Cpu36,
			"cpu_37":          ret.DtSystemDataCpuStats.Stats.Cpu37,
			"cpu_38":          ret.DtSystemDataCpuStats.Stats.Cpu38,
			"cpu_39":          ret.DtSystemDataCpuStats.Stats.Cpu39,
			"cpu_40":          ret.DtSystemDataCpuStats.Stats.Cpu40,
			"cpu_41":          ret.DtSystemDataCpuStats.Stats.Cpu41,
			"cpu_42":          ret.DtSystemDataCpuStats.Stats.Cpu42,
			"cpu_43":          ret.DtSystemDataCpuStats.Stats.Cpu43,
			"cpu_44":          ret.DtSystemDataCpuStats.Stats.Cpu44,
			"cpu_45":          ret.DtSystemDataCpuStats.Stats.Cpu45,
			"cpu_46":          ret.DtSystemDataCpuStats.Stats.Cpu46,
			"cpu_47":          ret.DtSystemDataCpuStats.Stats.Cpu47,
			"cpu_48":          ret.DtSystemDataCpuStats.Stats.Cpu48,
			"cpu_49":          ret.DtSystemDataCpuStats.Stats.Cpu49,
			"cpu_50":          ret.DtSystemDataCpuStats.Stats.Cpu50,
			"cpu_51":          ret.DtSystemDataCpuStats.Stats.Cpu51,
			"cpu_52":          ret.DtSystemDataCpuStats.Stats.Cpu52,
			"cpu_53":          ret.DtSystemDataCpuStats.Stats.Cpu53,
			"cpu_54":          ret.DtSystemDataCpuStats.Stats.Cpu54,
			"cpu_55":          ret.DtSystemDataCpuStats.Stats.Cpu55,
			"cpu_56":          ret.DtSystemDataCpuStats.Stats.Cpu56,
			"cpu_57":          ret.DtSystemDataCpuStats.Stats.Cpu57,
			"cpu_58":          ret.DtSystemDataCpuStats.Stats.Cpu58,
			"cpu_59":          ret.DtSystemDataCpuStats.Stats.Cpu59,
			"cpu_60":          ret.DtSystemDataCpuStats.Stats.Cpu60,
			"cpu_61":          ret.DtSystemDataCpuStats.Stats.Cpu61,
			"cpu_62":          ret.DtSystemDataCpuStats.Stats.Cpu62,
			"cpu_63":          ret.DtSystemDataCpuStats.Stats.Cpu63,
			"cpu_64":          ret.DtSystemDataCpuStats.Stats.Cpu64,
			"cpu_65":          ret.DtSystemDataCpuStats.Stats.Cpu65,
			"cpu_66":          ret.DtSystemDataCpuStats.Stats.Cpu66,
			"cpu_67":          ret.DtSystemDataCpuStats.Stats.Cpu67,
			"cpu_68":          ret.DtSystemDataCpuStats.Stats.Cpu68,
			"cpu_69":          ret.DtSystemDataCpuStats.Stats.Cpu69,
			"cpu_70":          ret.DtSystemDataCpuStats.Stats.Cpu70,
			"cpu_71":          ret.DtSystemDataCpuStats.Stats.Cpu71,
			"cpu_72":          ret.DtSystemDataCpuStats.Stats.Cpu72,
			"cpu_73":          ret.DtSystemDataCpuStats.Stats.Cpu73,
			"cpu_74":          ret.DtSystemDataCpuStats.Stats.Cpu74,
			"cpu_75":          ret.DtSystemDataCpuStats.Stats.Cpu75,
			"cpu_76":          ret.DtSystemDataCpuStats.Stats.Cpu76,
			"cpu_77":          ret.DtSystemDataCpuStats.Stats.Cpu77,
			"cpu_78":          ret.DtSystemDataCpuStats.Stats.Cpu78,
			"cpu_79":          ret.DtSystemDataCpuStats.Stats.Cpu79,
			"cpu_80":          ret.DtSystemDataCpuStats.Stats.Cpu80,
			"cpu_81":          ret.DtSystemDataCpuStats.Stats.Cpu81,
			"cpu_82":          ret.DtSystemDataCpuStats.Stats.Cpu82,
			"cpu_83":          ret.DtSystemDataCpuStats.Stats.Cpu83,
			"cpu_84":          ret.DtSystemDataCpuStats.Stats.Cpu84,
			"cpu_85":          ret.DtSystemDataCpuStats.Stats.Cpu85,
			"cpu_86":          ret.DtSystemDataCpuStats.Stats.Cpu86,
			"cpu_87":          ret.DtSystemDataCpuStats.Stats.Cpu87,
			"cpu_88":          ret.DtSystemDataCpuStats.Stats.Cpu88,
			"cpu_89":          ret.DtSystemDataCpuStats.Stats.Cpu89,
			"cpu_90":          ret.DtSystemDataCpuStats.Stats.Cpu90,
			"cpu_91":          ret.DtSystemDataCpuStats.Stats.Cpu91,
			"cpu_92":          ret.DtSystemDataCpuStats.Stats.Cpu92,
			"cpu_93":          ret.DtSystemDataCpuStats.Stats.Cpu93,
			"cpu_94":          ret.DtSystemDataCpuStats.Stats.Cpu94,
			"cpu_95":          ret.DtSystemDataCpuStats.Stats.Cpu95,
			"cpu_96":          ret.DtSystemDataCpuStats.Stats.Cpu96,
			"cpu_97":          ret.DtSystemDataCpuStats.Stats.Cpu97,
			"cpu_98":          ret.DtSystemDataCpuStats.Stats.Cpu98,
			"cpu_99":          ret.DtSystemDataCpuStats.Stats.Cpu99,
			"cpu_100":         ret.DtSystemDataCpuStats.Stats.Cpu100,
		},
	}
}

func getObjectSystemDataCpuStatsStats(d []interface{}) edpt.SystemDataCpuStatsStats {

	count1 := len(d)
	var ret edpt.SystemDataCpuStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DataCpuNumber = in["data_cpu_number"].(int)
		ret.Cpu1 = in["cpu_1"].(int)
		ret.Cpu2 = in["cpu_2"].(int)
		ret.Cpu3 = in["cpu_3"].(int)
		ret.Cpu4 = in["cpu_4"].(int)
		ret.Cpu5 = in["cpu_5"].(int)
		ret.Cpu6 = in["cpu_6"].(int)
		ret.Cpu7 = in["cpu_7"].(int)
		ret.Cpu8 = in["cpu_8"].(int)
		ret.Cpu9 = in["cpu_9"].(int)
		ret.Cpu10 = in["cpu_10"].(int)
		ret.Cpu11 = in["cpu_11"].(int)
		ret.Cpu12 = in["cpu_12"].(int)
		ret.Cpu13 = in["cpu_13"].(int)
		ret.Cpu14 = in["cpu_14"].(int)
		ret.Cpu15 = in["cpu_15"].(int)
		ret.Cpu16 = in["cpu_16"].(int)
		ret.Cpu17 = in["cpu_17"].(int)
		ret.Cpu18 = in["cpu_18"].(int)
		ret.Cpu19 = in["cpu_19"].(int)
		ret.Cpu20 = in["cpu_20"].(int)
		ret.Cpu21 = in["cpu_21"].(int)
		ret.Cpu22 = in["cpu_22"].(int)
		ret.Cpu23 = in["cpu_23"].(int)
		ret.Cpu24 = in["cpu_24"].(int)
		ret.Cpu25 = in["cpu_25"].(int)
		ret.Cpu26 = in["cpu_26"].(int)
		ret.Cpu27 = in["cpu_27"].(int)
		ret.Cpu28 = in["cpu_28"].(int)
		ret.Cpu29 = in["cpu_29"].(int)
		ret.Cpu30 = in["cpu_30"].(int)
		ret.Cpu31 = in["cpu_31"].(int)
		ret.Cpu32 = in["cpu_32"].(int)
		ret.Cpu33 = in["cpu_33"].(int)
		ret.Cpu34 = in["cpu_34"].(int)
		ret.Cpu35 = in["cpu_35"].(int)
		ret.Cpu36 = in["cpu_36"].(int)
		ret.Cpu37 = in["cpu_37"].(int)
		ret.Cpu38 = in["cpu_38"].(int)
		ret.Cpu39 = in["cpu_39"].(int)
		ret.Cpu40 = in["cpu_40"].(int)
		ret.Cpu41 = in["cpu_41"].(int)
		ret.Cpu42 = in["cpu_42"].(int)
		ret.Cpu43 = in["cpu_43"].(int)
		ret.Cpu44 = in["cpu_44"].(int)
		ret.Cpu45 = in["cpu_45"].(int)
		ret.Cpu46 = in["cpu_46"].(int)
		ret.Cpu47 = in["cpu_47"].(int)
		ret.Cpu48 = in["cpu_48"].(int)
		ret.Cpu49 = in["cpu_49"].(int)
		ret.Cpu50 = in["cpu_50"].(int)
		ret.Cpu51 = in["cpu_51"].(int)
		ret.Cpu52 = in["cpu_52"].(int)
		ret.Cpu53 = in["cpu_53"].(int)
		ret.Cpu54 = in["cpu_54"].(int)
		ret.Cpu55 = in["cpu_55"].(int)
		ret.Cpu56 = in["cpu_56"].(int)
		ret.Cpu57 = in["cpu_57"].(int)
		ret.Cpu58 = in["cpu_58"].(int)
		ret.Cpu59 = in["cpu_59"].(int)
		ret.Cpu60 = in["cpu_60"].(int)
		ret.Cpu61 = in["cpu_61"].(int)
		ret.Cpu62 = in["cpu_62"].(int)
		ret.Cpu63 = in["cpu_63"].(int)
		ret.Cpu64 = in["cpu_64"].(int)
		ret.Cpu65 = in["cpu_65"].(int)
		ret.Cpu66 = in["cpu_66"].(int)
		ret.Cpu67 = in["cpu_67"].(int)
		ret.Cpu68 = in["cpu_68"].(int)
		ret.Cpu69 = in["cpu_69"].(int)
		ret.Cpu70 = in["cpu_70"].(int)
		ret.Cpu71 = in["cpu_71"].(int)
		ret.Cpu72 = in["cpu_72"].(int)
		ret.Cpu73 = in["cpu_73"].(int)
		ret.Cpu74 = in["cpu_74"].(int)
		ret.Cpu75 = in["cpu_75"].(int)
		ret.Cpu76 = in["cpu_76"].(int)
		ret.Cpu77 = in["cpu_77"].(int)
		ret.Cpu78 = in["cpu_78"].(int)
		ret.Cpu79 = in["cpu_79"].(int)
		ret.Cpu80 = in["cpu_80"].(int)
		ret.Cpu81 = in["cpu_81"].(int)
		ret.Cpu82 = in["cpu_82"].(int)
		ret.Cpu83 = in["cpu_83"].(int)
		ret.Cpu84 = in["cpu_84"].(int)
		ret.Cpu85 = in["cpu_85"].(int)
		ret.Cpu86 = in["cpu_86"].(int)
		ret.Cpu87 = in["cpu_87"].(int)
		ret.Cpu88 = in["cpu_88"].(int)
		ret.Cpu89 = in["cpu_89"].(int)
		ret.Cpu90 = in["cpu_90"].(int)
		ret.Cpu91 = in["cpu_91"].(int)
		ret.Cpu92 = in["cpu_92"].(int)
		ret.Cpu93 = in["cpu_93"].(int)
		ret.Cpu94 = in["cpu_94"].(int)
		ret.Cpu95 = in["cpu_95"].(int)
		ret.Cpu96 = in["cpu_96"].(int)
		ret.Cpu97 = in["cpu_97"].(int)
		ret.Cpu98 = in["cpu_98"].(int)
		ret.Cpu99 = in["cpu_99"].(int)
		ret.Cpu100 = in["cpu_100"].(int)
	}
	return ret
}

func dataToEndpointSystemDataCpuStats(d *schema.ResourceData) edpt.SystemDataCpuStats {
	var ret edpt.SystemDataCpuStats

	ret.Stats = getObjectSystemDataCpuStatsStats(d.Get("stats").([]interface{}))
	return ret
}
