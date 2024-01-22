package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6SctpPermitPayloadProtocolProtocolName() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_sctp_permit_payload_protocol_protocol_name`: Configure SCTP permitted payload protocols\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6SctpPermitPayloadProtocolProtocolNameCreate,
		UpdateContext: resourceCgnv6SctpPermitPayloadProtocolProtocolNameUpdate,
		ReadContext:   resourceCgnv6SctpPermitPayloadProtocolProtocolNameRead,
		DeleteContext: resourceCgnv6SctpPermitPayloadProtocolProtocolNameDelete,

		Schema: map[string]*schema.Schema{
			"protocol": {
				Type: schema.TypeString, Required: true, Description: "'iua': IUA; 'm2ua': M2UA; 'm3ua': M3UA; 'sua': SUA; 'm2pa': M2PA; 'h.323': H.323;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceCgnv6SctpPermitPayloadProtocolProtocolNameCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6SctpPermitPayloadProtocolProtocolNameCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6SctpPermitPayloadProtocolProtocolName(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6SctpPermitPayloadProtocolProtocolNameRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6SctpPermitPayloadProtocolProtocolNameUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6SctpPermitPayloadProtocolProtocolNameUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6SctpPermitPayloadProtocolProtocolName(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6SctpPermitPayloadProtocolProtocolNameRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6SctpPermitPayloadProtocolProtocolNameDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6SctpPermitPayloadProtocolProtocolNameDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6SctpPermitPayloadProtocolProtocolName(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6SctpPermitPayloadProtocolProtocolNameRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6SctpPermitPayloadProtocolProtocolNameRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6SctpPermitPayloadProtocolProtocolName(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointCgnv6SctpPermitPayloadProtocolProtocolName(d *schema.ResourceData) edpt.Cgnv6SctpPermitPayloadProtocolProtocolName {
	var ret edpt.Cgnv6SctpPermitPayloadProtocolProtocolName
	ret.Inst.Protocol = d.Get("protocol").(string)
	//omit uuid
	return ret
}
