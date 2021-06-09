package kycaml

type KycAml struct {
}

func New() KycAml {
	k := KycAml{}

	return k
}

func (k *KycAml) InitMsg() string {
	res := "KycAml starting..."

	return res
}
