package main

import (
	"candy/assets"
	"candy/ui"
)

var _ ui.Component = (*app)(nil)

type app struct {
	*ui.Box
}

func newApp(assets *assets.Assets) *app {
	return &app{ui.NewBox(nil, []ui.Component{
		ui.NewText(&ui.TextProps{Text: `
I guess we could discuss the implications of the phrase "meant to be."
That is if we wanted to drown ourselves in a sea of backwardly referential 
semantics and other mumbo-jumbo. Maybe such a discussion would result in the 
determination that "meant to be" is exactly as meaningless a phrase as it 
seems to be, and that none of us is actually meant to be doing anything at all. 
But that's my existential underpants underpinnings showing. It's the way the 
cookie crumbles. And now I want a cookie.
`}, &ui.Style{FontStyle: ui.FontStyle{
			Family:     "Source Code Pro",
			Weight:     "ExtraLight",
			Italic:     false,
			Size:       20,
			LineHeight: 24,
			Color: ui.Color{
				Red:   255,
				Green: 255,
				Blue:  255,
				Alpha: 255,
			}}}),
		ui.NewImage(assets, &ui.ImageProps{ImagePath: "test/image3.png"}, nil),
		ui.NewImage(assets, &ui.ImageProps{ImagePath: "test/image1.jpg"}, nil),
		ui.NewImage(assets, &ui.ImageProps{ImagePath: "test/image2.jpg"}, nil),
	},
		nil)}
}