package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemAsicMmuFailSafeOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_asic_mmu_fail_safe_oper`: Operational Status for the object asic-mmu-fail-safe\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemAsicMmuFailSafeOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable_val": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"threshold_val": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"period_val": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"reboot_val": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"parity_unrecov": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"parity_error": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"parity_ecc": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"parity_ser": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"parity_start_err": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"mmu_ser": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"parity_bst": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"mmu_total_error": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceSystemAsicMmuFailSafeOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemAsicMmuFailSafeOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemAsicMmuFailSafeOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemAsicMmuFailSafeOperOper := setObjectSystemAsicMmuFailSafeOperOper(res)
		d.Set("oper", SystemAsicMmuFailSafeOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemAsicMmuFailSafeOperOper(ret edpt.DataSystemAsicMmuFailSafeOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"enable_val":       ret.DtSystemAsicMmuFailSafeOper.Oper.Enable_val,
			"threshold_val":    ret.DtSystemAsicMmuFailSafeOper.Oper.Threshold_val,
			"period_val":       ret.DtSystemAsicMmuFailSafeOper.Oper.Period_val,
			"reboot_val":       ret.DtSystemAsicMmuFailSafeOper.Oper.Reboot_val,
			"parity_unrecov":   ret.DtSystemAsicMmuFailSafeOper.Oper.Parity_unrecov,
			"parity_error":     ret.DtSystemAsicMmuFailSafeOper.Oper.Parity_error,
			"parity_ecc":       ret.DtSystemAsicMmuFailSafeOper.Oper.Parity_ecc,
			"parity_ser":       ret.DtSystemAsicMmuFailSafeOper.Oper.Parity_ser,
			"parity_start_err": ret.DtSystemAsicMmuFailSafeOper.Oper.Parity_start_err,
			"mmu_ser":          ret.DtSystemAsicMmuFailSafeOper.Oper.Mmu_ser,
			"parity_bst":       ret.DtSystemAsicMmuFailSafeOper.Oper.Parity_bst,
			"mmu_total_error":  ret.DtSystemAsicMmuFailSafeOper.Oper.Mmu_total_error,
		},
	}
}

func getObjectSystemAsicMmuFailSafeOperOper(d []interface{}) edpt.SystemAsicMmuFailSafeOperOper {

	count1 := len(d)
	var ret edpt.SystemAsicMmuFailSafeOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Enable_val = in["enable_val"].(string)
		ret.Threshold_val = in["threshold_val"].(int)
		ret.Period_val = in["period_val"].(int)
		ret.Reboot_val = in["reboot_val"].(string)
		ret.Parity_unrecov = in["parity_unrecov"].(int)
		ret.Parity_error = in["parity_error"].(int)
		ret.Parity_ecc = in["parity_ecc"].(int)
		ret.Parity_ser = in["parity_ser"].(int)
		ret.Parity_start_err = in["parity_start_err"].(int)
		ret.Mmu_ser = in["mmu_ser"].(int)
		ret.Parity_bst = in["parity_bst"].(int)
		ret.Mmu_total_error = in["mmu_total_error"].(int)
	}
	return ret
}

func dataToEndpointSystemAsicMmuFailSafeOper(d *schema.ResourceData) edpt.SystemAsicMmuFailSafeOper {
	var ret edpt.SystemAsicMmuFailSafeOper

	ret.Oper = getObjectSystemAsicMmuFailSafeOperOper(d.Get("oper").([]interface{}))
	return ret
}
