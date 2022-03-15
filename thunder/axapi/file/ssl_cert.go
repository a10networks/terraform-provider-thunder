package file

//pseudo endpoint for upload and delete SSL cert files

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	endpnt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi/oper"
)

type SslCert struct {
	Name            string
	Protocol        string
	Host            string
	Path            string
	Username        string
	Password        string
	UseMgmtPort     int
	Overwrite       int
	CertificateType string
	PfxPassword     string
	Secured         int
}

func (p *SslCert) GetId() string {
	return p.Name
}

func (p *SslCert) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("file.SslCert::Post")

	realInst := endpnt.Import{}
	realInst.Inst.SslCert = p.Name
	realInst.Inst.UseMgmtPort = p.UseMgmtPort
	realInst.Inst.Overwrite = p.Overwrite
	realInst.Inst.Secured = p.Secured
	realInst.Inst.CertificateType = p.CertificateType
	if len(p.Password) > 0 {
		realInst.Inst.Password = p.Password
	}
	if len(p.PfxPassword) > 0 {
		realInst.Inst.PfxPassword = p.PfxPassword
	}
	realInst.Inst.RemoteFile = axapi.GenUrlForDownload(p.Protocol, p.Host, p.Path, p.Username)

	return realInst.Post(authToken, host, logger)
}

func (p *SslCert) Get(authToken string, host, instId string, logger *axapi.ThunderLog) error {
	logger.Println("file.SslCert::Get")
	err := oper.FileSslCertGet(authToken, host, instId, logger)
	return err
}

func (p *SslCert) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("file.SslCert::Delet")
	obj := endpnt.PkiDelete{}
	obj.Inst.CertName = instId
	if p.CertificateType == "pfx" { //pfx=pkcs12 constains cert and key both
		logger.Println("Also delete the key of this pfx")
		obj.Inst.PrivateKey = instId
	}
	return obj.Post(authToken, host, logger)
}
