package ui

import (
	"context"
	"strings"
	"testing"

	"github.com/a-h/templ"
)

func renderToString(t *testing.T, c templ.Component) string {
	t.Helper()
	var sb strings.Builder
	if err := c.Render(context.Background(), &sb); err != nil {
		t.Fatalf("render: %v", err)
	}
	return sb.String()
}

// TestComponentsRender renders every component (and sub-component) in the
// library to make sure none of them panic or error, with both nil and
// populated attributes.
func TestComponentsRender(t *testing.T) {
	attrs := templ.Attributes{"data-test": "x", "disabled": true}

	cases := map[string]templ.Component{
		"accordion":            Accordion("max-w-md", attrs),
		"accordion-item":       AccordionItem("", nil),
		"accordion-trigger":    AccordionTrigger("", nil),
		"accordion-content":    AccorionContent("", nil),
		"alert-default":        Alert(AlertVariantDefault, "", nil),
		"alert-destructive":    Alert(AlertVariantDestructive, "", nil),
		"alert-title":          AlertTitle("", nil),
		"alert-description":    AlertDescription("", nil),
		"avatar":               Avatar("", nil),
		"avatar-image":         AvatarImage("https://example.com/a.png", "alt", "", nil),
		"avatar-fallback":      AvatarFallback("", nil),
		"badge-default":        Badge(BadgeVariantDefault, "", nil),
		"badge-secondary":      Badge(BadgeVariantSecondary, "", nil),
		"badge-destructive":    Badge(BadgeVariantDestructive, "", nil),
		"badge-outline":        Badge(BadgeVariantOutline, "", nil),
		"breadcrumb":           Breadcrumb("", nil),
		"breadcrumb-list":      BreadcrumbList("", nil),
		"breadcrumb-item":      BreadcrumbItem("", nil),
		"breadcrumb-link":      BreadcrumbLink("/", "", nil),
		"breadcrumb-separator": BreadcrumbSeparator("", nil),
		"button":               Button(ButtonTypeButton, ButtonVariantDefault, ButtonSizeDefault, "", nil),
		"button-submit-icon":   Button(ButtonTypeSubmit, ButtonVariantOutline, ButtonSizeIcon, "", attrs),
		"card":                 Card("", nil),
		"card-header":          CardHeader("", nil),
		"card-title":           CardTitle("", nil),
		"card-description":     CardDescription("", nil),
		"card-content":         CardContent("", nil),
		"card-footer":          CardFooter("", nil),
		"checkbox":             Checkbox("agree", "", attrs),
		"collapsible":          Collapsible(false, "", attrs),
		"collapsible-content":  CollapsibleContent("", nil),
		"dialog":               Dialog("", nil),
		"dialog-content":       DialogContent("", nil),
		"dialog-header":        DialogHeader("", nil),
		"dialog-footer":        DialogFooter("", nil),
		"dialog-title":         DialogTitle("", nil),
		"dialog-description":   DialogDescription("", nil),
		"dropdown":             DropdownMenu("", nil),
		"dropdown-content":     DropdownMenuContent(DropdownMenuAlignLeft, "", nil),
		"dropdown-item":        DropdownMenuItem("", nil),
		"dropdown-separator":   DropdownMenuSeparator("", nil),
		"dropdown-shortcut":    DropdownMenuShortcut("", nil),
		"dropdown-label":       DropdownMenuLabel("", nil),
		"field":                Field("", nil),
		"field-description":    FieldDescription("", nil),
		"field-error":          FieldError("err-id", "", nil),
		"input":                Input("name", InputTypeText, "", nil),
		"input-email":          Input("email", InputTypeEmail, "", attrs),
		"label":                Label("", nil),
		"separator-horizontal": Separator(SeparatorOrientationHorizontal, false, "", nil),
		"separator-vertical":   Separator(SeparatorOrientationVertical, true, "", nil),
		"skeleton":             Skeleton("", nil),
		"switch":               Switch("airplane", true, true, "", attrs),
		"table":                Table("", nil),
		"table-header":         TableHeader("", nil),
		"table-body":           TableBody("", nil),
		"table-footer":         TableFooter("", nil),
		"table-row":            TableRow("", nil),
		"table-head":           TableHead("", nil),
		"table-cell":           TableCell("", nil),
		"table-caption":        TableCaption("", nil),
		"tabs":                 Tabs("a", "", nil),
		"tabs-list":            TabsList("", nil),
		"tabs-trigger":         TabsTrigger("a", "", nil),
		"tabs-content":         TabsContent("a", "", nil),
		"textarea":             Textarea("msg", "", nil),
		"toggle":               Toggle(ToggleVariantDefault, ToggleSizeDefault, "pressed", "", nil),
		"toggle-secondary":     Toggle(ToggleVariantSecondary, ToggleSizeLarge, "pressed", "", attrs),
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			if out := strings.TrimSpace(renderToString(t, c)); out == "" {
				t.Fatal("rendered empty output")
			}
		})
	}
}

// TestComponentClasses checks a few components emit their expected base styles,
// confirming the tailwind-merge wiring works end-to-end.
func TestComponentClasses(t *testing.T) {
	checks := []struct {
		name string
		c    templ.Component
		want []string
	}{
		{"button-default", Button(ButtonTypeButton, ButtonVariantDefault, ButtonSizeDefault, "", nil), []string{"bg-primary", `type="button"`}},
		{"button-destructive", Button(ButtonTypeSubmit, ButtonVariantDestructive, ButtonSizeDefault, "", nil), []string{"bg-destructive", `type="submit"`}},
		{"badge-secondary", Badge(BadgeVariantSecondary, "", nil), []string{"bg-secondary"}},
		{"input-email", Input("email", InputTypeEmail, "", nil), []string{`type="email"`, `name="email"`}},
		{"separator", Separator(SeparatorOrientationHorizontal, false, "", nil), []string{`role="separator"`, "w-full"}},
		{"alert", Alert(AlertVariantDefault, "", nil), []string{`role="alert"`, "<svg"}},
	}
	for _, ch := range checks {
		t.Run(ch.name, func(t *testing.T) {
			out := renderToString(t, ch.c)
			for _, w := range ch.want {
				if !strings.Contains(out, w) {
					t.Errorf("output missing %q\n%s", w, out)
				}
			}
		})
	}
}

func TestClassOverrideMerges(t *testing.T) {
	// A custom rounded-* should override the base rounded-md via tailwind-merge.
	out := renderToString(t, Button(ButtonTypeButton, ButtonVariantDefault, ButtonSizeDefault, "rounded-full", nil))
	if strings.Contains(out, "rounded-md") {
		t.Errorf("expected rounded-md to be overridden by rounded-full:\n%s", out)
	}
	if !strings.Contains(out, "rounded-full") {
		t.Errorf("expected rounded-full in output:\n%s", out)
	}
}

func TestBoolAttr(t *testing.T) {
	if !boolAttr("x", templ.Attributes{"x": true}) {
		t.Error("expected true for bool true")
	}
	if !boolAttr("x", templ.Attributes{"x": "true"}) {
		t.Error("expected true for string \"true\"")
	}
	if boolAttr("x", templ.Attributes{"x": false}) {
		t.Error("expected false for bool false")
	}
	if boolAttr("missing", nil) {
		t.Error("expected false for nil attrs")
	}
}

func TestFmtBool(t *testing.T) {
	if fmtBool(true) != "true" || fmtBool(false) != "false" {
		t.Error("fmtBool returned unexpected value")
	}
}
