import (
	"bytes"
	"errors"
	"fmt"
	template2 "html/template"
	"net/url"
	"os"
	"path/filepath"
	"text/template"

	"github.com/avelino/awesome-go/pkg/markdown"
	cp "github.com/otiai10/copy"

	"github.com/PuerkitoBio/goquery"
	"github.com/avelino/awesome-go/pkg/slug"
)

// Link contains info about awesome url
type Link struct {
	Title       string
	URL         string
	Description string
}
