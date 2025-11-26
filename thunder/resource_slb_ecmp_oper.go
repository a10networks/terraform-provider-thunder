package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbEcmpOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_ecmp_oper`: Operational Status for the object ecmp\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbEcmpOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"total_path": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"ecmp_path": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"filter_type": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"source_addr_v4": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"dest_addr_v4": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"source_addr_v6": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"dest_addr_v6": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"dst_port": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"src_port": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceSlbEcmpOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbEcmpOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbEcmpOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbEcmpOperOper := setObjectSlbEcmpOperOper(res)
		d.Set("oper", SlbEcmpOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbEcmpOperOper(ret edpt.DataSlbEcmpOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"total_path":     ret.DtSlbEcmpOper.Oper.TotalPath,
			"ecmp_path":      ret.DtSlbEcmpOper.Oper.EcmpPath,
			"filter_type":    ret.DtSlbEcmpOper.Oper.Filter_type,
			"source_addr_v4": ret.DtSlbEcmpOper.Oper.SourceAddrV4,
			"dest_addr_v4":   ret.DtSlbEcmpOper.Oper.DestAddrV4,
			"source_addr_v6": ret.DtSlbEcmpOper.Oper.SourceAddrV6,
			"dest_addr_v6":   ret.DtSlbEcmpOper.Oper.DestAddrV6,
			"dst_port":       ret.DtSlbEcmpOper.Oper.DstPort,
			"src_port":       ret.DtSlbEcmpOper.Oper.SrcPort,
		},
	}
}

func getObjectSlbEcmpOperOper(d []interface{}) edpt.SlbEcmpOperOper {

	count1 := len(d)
	var ret edpt.SlbEcmpOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.TotalPath = in["total_path"].(int)
		ret.EcmpPath = in["ecmp_path"].(int)
		ret.Filter_type = in["filter_type"].(string)
		ret.SourceAddrV4 = in["source_addr_v4"].(string)
		ret.DestAddrV4 = in["dest_addr_v4"].(string)
		ret.SourceAddrV6 = in["source_addr_v6"].(string)
		ret.DestAddrV6 = in["dest_addr_v6"].(string)
		ret.DstPort = in["dst_port"].(int)
		ret.SrcPort = in["src_port"].(int)
	}
	return ret
}

func dataToEndpointSlbEcmpOper(d *schema.ResourceData) edpt.SlbEcmpOper {
	var ret edpt.SlbEcmpOper

	ret.Oper = getObjectSlbEcmpOperOper(d.Get("oper").([]interface{}))
	return ret
}
