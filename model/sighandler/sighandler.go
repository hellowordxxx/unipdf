//
// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// This is a commercial product and requires a license to operate.
// A trial license can be obtained at https://unidoc.io
//
// DO NOT EDIT: generated by unitwist Go source code obfuscator.
//
// Use of this source code is governed by the UniDoc End User License Agreement
// terms that can be accessed at https://unidoc.io/eula/

// Package sighandler implements digital signature handlers for PDF signature validation and signing.
package sighandler ;import (_b "bytes";_ce "crypto";_cea "crypto/rand";_fbe "crypto/rsa";_fd "crypto/x509";_be "crypto/x509/pkix";_gg "encoding/asn1";_gc "errors";_fe "fmt";_d "github.com/unidoc/pkcs7";_fg "github.com/unidoc/timestamp";_a "github.com/unidoc/unipdf/v3/core";_ec "github.com/unidoc/unipdf/v3/model";_fc "hash";_g "io";_e "io/ioutil";_fb "net/http";_c "time";);

// NewAdobeX509RSASHA1 creates a new Adobe.PPKMS/Adobe.PPKLite adbe.x509.rsa_sha1 signature handler.
// Both parameters may be nil for the signature validation.
func NewAdobeX509RSASHA1 (privateKey *_fbe .PrivateKey ,certificate *_fd .Certificate )(_ec .SignatureHandler ,error ){return &adobeX509RSASHA1 {_cbf :certificate ,_ccd :privateKey },nil ;};type timestampInfo struct{Version int ;Policy _gg .RawValue ;MessageImprint struct{HashAlgorithm _be .AlgorithmIdentifier ;HashedMessage []byte ;};SerialNumber _gg .RawValue ;GeneralizedTime _c .Time ;};type docTimeStamp struct{_dcf string ;_eff _ce .Hash ;};

// Sign sets the Contents fields for the PdfSignature.
func (_dbcf *adobeX509RSASHA1 )Sign (sig *_ec .PdfSignature ,digest _ec .Hasher )error {var _de []byte ;var _gac error ;if _dbcf ._gf !=nil {_de ,_gac =_dbcf ._gf (sig ,digest );if _gac !=nil {return _gac ;};}else {_ege ,_gfd :=digest .(_fc .Hash );if !_gfd {return _gc .New ("\u0068a\u0073h\u0020\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072");};_bgc ,_ :=_cfd (_dbcf ._cbf .SignatureAlgorithm );_de ,_gac =_fbe .SignPKCS1v15 (_cea .Reader ,_dbcf ._ccd ,_bgc ,_ege .Sum (nil ));if _gac !=nil {return _gac ;};};_de ,_gac =_gg .Marshal (_de );if _gac !=nil {return _gac ;};sig .Contents =_a .MakeHexString (string (_de ));return nil ;};

// SignFunc represents a custom signing function. The function should return
// the computed signature.
type SignFunc func (_cfe *_ec .PdfSignature ,_cae _ec .Hasher )([]byte ,error );

// NewDigest creates a new digest.
func (_edg *docTimeStamp )NewDigest (sig *_ec .PdfSignature )(_ec .Hasher ,error ){return _b .NewBuffer (nil ),nil ;};

// NewAdobeX509RSASHA1Custom creates a new Adobe.PPKMS/Adobe.PPKLite adbe.x509.rsa_sha1 signature handler
// with a custom signing function. Both parameters may be nil for the signature validation.
func NewAdobeX509RSASHA1Custom (certificate *_fd .Certificate ,signFunc SignFunc )(_ec .SignatureHandler ,error ){return &adobeX509RSASHA1 {_cbf :certificate ,_gf :signFunc },nil ;};func (_bd *adobeX509RSASHA1 )getCertificate (_fbb *_ec .PdfSignature )(*_fd .Certificate ,error ){if _bd ._cbf !=nil {return _bd ._cbf ,nil ;};var _gga []byte ;switch _cca :=_fbb .Cert .(type ){case *_a .PdfObjectString :_gga =_cca .Bytes ();case *_a .PdfObjectArray :if _cca .Len ()==0{return nil ,_gc .New ("\u006e\u006f\u0020s\u0069\u0067\u006e\u0061t\u0075\u0072\u0065\u0020\u0063\u0065\u0072t\u0069\u0066\u0069\u0063\u0061\u0074\u0065\u0073\u0020\u0066\u006f\u0075\u006e\u0064");};for _ ,_gb :=range _cca .Elements (){_cda ,_gce :=_a .GetString (_gb );if !_gce {return nil ,_fe .Errorf ("\u0069\u006ev\u0061\u006c\u0069\u0064\u0020\u0063\u0065\u0072\u0074\u0069\u0066\u0069\u0063\u0061\u0074\u0065\u0020\u006f\u0062j\u0065\u0063\u0074\u0020\u0074\u0079p\u0065\u0020\u0069\u006e\u0020\u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065 \u0063\u0065r\u0074\u0069\u0066\u0069c\u0061\u0074\u0065\u0020\u0063h\u0061\u0069\u006e\u003a\u0020\u0025\u0054",_gb );};_gga =append (_gga ,_cda .Bytes ()...);};default:return nil ,_fe .Errorf ("\u0069n\u0076\u0061l\u0069\u0064\u0020s\u0069\u0067\u006e\u0061\u0074\u0075\u0072e\u0020\u0063\u0065\u0072\u0074\u0069f\u0069\u0063\u0061\u0074\u0065\u0020\u006f\u0062\u006a\u0065\u0063t\u0020\u0074\u0079\u0070\u0065\u003a\u0020\u0025\u0054",_cca );};_eed ,_ced :=_fd .ParseCertificates (_gga );if _ced !=nil {return nil ,_ced ;};return _eed [0],nil ;};

// Sign sets the Contents fields.
func (_baab *adobePKCS7Detached )Sign (sig *_ec .PdfSignature ,digest _ec .Hasher )error {if _baab ._fcc {_ee :=_baab ._ae ;if _ee <=0{_ee =8192;};sig .Contents =_a .MakeHexString (string (make ([]byte ,_ee )));return nil ;};_dba :=digest .(*_b .Buffer );_ab ,_ecf :=_d .NewSignedData (_dba .Bytes ());if _ecf !=nil {return _ecf ;};if _fcd :=_ab .AddSigner (_baab ._df ,_baab ._ac ,_d .SignerInfoConfig {});_fcd !=nil {return _fcd ;};_ab .Detach ();_bc ,_ecf :=_ab .Finish ();if _ecf !=nil {return _ecf ;};_cbb :=make ([]byte ,8192);copy (_cbb ,_bc );sig .Contents =_a .MakeHexString (string (_cbb ));return nil ;};

// NewDocTimeStamp creates a new DocTimeStamp signature handler.
// The timestampServerURL parameter can be empty string for the signature validation.
// The hashAlgorithm parameter can be crypto.SHA1, crypto.SHA256, crypto.SHA384, crypto.SHA512.
func NewDocTimeStamp (timestampServerURL string ,hashAlgorithm _ce .Hash )(_ec .SignatureHandler ,error ){return &docTimeStamp {_dcf :timestampServerURL ,_eff :hashAlgorithm },nil ;};

// IsApplicable returns true if the signature handler is applicable for the PdfSignature.
func (_dc *adobeX509RSASHA1 )IsApplicable (sig *_ec .PdfSignature )bool {if sig ==nil ||sig .Filter ==nil ||sig .SubFilter ==nil {return false ;};return (*sig .Filter =="A\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004d\u0053"||*sig .Filter =="\u0041\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004c\u0069\u0074\u0065")&&*sig .SubFilter =="\u0061d\u0062e\u002e\u0078\u0035\u0030\u0039.\u0072\u0073a\u005f\u0073\u0068\u0061\u0031";};

// Validate validates PdfSignature.
func (_ggb *docTimeStamp )Validate (sig *_ec .PdfSignature ,digest _ec .Hasher )(_ec .SignatureValidationResult ,error ){_bdb :=sig .Contents .Bytes ();_cge ,_gdge :=_d .Parse (_bdb );if _gdge !=nil {return _ec .SignatureValidationResult {},_gdge ;};if _gdge =_cge .Verify ();_gdge !=nil {return _ec .SignatureValidationResult {},_gdge ;};var _bcc timestampInfo ;_ ,_gdge =_gg .Unmarshal (_cge .Content ,&_bcc );if _gdge !=nil {return _ec .SignatureValidationResult {},_gdge ;};_gbd ,_gdge :=_gdg (_bcc .MessageImprint .HashAlgorithm .Algorithm );if _gdge !=nil {return _ec .SignatureValidationResult {},_gdge ;};_dd :=_gbd .New ();_ace :=digest .(*_b .Buffer );_dd .Write (_ace .Bytes ());_fgdf :=_dd .Sum (nil );_eab :=_ec .SignatureValidationResult {IsSigned :true ,IsVerified :_b .Equal (_fgdf ,_bcc .MessageImprint .HashedMessage ),GeneralizedTime :_bcc .GeneralizedTime };return _eab ,nil ;};

// InitSignature initialises the PdfSignature.
func (_ea *docTimeStamp )InitSignature (sig *_ec .PdfSignature )error {_eba :=*_ea ;sig .Handler =&_eba ;sig .Filter =_a .MakeName ("\u0041\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004c\u0069\u0074\u0065");sig .SubFilter =_a .MakeName ("\u0045\u0054\u0053I\u002e\u0052\u0046\u0043\u0033\u0031\u0036\u0031");sig .Reference =nil ;_bbe ,_afc :=_ea .NewDigest (sig );if _afc !=nil {return _afc ;};_bbe .Write ([]byte ("\u0063\u0061\u006c\u0063\u0075\u006ca\u0074\u0065\u0020\u0074\u0068\u0065\u0020\u0043\u006f\u006e\u0074\u0065\u006et\u0073\u0020\u0066\u0069\u0065\u006c\u0064 \u0073\u0069\u007a\u0065"));return _eba .Sign (sig ,_bbe );};

// NewAdobePKCS7Detached creates a new Adobe.PPKMS/Adobe.PPKLite adbe.pkcs7.detached signature handler.
// Both parameters may be nil for the signature validation.
func NewAdobePKCS7Detached (privateKey *_fbe .PrivateKey ,certificate *_fd .Certificate )(_ec .SignatureHandler ,error ){return &adobePKCS7Detached {_df :certificate ,_ac :privateKey },nil ;};func (_cb *adobePKCS7Detached )getCertificate (_dg *_ec .PdfSignature )(*_fd .Certificate ,error ){if _cb ._df !=nil {return _cb ._df ,nil ;};var _eg []byte ;switch _aed :=_dg .Cert .(type ){case *_a .PdfObjectString :_eg =_aed .Bytes ();case *_a .PdfObjectArray :if _aed .Len ()==0{return nil ,_gc .New ("\u006e\u006f\u0020s\u0069\u0067\u006e\u0061t\u0075\u0072\u0065\u0020\u0063\u0065\u0072t\u0069\u0066\u0069\u0063\u0061\u0074\u0065\u0073\u0020\u0066\u006f\u0075\u006e\u0064");};for _ ,_fee :=range _aed .Elements (){_db ,_ef :=_a .GetString (_fee );if !_ef {return nil ,_fe .Errorf ("\u0069\u006ev\u0061\u006c\u0069\u0064\u0020\u0063\u0065\u0072\u0074\u0069\u0066\u0069\u0063\u0061\u0074\u0065\u0020\u006f\u0062j\u0065\u0063\u0074\u0020\u0074\u0079p\u0065\u0020\u0069\u006e\u0020\u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065 \u0063\u0065r\u0074\u0069\u0066\u0069c\u0061\u0074\u0065\u0020\u0063h\u0061\u0069\u006e\u003a\u0020\u0025\u0054",_fee );};_eg =append (_eg ,_db .Bytes ()...);};default:return nil ,_fe .Errorf ("\u0069n\u0076\u0061l\u0069\u0064\u0020s\u0069\u0067\u006e\u0061\u0074\u0075\u0072e\u0020\u0063\u0065\u0072\u0074\u0069f\u0069\u0063\u0061\u0074\u0065\u0020\u006f\u0062\u006a\u0065\u0063t\u0020\u0074\u0079\u0070\u0065\u003a\u0020\u0025\u0054",_aed );};_bb ,_gde :=_fd .ParseCertificates (_eg );if _gde !=nil {return nil ,_gde ;};return _bb [0],nil ;};type adobeX509RSASHA1 struct{_ccd *_fbe .PrivateKey ;_cbf *_fd .Certificate ;_gf SignFunc ;};func _cfd (_cfdd _fd .SignatureAlgorithm )(_ce .Hash ,bool ){var _fgd _ce .Hash ;switch _cfdd {case _fd .SHA1WithRSA :_fgd =_ce .SHA1 ;case _fd .SHA256WithRSA :_fgd =_ce .SHA256 ;case _fd .SHA384WithRSA :_fgd =_ce .SHA384 ;case _fd .SHA512WithRSA :_fgd =_ce .SHA512 ;default:return _ce .SHA1 ,false ;};return _fgd ,true ;};func (_ece *docTimeStamp )getCertificate (_gced *_ec .PdfSignature )(*_fd .Certificate ,error ){var _cbfc []byte ;switch _bgcg :=_gced .Cert .(type ){case *_a .PdfObjectString :_cbfc =_bgcg .Bytes ();case *_a .PdfObjectArray :if _bgcg .Len ()==0{return nil ,_gc .New ("\u006e\u006f\u0020s\u0069\u0067\u006e\u0061t\u0075\u0072\u0065\u0020\u0063\u0065\u0072t\u0069\u0066\u0069\u0063\u0061\u0074\u0065\u0073\u0020\u0066\u006f\u0075\u006e\u0064");};for _ ,_gfe :=range _bgcg .Elements (){_fde ,_bdc :=_a .GetString (_gfe );if !_bdc {return nil ,_fe .Errorf ("\u0069\u006ev\u0061\u006c\u0069\u0064\u0020\u0063\u0065\u0072\u0074\u0069\u0066\u0069\u0063\u0061\u0074\u0065\u0020\u006f\u0062j\u0065\u0063\u0074\u0020\u0074\u0079p\u0065\u0020\u0069\u006e\u0020\u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065 \u0063\u0065r\u0074\u0069\u0066\u0069c\u0061\u0074\u0065\u0020\u0063h\u0061\u0069\u006e\u003a\u0020\u0025\u0054",_gfe );};_cbfc =append (_cbfc ,_fde .Bytes ()...);};default:return nil ,_fe .Errorf ("\u0069n\u0076\u0061l\u0069\u0064\u0020s\u0069\u0067\u006e\u0061\u0074\u0075\u0072e\u0020\u0063\u0065\u0072\u0074\u0069f\u0069\u0063\u0061\u0074\u0065\u0020\u006f\u0062\u006a\u0065\u0063t\u0020\u0074\u0079\u0070\u0065\u003a\u0020\u0025\u0054",_bgcg );};_eee ,_cdg :=_fd .ParseCertificates (_cbfc );if _cdg !=nil {return nil ,_cdg ;};return _eee [0],nil ;};

// IsApplicable returns true if the signature handler is applicable for the PdfSignature
func (_cg *adobePKCS7Detached )IsApplicable (sig *_ec .PdfSignature )bool {if sig ==nil ||sig .Filter ==nil ||sig .SubFilter ==nil {return false ;};return (*sig .Filter =="A\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004d\u0053"||*sig .Filter =="\u0041\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004c\u0069\u0074\u0065")&&*sig .SubFilter =="\u0061\u0064\u0062\u0065.p\u006b\u0063\u0073\u0037\u002e\u0064\u0065\u0074\u0061\u0063\u0068\u0065\u0064";};type adobePKCS7Detached struct{_ac *_fbe .PrivateKey ;_df *_fd .Certificate ;_fcc bool ;_ae int ;};

// IsApplicable returns true if the signature handler is applicable for the PdfSignature.
func (_gag *docTimeStamp )IsApplicable (sig *_ec .PdfSignature )bool {if sig ==nil ||sig .Filter ==nil ||sig .SubFilter ==nil {return false ;};return (*sig .Filter =="A\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004d\u0053"||*sig .Filter =="\u0041\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004c\u0069\u0074\u0065")&&*sig .SubFilter =="\u0045\u0054\u0053I\u002e\u0052\u0046\u0043\u0033\u0031\u0036\u0031";};

// NewDigest creates a new digest.
func (_ga *adobeX509RSASHA1 )NewDigest (sig *_ec .PdfSignature )(_ec .Hasher ,error ){_dbc ,_fcg :=_ga .getCertificate (sig );if _fcg !=nil {return nil ,_fcg ;};_gdb ,_ :=_cfd (_dbc .SignatureAlgorithm );return _gdb .New (),nil ;};

// NewEmptyAdobePKCS7Detached creates a new Adobe.PPKMS/Adobe.PPKLite adbe.pkcs7.detached
// signature handler. The generated signature is empty and of size signatureLen.
// The signatureLen parameter can be 0 for the signature validation.
func NewEmptyAdobePKCS7Detached (signatureLen int )(_ec .SignatureHandler ,error ){return &adobePKCS7Detached {_fcc :true ,_ae :signatureLen },nil ;};

// Validate validates PdfSignature.
func (_cd *adobePKCS7Detached )Validate (sig *_ec .PdfSignature ,digest _ec .Hasher )(_ec .SignatureValidationResult ,error ){_cc :=sig .Contents .Bytes ();_ff ,_dbd :=_d .Parse (_cc );if _dbd !=nil {return _ec .SignatureValidationResult {},_dbd ;};_ca :=digest .(*_b .Buffer );_ff .Content =_ca .Bytes ();if _dbd =_ff .Verify ();_dbd !=nil {return _ec .SignatureValidationResult {},_dbd ;};return _ec .SignatureValidationResult {IsSigned :true ,IsVerified :true },nil ;};

// NewDigest creates a new digest.
func (_baa *adobePKCS7Detached )NewDigest (sig *_ec .PdfSignature )(_ec .Hasher ,error ){return _b .NewBuffer (nil ),nil ;};

// InitSignature initialises the PdfSignature.
func (_fccd *adobeX509RSASHA1 )InitSignature (sig *_ec .PdfSignature )error {if _fccd ._cbf ==nil {return _gc .New ("c\u0065\u0072\u0074\u0069\u0066\u0069c\u0061\u0074\u0065\u0020\u006d\u0075\u0073\u0074\u0020n\u006f\u0074\u0020b\u0065 \u006e\u0069\u006c");};if _fccd ._ccd ==nil &&_fccd ._gf ==nil {return _gc .New ("\u006d\u0075\u0073\u0074\u0020\u0070\u0072o\u0076\u0069\u0064e\u0020\u0065\u0069t\u0068\u0065r\u0020\u0061\u0020\u0070\u0072\u0069v\u0061te\u0020\u006b\u0065\u0079\u0020\u006f\u0072\u0020\u0061\u0020\u0073\u0069\u0067\u006e\u0069\u006e\u0067\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e");};_eb :=*_fccd ;sig .Handler =&_eb ;sig .Filter =_a .MakeName ("\u0041\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004c\u0069\u0074\u0065");sig .SubFilter =_a .MakeName ("\u0061d\u0062e\u002e\u0078\u0035\u0030\u0039.\u0072\u0073a\u005f\u0073\u0068\u0061\u0031");sig .Cert =_a .MakeString (string (_eb ._cbf .Raw ));sig .Reference =nil ;_ed ,_dbe :=_eb .NewDigest (sig );if _dbe !=nil {return _dbe ;};_ed .Write ([]byte ("\u0063\u0061\u006c\u0063\u0075\u006ca\u0074\u0065\u0020\u0074\u0068\u0065\u0020\u0043\u006f\u006e\u0074\u0065\u006et\u0073\u0020\u0066\u0069\u0065\u006c\u0064 \u0073\u0069\u007a\u0065"));return _eb .Sign (sig ,_ed );};

// Validate validates PdfSignature.
func (_fff *adobeX509RSASHA1 )Validate (sig *_ec .PdfSignature ,digest _ec .Hasher )(_ec .SignatureValidationResult ,error ){_aa ,_bab :=_fff .getCertificate (sig );if _bab !=nil {return _ec .SignatureValidationResult {},_bab ;};_da :=sig .Contents .Bytes ();var _af []byte ;if _ ,_ebc :=_gg .Unmarshal (_da ,&_af );_ebc !=nil {return _ec .SignatureValidationResult {},_ebc ;};_dgg ,_ccc :=digest .(_fc .Hash );if !_ccc {return _ec .SignatureValidationResult {},_gc .New ("\u0068a\u0073h\u0020\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072");};_bg ,_ :=_cfd (_aa .SignatureAlgorithm );if _fbf :=_fbe .VerifyPKCS1v15 (_aa .PublicKey .(*_fbe .PublicKey ),_bg ,_dgg .Sum (nil ),_af );_fbf !=nil {return _ec .SignatureValidationResult {},_fbf ;};return _ec .SignatureValidationResult {IsSigned :true ,IsVerified :true },nil ;};func _gdg (_geg _gg .ObjectIdentifier )(_ce .Hash ,error ){switch {case _geg .Equal (_d .OIDDigestAlgorithmSHA1 ),_geg .Equal (_d .OIDDigestAlgorithmECDSASHA1 ),_geg .Equal (_d .OIDDigestAlgorithmDSA ),_geg .Equal (_d .OIDDigestAlgorithmDSASHA1 ),_geg .Equal (_d .OIDEncryptionAlgorithmRSA ):return _ce .SHA1 ,nil ;case _geg .Equal (_d .OIDDigestAlgorithmSHA256 ),_geg .Equal (_d .OIDDigestAlgorithmECDSASHA256 ):return _ce .SHA256 ,nil ;case _geg .Equal (_d .OIDDigestAlgorithmSHA384 ),_geg .Equal (_d .OIDDigestAlgorithmECDSASHA384 ):return _ce .SHA384 ,nil ;case _geg .Equal (_d .OIDDigestAlgorithmSHA512 ),_geg .Equal (_d .OIDDigestAlgorithmECDSASHA512 ):return _ce .SHA512 ,nil ;};return _ce .Hash (0),_d .ErrUnsupportedAlgorithm ;};

// Sign sets the Contents fields for the PdfSignature.
func (_ffc *docTimeStamp )Sign (sig *_ec .PdfSignature ,digest _ec .Hasher )error {_fgc :=digest .(*_b .Buffer );_agb :=_ffc ._eff .New ();if _ ,_aee :=_g .Copy (_agb ,_fgc );_aee !=nil {return _aee ;};_dae :=_agb .Sum (nil );_fbc :=_fg .Request {HashAlgorithm :_ffc ._eff ,HashedMessage :_dae ,Certificates :true ,Extensions :nil ,ExtraExtensions :nil };_cega ,_caf :=_fbc .Marshal ();if _caf !=nil {return _caf ;};_daa ,_caf :=_fb .Post (_ffc ._dcf ,"a\u0070\u0070\u006c\u0069\u0063\u0061t\u0069\u006f\u006e\u002f\u0074\u0069\u006d\u0065\u0073t\u0061\u006d\u0070-\u0071u\u0065\u0072\u0079",_b .NewBuffer (_cega ));if _caf !=nil {return _caf ;};defer _daa .Body .Close ();_gfdf ,_caf :=_e .ReadAll (_daa .Body );if _caf !=nil {return _caf ;};if _daa .StatusCode !=_fb .StatusOK {return _fe .Errorf ("\u0068\u0074\u0074\u0070\u0020\u0073\u0074\u0061\u0074\u0075\u0073\u0020\u0063\u006f\u0064e\u0020n\u006f\u0074\u0020\u006f\u006b\u0020\u0028\u0067\u006f\u0074\u0020\u0025\u0064\u0029",_daa .StatusCode );};var _ad struct{Version _gg .RawValue ;Content _gg .RawValue ;};_ ,_caf =_gg .Unmarshal (_gfdf ,&_ad );if _caf !=nil {return _caf ;};sig .Contents =_a .MakeHexString (string (_ad .Content .FullBytes ));return nil ;};

// InitSignature initialises the PdfSignature.
func (_ba *adobePKCS7Detached )InitSignature (sig *_ec .PdfSignature )error {if !_ba ._fcc {if _ba ._df ==nil {return _gc .New ("c\u0065\u0072\u0074\u0069\u0066\u0069c\u0061\u0074\u0065\u0020\u006d\u0075\u0073\u0074\u0020n\u006f\u0074\u0020b\u0065 \u006e\u0069\u006c");};if _ba ._ac ==nil {return _gc .New ("\u0070\u0072\u0069\u0076\u0061\u0074\u0065\u004b\u0065\u0079\u0020m\u0075\u0073\u0074\u0020\u006e\u006f\u0074\u0020\u0062\u0065 \u006e\u0069\u006c");};};_cf :=*_ba ;sig .Handler =&_cf ;sig .Filter =_a .MakeName ("\u0041\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004c\u0069\u0074\u0065");sig .SubFilter =_a .MakeName ("\u0061\u0064\u0062\u0065.p\u006b\u0063\u0073\u0037\u002e\u0064\u0065\u0074\u0061\u0063\u0068\u0065\u0064");sig .Reference =nil ;_cfc ,_fbed :=_cf .NewDigest (sig );if _fbed !=nil {return _fbed ;};_cfc .Write ([]byte ("\u0063\u0061\u006c\u0063\u0075\u006ca\u0074\u0065\u0020\u0074\u0068\u0065\u0020\u0043\u006f\u006e\u0074\u0065\u006et\u0073\u0020\u0066\u0069\u0065\u006c\u0064 \u0073\u0069\u007a\u0065"));return _cf .Sign (sig ,_cfc );};