package protocols

import (
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
	"strconv"

	"github.com/facebookgo/structtag"
	"github.com/stellar/go/build"
	"github.com/stellar/go/support/errors"
)

var federationDestinationFieldName = regexp.MustCompile("forward_destination\\[fields\\]\\[([a-z_-]+)\\]")

// Asset represents native or credit asset
type Asset struct {
	Code   string `name:"asset_code" json:"code"`
	Issuer string `name:"asset_issuer" json:"issuer"`
}

// ToBaseAsset transforms Asset to github.com/stellar/go-stellar-base/build.Asset
func (a Asset) ToBaseAsset() build.Asset {
	if a.Code == "" && a.Issuer == "" {
		return build.NativeAsset()
	}
	return build.CreditAsset(a.Code, a.Issuer)
}

// String returns string representation of this asset
func (a Asset) String() string {
	return fmt.Sprintf("Code: %s, Issuer: %s", a.Code, a.Issuer)
}

// Validate checks if asset params are correct.
func (a Asset) Validate() bool {
	if a.Code != "" && a.Issuer != "" {
		// Credit
		return IsValidAssetCode(a.Code) && IsValidAccountID(a.Issuer)
	} else if a.Code == "" && a.Issuer == "" {
		// Native
		return true
	} else {
		return false
	}
}

// ForwardDestination contains fields required to create forward federation request
type ForwardDestination struct {
	Domain string     `name:"domain"`
	Fields url.Values `name:"fields"`
}

// FormRequest allows transforming http.Request url.Values from/to request structs
type FormRequest struct {
	HTTPRequest *http.Request
}

const (
	pathCodeField   = "path[%d][asset_code]"
	pathIssuerField = "path[%d][asset_issuer]"
)

// FromRequest transforms http.Request to request struct object
func (request *FormRequest) FromRequest(r *http.Request, destination interface{}) error {
	request.HTTPRequest = r

	rvalue := reflect.ValueOf(destination).Elem()
	typ := rvalue.Type()
	for i := 0; i < rvalue.NumField(); i++ {
		tag := typ.Field(i).Tag.Get("name")
		switch tag {
		case "":
			continue
		case "forward_destination":
			var destination ForwardDestination
			destination.Domain = r.PostFormValue("forward_destination[domain]")
			destination.Fields = make(url.Values)

			err := r.ParseForm()
			if err != nil {
				return err
			}

			for key := range r.PostForm {
				matches := federationDestinationFieldName.FindStringSubmatch(key)
				if len(matches) < 2 {
					continue
				}

				fieldName := matches[1]
				destination.Fields.Add(fieldName, r.PostFormValue(key))
			}

			ptr := rvalue.Field(i).Addr().Interface().(**ForwardDestination)
			if destination.Domain != "" && len(destination.Fields) > 0 {
				*ptr = &destination
			} else {
				*ptr = nil
			}
		case "path":
			var path []Asset

			for i := 0; i < 5; i++ {
				codeFieldName := fmt.Sprintf(pathCodeField, i)
				issuerFieldName := fmt.Sprintf(pathIssuerField, i)

				// If the element does not exist in PostForm break the loop
				if _, exists := r.PostForm[codeFieldName]; !exists {
					break
				}

				code := r.PostFormValue(codeFieldName)
				issuer := r.PostFormValue(issuerFieldName)

				if code == "" && issuer == "" {
					path = append(path, Asset{})
				} else {
					path = append(path, Asset{code, issuer})
				}
			}

			ptr := rvalue.Field(i).Addr().Interface().(*[]Asset)
			*ptr = path
		default:
			value := r.PostFormValue(tag)
			if value == "" {
				continue
			}

			switch rvalue.Field(i).Kind() {
			case reflect.Bool:
				b, err := strconv.ParseBool(value)
				if err != nil {
					return err
				}
				rvalue.Field(i).SetBool(b)
			case reflect.String:
				rvalue.Field(i).SetString(value)
			default:
				return errors.New("Invalid value: " + value + " type for type: " + tag)
			}
		}
	}
	return nil
}

// CheckRequired checks whether all fields marked as required have value
func (request *FormRequest) CheckRequired(destination interface{}) error {
	rvalue := reflect.ValueOf(destination).Elem()
	typ := rvalue.Type()
	for i := 0; i < rvalue.NumField(); i++ {
		required, _, err := structtag.Extract("required", string(typ.Field(i).Tag))

		if err != nil {
			return NewInternalServerError(
				"Error extracting tag using structtag",
				map[string]interface{}{"error": err},
			)
		}

		if required {
			name := typ.Field(i).Tag.Get("name")
			if request.HTTPRequest.PostFormValue(name) == "" {
				return NewMissingParameter(name)
			}
		}
	}
	return nil
}

// ToValues transforms request object to url.Values
func (request *FormRequest) ToValues(object interface{}) (values url.Values) {
	values = make(map[string][]string)
	rvalue := reflect.ValueOf(object).Elem()
	typ := rvalue.Type()
	for i := 0; i < rvalue.NumField(); i++ {
		field := rvalue.Field(i)
		tag := typ.Field(i).Tag.Get("name")
		if tag == "" {
			continue
		}
		switch field.Interface().(type) {
		case bool:
			value := rvalue.Field(i).Bool()
			values.Set(tag, strconv.FormatBool(value))
		case string:
			value := rvalue.Field(i).String()
			if value == "" {
				continue
			}
			values.Set(tag, value)
		case []Asset:
			assets := rvalue.Field(i).Interface().([]Asset)
			for i, asset := range assets {
				values.Set(fmt.Sprintf(pathCodeField, i), asset.Code)
				values.Set(fmt.Sprintf(pathIssuerField, i), asset.Issuer)
			}
		case *ForwardDestination:
			destination := rvalue.Field(i).Interface().(*ForwardDestination)
			if destination == nil {
				continue
			}

			values.Add("forward_destination[domain]", destination.Domain)
			for key := range destination.Fields {
				values.Add(fmt.Sprintf("forward_destination[fields][%s]", key), destination.Fields.Get(key))
			}
		}
	}
	return
}

// SuccessResponse is embedded in all success responses and implements server.Response interface
type SuccessResponse struct{}

// HTTPStatus returns http.StatusOK
func (response *SuccessResponse) HTTPStatus() int {
	return http.StatusOK
}
