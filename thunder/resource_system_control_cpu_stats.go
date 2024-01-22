package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemControlCpuStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_control_cpu_stats`: Statistics for the object control-cpu\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemControlCpuStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ctrl_cpu_number": {
							Type: schema.TypeInt, Optional: true, Description: "Number of ctrl cpus",
						},
						"cpu_1": {
							Type: schema.TypeInt, Optional: true, Description: "Control CPU-1",
						},
						"cpu_2": {
							Type: schema.TypeInt, Optional: true, Description: "Control CPU-2",
						},
						"cpu_3": {
							Type: schema.TypeInt, Optional: true, Description: "Control CPU-3",
						},
						"cpu_4": {
							Type: schema.TypeInt, Optional: true, Description: "Control CPU-4",
						},
						"cpu_5": {
							Type: schema.TypeInt, Optional: true, Description: "Control CPU-5",
						},
						"cpu_6": {
							Type: schema.TypeInt, Optional: true, Description: "Control CPU-6",
						},
						"cpu_7": {
							Type: schema.TypeInt, Optional: true, Description: "Control CPU-7",
						},
						"cpu_8": {
							Type: schema.TypeInt, Optional: true, Description: "Control CPU-8",
						},
						"cpu_9": {
							Type: schema.TypeInt, Optional: true, Description: "Control CPU-9",
						},
						"cpu_10": {
							Type: schema.TypeInt, Optional: true, Description: "Control CPU-10",
						},
						"cpu_11": {
							Type: schema.TypeInt, Optional: true, Description: "Control CPU-11",
						},
						"cpu_12": {
							Type: schema.TypeInt, Optional: true, Description: "Control CPU-12",
						},
						"cpu_13": {
							Type: schema.TypeInt, Optional: true, Description: "Control CPU-13",
						},
						"cpu_14": {
							Type: schema.TypeInt, Optional: true, Description: "Control CPU-14",
						},
						"cpu_15": {
							Type: schema.TypeInt, Optional: true, Description: "Control CPU-15",
						},
						"cpu_16": {
							Type: schema.TypeInt, Optional: true, Description: "Control CPU-16",
						},
						"cpu_17": {
							Type: schema.TypeInt, Optional: true, Description: "Control CPU-17",
						},
						"cpu_18": {
							Type: schema.TypeInt, Optional: true, Description: "Control CPU-18",
						},
						"cpu_19": {
							Type: schema.TypeInt, Optional: true, Description: "Control CPU-19",
						},
						"cpu_20": {
							Type: schema.TypeInt, Optional: true, Description: "Control CPU-20",
						},
						"cpu_21": {
							Type: schema.TypeInt, Optional: true, Description: "Control CPU-21",
						},
						"cpu_22": {
							Type: schema.TypeInt, Optional: true, Description: "Control CPU-22",
						},
						"cpu_23": {
							Type: schema.TypeInt, Optional: true, Description: "Control CPU-23",
						},
						"cpu_24": {
							Type: schema.TypeInt, Optional: true, Description: "Control CPU-24",
						},
						"cpu_25": {
							Type: schema.TypeInt, Optional: true, Description: "Control CPU-25",
						},
						"cpu_26": {
							Type: schema.TypeInt, Optional: true, Description: "Control CPU-26",
						},
						"cpu_27": {
							Type: schema.TypeInt, Optional: true, Description: "Control CPU-27",
						},
						"cpu_28": {
							Type: schema.TypeInt, Optional: true, Description: "Control CPU-28",
						},
						"cpu_29": {
							Type: schema.TypeInt, Optional: true, Description: "Control CPU-29",
						},
						"cpu_30": {
							Type: schema.TypeInt, Optional: true, Description: "Control CPU-30",
						},
						"cpu_31": {
							Type: schema.TypeInt, Optional: true, Description: "Control CPU-31",
						},
						"cpu_32": {
							Type: schema.TypeInt, Optional: true, Description: "Control CPU-32",
						},
						"cpu_33": {
							Type: schema.TypeInt, Optional: true, Description: "Control CPU-33",
						},
						"cpu_34": {
							Type: schema.TypeInt, Optional: true, Description: "Control CPU-34",
						},
						"cpu_35": {
							Type: schema.TypeInt, Optional: true, Description: "Control CPU-35",
						},
						"cpu_36": {
							Type: schema.TypeInt, Optional: true, Description: "Control CPU-36",
						},
						"cpu_37": {
							Type: schema.TypeInt, Optional: true, Description: "Control CPU-37",
						},
						"cpu_38": {
							Type: schema.TypeInt, Optional: true, Description: "Control CPU-38",
						},
						"cpu_39": {
							Type: schema.TypeInt, Optional: true, Description: "Control CPU-39",
						},
						"cpu_40": {
							Type: schema.TypeInt, Optional: true, Description: "Control CPU-40",
						},
						"cpu_41": {
							Type: schema.TypeInt, Optional: true, Description: "Control CPU-41",
						},
						"cpu_42": {
							Type: schema.TypeInt, Optional: true, Description: "Control CPU-42",
						},
						"cpu_43": {
							Type: schema.TypeInt, Optional: true, Description: "Control CPU-43",
						},
						"cpu_44": {
							Type: schema.TypeInt, Optional: true, Description: "Control CPU-44",
						},
						"cpu_45": {
							Type: schema.TypeInt, Optional: true, Description: "Control CPU-45",
						},
						"cpu_46": {
							Type: schema.TypeInt, Optional: true, Description: "Control CPU-46",
						},
						"cpu_47": {
							Type: schema.TypeInt, Optional: true, Description: "Control CPU-47",
						},
						"cpu_48": {
							Type: schema.TypeInt, Optional: true, Description: "Control CPU-48",
						},
						"cpu_49": {
							Type: schema.TypeInt, Optional: true, Description: "Control CPU-49",
						},
						"cpu_50": {
							Type: schema.TypeInt, Optional: true, Description: "Control CPU-50",
						},
						"cpu_51": {
							Type: schema.TypeInt, Optional: true, Description: "Control CPU-51",
						},
						"cpu_52": {
							Type: schema.TypeInt, Optional: true, Description: "Control CPU-52",
						},
						"cpu_53": {
							Type: schema.TypeInt, Optional: true, Description: "Control CPU-53",
						},
						"cpu_54": {
							Type: schema.TypeInt, Optional: true, Description: "Control CPU-54",
						},
						"cpu_55": {
							Type: schema.TypeInt, Optional: true, Description: "Control CPU-55",
						},
						"cpu_56": {
							Type: schema.TypeInt, Optional: true, Description: "Control CPU-56",
						},
						"cpu_57": {
							Type: schema.TypeInt, Optional: true, Description: "Control CPU-57",
						},
						"cpu_58": {
							Type: schema.TypeInt, Optional: true, Description: "Control CPU-58",
						},
						"cpu_59": {
							Type: schema.TypeInt, Optional: true, Description: "Control CPU-59",
						},
						"cpu_60": {
							Type: schema.TypeInt, Optional: true, Description: "Control CPU-60",
						},
						"cpu_61": {
							Type: schema.TypeInt, Optional: true, Description: "Control CPU-61",
						},
						"cpu_62": {
							Type: schema.TypeInt, Optional: true, Description: "Control CPU-62",
						},
						"cpu_63": {
							Type: schema.TypeInt, Optional: true, Description: "Control CPU-63",
						},
						"cpu_64": {
							Type: schema.TypeInt, Optional: true, Description: "Control CPU-64",
						},
					},
				},
			},
		},
	}
}

func resourceSystemControlCpuStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemControlCpuStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemControlCpuStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemControlCpuStatsStats := setObjectSystemControlCpuStatsStats(res)
		d.Set("stats", SystemControlCpuStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemControlCpuStatsStats(ret edpt.DataSystemControlCpuStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"ctrl_cpu_number": ret.DtSystemControlCpuStats.Stats.CtrlCpuNumber,
			"cpu_1":           ret.DtSystemControlCpuStats.Stats.Cpu1,
			"cpu_2":           ret.DtSystemControlCpuStats.Stats.Cpu2,
			"cpu_3":           ret.DtSystemControlCpuStats.Stats.Cpu3,
			"cpu_4":           ret.DtSystemControlCpuStats.Stats.Cpu4,
			"cpu_5":           ret.DtSystemControlCpuStats.Stats.Cpu5,
			"cpu_6":           ret.DtSystemControlCpuStats.Stats.Cpu6,
			"cpu_7":           ret.DtSystemControlCpuStats.Stats.Cpu7,
			"cpu_8":           ret.DtSystemControlCpuStats.Stats.Cpu8,
			"cpu_9":           ret.DtSystemControlCpuStats.Stats.Cpu9,
			"cpu_10":          ret.DtSystemControlCpuStats.Stats.Cpu10,
			"cpu_11":          ret.DtSystemControlCpuStats.Stats.Cpu11,
			"cpu_12":          ret.DtSystemControlCpuStats.Stats.Cpu12,
			"cpu_13":          ret.DtSystemControlCpuStats.Stats.Cpu13,
			"cpu_14":          ret.DtSystemControlCpuStats.Stats.Cpu14,
			"cpu_15":          ret.DtSystemControlCpuStats.Stats.Cpu15,
			"cpu_16":          ret.DtSystemControlCpuStats.Stats.Cpu16,
			"cpu_17":          ret.DtSystemControlCpuStats.Stats.Cpu17,
			"cpu_18":          ret.DtSystemControlCpuStats.Stats.Cpu18,
			"cpu_19":          ret.DtSystemControlCpuStats.Stats.Cpu19,
			"cpu_20":          ret.DtSystemControlCpuStats.Stats.Cpu20,
			"cpu_21":          ret.DtSystemControlCpuStats.Stats.Cpu21,
			"cpu_22":          ret.DtSystemControlCpuStats.Stats.Cpu22,
			"cpu_23":          ret.DtSystemControlCpuStats.Stats.Cpu23,
			"cpu_24":          ret.DtSystemControlCpuStats.Stats.Cpu24,
			"cpu_25":          ret.DtSystemControlCpuStats.Stats.Cpu25,
			"cpu_26":          ret.DtSystemControlCpuStats.Stats.Cpu26,
			"cpu_27":          ret.DtSystemControlCpuStats.Stats.Cpu27,
			"cpu_28":          ret.DtSystemControlCpuStats.Stats.Cpu28,
			"cpu_29":          ret.DtSystemControlCpuStats.Stats.Cpu29,
			"cpu_30":          ret.DtSystemControlCpuStats.Stats.Cpu30,
			"cpu_31":          ret.DtSystemControlCpuStats.Stats.Cpu31,
			"cpu_32":          ret.DtSystemControlCpuStats.Stats.Cpu32,
			"cpu_33":          ret.DtSystemControlCpuStats.Stats.Cpu33,
			"cpu_34":          ret.DtSystemControlCpuStats.Stats.Cpu34,
			"cpu_35":          ret.DtSystemControlCpuStats.Stats.Cpu35,
			"cpu_36":          ret.DtSystemControlCpuStats.Stats.Cpu36,
			"cpu_37":          ret.DtSystemControlCpuStats.Stats.Cpu37,
			"cpu_38":          ret.DtSystemControlCpuStats.Stats.Cpu38,
			"cpu_39":          ret.DtSystemControlCpuStats.Stats.Cpu39,
			"cpu_40":          ret.DtSystemControlCpuStats.Stats.Cpu40,
			"cpu_41":          ret.DtSystemControlCpuStats.Stats.Cpu41,
			"cpu_42":          ret.DtSystemControlCpuStats.Stats.Cpu42,
			"cpu_43":          ret.DtSystemControlCpuStats.Stats.Cpu43,
			"cpu_44":          ret.DtSystemControlCpuStats.Stats.Cpu44,
			"cpu_45":          ret.DtSystemControlCpuStats.Stats.Cpu45,
			"cpu_46":          ret.DtSystemControlCpuStats.Stats.Cpu46,
			"cpu_47":          ret.DtSystemControlCpuStats.Stats.Cpu47,
			"cpu_48":          ret.DtSystemControlCpuStats.Stats.Cpu48,
			"cpu_49":          ret.DtSystemControlCpuStats.Stats.Cpu49,
			"cpu_50":          ret.DtSystemControlCpuStats.Stats.Cpu50,
			"cpu_51":          ret.DtSystemControlCpuStats.Stats.Cpu51,
			"cpu_52":          ret.DtSystemControlCpuStats.Stats.Cpu52,
			"cpu_53":          ret.DtSystemControlCpuStats.Stats.Cpu53,
			"cpu_54":          ret.DtSystemControlCpuStats.Stats.Cpu54,
			"cpu_55":          ret.DtSystemControlCpuStats.Stats.Cpu55,
			"cpu_56":          ret.DtSystemControlCpuStats.Stats.Cpu56,
			"cpu_57":          ret.DtSystemControlCpuStats.Stats.Cpu57,
			"cpu_58":          ret.DtSystemControlCpuStats.Stats.Cpu58,
			"cpu_59":          ret.DtSystemControlCpuStats.Stats.Cpu59,
			"cpu_60":          ret.DtSystemControlCpuStats.Stats.Cpu60,
			"cpu_61":          ret.DtSystemControlCpuStats.Stats.Cpu61,
			"cpu_62":          ret.DtSystemControlCpuStats.Stats.Cpu62,
			"cpu_63":          ret.DtSystemControlCpuStats.Stats.Cpu63,
			"cpu_64":          ret.DtSystemControlCpuStats.Stats.Cpu64,
		},
	}
}

func getObjectSystemControlCpuStatsStats(d []interface{}) edpt.SystemControlCpuStatsStats {

	count1 := len(d)
	var ret edpt.SystemControlCpuStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.CtrlCpuNumber = in["ctrl_cpu_number"].(int)
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
	}
	return ret
}

func dataToEndpointSystemControlCpuStats(d *schema.ResourceData) edpt.SystemControlCpuStats {
	var ret edpt.SystemControlCpuStats

	ret.Stats = getObjectSystemControlCpuStatsStats(d.Get("stats").([]interface{}))
	return ret
}
