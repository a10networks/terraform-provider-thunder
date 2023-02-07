package thunder

//Thunder resource SlbTemplateDynamicService

import (
	"context"
	"fmt"
	"log"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateDynamicService() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSlbTemplateDynamicServiceCreate,
		UpdateContext: resourceSlbTemplateDynamicServiceUpdate,
		ReadContext:   resourceSlbTemplateDynamicServiceRead,
		DeleteContext: resourceSlbTemplateDynamicServiceDelete,
		Schema: map[string]*schema.Schema{
			"user_tag": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"dns_server": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ipv4_dns_server": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"ipv6_dns_server": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceSlbTemplateDynamicServiceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating SlbTemplateDynamicService (Inside resourceSlbTemplateDynamicServiceCreate) ")
		name := d.Get("name").(string)
		data := dataToSlbTemplateDynamicService(d)
		logger.Println("[INFO] received formatted data from method data to SlbTemplateDynamicService --")
		d.SetId(name)
		err := go_thunder.PostSlbTemplateDynamicService(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSlbTemplateDynamicServiceRead(ctx, d, meta)

	}
	return diags
}

func resourceSlbTemplateDynamicServiceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading SlbTemplateDynamicService (Inside resourceSlbTemplateDynamicServiceRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_thunder.GetSlbTemplateDynamicService(client.Token, name, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return diags
	}
	return nil
}

func resourceSlbTemplateDynamicServiceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Modifying SlbTemplateDynamicService   (Inside resourceSlbTemplateDynamicServiceUpdate) ")
		name := d.Get("name").(string)
		data := dataToSlbTemplateDynamicService(d)
		logger.Println("[INFO] received formatted data from method data to SlbTemplateDynamicService ")
		d.SetId(name)
		err := go_thunder.PutSlbTemplateDynamicService(client.Token, name, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSlbTemplateDynamicServiceRead(ctx, d, meta)

	}
	return diags
}

func resourceSlbTemplateDynamicServiceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceSlbTemplateDynamicServiceDelete) " + name)
		err := go_thunder.DeleteSlbTemplateDynamicService(client.Token, name, client.Host)
		if err != nil {
			log.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name, err)
			return diags
		}
		d.SetId("")
		return nil
	}
	return nil
}

func dataToSlbTemplateDynamicService(d *schema.ResourceData) go_thunder.DynamicService {
	var vc go_thunder.DynamicService
	var c go_thunder.DynamicServiceInstance

	DNSServerCount := d.Get("dns_server.#").(int)
	c.Ipv6DNSServer = make([]go_thunder.DNSServer, 0, DNSServerCount)

	for i := 0; i < DNSServerCount; i++ {
		var obj1 go_thunder.DNSServer
		prefix := fmt.Sprintf("dns_server.%d.", i)
		obj1.Ipv6DNSServer = d.Get(prefix + "ipv6_dns_server").(string)
		obj1.Ipv4DNSServer = d.Get(prefix + "ipv4_dns_server").(string)
		c.Ipv6DNSServer = append(c.Ipv6DNSServer, obj1)
	}

	c.Name = d.Get("name").(string)
	c.UserTag = d.Get("user_tag").(string)

	vc.UUID = c
	return vc
}
