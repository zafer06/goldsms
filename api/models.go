package api

type (
	// LoginRes struct
	LoginRes struct {
		Status            string `json:"status"`
		Error             string `json:"error"`
		Result            string `json:"result"`
		Sonuc             bool   `json:"sonuc"`
		Mesaj             string `json:"mesaj"`
		DictCount         string `json:"dictCount"`
		ID                string `json:"ID"`
		AdSoyad           string `json:"adSoyad"`
		User              string `json:"user"`
		Email             string `json:"email"`
		IsSuperAdmin      bool   `json:"isSuperAdmin"`
		IsCustomerAdmin   bool   `json:"isCustomerAdmin"`
		Token             string `json:"token"`
		TokenTime         string `json:"tokenTime"`
		CustomerID        string `json:"customer.ID"`
		CustomerFirmaName string `json:"customer.firmaName"`
	}

	// ChangePasswordRes struct
	ChangePasswordRes struct {
		Status string `json:"status"`
		Error  string `json:"error"`
		Result string `json:"result"`
	}

	// AlfanumerikRes struct
	AlfanumerikRes struct {
		Status string   `json:"status"`
		Error  string   `json:"error"`
		Result []string `json:"result"`
	}

	// SendSmsRes struct
	SendSmsRes struct {
		Sonuc          bool
		Mesaj          string        `json:"mesaj"`
		Status         string        `json:"status"`
		Error          string        `json:"error"`
		GonderilerAdet int           `json:"gonderilenAdet"`
		HataliAdet     int           `json:"hataliAdet"`
		HataliGsm      string        `json:"hataliGsm"`
		HataliTckNo    string        `json:"hataliTckNo"`
		Result         ResultSendSms `json:"result"`
	}

	// ResultSendSms struct
	ResultSendSms struct {
		MessageID string `json:"messageid"`
		Count     string `json:"count"`
		Kredi     string `json:"kredi"`
	}

	// ReportRes struct
	ReportRes struct {
		Status    string          `json:"status"`
		Error     string          `json:"error"`
		Result    string          `json:"result"`
		Sonuc     bool            `json:"sonuc"`
		Mesaj     string          `json:"mesaj"`
		ListCount string          `json:"listCount"`
		List      []ListReportRes `json:"list"`
	}

	// ListReportRes struct
	ListReportRes struct {
		Index             string `json:"i"`
		Alfa              string `json:"alfa"`
		GUID              string `json:"guid"`
		Zaman             string `json:"zaman"`
		IsScheduled       string `json:"isScheduled"`
		IsCanceled        string `json:"isCanceled"`
		SmsAdet           string `json:"smsAdet"`
		Uzunluk           string `json:"uzunluk"`
		Iletildi          string `json:"iletildi"`
		Bekleyen          string `json:"bekleyen"`
		Iletilmedi        string `json:"iletilmedi"`
		Mesaj             string `json:"mesaj"`
		KarakterAdet      string `json:"karakterAdet"`
		MesajIptalEdilmis string `json:"mesajIptalEdilmis"`
		IsLoadedProgress  string `json:"isLoadedProgress"`
		GecerlilikSuresi  string `json:"gecerlilikSuresi"`
	}

	// ReportByIDRes struct
	ReportByIDRes struct {
		Status string                `json:"status"`
		Error  string                `json:"error"`
		Result []ResultReportByIDRes `json:"result"`
	}

	// ResultReportByIDRes struct
	ResultReportByIDRes struct {
		Gsm      string `json:"gsm"`
		ReportID string `json:"reportid"`
		Status   string `json:"status"`
	}

	// KrediRes struct
	KrediRes struct {
		Status string `json:"status"`
		Error  string `json:"error"`
		Result string `json:"result"`
	}

	// MultiSendRes struct
	MultiSendRes struct {
		Sonuc          string             `json:"sonuc"`
		Mesaj          string             `json:"mesaj"`
		Status         string             `json:"status"`
		Error          string             `json:"error"`
		GonderilenAdet string             `json:"gonderilenAdet"`
		HataliAdet     string             `json:"hataliAdet"`
		HataliGsm      string             `json:"hataliGsm"`
		HataliTckNo    string             `json:"hataliTckNo"`
		Result         ResultMultiSendRes `json:"result"`
	}

	// ResultMultiSendRes struct
	ResultMultiSendRes struct {
		MessageID string `json:"array"`
		Count     int    `json:"count"`
		Kredi     int    `json:"kredi"`
	}
)
