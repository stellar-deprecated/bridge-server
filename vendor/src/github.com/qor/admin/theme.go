package admin

// ThemeInterface theme interface
type ThemeInterface interface {
	GetName() string
	GetViewPaths() []string
	ConfigAdminTheme(*Resource)
}

// Theme base theme config struct
type Theme struct {
	Name string
}

// GetName get name from theme
func (theme Theme) GetName() string {
	return theme.Name
}

// GetViewPaths get view paths from theme
func (Theme) GetViewPaths() []string {
	return []string{}
}

// ConfigAdminTheme config theme for admin resource
func (Theme) ConfigAdminTheme(*Resource) {
	return
}
