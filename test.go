package main

import (
	"context"
	"crypto/md5"
	"encoding/json"
	"flag"
	"fmt"
	lib "github.com/qingcc/yi/lib/workpool"
	"github.com/qingcc/yi/utils"
	"github.com/qingcc/yi/utils/rpcx/trans_server/service"
	"github.com/tealeg/xlsx"
	"io"
	"log"
	"net/url"
	"strings"
	"time"
)

var (
	addr = flag.String("addr", "localhost:10003", "server address")
)

const (
	AUTHKEY = "8x76w362uc9v6k2hm9tc"
)

func f() {
	flag.Parse()
	//res := service.Transfer("中文", "zh", "en", false)
	res := service.Trans("中文", "zh", "en", false)
	fmt.Println(res)
}

type CreateOrderRequest struct {
	GuestRemarks     string                  `json:"GuestRemarks"`     //非必填， 宾客特殊要求"Non-Smoking", "Smoking Room", "Higher Floor", "Lower Floor", "King-size Bed", "Twin bed"
	RateKey          string                  `json:"RateKey"`          //套餐唯一标识
	TotalPrice       float64                 `json:"TotalPrice"`       //套餐总价
	PartnerBookingID string                  `json:"PartnerBookingID"` //合作伙伴订单号，必须保证唯一
	RoomList         []ApiBookingRoomRequest `json:"RoomList"`         //入住客人信息，每间房间至少要填一个入住人信息
}

type ApiBookingRoomRequest struct {
	RoomID  int                        `json:"RoomID"`  //房间号，从1开始
	PaxList []ApiBookingRoomPaxRequest `json:"PaxList"` //客人信息
}

type ApiBookingRoomPaxRequest struct {
	LastName  string `json:"LastName"`
	FirstName string `json:"FirstName"`
	PaxType   int    `json:"PaxType"` //1成人，2儿童 默认成人
}

func init1() {
	//Partner=2019001113&RequestData=%7b%22HotelID%22%3a27889%2c%22CheckInDate%22%3a%222019-12-28%22%2c%22CheckOutDate%22%3a%222019-12-29%22%2c%22AdultCount%22%3a2%2c%22RoomCount%22%3a1%2c%22CountryCode%22%3a%22CN%22%7d
	// &Sign=2b10262f7111a9bc36553cd23f916ab8
	//Partner=1000000010&RequestData=%7b%22RateKey%22%3a%22JTAPI4b281a199d297c4ddf695e2dc89704da%22%2c%22TotalPrice%22%3a1156.0%2c%22GuestRemarks%22%3a%22ceshi%22%2c%22PartnerBookingID%22%3a%22637095076042452695%22%2c%22RoomList%22%3a%5b%7b%22RoomID%22%3a1%2c%22PaxList%22%3a%5b%7b%22LastName%22%3a%22jjjjj%22%2c%22FirstName%22%3a%22xxxx%22%2c%22PaxType%22%3a1%7d%2c%7b%22LastName%22%3a%22tab%22%2c%22FirstName%22%3a%22tab%22%2c%22PaxType%22%3a1%7d%5d%7d%5d%7d
	//&Sign=58ce1ba657ae20842b8a0224afb0fd47
	//	sig := "Partner=2019001113&RequestData=%7b%22HotelID%22%3a27889%2c%22CheckInDate%22%3a%222019-12-28%22%2c%22CheckOutDate%22%3a%222019-12-29%22%2c%22AdultCount%22%3a2%2c%22RoomCount%22%3a1%2c%22CountryCode%22%3a%22CN%22%7d"
	sig := "Partner=2019001113&Sign=c366f56b8c65984e7a1531c445483c31&RequestData=%7b%22HotelID%22%3a%2227889%22%2c%22CheckInDate%22%3a%222019-12-28+00%3a00%3a00%22%2c%22CheckOutDate%22%3a%222019-12-29+00%3a00%3a00%22%2c%22AdultCount%22%3a2%2c%22ChildrenCount%22%3a0%2c%22ChildAgeList%22%3anull%2c%22RoomCount%22%3a1%2c%22CountryCode%22%3a%22CN%22%7d"
	rd := "Partner=2019001113&RequestData=%7b%22HotelID%22%3a27889%2c%22CheckInDate%22%3a%222019-12-28%22%2c%22CheckOutDate%22%3a%222019-12-29%22%2c%22AdultCount%22%3a2%2c%22RoomCount%22%3a1%2c%22CountryCode%22%3a%22CN%22%7d"
	signMd5 := md5.New()
	_, err := io.WriteString(signMd5, rd+AUTHKEY)
	if err != nil {
		log.Printf("[ERROR]DAIEI common sign md5 failed:%s, source:%s", err.Error(), sig)
		return
	}
	sign := fmt.Sprintf("%x", signMd5.Sum(nil))

	si, _ := url.QueryUnescape(sig)
	fmt.Println(sig)
	fmt.Println(si)
	fmt.Println(sign)

	fmt.Println(url.QueryEscape(si))
}

func marshal(a interface{}) {
	fmt.Println()
	fmt.Println("---------------------------------")
	abyte, _ := json.Marshal(a)
	url.QueryEscape(string(abyte))
	fmt.Printf("%s", string(abyte))
}

var (
	hotelids = make(map[string]bool)
	dumIds   = make(map[string]int)
)

func main12() {
	//readfile("hotelinTG.xlsx")
	//	getHotelIds("hotelinTG.xlsx")
	//	getHotelIds("sellhotellist.xlsx")
	//
	//	str := ""
	//	for i, _ := range hotelids {
	//		str += "," + i
	//	}
	//	unStr := ""
	//	for id, _ := range dumIds {
	//		unStr += "," + id
	//	}
	//	fmt.Println(unStr)
	//	utils.Tracefile(str, "log.log")
	time.Sleep(time.Hour)
}

func getHotelIds(file string) {
	// 打开文件
	xlFile, err := xlsx.OpenFile(file)
	if err != nil {
		log.Println("打开文件失败！" + err.Error())
		return
	}
	if len(xlFile.Sheets) == 0 {
		log.Println("没有找到工作表")
	}
	for k, row := range xlFile.Sheets[0].Rows {
		if k == 0 {
			continue
		}
		id := row.Cells[0].String()
		if _, ok := hotelids[id]; !ok {
			hotelids[id] = true
		} else {
			if _, b := dumIds[id]; !b {
				dumIds[id] = 1
			} else {
				dumIds[id] = dumIds[id] + 1
			}
		}
	}
}

func readfile(file string) {
	// 打开文件
	xlFile, err := xlsx.OpenFile(file)
	if err != nil {
		log.Println("打开文件失败！" + err.Error())
		return
	}
	if len(xlFile.Sheets) == 0 {
		log.Println("没有找到工作表")
	}

	unMap := make(map[string]bool)
	leftMap := make(map[string]string)
	for k, row := range xlFile.Sheets[0].Rows {
		if k == 0 {
			continue
		}
		countryname := row.Cells[2].String()
		cen := strings.ToLower(row.Cells[3].String())
		if _, ok := unMap[countryname]; !ok {
			if _, b := Country2Code[cen]; !b {
				leftMap[countryname] = cen
				continue
			}
			str := "\"" + countryname + "\"" + ":\"" + Country2Code[cen] + "\","
			utils.Tracefile(str, "log.log")
			unMap[countryname] = true
		}
	}
	for cn, val := range leftMap {
		println(cn, " : ", val)
	}
	fmt.Println("\n\nimport success")
	//for country, _ := range unMap {
	//	fmt.Println("\""+country +"\""+":"+"\"\"")
	//}
	fmt.Println(len(unMap))

}

var Country2Code = map[string]string{
	"afghanistan":                      "AF",
	"albania":                          "AL",
	"algeria":                          "DZ",
	"american samoa":                   "AS",
	"andorra":                          "AD",
	"angola":                           "AO",
	"anguilla":                         "AI",
	"antarctica":                       "AQ",
	"antigua and barbuda":              "AG",
	"argentina":                        "AR",
	"armenia":                          "AM",
	"aruba":                            "AW",
	"australia":                        "AU",
	"austria":                          "AT",
	"azerbaijan":                       "AZ",
	"bahamas":                          "BS",
	"bahrain":                          "BH",
	"bangladesh":                       "BD",
	"barbados":                         "BB",
	"belarus":                          "BY",
	"belgium":                          "BE",
	"belize":                           "BZ",
	"benin":                            "BJ",
	"bermuda":                          "BM",
	"bhutan":                           "BT",
	"bolivia":                          "BO",
	"bosnia and herzegovina":           "BA",
	"botswana":                         "BW",
	"brazil":                           "BR",
	"british indian ocean territory":   "IO",
	"british virgin islands":           "VG",
	"brunei":                           "BN",
	"bulgaria":                         "BG",
	"burkina faso":                     "BF",
	"burundi":                          "BI",
	"cambodia":                         "KH",
	"cameroon":                         "CM",
	"canada":                           "CA",
	"cape verde":                       "CV",
	"cayman islands":                   "KY",
	"central african republic":         "CF",
	"chad":                             "TD",
	"chile":                            "CL",
	"china":                            "CN",
	"christmas island":                 "CX",
	"cocos islands":                    "CC",
	"colombia":                         "CO",
	"comoros":                          "KM",
	"cook islands":                     "CK",
	"costa rica":                       "CR",
	"croatia":                          "HR",
	"cuba":                             "CU",
	"curacao":                          "CW",
	"cyprus":                           "CY",
	"czech republic":                   "CZ",
	"democratic republic of the congo": "CD",
	"denmark":                          "DK",
	"djibouti":                         "DJ",
	"dominica":                         "DM",
	"dominican republic":               "DO",
	"east timor":                       "TL",
	"ecuador":                          "EC",
	"egypt":                            "EG",
	"el salvador":                      "SV",
	"equatorial guinea":                "GQ",
	"eritrea":                          "ER",
	"estonia":                          "EE",
	"ethiopia":                         "ET",
	"falkland islands":                 "FK",
	"faroe islands":                    "FO",
	"fiji":                             "FJ",
	"finland":                          "FI",
	"france":                           "FR",
	"french polynesia":                 "PF",
	"gabon":                            "GA",
	"gambia":                           "GM",
	"georgia":                          "GE",
	"germany":                          "DE",
	"ghana":                            "GH",
	"gibraltar":                        "GI",
	"greece":                           "GR",
	"greenland":                        "GL",
	"grenada":                          "GD",
	"guam":                             "GU",
	"guatemala":                        "GT",
	"guernsey":                         "GG",
	"guinea":                           "GN",
	"guinea-bissau":                    "GW",
	"guyana":                           "GY",
	"haiti":                            "HT",
	"honduras":                         "HN",
	"hong kong":                        "HK",
	"hungary":                          "HU",
	"iceland":                          "IS",
	"india":                            "IN",
	"indonesia":                        "ID",
	"iran":                             "IR",
	"iraq":                             "IQ",
	"ireland":                          "IE",
	"isle of man":                      "IM",
	"israel":                           "IL",
	"italy":                            "IT",
	"ivory coast":                      "CI",
	"jamaica":                          "JM",
	"japan":                            "JP",
	"jersey":                           "JE",
	"jordan":                           "JO",
	"kazakhstan":                       "KZ",
	"kenya":                            "KE",
	"kiribati":                         "KI",
	"kosovo":                           "XK",
	"kuwait":                           "KW",
	"kyrgyzstan":                       "KG",
	"laos":                             "LA",
	"latvia":                           "LV",
	"lebanon":                          "LB",
	"lesotho":                          "LS",
	"liberia":                          "LR",
	"libya":                            "LY",
	"liechtenstein":                    "LI",
	"lithuania":                        "LT",
	"luxembourg":                       "LU",
	"macau":                            "MO",
	"macedonia":                        "MK",
	"madagascar":                       "MG",
	"malawi":                           "MW",
	"malaysia":                         "MY",
	"maldives":                         "MV",
	"mali":                             "ML",
	"malta":                            "MT",
	"marshall islands":                 "MH",
	"mauritania":                       "MR",
	"mauritius":                        "MU",
	"mayotte":                          "YT",
	"mexico":                           "MX",
	"micronesia":                       "FM",
	"moldova":                          "MD",
	"monaco":                           "MC",
	"mongolia":                         "MN",
	"montenegro":                       "ME",
	"montserrat":                       "MS",
	"morocco":                          "MA",
	"mozambique":                       "MZ",
	"myanmar":                          "MM",
	"namibia":                          "NA",
	"nauru":                            "NR",
	"nepal":                            "NP",
	"netherlands":                      "NL",
	"netherlands antilles":             "AN",
	"new caledonia":                    "NC",
	"new zealand":                      "NZ",
	"nicaragua":                        "NI",
	"niger":                            "NE",
	"nigeria":                          "NG",
	"niue":                             "NU",
	"north korea":                      "KP",
	"northern mariana islands":         "MP",
	"norway":                           "NO",
	"oman":                             "OM",
	"pakistan":                         "PK",
	"palau":                            "PW",
	"palestine":                        "PS",
	"panama":                           "PA",
	"papua new guinea":                 "PG",
	"paraguay":                         "PY",
	"peru":                             "PE",
	"philippines":                      "PH",
	"pitcairn":                         "PN",
	"poland":                           "PL",
	"portugal":                         "PT",
	"puerto rico":                      "PR",
	"qatar":                            "QA",
	"republic of the congo":            "CG",
	"reunion":                          "RE",
	"romania":                          "RO",
	"russia":                           "RU",
	"rwanda":                           "RW",
	"saint barthelemy":                 "BL",
	"saint helena":                     "SH",
	"saint kitts and nevis":            "KN",
	"saint lucia":                      "LC",
	"saint martin":                     "MF",
	"saint pierre and miquelon":        "PM",
	"saint vincent and the grenadines": "VC",
	"samoa":                            "WS",
	"san marino":                       "SM",
	"sao tome and principe":            "ST",
	"saudi arabia":                     "SA",
	"senegal":                          "SN",
	"serbia":                           "RS",
	"seychelles":                       "SC",
	"sierra leone":                     "SL",
	"singapore":                        "SG",
	"sint maarten":                     "SX",
	"slovakia":                         "SK",
	"slovenia":                         "SI",
	"solomon islands":                  "SB",
	"somalia":                          "SO",
	"south africa":                     "ZA",
	"south korea":                      "KR",
	"south sudan":                      "SS",
	"spain":                            "ES",
	"sri lanka":                        "LK",
	"sudan":                            "SD",
	"suriname":                         "SR",
	"svalbard and jan mayen":           "SJ",
	"swaziland":                        "SZ",
	"sweden":                           "SE",
	"switzerland":                      "CH",
	"syria":                            "SY",
	"taiwan":                           "TW",
	"tajikistan":                       "TJ",
	"tanzania":                         "TZ",
	"thailand":                         "TH",
	"togo":                             "TG",
	"tokelau":                          "TK",
	"tonga":                            "TO",
	"trinidad and tobago":              "TT",
	"tunisia":                          "TN",
	"turkey":                           "TR",
	"turkmenistan":                     "TM",
	"turks and caicos islands":         "TC",
	"tuvalu":                           "TV",
	"u.s. virgin islands":              "VI",
	"uganda":                           "UG",
	"ukraine":                          "UA",
	"united arab emirates":             "AE",
	"united kingdom":                   "GB",
	"united states":                    "US",
	"uruguay":                          "UY",
	"uzbekistan":                       "UZ",
	"vanuatu":                          "VU",
	"vatican":                          "VA",
	"venezuela":                        "VE",
	"vietnam":                          "VN",
	"wallis and futuna":                "WF",
	"western sahara":                   "EH",
	"yemen":                            "YE",
	"zambia":                           "ZM",
	"zimbabwe":                         "ZW",
}

//
//func init()  {
//	go func() {
//		for {
//			if res, err := redis_utils.RetrieveStringFromRedis("test_key", rediscli.RedisDBIdx_Common); err == nil {
//				log.Println(res)
//			}else {
//				log.Println("get:", err.Error())
//			}
//		}
//	}()
//	for i:=3; i< 10; i++  {
//		if err := redis_utils.SetKeyNotExistEx("test_key", rediscli.RedisDBIdx_Common, strconv.Itoa(i), time.Minute); err != nil {
//			log.Println("set err:", err.Error())
//		}
//	}
//}
//

func main123() {
	wp := lib.WorkPool{}
	ids := []string{"1", "2", "3", "4", "5", "6", "7", "8"}
	tasks := make([]interface{}, 0, len(ids))
	for _, key := range ids {
		tasks = append(tasks, key)
	}
	log.Printf("tasks len:%d", len(tasks))

	wp.WorkerCount = 2
	wp.Tasks = tasks

	wp.DoFunc = func(t interface{}) {
		log.Printf("num:%s", t.(string))
	}
	wp.Process()
	if _, isClose := <-wp.Done; !isClose {

	}
}
func main() {
	//ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	//defer cancel()
	//
	//go handle(ctx, 1500*time.Millisecond)
	//
	//select {
	//case <-ctx.Done():
	//	fmt.Println("main", ctx.Err())
	//}
	time.Sleep(time.Second)
}

func handle(ctx context.Context, duration time.Duration) {
	select {
	case <-ctx.Done():
		fmt.Println("handle", ctx.Err())

	case <-time.After(duration):
		fmt.Println("process request with", duration)
	}
}

func init() {
	t := time.Now().UTC()
	date := fmt.Sprintf("%s UTC", t.Format("Mon,02 Jan 2006 15:04:05"))
	fmt.Printf(date)

}
