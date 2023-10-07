package main

import "fmt"

type TextComponent interface {
	GetText() string
}

type text struct {
	Text string
}

func (p text) GetText() string {
	return p.Text
}

type TextDecorator struct {
	TextComponent TextComponent
}

func (td TextDecorator) GetText() string {
	return td.TextComponent.GetText()
}

type BoldTextDecorator struct {
	TextDecorator
}

func (b BoldTextDecorator) GetText() string {
	return "**b**" + b.TextComponent.GetText() + "**b**"
}

type ItalicTextDecorator struct {
	TextDecorator
}

func (i ItalicTextDecorator) GetText() string {
	return "*i*" + i.TextComponent.GetText() + "*i*"
}

type UnderlineTextDecorator struct {
	TextDecorator
}

func (u UnderlineTextDecorator) GetText() string {
	return "_u_" + u.TextComponent.GetText() + "_u_"
}

func main() {
	text := text{Text: "Golang - True?"}
	boldText := BoldTextDecorator{TextDecorator: TextDecorator{TextComponent: text}}
	italicBoldText := ItalicTextDecorator{TextDecorator: TextDecorator{TextComponent: boldText}}
	underlineItalicBoldText := UnderlineTextDecorator{TextDecorator: TextDecorator{TextComponent: italicBoldText}}

	fmt.Println("Text:", text.GetText())
	fmt.Println("Bold:", boldText.GetText())
	fmt.Println("Italic and bold:", italicBoldText.GetText())
	fmt.Println("Underline, italic, bold:", underlineItalicBoldText.GetText())
}
