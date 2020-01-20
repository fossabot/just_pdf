package main

import (
	"fmt"
	"github.com/muhammadmuhlas/just_pdf/pkg/consts"
	"github.com/muhammadmuhlas/just_pdf/pkg/pdf"
	"github.com/muhammadmuhlas/just_pdf/pkg/props"
	"os"
	"time"
)

func main() {
	begin := time.Now()
	m := pdf.NewJustPdf(consts.Landscape, consts.A4)
	//m.SetBorder(true)

	m.Row(20, func() {
		m.Col(func() {
			_ = m.FileImage("internal/assets/images/frontpage.png", props.Rect{
				Percent: 88,
				Center:  true,
			})
		})
		m.Col(func() {
			m.Text("Golang Certificate", props.Text{
				Top:   12,
				Align: consts.Center,
				Size:  20,
				Style: consts.BoldItalic,
			})
		})
		m.Col(func() {
			_ = m.FileImage("internal/assets/images/frontpage.png", props.Rect{
				Percent: 90,
				Center:  true,
			})
		})
	})

	m.Row(130, func() {
		m.Col(func() {
			text := "Lorem Ipsum is simply dummy textá of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum."
			m.Text(text, props.Text{
				Size:            13,
				Align:           consts.Center,
				Top:             60,
				VerticalPadding: 2.0,
			})
		})
	})

	m.Row(25, func() {
		m.Col(func() {
			m.Signature("Gopher Senior")
		})
		m.Col(func() {
			m.Signature("Gopheroid")
		})
		m.Col(func() {
			m.Signature("Sign Here")
		})
	})

	err := m.OutputFileAndClose("internal/examples/pdfs/certificate.pdf")
	if err != nil {
		fmt.Println("Could not save PDF:", err)
		os.Exit(1)
	}

	end := time.Now()
	fmt.Println(end.Sub(begin))
}
