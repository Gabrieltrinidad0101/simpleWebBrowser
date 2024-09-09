package css

import (
	"simpleWebBrowser/render"
)

func (c *CSS) padding(tag *render.Tag) {
	*tag.Height += tag.PaddingTop
}
