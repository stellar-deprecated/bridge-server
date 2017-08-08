package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/qor/admin"
	"github.com/qor/qor"
	"github.com/qor/roles"
	templates "github.com/stellar/gateway/bridge/handlers/admin"
	"github.com/stellar/gateway/horizon"
	"github.com/stellar/gateway/protocols/compliance"
	monorepoCompliance "github.com/stellar/go/protocols/compliance"
)

type ReceivedPayment struct {
	ID              int64     `db:"id"`
	OperationID     string    `db:"operation_id"`
	ProcessedAt     time.Time `db:"processed_at"`
	PagingToken     string    `db:"paging_token"`
	Status          string    `db:"status"`
	ComplianceError string
	SenderInfoText  string
	monorepoCompliance.AuthData
	monorepoCompliance.Attachment
	horizon.PaymentResponse
}

func (ReceivedPayment) TableName() string {
	return "ReceivedPayment"
}

var (
	initOnce sync.Once
	client   http.Client
	handler  http.Handler
	// Dirty hack to get Config.Compliance in gorm callback
	complianceURL string
	// Dirty hack to get horizon.Client in gorm callback
	horizonClient horizon.HorizonInterface
)

func init() {
	client = http.Client{
		Timeout: 5 * time.Second,
	}
}

func (r *ReceivedPayment) AfterFind(scope *gorm.Scope) error {
	// Dirty hack to load compliance data when loading a single payment
	if strings.Contains(scope.SQL, "LIMIT 1 ") {
		var err error
		r.PaymentResponse, err = horizonClient.LoadOperation(r.OperationID)
		if err != nil {
			r.ComplianceError = "Error getting payment data: " + err.Error()
			return nil
		}

		err = horizonClient.LoadMemo(&r.PaymentResponse)
		if err != nil {
			r.ComplianceError = "Error loading payment memo: " + err.Error()
			return nil
		}

		if complianceURL != "" && r.PaymentResponse.Memo.Value != "" {
			resp, err := client.PostForm(complianceURL+"/receive", url.Values{"memo": []string{r.PaymentResponse.Memo.Value}})
			if err != nil {
				r.ComplianceError = err.Error()
				return nil
			}
			if resp.StatusCode == http.StatusNotFound {
				r.ComplianceError = "Transaction not found"
				return nil
			}
			if resp.StatusCode != http.StatusOK {
				r.ComplianceError = "Invalid response from compliance server"
				return nil
			}

			var response compliance.ReceiveResponse
			err = json.NewDecoder(resp.Body).Decode(&response)
			if err != nil {
				r.ComplianceError = "Error decoding compliance response"
				return nil
			}

			err = json.Unmarshal([]byte(response.Data), &r.AuthData)
			if err != nil {
				r.ComplianceError = "Error decoding auth data"
				return nil
			}

			r.Attachment, err = r.AuthData.Attachment()
			if err == nil {
				for key, value := range r.Attachment.SenderInfo {
					r.SenderInfoText += fmt.Sprintf("%s: %s\n", key, value)
				}
			}
		}
	}
	return nil
}

func (rh *RequestHandler) Admin(w http.ResponseWriter, r *http.Request) {
	complianceURL = rh.Config.Compliance
	horizonClient = rh.Horizon
	initOnce.Do(rh.createAdmin)
	handler.ServeHTTP(w, r)
}

func (rh *RequestHandler) createAdmin() {
	DB, err := gorm.Open(rh.Config.Database.Type, rh.Config.Database.URL)
	if err != nil {
		panic(err)
	}

	qorAdmin := admin.New(&qor.Config{DB: DB})
	qorAdmin.SetAssetFS(templates.AssetFS)
	qorAdmin.RegisterViewPath("vendor/src/github.com/qor/admin/views")

	receivedPaymentResource := qorAdmin.AddResource(
		&ReceivedPayment{},
		&admin.Config{
			// Allow reading only
			Permission: roles.Allow(roles.Read, roles.Anyone),
		},
	)
	receivedPaymentResource.ShowAttrs(
		&admin.Section{
			Title: "Bridge",
			Rows:  [][]string{{"OperationID"}, {"Status"}, {"ProcessedAt"}},
		},
		&admin.Section{
			Title: "Horizon",
			Rows:  [][]string{{"From"}, {"To"}, {"Amount"}, {"AssetCode"}, {"AssetIssuer"}, {"Memo.Value"}},
		},
		&admin.Section{
			Title: "Compliance",
			Rows:  [][]string{{"ComplianceError"}, {"Sender"}, {"NeedInfo"}, {"Tx"}},
		},
		&admin.Section{
			Title: "Attachment",
			Rows:  [][]string{{"Nonce"}, {"SenderInfoText"}, {"Route"}, {"Note"}, {"Extra"}},
		},
	)
	receivedPaymentResource.Action(
		&admin.Action{
			Name:       "reprocess",
			Modes:      []string{"show"},
			Permission: roles.Allow(roles.Update, roles.Anyone),
			Handler: func(actionArgument *admin.ActionArgument) error {
				id, err := strconv.ParseInt(actionArgument.PrimaryValues[0], 10, 64)
				if err != nil {
					return err
				}

				receivedPayment, err := rh.Repository.GetReceivedPaymentByPrimaryKey(id)
				if err != nil {
					return err
				}

				operation, err := rh.Horizon.LoadOperation(receivedPayment.OperationID)
				if err != nil {
					return err
				}

				err = rh.PaymentListener.ReprocessPayment(operation, false)
				return err
			},
			Visible: func(record interface{}, context *admin.Context) bool {
				if payment, ok := record.(*ReceivedPayment); ok {
					return payment.Status != "Success"
				}
				return false
			},
		},
	)
	receivedPaymentResource.Meta(&admin.Meta{Name: "NeedInfo", Type: "text"})
	receivedPaymentResource.IndexAttrs("ID", "OperationID", "ProcessedAt", "Status")

	handler = qorAdmin.NewServeMux("/admin")
}
