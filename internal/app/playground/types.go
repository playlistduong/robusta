package playground

type (
	Request struct {
		Code string `json:"code"`
	}

	Response struct {
		Code   int    `json:"code"`
		Errors string `json:"errors"`
		Events []struct {
			Message string `json:"message"`
			Kind    string `json:"kind"`
			Delay   int64  `json:"delay"`
		} `json:"events"`
		Status      int  `json:"status"`
		IsTest      bool `json:"is_test"`
		TestsFailed int  `json:"tests_failed"`
	}
)
