package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceScmLicensecpeOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_scm_licensecpe_oper`: Operational Status for the object licensecpe\n\n__PLACEHOLDER__",
		ReadContext: resourceScmLicensecpeOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "",
						},
						"product": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"platform": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"cpe": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceScmLicensecpeOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScmLicensecpeOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScmLicensecpeOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		ScmLicensecpeOperOper := setObjectScmLicensecpeOperOper(res)
		d.Set("oper", ScmLicensecpeOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectScmLicensecpeOperOper(ret edpt.DataScmLicensecpeOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			//omit uuid
			"product":  ret.DtScmLicensecpeOper.Oper.Product,
			"platform": ret.DtScmLicensecpeOper.Oper.Platform,
			"cpe":      ret.DtScmLicensecpeOper.Oper.Cpe,
		},
	}
}

func getObjectScmLicensecpeOperOper(d []interface{}) edpt.ScmLicensecpeOperOper {

	count1 := len(d)
	var ret edpt.ScmLicensecpeOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.Product = in["product"].(string)
		ret.Platform = in["platform"].(string)
		ret.Cpe = in["cpe"].(string)
	}
	return ret
}

func dataToEndpointScmLicensecpeOper(d *schema.ResourceData) edpt.ScmLicensecpeOper {
	var ret edpt.ScmLicensecpeOper

	ret.Oper = getObjectScmLicensecpeOperOper(d.Get("oper").([]interface{}))
	return ret
}
