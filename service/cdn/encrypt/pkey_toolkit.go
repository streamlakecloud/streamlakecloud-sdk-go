package encrypt

const PKEY = "pkey"

func buildPkeyParameter(request *PkeyRequest) map[string]string {
	parameters := make(map[string]string)
	pkey, err := NewDefaultPkeyFormat().format(request)
	if err != nil {
		panic("failed to generate pkey")
	}
	parameters[PKEY] = pkey
	return parameters
}
