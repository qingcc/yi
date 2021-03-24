package utils

import (
	js "encoding/json"
	"testing"
)

var str = `{"auditData":{"processTime":"816","timestamp":"2021-02-24 22:49:57.692","requestHost":"198.245.173.4, 34.120.75.158, 130.211.1.144, 10.197.2.79","serverId":"ip-10-185-10-60.eu-west-1.compute.internal#A+","environment":"[awseuwest1, awseuwest1a, ip_10_185_10_60]","release":"","token":"BF90F2F2C8D04C51B55CCC6B61AE3A0B","internal":"0|06~A-SIC~24d40~1278844969~N~~~NOR~76BDC17A786247C161420339604304AWBR0000001000000000524d40|BR|05|1|1|||||||AT_WEB||||R|1|3|~1~2~0|0|0||0|cg9xqnjx8t8dx8ceb6yq2qtn||||"},"booking":{"reference":"258-1494511","clientReference":"11178696-872547","creationDate":"2021-02-24","status":"CONFIRMED","modificationPolicies":{"cancellation":true,"modification":true},"creationUser":"cg9xqnjx8t8dx8ceb6yq2qtn","holder":{"name":"GABRIELE","surname":"CARNEIRO DE LIMA"},"hotel":{"checkOut":"2021-03-08","checkIn":"2021-03-05","code":162670,"name":"Americas Granada Hotel","categoryCode":"4EST","categoryName":"4 STARS","destinationCode":"RIO","destinationName":"Rio de Janeiro","zoneCode":1,"zoneName":"Rio de Janeiro - Centro","latitude":"-22.911513","longitude":"-43.183629","rooms":[{"status":"CONFIRMED","id":1,"code":"TWN.ST","name":"TWIN STANDARD","paxes":[{"roomId":1,"type":"AD","name":"Gabriele","surname":"Carneiro de lima"},{"roomId":1,"type":"AD"}],"rates":[{"rateClass":"NOR","net":"64.77","rateComments":"Car park NO. Check-in hour 14:00 - . Following the Statute of Children and Adolescents (Law Nº. 8.069/90), the accommodation of people under 18 unaccompanied or with their guardians will be accepted only with written authorization of their parents on a notarized document. Parenthood must be proved through ID or Birth Certificate. (22/03/2015-31/12/2040) As a result of local government measures and guidelines put in place by services providers – including hotels and ancillaries – guests may find that some facilities or services are not available.Please visit https://static-sources.s3-eu-west-1.amazonaws.com/policy/index.html for further information (18/05/2020-30/04/2021).","paymentType":"AT_WEB","packaging":false,"boardCode":"BB","boardName":"BED AND BREAKFAST","cancellationPolicies":[{"amount":"21.59","from":"2021-03-02T23:59:00-03:00"}],"rooms":1,"adults":2,"children":0}]}],"totalNet":"64.77","currency":"USD","supplier":{"name":"ADVANTOS BRASIL OPERADORA DE TURISMO, LTDA","vatNumber":"16.847.249/0001-77"}},"invoiceCompany":{"code":"SG1","company":"HOTELBEDS PTE. LTD","registrationNumber":"M2-0084578-1"},"totalNet":64.77,"pendingAmount":64.77,"currency":"USD"}}`

// region type
type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type AuditData struct {
	ProcessTime string `json:"processTime"`
	Timestamp   string `json:"time"`
	RequestHost string `json:"requestHost"`
	ServerId    string `json:"serverId"`
	Environment string `json:"environment"`
	Release     string `json:"release"`
	Token       string `json:"token"`
	Internal    string `json:"internal"`
}

type Response struct {
	AuditData
	Booking Booking `json:"booking"`
	Error   Error   `json:"error"`
}

type Holder struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
}
type Booking struct {
	Reference            string               `json:"reference"`
	ClientReference      string               `json:"clientReference"`
	CreationDate         string               `json:"creationDate"`
	Status               string               `json:"status"`
	CreationUser         string               `json:"creationUser"`
	TotalNet             float64              `json:"totalNet"`
	PendingAmount        float64              `json:"pendingAmount"`
	Currency             string               `json:"currency"`
	ModificationPolicies ModificationPolicies `json:"modificationPolicies"`
	Holder               Holder               `json:"holder"`
	Hotel                Hotel                `json:"hotel"`
	InvoiceCompany       InvoiceCompany       `json:"invoiceCompany"`
}

type InvoiceCompany struct {
	Code               string `json:"code"`
	Name               string `json:"name"`
	RegistrationNumber string `json:"registrationNumber"`
}

type Hotel struct {
	CheckIn          string    `json:"checkIn"`
	CheckOut         string    `json:"checkOut"`
	Name             string    `json:"name"`
	Code             int       `json:"code"`
	CategoryCode     string    `json:"categoryCode"`
	CategoryName     string    `json:"categoryName"`
	DestinationCode  string    `json:"destinationCode"`
	DestinationName  string    `json:"destinationName"`
	ZoneCode         int       `json:"zoneCode"`
	ZoneName         string    `json:"zoneName"`
	Latitude         string    `json:"latitude"`
	Longitude        string    `json:"longitude"`
	Currency         string    `json:"currency"`
	TotalNet         string    `json:"totalNet"`
	TotalSellingRate float64   `json:"totalSellingRate"`
	Rooms            []ResRoom `json:"rooms"`
	Supplier         Supplier  `json:"supplier"`
}

type Supplier struct {
	Name      string `json:"name"`
	VatNumber string `json:"vatNumber"`
}

type ResRoom struct {
	Code              string `json:"code"`
	Name              string `json:"name"`
	SupplierReference string `json:"supplierReference"`
	Paxes             []Pax  `json:"paxes"`
	Rates             []Rate `json:"rates"`
}

type Pax struct {
	RoomId  int    `json:"roomId"`
	Type    string `json:"type"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Age     int    `json:"age,omitempty"`
}

type Rate struct {
	RateClass            string               `json:"rateClass"`
	Net                  string               `json:"net"`
	RateComments         string               `json:"rateComments"`
	PaymentType          string               `json:"paymentType"`
	Packaging            bool                 `json:"packaging"`
	BoardCode            string               `json:"boardCode"`
	BoardName            string               `json:"boardName"`
	SellingPrice         string               `json:"sellingPrice"`
	Comission            string               `json:"comission"`
	HotelSellingRate     string               `json:"hotelSellingRate"`
	HotelCurrency        string               `json:"hotelCurrency"`
	HotelMandatory       bool                 `json:"hotelMandatory"`
	Rooms                int                  `json:"rooms"`
	Adults               int                  `json:"adults"`
	Children             int                  `json:"children"`
	CancellationPolicies []CancellationPolicy `json:"cancellationPolicies"`
	RateBreakDown        RateBreakDown        `json:"rateBreakDown"`
	Taxes                Taxes                `json:"taxes"`
}

type Taxes struct {
	AllIncluded bool  `json:"allIncluded"`
	Taxes       []Tax `json:"taxes"`
}

type Tax struct {
	Included       bool   `json:"Included"`
	Percent        string `json:"percent"`
	Amount         string `json:"amount"`
	Currency       string `json:"currency"`
	Type           string `json:"type"`
	ClientAmount   string `json:"clientAmount"`
	ClientCurrency string `json:"clientCurrency"`
}

type RateBreakDown struct {
	AgComission     string           `json:"agComission"`
	ComissionPct    string           `json:"comissionPct"`
	ComissionVat    string           `json:"comissionVat"`
	RateDiscounts   []RateDiscount   `json:"rateDiscounts"`
	RateSupplements []RateSupplement `json:"rateSupplement"`
}

type RateSupplement struct {
	Name      string `json:"name"`
	Amount    string `json:"amount"`
	Code      int    `json:"code"`
	From      string `json:"name"`
	To        string `json:"to"`
	Nights    int    `json:"nights"`
	PaxNumber int    `json:"paxNumber"`
	PaxType   string `json:"paxType"`
}

type RateDiscount struct {
	Amount string `json:"amount"`
	Code   string `json:"code"`
	Name   string `json:"name"`
}

type CancellationPolicy struct {
	Amount        string `json:"amount"`
	From          string `json:"from"`
	HotelAmount   string `json:"hotelAmount"`
	HotelCurrency string `json:"hotelCurrency"`
}

type ModificationPolicies struct {
	Cancellation bool `json:"cancellation"`
	Modification bool `json:"modification"`
}

// endregion

func BenchmarkToJson(b *testing.B) {
	r := new(Response)
	for n := 0; n < b.N; n++ {
		r = new(Response)
		if err := JsonUnmarshal([]byte(str), r); err != nil {
			b.Errorf("unmarshal failed, err:%s", err.Error())
		}
	}
	b.StopTimer()
	if r.Booking.Status == "" {
		b.Errorf("r.Booking.Status is 0")
	}
}

func BenchmarkToJson2(b *testing.B) {
	r := new(Response)
	for n := 0; n < b.N; n++ {
		r = new(Response)
		if err := js.Unmarshal([]byte(str), r); err != nil {
			b.Errorf("unmarshal failed, err:%s", err.Error())
		}
	}
	b.StopTimer()
	if r.Booking.Status == "" {
		b.Errorf("r.Booking.Status is 0")
	}
}
