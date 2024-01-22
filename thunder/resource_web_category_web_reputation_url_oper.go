package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWebCategoryWebReputationUrlOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_web_category_web_reputation_url_oper`: Operational Status for the object url\n\n__PLACEHOLDER__",
		ReadContext: resourceWebCategoryWebReputationUrlOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"reputation_score": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"name": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"local_db_only": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceWebCategoryWebReputationUrlOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceWebCategoryWebReputationUrlOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointWebCategoryWebReputationUrlOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		WebCategoryWebReputationUrlOperOper := setObjectWebCategoryWebReputationUrlOperOper(res)
		d.Set("oper", WebCategoryWebReputationUrlOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectWebCategoryWebReputationUrlOperOper(ret edpt.DataWebCategoryWebReputationUrlOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"reputation_score": ret.DtWebCategoryWebReputationUrlOper.Oper.ReputationScore,
			"name":             ret.DtWebCategoryWebReputationUrlOper.Oper.Name,
			"local_db_only":    ret.DtWebCategoryWebReputationUrlOper.Oper.LocalDbOnly,
		},
	}
}

func getObjectWebCategoryWebReputationUrlOperOper(d []interface{}) edpt.WebCategoryWebReputationUrlOperOper {

	count1 := len(d)
	var ret edpt.WebCategoryWebReputationUrlOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ReputationScore = in["reputation_score"].(string)
		ret.Name = in["name"].(string)
		ret.LocalDbOnly = in["local_db_only"].(int)
	}
	return ret
}

func dataToEndpointWebCategoryWebReputationUrlOper(d *schema.ResourceData) edpt.WebCategoryWebReputationUrlOper {
	var ret edpt.WebCategoryWebReputationUrlOper

	ret.Oper = getObjectWebCategoryWebReputationUrlOperOper(d.Get("oper").([]interface{}))
	return ret
}
