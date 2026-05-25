package ui

import (
	"strings"
	"testing"
	"time"

	"github.com/a-h/templ"
)

// TestNewComponentsRender directly renders every component added in the full
// port (and its sub-components) so the ui package is tested independently of
// the docs site.
func TestNewComponentsRender(t *testing.T) {
	cases := map[string]templ.Component{
		// display / primitives
		"aspect-ratio":      AspectRatio(16.0/9.0, "", nil),
		"progress":          Progress(60, "", nil),
		"kbd":               Kbd("", nil),
		"spinner":           Spinner("", nil),
		"empty":             Empty("", nil),
		"empty-icon":        EmptyIcon("", nil),
		"empty-title":       EmptyTitle("", nil),
		"empty-description": EmptyDescription("", nil),
		"empty-action":      EmptyAction("", nil),
		"scroll-area":       ScrollArea("", nil),
		"typography-h1":     TypographyH1("", nil),
		"typography-p":      TypographyP("", nil),
		"typography-code":   TypographyCode("", nil),
		// form controls
		"radio-group":       RadioGroup("{ value: 'a' }", "", nil),
		"radio-group-item":  RadioGroupItem("g", "a", "value", "", nil),
		"slider":            Slider("s", 0, 100, 1, 50, "", nil),
		"select":            Select("fruit", "val", "Pick one", []SelectOption{{Value: "a", Label: "Apple"}, {Value: "b", Label: "Banana", Disabled: true}}, "", nil),
		"select-label":      SelectLabel("", nil),
		"select-separator":  SelectSeparator("", nil),
		"native-select":     NativeSelect("s", "", nil),
		"native-option":     NativeSelectOption("a", "", nil),
		"input-group":       InputGroup("", nil),
		"input-group-addon": InputGroupAddon("leading", "", nil),
		"input-group-input": InputGroupInput("n", InputTypeText, true, false, "", nil),
		"button-group":      ButtonGroup("", nil),
		"toggle-group":      ToggleGroup("{ value: 'a' }", "", nil),
		"toggle-group-item": ToggleGroupItem("value", "a", "single", ToggleVariantDefault, ToggleSizeDefault, "", nil),
		"input-otp":         InputOTP(6, "otp", "", nil),
		"input-otp-sep":     InputOTPWithSeparator(3, 2, "otp", "", nil),
		// layout / list
		"item":             Item("", nil),
		"item-media":       ItemMedia("", nil),
		"item-content":     ItemContent("", nil),
		"item-title":       ItemTitle("", nil),
		"item-description": ItemDescription("", nil),
		"item-actions":     ItemActions("", nil),
		"pagination":       Pagination("", nil),
		"pagination-list":  PaginationContent("", nil),
		"pagination-item":  PaginationItem("", nil),
		"pagination-link":  PaginationLink("/", true, "", nil),
		"pagination-prev":  PaginationPrevious("/", "", nil),
		"pagination-next":  PaginationNext("/", "", nil),
		"pagination-elip":  PaginationEllipsis("", nil),
		// overlays
		"tooltip":            Tooltip("", nil),
		"tooltip-trigger":    TooltipTrigger("", nil),
		"tooltip-content":    TooltipContent("top", "", nil),
		"popover":            Popover("", nil),
		"popover-content":    PopoverContent("", nil),
		"hover-card":         HoverCard("", nil),
		"hover-card-content": HoverCardContent("", nil),
		"alert-dialog":       AlertDialog("", nil),
		"alert-dialog-body":  AlertDialogContent("", nil),
		"alert-dialog-head":  AlertDialogHeader("", nil),
		"alert-dialog-foot":  AlertDialogFooter("", nil),
		"alert-dialog-title": AlertDialogTitle("", nil),
		"alert-dialog-desc":  AlertDialogDescription("", nil),
		"sheet":              Sheet("", nil),
		"sheet-content":      SheetContent(SheetSideRight, "", nil),
		"sheet-left":         SheetContent(SheetSideLeft, "", nil),
		"sheet-header":       SheetHeader("", nil),
		"sheet-footer":       SheetFooter("", nil),
		"sheet-title":        SheetTitle("", nil),
		"sheet-desc":         SheetDescription("", nil),
		"drawer":             Drawer("", nil),
		"drawer-content":     DrawerContent("", nil),
		"drawer-header":      DrawerHeader("", nil),
		"drawer-footer":      DrawerFooter("", nil),
		"drawer-title":       DrawerTitle("", nil),
		"drawer-desc":        DrawerDescription("", nil),
		"context-menu":       ContextMenu("", nil),
		"context-trigger":    ContextMenuTrigger("", nil),
		"context-content":    ContextMenuContent("", nil),
		"context-item":       ContextMenuItem("", nil),
		"context-separator":  ContextMenuSeparator("", nil),
		"context-shortcut":   ContextMenuShortcut("", nil),
		"context-label":      ContextMenuLabel("", nil),
		// menus / nav
		"menubar":            Menubar("", nil),
		"menubar-menu":       MenubarMenu("", nil),
		"menubar-trigger":    MenubarTrigger("", nil),
		"menubar-content":    MenubarContent("", nil),
		"menubar-item":       MenubarItem("", nil),
		"menubar-separator":  MenubarSeparator("", nil),
		"menubar-shortcut":   MenubarShortcut("", nil),
		"menubar-label":      MenubarLabel("", nil),
		"nav-menu":           NavigationMenu("", nil),
		"nav-menu-list":      NavigationMenuList("", nil),
		"nav-menu-item":      NavigationMenuItem("", nil),
		"nav-menu-trigger":   NavigationMenuTrigger("", nil),
		"nav-menu-content":   NavigationMenuContent("", nil),
		"nav-menu-link":      NavigationMenuLink("/", "", nil),
		"nav-menu-title":     NavigationMenuLinkTitle("", nil),
		"nav-menu-desc":      NavigationMenuLinkDescription("", nil),
		"command":            Command("", nil),
		"command-input":      CommandInput("Search", "", nil),
		"command-list":       CommandList("", nil),
		"command-empty":      CommandEmpty("", nil),
		"command-group":      CommandGroup("Group", "", nil),
		"command-item":       CommandItem("a", "", nil),
		"command-separator":  CommandSeparator("", nil),
		"combobox":           Combobox("Pick one", []ComboboxOption{{Value: "a", Label: "Apple"}}, "", nil),
		"sidebar-provider":   SidebarProvider(true, "", nil),
		"sidebar":            Sidebar("", nil),
		"sidebar-header":     SidebarHeader("", nil),
		"sidebar-content":    SidebarContent("", nil),
		"sidebar-footer":     SidebarFooter("", nil),
		"sidebar-group":      SidebarGroup("", nil),
		"sidebar-grouplabel": SidebarGroupLabel("", nil),
		"sidebar-menu":       SidebarMenu("", nil),
		"sidebar-menuitem":   SidebarMenuItem("", nil),
		"sidebar-menubtn":    SidebarMenuButton(true, "", nil),
		"sidebar-menulink":   SidebarMenuButtonLink("/", true, "", nil),
		"sidebar-trigger":    SidebarTrigger("", nil),
		"sidebar-inset":      SidebarInset("", nil),
		// complex
		"calendar":     Calendar(2026, time.May, "", nil),
		"date-picker":  DatePicker(2026, time.May, "", nil),
		"carousel":     CarouselRoot(3, "", nil),
		"carousel-cnt": CarouselContent("", nil),
		"carousel-itm": CarouselItem("", nil),
		"carousel-prv": CarouselPrevious("", nil),
		"carousel-nxt": CarouselNext("", nil),
		"data-table":   DataTable([]DataTableColumn{{Key: "name", Label: "Name", Sortable: true}}, []DataTableRow{{"name": "Ada"}}, 5, "", nil),
		"bar-chart":    BarChart([]ChartDataPoint{{Label: "Jan", Value: 10}, {Label: "Feb", Value: 20}}, "Visitors", "", nil),
		"line-chart":   LineChart([]ChartDataPoint{{Label: "Jan", Value: 10}, {Label: "Feb", Value: 20}}, "Visitors", "", nil),
		"sonner":       SonnerToaster("", nil),
		"toast-vp":     ToastViewport("", nil),
		"toast":        Toast("default", "", nil),
		"toast-title":  ToastTitle("", nil),
		"toast-desc":   ToastDescription("", nil),
		"toast-action": ToastAction("", nil),
		"resizable":    ResizablePanelGroup("horizontal", "", nil),
		"resizable-pn": ResizablePanel(50, "", nil),
		"resizable-hd": ResizableHandle(true, false, "", nil),
		"direction":    Direction("rtl", "", nil),
		// sub-components added to match shadcn's component APIs
		"card-action":         CardAction("", nil),
		"breadcrumb-page":     BreadcrumbPage("", nil),
		"breadcrumb-ellipsis": BreadcrumbEllipsis("", nil),
		"select-group":        SelectGroup("", nil),
		"dropdown-group":      DropdownMenuGroup("", nil),
		"dropdown-checkbox":   DropdownMenuCheckboxItem("checked", "", nil),
		"dropdown-radiogroup": DropdownMenuRadioGroup("", nil),
		"dropdown-radio-item": DropdownMenuRadioItem("val", "a", "", nil),
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			if out := strings.TrimSpace(renderToString(t, c)); out == "" {
				t.Fatal("rendered empty output")
			}
		})
	}
}

// TestNewComponentStructure verifies high-confidence structural details of the
// more complex components.
func TestNewComponentStructure(t *testing.T) {
	checks := []struct {
		name string
		c    templ.Component
		want []string
	}{
		{"calendar", Calendar(2026, time.May, "", nil), []string{"May", "2026"}},
		{"bar-chart", BarChart([]ChartDataPoint{{Label: "Jan", Value: 10}}, "T", "", nil), []string{"<svg"}},
		{"line-chart", LineChart([]ChartDataPoint{{Label: "Jan", Value: 10}}, "T", "", nil), []string{"<svg"}},
		{"data-table", DataTable([]DataTableColumn{{Key: "name", Label: "Name"}}, []DataTableRow{{"name": "Ada"}}, 5, "", nil), []string{"<table", "Ada"}},
		{"progress", Progress(40, "", nil), []string{"40"}},
		{"select", Select("f", "v", "Pick", []SelectOption{{Value: "a", Label: "Apple"}}, "", nil), []string{"Apple", "x-data"}},
		{"sheet", SheetContent(SheetSideRight, "", nil), []string{"x-show"}},
	}
	for _, ch := range checks {
		t.Run(ch.name, func(t *testing.T) {
			out := renderToString(t, ch.c)
			for _, w := range ch.want {
				if !strings.Contains(out, w) {
					t.Errorf("output missing %q", w)
				}
			}
		})
	}
}

// TestTriggerAttrVarsPopulated checks the exported "*Attrs" trigger maps exist
// and carry Alpine wiring.
func TestTriggerAttrVarsPopulated(t *testing.T) {
	vars := map[string]templ.Attributes{
		"PopoverTriggerAttrs":     PopoverTriggerAttrs,
		"PopoverCloseAttrs":       PopoverCloseAttrs,
		"SheetTriggerAttrs":       SheetTriggerAttrs,
		"SheetCloseAttrs":         SheetCloseAttrs,
		"AlertDialogTriggerAttrs": AlertDialogTriggerAttrs,
		"AlertDialogActionAttrs":  AlertDialogActionAttrs,
		"AlertDialogCancelAttrs":  AlertDialogCancelAttrs,
		"DrawerTriggerAttrs":      DrawerTriggerAttrs,
		"DrawerCloseAttrs":        DrawerCloseAttrs,
	}
	for name, v := range vars {
		if len(v) == 0 {
			t.Errorf("%s is empty; trigger wiring missing", name)
		}
	}
}
