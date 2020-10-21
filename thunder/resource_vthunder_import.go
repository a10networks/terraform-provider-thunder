package thunder

import (
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"util"
)

func resourceImport() *schema.Resource {
	return &schema.Resource{
		Create: resourceImportCreate,
		Update: resourceImportUpdate,
		Read:   resourceImportRead,
		Delete: resourceImportDelete,

		Schema: map[string]*schema.Schema{
			"ssl_cert": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"usb_license": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"ca_cert": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"store_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"dnssec_dnskey": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"ip_map_list": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"glm_cert": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"health_postfile": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"use_mgmt_port": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"postfilename": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"password": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"overwrite": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"remote_file": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"remote_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
				Computed:    true,
			},
			"web_category_license": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"geo_location": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
				Computed:    true,
			},
			"pfx_password": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"wsdl": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"user_tag": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"dnssec_ds": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"file_inspection_bw_list": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"local_uri_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"auth_portal_image": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"auth_saml_idp": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"use_mgmt_port": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"password": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"saml_idp_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"verify_xml_signature": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"overwrite": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"remote_file": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"aflex": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"overwrite": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
				Computed:    true,
			},
			"thales_kmdata": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"policy": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"ssl_cert_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"cloud_config": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
				Computed:    true,
			},
			"to_device": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"use_mgmt_port": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"web_category_license": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"glm_license": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"glm_cert": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"device": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"overwrite": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"remote_file": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"class_list_convert": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"cloud_creds": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
				Computed:    true,
			},
			"terminal": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"thales_secworld": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"store": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"create": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"delete": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"remote_file": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"auth_portal": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"ssl_crl": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"certificate_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"health_external": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"use_mgmt_port": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"password": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"description": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"externalfilename": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"overwrite": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"remote_file": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"ssl_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
				Computed:    true,
			},
			"use_mgmt_port": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
				Computed:    true,
			},
			"lw_4o6": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"glm_license": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"auth_jwks": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"bw_list": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"class_list_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"secured": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"class_list": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"xml_schema": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}

}

func resourceImportCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	logger.Println("[INFO] Creating Import (Inside resourceImportCreate)")

	if client.Host != "" {
		name := d.Get("remote_file").(string)
		vc := dataToImport(d)
		d.SetId(name)
		go_thunder.PostImport(client.Token, vc, client.Host)

		return resourceImportRead(d, meta)
	}
	return nil
}

func resourceImportRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	logger.Println("[INFO] Reading Import (Inside resourceImportRead)")

	client := meta.(Thunder)

	if client.Host != "" {

		name := d.Id()

		logger.Println("[INFO] Fetching Import Read - " + name)

		vc, err := go_thunder.GetImport(client.Token, client.Host)

		if vc == nil {
			logger.Println("[INFO] No Import found")
			//d.SetId("")
			return nil
		}

		return err
	}
	return nil
}

func resourceImportUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceImportRead(d, meta)
}

func resourceImportDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceImportRead(d, meta)
}

//Utility method to instantiate Import Structure
//Configuration for basic required parameters
//TODO: Conf. for all the parameters

func dataToImport(d *schema.ResourceData) go_thunder.Import {
	var vc go_thunder.Import

	var c go_thunder.ImportInstance

	c.CloudCreds = d.Get("cloud_creds").(string)
	c.CloudConfig = d.Get("cloud_config").(string)
	c.UseMgmtPort = d.Get("use_mgmt_port").(int)
	c.SslKey = d.Get("ssl_key").(string)
	c.RemoteFile = d.Get("remote_file").(string)
	c.Password = d.Get("password").(string)
	c.Overwrite = d.Get("overwrite").(int)

	c.Aflex = d.Get("aflex").(string)



	vc.GeoLocation = c

	return vc
}
