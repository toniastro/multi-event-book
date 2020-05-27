package models

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/iamt-chadwick/multi-event-book/pdfGenerator"
	u "github.com/iamt-chadwick/multi-event-book/utils"
	"github.com/ichtrojan/thoth"
	_ "github.com/ichtrojan/thoth"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/rs/xid"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	_ "regexp"
	_ "strings"
	"time"
)
//Post Type details
type Users struct{
	gorm.Model
	Name         string   `gorm:"type:varchar(100)" json:"name"`
	Email        string   `gorm:"type:varchar(100)" json:"email"`
	EventID		 uint	  `json:"event_id"`
	Reference    string   `gorm:"not null"`
	Initiated    bool     `gorm:"default:true"`
	Completed    bool     `gorm:"default: false"`
	Verified     bool     `gorm:"default: false"`
	Verified_by  string   `gorm:"null;type:varchar(100)"`
}

type Price struct{
	Price        float64       `json:"amount"`
	EventsID     int		   `json:"events_id"`
}
type Payload struct {
	Txref        string    `json:"txref"`
	SECKEY       string    `json:"SECKEY"`
}
type TxnVerificationResponse struct {
	Data    txnVerificationResponseData `json:"data"`
	Message string                      `json:"message"`
	Status  string                      `json:"status"`
}

type txnVerificationResponseData struct {
	Txid                            int         `json:"txid"`
	Txref                           string      `json:"txref"`
	Flwref                          string      `json:"flwref"`
	Devicefingerprint               string      `json:"devicefingerprint"`
	Cycle                           string      `json:"cycle"`
	Amount                          int         `json:"amount"`
	Currency                        string      `json:"currency"`
	Chargedamount                   int         `json:"chargedamount"`
	Appfee                          int         `json:"appfee"`
	Merchantfee                     int         `json:"merchantfee"`
	Merchantbearsfee                int         `json:"merchantbearsfee"`
	Chargecode                      string      `json:"chargecode"`
	Chargemessage                   string      `json:"chargemessage"`
	Authmodel                       string      `json:"authmodel"`
	IP                              string      `json:"ip"`
	Narration                       string      `json:"narration"`
	Status                          string      `json:"status"`
	Vbvcode                         string      `json:"vbvcode"`
	Vbvmessage                      string      `json:"vbvmessage"`
	Authurl                         string      `json:"authurl"`
	Acctcode                        interface{} `json:"acctcode"`
	Acctmessage                     interface{} `json:"acctmessage"`
	Paymenttype                     string      `json:"paymenttype"`
	Paymentid                       string      `json:"paymentid"`
	Fraudstatus                     string      `json:"fraudstatus"`
	Chargetype                      string      `json:"chargetype"`
	Createdday                      int         `json:"createdday"`
	Createddayname                  string      `json:"createddayname"`
	Createdweek                     int         `json:"createdweek"`
	Createdmonth                    int         `json:"createdmonth"`
	Createdmonthname                string      `json:"createdmonthname"`
	Createdquarter                  int         `json:"createdquarter"`
	Createdyear                     int         `json:"createdyear"`
	Createdyearisleap               bool        `json:"createdyearisleap"`
	Createddayispublicholiday       int         `json:"createddayispublicholiday"`
	Createdhour                     int         `json:"createdhour"`
	Createdminute                   int         `json:"createdminute"`
	Createdpmam                     string      `json:"createdpmam"`
	Created                         time.Time   `json:"created"`
	Customerid                      int         `json:"customerid"`
	Custphone                       interface{} `json:"custphone"`
	Custnetworkprovider             string      `json:"custnetworkprovider"`
	Custname                        string      `json:"custname"`
	Custemail                       string      `json:"custemail"`
	Custemailprovider               string      `json:"custemailprovider"`
	Custcreated                     time.Time   `json:"custcreated"`
	Accountid                       int         `json:"accountid"`
	Acctbusinessname                string      `json:"acctbusinessname"`
	Acctcontactperson               string      `json:"acctcontactperson"`
	Acctcountry                     string      `json:"acctcountry"`
	Acctbearsfeeattransactiontime   int         `json:"acctbearsfeeattransactiontime"`
	Acctparent                      int         `json:"acctparent"`
	Acctvpcmerchant                 string      `json:"acctvpcmerchant"`
	Acctalias                       interface{} `json:"acctalias"`
	Acctisliveapproved              int         `json:"acctisliveapproved"`
	Orderref                        string      `json:"orderref"`
	Paymentplan                     interface{} `json:"paymentplan"`
	Paymentpage                     interface{} `json:"paymentpage"`
	Raveref                         string      `json:"raveref"`
	Amountsettledforthistransaction int         `json:"amountsettledforthistransaction"`
	Card                            struct {
		Expirymonth    string `json:"expirymonth"`
		Expiryyear     string `json:"expiryyear"`
		CardBIN        string `json:"cardBIN"`
		Last4Digits    string `json:"last4digits"`
		Brand          string `json:"brand"`
		IssuingCountry string `json:"issuing_country"`
		CardTokens     []struct {
			Embedtoken string `json:"embedtoken"`
			Shortcode  string `json:"shortcode"`
			Expiry     string `json:"expiry"`
		} `json:"card_tokens"`
		Type          string `json:"type"`
		LifeTimeToken string `json:"life_time_token"`
	} `json:"card"`
	Meta []struct {
		ID                   int         `json:"id"`
		Metaname             string      `json:"metaname"`
		Metavalue            string      `json:"metavalue"`
		CreatedAt            time.Time   `json:"createdAt"`
		UpdatedAt            time.Time   `json:"updatedAt"`
		DeletedAt            interface{} `json:"deletedAt"`
		GetpaidTransactionID int         `json:"getpaidTransactionId"`
	} `json:"meta"`
}

type Error struct {
   Mesage string `json:"message"`
}
var count int64

var logger, _ = thoth.Init("log")

var reference string


func init(){
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}
}

func (a Users) Validate() error {

	return validation.ValidateStruct(&a,
		validation.Field(&a.Name, validation.Required, validation.Length(3, 50)),
		validation.Field(&a.Email, validation.Required, is.Email),
		//validation.Field(&a.Phone, validation.Required, validation.Match(regexp.MustCompile("^[0-9]{11}$"))),
	)
}




func (detail *Users) Create(w http.ResponseWriter) map[string] interface{} {

	err := detail.Validate()

	if err != nil {

		u.Errors(w, http.StatusBadRequest, "", err)

		return nil
	}

	temp := &Users{}

	price := &Price{}

	GetDB().Table("users").Where("email = ?", detail.Email).Where("event_id = ?", detail.EventID).First(temp)

	GetDB().Table("event_kinds").Select("price").Where("id = ?", detail.EventID).First(price)

	if temp.Email != "" {

		if temp.Completed == true {

			u.Errors(w, http.StatusAlreadyReported, "This email has been used to complete a payment for this event.", nil)

			return nil

		}else{

			detail.Reference = checkValidReference()

			err := GetDB().Table("users").Where("email = ?", detail.Email).Update("reference", detail.Reference).Error

			if err != nil && err != gorm.ErrRecordNotFound {

				u.Errors(w, http.StatusInternalServerError, "Could not update reference ",nil)

				return nil

			}

			mapD := map[string] interface{} {"email": detail.Email, "reference": detail.Reference,"amount" : price.Price}

			u.Response(w, http.StatusOK, "", mapD)
			return nil
		}
	}

	detail.Reference = checkValidReference()

	GetDB().Create(detail)

	if detail.ID <= 0 {

		u.Response(w, http.StatusBadRequest,"Failed to initiate payment.", nil)

		return nil

	}

	mapD := map[string] interface{} {"email": detail.Email, "reference": detail.Reference,"amount" : price.Price}

	u.Response(w, http.StatusOK, "", mapD)
	return nil
}

func (detail *Payload) Confirm(w http.ResponseWriter) (map[string] interface{}) {

	user := &Users{}

	if detail.Txref == "" {

		  u.Response(w, http.StatusNotFound, "No reference code passed.", nil)

		  return nil


	}else{

		err := GetDB().Table("users").Where("reference = ?", detail.Txref).First(user).Error

		if err != nil || err == gorm.ErrRecordNotFound {

			 u.Response(w, http.StatusNotFound, "Reference code does not exist", nil)

			 return nil

		}
		if user.Completed == true {

			 u.Response(w, http.StatusAlreadyReported, "This payment has been verified before ;) .", nil)

			 return nil
		}

		transactionReference := detail.Txref

		results := confirmReferenceCode(transactionReference);

		data := TxnVerificationResponse{}

		_ = json.Unmarshal([]byte(results), &data)

		price := &Price{}

		GetDB().Table("event_kinds").Select("price,events_id").Where("id = ?", user.EventID).First(price)

		if data.Status == "success" && data.Data.Currency == "NGN" && data.Data.Chargedamount == int(price.Price) {

			err := GetDB().Table("users").Where("reference = ? AND email = ? AND event_id = ?" , transactionReference,data.Data.Custemail,user.EventID).Update("completed", true).Error

			if err != nil && err != gorm.ErrRecordNotFound {

				u.Response(w, http.StatusNotFound, "Something went wrong with verifying reference code", nil)

				return nil
			}

			event := &Events{}

			GetDB().Table("events").Select("name,address,details").Where("id = ?", user.EventID).First(event)

			templateData := pdfGenerator.Receipt{
				Currency: data.Data.Currency,
				Amount  : data.Data.Chargedamount,
				Reference: transactionReference,
				Paid : data.Data.Custcreated,
				FullName: data.Data.Custname,
				Email : data.Data.Custemail,
				Address : event.Address,
				EventName : event.Name,
				Details : event.Details,
			}

			if templateData.GeneratePDF() {

				domain, domainExists := os.LookupEnv("DOMAIN_HOST")

				if !domainExists {

					domain = "http://localhost/"

				}

				u.Response(w, http.StatusOK, "This payment has been verified","pdf/"+domain+transactionReference + ".pdf")

				return nil
			}

		}

		logger.Log(errors.New(results))

		u.Response(w, http.StatusBadRequest, "This payment doesn't look legit my friend.", nil)

		return nil
	}

}

func confirmReferenceCode(txref string) string {

	secKey, found := os.LookupEnv("RAVE_SECRET_KEY")

	if !found {
		log.Fatal("You need to set the \"RAVE_SECRET_KEY\" environment variable")
	}
	data := Payload{txref,secKey}

	payloadBytes, err := json.Marshal(data)

	if err != nil {

		logger.Log(err)

		panic(err)

	}
	body := bytes.NewBuffer(payloadBytes)

	timeout := time.Duration(5 * time.Second)

	client := http.Client{
		Timeout: timeout,
	}

	req, err := http.NewRequest("POST",returnAPIKey(), body)

	req.Header.Add("Content-Type", "application/json")

	if err != nil {

		logger.Log(err)

	}

	resp, err := client.Do(req)

	if err != nil {

		logger.Log(err)

		return "Something went wrong here sha 3"
	}

	defer resp.Body.Close()

	result,err := ioutil.ReadAll(resp.Body)

	if err != nil{

		logger.Log(err)

		return "Something went wrong here sha 4"
	}

	return string(result)
}


func checkValidReference() string{

	reference := generateReferenceCode()

	GetDB().Table("users").Where("reference = ?", reference).Count(&count)

	if count >= 1 {

		return generateReferenceCode()

	} else{

		return reference

	}
}

func generateReferenceCode() string {

	guid := xid.New()

	reference := guid.String()

	return reference
}

func returnAPIKey() string {

	if os.Getenv("RAVE_MODE") != "live" {

		return os.Getenv("RAVE_API_TEST")

	}

	return os.Getenv("RAVE_API_LIVE")
}

