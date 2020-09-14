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

package pdfutil ;import (_a "github.com/unidoc/unipdf/v3/common";_b "github.com/unidoc/unipdf/v3/contentstream";_ge "github.com/unidoc/unipdf/v3/contentstream/draw";_g "github.com/unidoc/unipdf/v3/model";);

// NormalizePage rotates the contents of the passed in page according to its
// Rotate entry (i.e. flattens the rotation). If rotation is applied, the
// Rotate entry of the page is set to nil.
// After normalization, the page should look the same if openend using a
// PDF viewer.
func NormalizePage (page *_g .PdfPage )error {if _ae :=page .Rotate ;_ae ==nil ||*_ae ==0||*_ae %90!=0{return nil ;};_gg ,_f :=page .GetMediaBox ();if _f !=nil {return _f ;};_e ,_f :=page .GetContentStreams ();if _f !=nil {return _f ;};_gc ,_ff :=_gg .Width (),_gg .Height ();_ag :=-float64 (*page .Rotate );_cc :=_ge .Path {Points :[]_ge .Point {_ge .NewPoint (0,0).Rotate (_ag ),_ge .NewPoint (_gc ,0).Rotate (_ag ),_ge .NewPoint (0,_ff ).Rotate (_ag ),_ge .NewPoint (_gc ,_ff ).Rotate (_ag )}}.GetBoundingBox ();_bb :=-_gg .Llx +(_cc .Width -_gc )/2+_gc /2;_ef :=_gg .Lly +(_cc .Height -_ff )/2+_ff /2;_d :=_b .NewContentCreator ();_d .Translate (_bb ,_ef );_d .RotateDeg (_ag );_d .Translate (-_gc /2,-_ff /2);_fg :=_d .Operations ().String ();*_gg =_g .PdfRectangle {Urx :_cc .Width ,Ury :_cc .Height };_e =append ([]string {_fg },_e ...);if _f =page .SetContentStreams (_e ,nil );_f !=nil {return _f ;};_a .Log .Debug ("\u0052o\u0074\u0061\u0074\u0065\u003d\u0025\u0066\u00b0\u0020\u0025\u0071 \u006d\u0062\u006f\u0078\u003d\u0025\u002e\u0032\u0066",_ag ,_fg ,*_gg );page .Rotate =nil ;return nil ;};