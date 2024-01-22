package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePkiDeleteOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_pki_delete_oper`: \n\n__PLACEHOLDER__",
		ReadContext: resourcePkiDeleteOperRead,

		Schema: map[string]*schema.Schema{
			"filename": {
				Type: schema.TypeString, Optional: true, Description: "",
			},
		},
	}
}

func resourcePkiDeleteOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourcePkiDeleteOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointPkiDeleteOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointPkiDeleteOper(d *schema.ResourceData) edpt.PkiDeleteOper {
	var ret edpt.PkiDeleteOper

	ret.Filename = d.Get("filename").(string)
	return ret
}
