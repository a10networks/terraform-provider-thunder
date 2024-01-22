package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6MapTranslationDomainHealthCheckGateway() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_map_translation_domain_health_check_gateway`: Health-check gateway for route withdrawn\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6MapTranslationDomainHealthCheckGatewayCreate,
		UpdateContext: resourceCgnv6MapTranslationDomainHealthCheckGatewayUpdate,
		ReadContext:   resourceCgnv6MapTranslationDomainHealthCheckGatewayRead,
		DeleteContext: resourceCgnv6MapTranslationDomainHealthCheckGatewayDelete,

		Schema: map[string]*schema.Schema{
			"address_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ipv4_gateway": {
							Type: schema.TypeString, Optional: true, Description: "IPv4 Gateway",
						},
					},
				},
			},
			"ipv6_address_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ipv6_gateway": {
							Type: schema.TypeString, Optional: true, Description: "IPv6 Gateway",
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"withdraw_route": {
				Type: schema.TypeString, Optional: true, Default: "any-link-failure", Description: "'all-link-failure': Withdraw routes on health-check failure of all IPv4 gateways or all IPv6 gateways; 'any-link-failure': Withdraw routes on health-check failure of any gateway (default);",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceCgnv6MapTranslationDomainHealthCheckGatewayCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6MapTranslationDomainHealthCheckGatewayCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6MapTranslationDomainHealthCheckGateway(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6MapTranslationDomainHealthCheckGatewayRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6MapTranslationDomainHealthCheckGatewayUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6MapTranslationDomainHealthCheckGatewayUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6MapTranslationDomainHealthCheckGateway(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6MapTranslationDomainHealthCheckGatewayRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6MapTranslationDomainHealthCheckGatewayDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6MapTranslationDomainHealthCheckGatewayDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6MapTranslationDomainHealthCheckGateway(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6MapTranslationDomainHealthCheckGatewayRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6MapTranslationDomainHealthCheckGatewayRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6MapTranslationDomainHealthCheckGateway(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceCgnv6MapTranslationDomainHealthCheckGatewayAddressList(d []interface{}) []edpt.Cgnv6MapTranslationDomainHealthCheckGatewayAddressList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6MapTranslationDomainHealthCheckGatewayAddressList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6MapTranslationDomainHealthCheckGatewayAddressList
		oi.Ipv4Gateway = in["ipv4_gateway"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceCgnv6MapTranslationDomainHealthCheckGatewayIpv6AddressList(d []interface{}) []edpt.Cgnv6MapTranslationDomainHealthCheckGatewayIpv6AddressList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6MapTranslationDomainHealthCheckGatewayIpv6AddressList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6MapTranslationDomainHealthCheckGatewayIpv6AddressList
		oi.Ipv6Gateway = in["ipv6_gateway"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCgnv6MapTranslationDomainHealthCheckGateway(d *schema.ResourceData) edpt.Cgnv6MapTranslationDomainHealthCheckGateway {
	var ret edpt.Cgnv6MapTranslationDomainHealthCheckGateway
	ret.Inst.AddressList = getSliceCgnv6MapTranslationDomainHealthCheckGatewayAddressList(d.Get("address_list").([]interface{}))
	ret.Inst.Ipv6AddressList = getSliceCgnv6MapTranslationDomainHealthCheckGatewayIpv6AddressList(d.Get("ipv6_address_list").([]interface{}))
	//omit uuid
	ret.Inst.WithdrawRoute = d.Get("withdraw_route").(string)
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
