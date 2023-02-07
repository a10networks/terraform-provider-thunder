package thunder

//Thunder resource SlbTemplateSNMP

import (
	"context"
	"log"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateSNMP() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSlbTemplateSNMPCreate,
		UpdateContext: resourceSlbTemplateSNMPUpdate,
		ReadContext:   resourceSlbTemplateSNMPRead,
		DeleteContext: resourceSlbTemplateSNMPDelete,
		Schema: map[string]*schema.Schema{
			"security_level": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"context_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"snmp_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"priv_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"oid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"context_engine_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"community": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"interface": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"version": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"security_engine_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"user_tag": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"priv_proto": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"port": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"auth_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"auth_proto": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"host": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"interval": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"username": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceSlbTemplateSNMPCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating SlbTemplateSNMP (Inside resourceSlbTemplateSNMPCreate) ")
		snmp_name := d.Get("snmp_name").(string)
		data := dataToSlbTemplateSNMP(d)
		logger.Println("[INFO] received formatted data from method data to SlbTemplateSNMP --")
		d.SetId(snmp_name)
		err := go_thunder.PostSlbTemplateSNMP(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSlbTemplateSNMPRead(ctx, d, meta)

	}
	return diags
}

func resourceSlbTemplateSNMPRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading SlbTemplateSNMP (Inside resourceSlbTemplateSNMPRead)")

	if client.Host != "" {
		snmp_name := d.Id()
		logger.Println("[INFO] Fetching service Read" + snmp_name)
		data, err := go_thunder.GetSlbTemplateSNMP(client.Token, snmp_name, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		if data == nil {
			logger.Println("[INFO] No data found " + snmp_name)
			d.SetId("")
			return nil
		}
		return diags
	}
	return nil
}

func resourceSlbTemplateSNMPUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Modifying SlbTemplateSNMP   (Inside resourceSlbTemplateSNMPUpdate) ")
		snmp_name := d.Get("snmp_name").(string)
		data := dataToSlbTemplateSNMP(d)
		logger.Println("[INFO] received formatted data from method data to SlbTemplateSNMP ")
		d.SetId(snmp_name)
		err := go_thunder.PutSlbTemplateSNMP(client.Token, snmp_name, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSlbTemplateSNMPRead(ctx, d, meta)

	}
	return diags
}

func resourceSlbTemplateSNMPDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		snmp_name := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceSlbTemplateSNMPDelete) " + snmp_name)
		err := go_thunder.DeleteSlbTemplateSNMP(client.Token, snmp_name, client.Host)
		if err != nil {
			log.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", snmp_name, err)
			return diags
		}
		d.SetId("")
		return nil
	}
	return nil
}

func dataToSlbTemplateSNMP(d *schema.ResourceData) go_thunder.SNMP {
	var vc go_thunder.SNMP
	var c go_thunder.SNMPInstance

	c.Username = d.Get("username").(string)
	c.Oid = d.Get("oid").(string)
	c.PrivProto = d.Get("priv_proto").(string)
	c.ContextName = d.Get("context_name").(string)
	c.AuthKey = d.Get("auth_key").(string)
	c.Interval = d.Get("interval").(int)
	c.ContextEngineID = d.Get("context_engine_id").(string)
	c.SecurityLevel = d.Get("security_level").(string)
	c.Community = d.Get("community").(string)
	c.AuthProto = d.Get("auth_proto").(string)
	c.Host = d.Get("host").(string)
	c.Version = d.Get("version").(string)
	c.UserTag = d.Get("user_tag").(string)
	c.Interface = d.Get("interface").(int)
	c.PrivKey = d.Get("priv_key").(string)
	c.SecurityEngineID = d.Get("security_engine_id").(string)
	c.Port = d.Get("port").(int)
	c.SnmpName = d.Get("snmp_name").(string)
	vc.UUID = c
	return vc
}
