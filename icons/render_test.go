package icons

import (
	"context"
	"strings"
	"testing"

	"github.com/a-h/templ"
)

func render(t *testing.T, c templ.Component) string {
	t.Helper()
	var sb strings.Builder
	if err := c.Render(context.Background(), &sb); err != nil {
		t.Fatalf("render: %v", err)
	}
	return sb.String()
}

// allIcons is every icon in the package. Keeping it exhaustive means a new icon
// must be added here, and guarantees each one renders valid SVG.
func allIcons(p Props) map[string]templ.Component {
	return map[string]templ.Component{
		"Activity": Activity(p), "Archive": Archive(p), "ArrowUpDown": ArrowUpDown(p),
		"Bell": Bell(p), "Bold": Bold(p), "Calendar": Calendar(p), "Check": Check(p),
		"ChevronDown": ChevronDown(p), "ChevronLeft": ChevronLeft(p), "ChevronRight": ChevronRight(p),
		"ChevronsUpDown": ChevronsUpDown(p), "ChevronUp": ChevronUp(p), "Circle": Circle(p),
		"CircleAlert": CircleAlert(p), "CircleCheck": CircleCheck(p), "CircleX": CircleX(p),
		"Clipboard": Clipboard(p), "Copy": Copy(p), "CreditCard": CreditCard(p),
		"DollarSign": DollarSign(p), "Download": Download(p), "ExternalLink": ExternalLink(p),
		"File": File(p), "Film": Film(p), "FolderOpen": FolderOpen(p), "Github": Github(p),
		"Google": Google(p), "GripVertical": GripVertical(p), "Home": Home(p), "Inbox": Inbox(p),
		"Info": Info(p), "Italic": Italic(p), "LayoutGrid": LayoutGrid(p), "Link": Link(p),
		"ListMusic": ListMusic(p), "Loader": Loader(p), "LogIn": LogIn(p), "LogOut": LogOut(p),
		"Mail": Mail(p), "Megaphone": Megaphone(p), "Moon": Moon(p), "MoreHorizontal": MoreHorizontal(p),
		"Package": Package(p), "PanelLeft": PanelLeft(p), "Play": Play(p), "Plus": Plus(p),
		"PlusCircle": PlusCircle(p), "Radio": Radio(p), "Scissors": Scissors(p), "Search": Search(p),
		"Send": Send(p), "Settings": Settings(p), "Sun": Sun(p), "SunMoon": SunMoon(p),
		"Terminal": Terminal(p), "Trash": Trash(p), "TrendingUp": TrendingUp(p), "Underline": Underline(p),
		"User": User(p), "Users": Users(p), "X": X(p),
	}
}

// TestIconsRender renders every icon and checks it produces valid <svg> markup
// and honors the Class prop.
func TestIconsRender(t *testing.T) {
	p := Props{Class: "h-4 w-4", Attrs: templ.Attributes{"data-test": "x"}}
	for name, c := range allIcons(p) {
		t.Run(name, func(t *testing.T) {
			out := render(t, c)
			if !strings.Contains(out, "<svg") || !strings.Contains(out, "</svg>") {
				t.Errorf("expected a complete <svg> element:\n%s", out)
			}
			// tailwind-merge may canonicalize ordering ("h-4 w-4" -> "w-4 h-4"),
			// so check both utility tokens are present rather than the exact order.
			if !strings.Contains(out, "h-4") || !strings.Contains(out, "w-4") {
				t.Errorf("expected the Class prop (h-4 w-4) to be applied:\n%s", out)
			}
			if !strings.Contains(out, `data-test="x"`) {
				t.Errorf("expected Attrs to be spread:\n%s", out)
			}
		})
	}
}

// TestGoogleBrandColors checks the multi-color Google mark renders fills.
func TestGoogleBrandColors(t *testing.T) {
	out := render(t, Google(Props{}))
	for _, c := range []string{"#4285F4", "#34A853", "#FBBC05", "#EA4335"} {
		if !strings.Contains(out, c) {
			t.Errorf("Google icon missing brand color %s", c)
		}
	}
}
