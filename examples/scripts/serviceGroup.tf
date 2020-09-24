provider "vthunder" {
  address = ""
  username = ""
  password = ""
}

resource "vthunder_service_group" "sg3" {
name="SG3"
protocol=""
member_list {
		name="rs1"
		port=
	}
member_list {
		name="rs1"
		port=
		}
}

resource "vthunder_slb_template_client_ssl" "client_ssl1" {
	name = "testclientssl"
	user_tag = "test_tag"
	forward_proxy_ssl_version = 33
	forward_proxy_crl_disable = 1
	fp_cert_ext_aia_ocsp = "ocp" 
	ca_certs {
		ca_cert = "file_name"	
	}
}