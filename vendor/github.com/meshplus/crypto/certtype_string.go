// Code generated by "stringer -type CertType -linecomment"; DO NOT EDIT.

package crypto

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ECert-0]
	_ = x[RCert-1]
	_ = x[SDKCert-2]
	_ = x[TCert-3]
	_ = x[ERCert-4]
	_ = x[IDCert-5]
	_ = x[RAWPub-6]
	_ = x[UnknownCertType-7]
}

const _CertType_name = "ecertrcertsdkcerttcertercertidcertrawpubunknown_cert_type"

var _CertType_index = [...]uint8{0, 5, 10, 17, 22, 28, 34, 40, 57}

func (i CertType) String() string {
	if i < 0 || i >= CertType(len(_CertType_index)-1) {
		return "CertType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _CertType_name[_CertType_index[i]:_CertType_index[i+1]]
}