package admin

import "github.com/qor/roles"

// Config admin config struct
type Config struct {
	Name       string
	Menu       []string
	Invisible  bool
	Priority   int
	PageCount  int
	Singleton  bool
	Permission *roles.Permission
	Themes     []ThemeInterface
}
