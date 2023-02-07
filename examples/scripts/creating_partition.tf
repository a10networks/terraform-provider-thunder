provider "thunder" {
  address  = "${var.address}"
  username = "${var.username}"
  password = "${var.password}"
  #partition = "shared"  # default partition is shared 
}


resource "thunder_partition" "test" {
	user_tag = "tag1"
	partition_name = "newpartition"
	id1 = 16
	application_type = "adc"
}