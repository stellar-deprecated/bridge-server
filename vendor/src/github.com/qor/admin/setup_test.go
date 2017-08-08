package admin_test

import (
	"net/http/httptest"

	"github.com/jinzhu/gorm"
	"github.com/qor/admin"
	. "github.com/qor/admin/tests/dummy"
)

var (
	server *httptest.Server
	db     *gorm.DB
	Admin  *admin.Admin
)

func init() {
	Admin = NewDummyAdmin()
	db = Admin.Config.DB
	server = httptest.NewServer(Admin.NewServeMux("/admin"))
}
