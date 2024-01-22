package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstEntrySrcDstPairClassListCidAppTypeSrcDstCid() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dst_entry_src_dst_pair_class_list_cid_app_type_src_dst_cid`: DDOS APP type\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDstEntrySrcDstPairClassListCidAppTypeSrcDstCidCreate,
		UpdateContext: resourceDdosDstEntrySrcDstPairClassListCidAppTypeSrcDstCidUpdate,
		ReadContext:   resourceDdosDstEntrySrcDstPairClassListCidAppTypeSrcDstCidRead,
		DeleteContext: resourceDdosDstEntrySrcDstPairClassListCidAppTypeSrcDstCidDelete,

		Schema: map[string]*schema.Schema{
			"protocol": {
				Type: schema.TypeString, Required: true, Description: "'dns': dns; 'http': http; 'ssl-l4': ssl-l4; 'sip': sip;",
			},
			"template": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ssl_l4": {
							Type: schema.TypeString, Optional: true, Description: "DDOS SSL-L4 template",
						},
						"dns": {
							Type: schema.TypeString, Optional: true, Description: "DDOS dns template",
						},
						"http": {
							Type: schema.TypeString, Optional: true, Description: "DDOS http template",
						},
						"sip": {
							Type: schema.TypeString, Optional: true, Description: "DDOS sip template",
						},
					},
				},
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"cid_num": {
				Type: schema.TypeString, Required: true, Description: "CidNum",
			},
			"class_list_name": {
				Type: schema.TypeString, Required: true, Description: "ClassListName",
			},
			"dst_entry_name": {
				Type: schema.TypeString, Required: true, Description: "DstEntryName",
			},
		},
	}
}
func resourceDdosDstEntrySrcDstPairClassListCidAppTypeSrcDstCidCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntrySrcDstPairClassListCidAppTypeSrcDstCidCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntrySrcDstPairClassListCidAppTypeSrcDstCid(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstEntrySrcDstPairClassListCidAppTypeSrcDstCidRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDstEntrySrcDstPairClassListCidAppTypeSrcDstCidUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntrySrcDstPairClassListCidAppTypeSrcDstCidUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntrySrcDstPairClassListCidAppTypeSrcDstCid(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstEntrySrcDstPairClassListCidAppTypeSrcDstCidRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDstEntrySrcDstPairClassListCidAppTypeSrcDstCidDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntrySrcDstPairClassListCidAppTypeSrcDstCidDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntrySrcDstPairClassListCidAppTypeSrcDstCid(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDstEntrySrcDstPairClassListCidAppTypeSrcDstCidRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntrySrcDstPairClassListCidAppTypeSrcDstCidRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntrySrcDstPairClassListCidAppTypeSrcDstCid(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectDdosDstEntrySrcDstPairClassListCidAppTypeSrcDstCidTemplate(d []interface{}) edpt.DdosDstEntrySrcDstPairClassListCidAppTypeSrcDstCidTemplate {

	count1 := len(d)
	var ret edpt.DdosDstEntrySrcDstPairClassListCidAppTypeSrcDstCidTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SslL4 = in["ssl_l4"].(string)
		ret.Dns = in["dns"].(string)
		ret.Http = in["http"].(string)
		ret.Sip = in["sip"].(string)
	}
	return ret
}

func dataToEndpointDdosDstEntrySrcDstPairClassListCidAppTypeSrcDstCid(d *schema.ResourceData) edpt.DdosDstEntrySrcDstPairClassListCidAppTypeSrcDstCid {
	var ret edpt.DdosDstEntrySrcDstPairClassListCidAppTypeSrcDstCid
	ret.Inst.Protocol = d.Get("protocol").(string)
	ret.Inst.Template = getObjectDdosDstEntrySrcDstPairClassListCidAppTypeSrcDstCidTemplate(d.Get("template").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.CidNum = d.Get("cid_num").(string)
	ret.Inst.ClassListName = d.Get("class_list_name").(string)
	ret.Inst.DstEntryName = d.Get("dst_entry_name").(string)
	return ret
}
