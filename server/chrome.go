package main

import (
	"context"

	"github.com/chromedp/chromedp"
)

type ChromeHandler struct {
	ctx    context.Context
	cancel []context.CancelFunc
}

func NewChromeHandler() *ChromeHandler {
	var c ChromeHandler
	var opts = append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
		chromedp.Flag("hide-scrollbars", false),
		chromedp.Flag("mute-audio", false),
	)
	var cancel context.CancelFunc

	c.ctx, cancel = chromedp.NewExecAllocator(
		context.Background(), opts...,
	)
	c.cancel = append(c.cancel, cancel)

	c.ctx, cancel = chromedp.NewContext(c.ctx)
	c.cancel = append(c.cancel, cancel)

	return &c
}

func (c *ChromeHandler) Show() (err error) {
	return chromedp.Run(c.ctx,
		chromedp.Navigate("http://127.0.0.1:"+Settings.ListenPort),
	)
}

func (c *ChromeHandler) Cancel() {
	for i := range c.cancel {
		c.cancel[len(c.cancel)-i-1]()
	}
	return
}
