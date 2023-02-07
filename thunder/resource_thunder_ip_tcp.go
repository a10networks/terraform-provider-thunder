package thunder

//Thunder resource IpTcp

import (
	"context"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpTcp() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIpTcpCreate,
		UpdateContext: resourceIpTcpUpdate,
		ReadContext:   resourceIpTcpRead,
		DeleteContext: resourceIpTcpDelete,
		Schema: map[string]*schema.Schema{
			"syn_cookie": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"threshold": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceIpTcpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println(d)
	if client.Host != "" {
		logger.Println("[INFO] Creating IpTcp (Inside resourceIpTcpCreate) ")

		data := dataToIpTcp(d)
		logger.Println("[INFO] received V from method data to IpTcp --")
		d.SetId("1")
		err := go_thunder.PostIpTcp(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceIpTcpRead(ctx, d, meta)

	}
	return diags
}

func resourceIpTcpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading IpTcp (Inside resourceIpTcpRead)")

	if client.Host != "" {
		data, err := go_thunder.GetIpTcp(client.Token, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		if data == nil {
			d.SetId("")
			return nil
		}
		return diags
	}
	return nil
}
func resourceIpTcpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceIpTcpRead(ctx, d, meta)
}

func resourceIpTcpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceIpTcpRead(ctx, d, meta)
}

func dataToIpTcp(d *schema.ResourceData) go_thunder.IpTcp {
	var vc go_thunder.IpTcp
	var c go_thunder.IpTcpInstance
	logger := util.GetLoggerInstance()
	logger.Println(d.Get("syn_cookie"))
	var obj1 go_thunder.SynCookie
	obj1.Threshold = d.Get("syn_cookie.0.threshold").(int)
	//obj1.Threshold = 4
	c.Threshold = obj1

	vc.UUID = c
	return vc
}
