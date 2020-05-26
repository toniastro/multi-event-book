package pdfGenerator

import (
	"encoding/base64"
	"fmt"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
	"io/ioutil"
	"os"
	"time"
)

type Receipt struct {
	Currency    string
	Amount      int
	Reference   string
	Paid        time.Time
	FullName    string
	Email       string
	Address     string
	EventName   string
	Details     string
}

func (r Receipt) GeneratePDF() bool {
	data := r
	begin := time.Now()
	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	m.SetPageMargins(10, 15, 10)
	//m.SetBorder(true)

	byteSlices, err := ioutil.ReadFile("storage/pdf/go.jpg")
	if err != nil {
		fmt.Println("Got error while opening file:", err)
		os.Exit(1)
	}

	base64 := base64.StdEncoding.EncodeToString(byteSlices)

	m.RegisterHeader(func() {
		m.Row(20, func() {
			m.Col(3, func() {
				m.Base64Image(base64, consts.Jpg, props.Rect{
					Center:  true,
					Percent: 70,
				})
			})

			m.ColSpace(3)

			m.Col(3, func() {
				m.QrCode("https://github.com/iamt-chadwick/multi-event-book", props.Rect{
					Center:  true,
					Percent: 75,
				})
			})
		})

		m.Line(1.0)

	})


	m.Row(15, func() {
		m.Col(12, func() {
			m.Text(fmt.Sprintf("Receipt of Payment for Event"), props.Text{
				Top:   8,
				Style: consts.Bold,
			})
		})
	})
	m.Row(15, func() {
		m.Col(12, func() {

			m.Text(fmt.Sprintf("Reference ID for payment: " + data.Reference), props.Text{
				Top:   8,
				Style: consts.Bold,
			})
		})
	})
	m.Row(15, func() {
		m.Col(12, func() {

			m.Text(fmt.Sprintf("%d %s in total \n", data.Amount ,data. Currency), props.Text{
				Top:   8,
				Style: consts.Bold,
			})
		})
	})
	m.Row(15, func() {
		m.Col(12, func() {

			m.Text(fmt.Sprintf("Address of Event: " + data.Address), props.Text{
				Top:   8,
				Style: consts.Bold,
			})
		})
	})
	m.Row(15, func() {
		m.Col(12, func() {

			m.Text(fmt.Sprintf("Event Name: " + data.EventName), props.Text{
				Top:   8,
				Style: consts.Bold,
			})
		})
	})
	m.Row(15, func() {
		m.Col(12, func() {

			m.Text(fmt.Sprintf("Full Name: " + data.FullName), props.Text{
				Top:   8,
				Style: consts.Bold,
			})
		})
	})
	m.Row(15, func() {
		m.Col(12, func() {

			m.Text(fmt.Sprintf("Email Address: " + data.Email), props.Text{
				Top:   8,
				Style: consts.Bold,
			})
		})
	})

	m.Row(15, func() {
		m.Col(12, func() {

			m.Text(fmt.Sprintf("I'd have loved to send this to your email, but no stress :)"), props.Text{
				Top:   8,
				Style: consts.Bold,
			})
		})
	})

	m.Line(1.0)

	m.AddPage()

	err = m.OutputFileAndClose("storage/pdf/"+ data.Reference + ".pdf")
	if err != nil {
		fmt.Println("Could not save PDF:", err)
		os.Exit(1)
	}

	end := time.Now()

	fmt.Println(end.Sub(begin))

	return true
}