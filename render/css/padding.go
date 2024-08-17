package css

import (
	"simpleWebBrowser/render"
)

func (c *CSS) padding(tag *render.Tag) {
	c.y += tag.PaddingTop
	tag.Height += tag.PaddingTop
}
