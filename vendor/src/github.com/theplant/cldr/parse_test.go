package cldr_test

import (
	"html/template"

	"github.com/theplant/cldr"
	_ "github.com/theplant/cldr/resources/locales/en"

	"testing"
)

type testCart struct {
	Name          string
	Items         []string
	NumberOfItems int32
}

func TestParse(t *testing.T) {
	cases := []struct {
		locale string
		text   string
		data   []interface{}
		want   string
	}{
		{
			locale: "en",
			text:   `{{p "Count" (one "{{.Count}} item") (other "{{.Count}} items in Your Cart")}} in Your Cart`,
			data:   []interface{}{map[string]int{"Count": 1}},
			want:   "1 item in Your Cart",
		},
		{
			locale: "en",
			text:   `{{p "Count" (one "{{.Count}} item") (other "{{.Count}} items in Your Cart")}} in Your Cart`,
			data:   []interface{}{map[string]int{"Count": 2}},
			want:   "2 items in Your Cart in Your Cart",
		},
		{
			locale: "en",
			text:   `{{p "Count" (one "{{.Count}} item") (other "{{.Count}} items in Your Cart")}} in Your Cart; {{p "Count2" (one "{{.Count2}} item") (other "{{.Count2}} items in Your Cart")}} in Your Cart`,
			data:   []interface{}{map[string]int{"Count": 2, "Count2": 1}},
			want:   "2 items in Your Cart in Your Cart; 1 item in Your Cart",
		},
		{
			locale: "en",
			text:   `{{p "Cart.Items" (one "1 item") (other "{{len .Cart.Items}} items")}} in your cart; {{p "Cart.NumberOfItems" (one "1 item") (other "{{.Cart.NumberOfItems}} items")}} in your cart.`,
			data: []interface{}{map[string]interface{}{"Cart": struct {
				Name          string
				Items         []string
				NumberOfItems int32
			}{Name: "Mr Someone", Items: []string{"Item 1", "Item 2"}, NumberOfItems: 4}}},
			want: "2 items in your cart; 4 items in your cart.",
		},
		{
			locale: "en",
			text:   `<div>Your search "{{ .Keyword }}" returned {{ .Count }} results</div>`,
			data: []interface{}{struct {
				Keyword string
				Count   template.HTML
			}{
				Keyword: `<img src="https://getqor.com/imageproxy?link=error" onerror="alert('hello')" >`,
				Count:   template.HTML("<b>1</b>"),
			}},
			want: "<div>Your search \"&lt;img src=&#34;https://getqor.com/imageproxy?link=error&#34; onerror=&#34;alert(&#39;hello&#39;)&#34; &gt;\" returned <b>1</b> results</div>",
		},
		{
			locale: "en",
			text:   `{{p "Cart2.Items" (one "1 item") (other "{{len .Cart2.Items}} items")}} in your cart; {{p "Cart2.NumberOfItems" (one "1 item") (other "{{.Cart2.NumberOfItems}} items")}} in your cart.`,
			data:   []interface{}{map[string]interface{}{"Cart2": &testCart{Name: "Test cart", Items: []string{"Item 3", "Item 4", "Item 5"}, NumberOfItems: 6}}},
			want:   "3 items in your cart; 6 items in your cart.",
		},
		{
			locale: "en",
			text:   `{{p "." (one "1 item") (other "{{.}} items")}} in your cart.`,
			data:   []interface{}{4},
			want:   "4 items in your cart.",
		},
		{
			locale: "en",
			text:   `{{p "." (one "1 item") (other "{{len .}} items")}} in your cart.`,
			data:   []interface{}{[]string{"Item 1", "Item 2"}},
			want:   "2 items in your cart.",
		},
		{
			locale: "en",
			text:   `{{$1}} {{$2}} {{$1}}`,
			data:   []interface{}{"string1", "string2"},
			want:   "string1 string2 string1",
		},
	}
	for i := 0; i < 2; i++ {
		for _, c := range cases {
			got, err := cldr.Parse(c.locale, c.text, c.data...)
			if err != nil {
				t.Error(err)
			}
			if got != c.want {
				t.Errorf("got %q; want %q", got, c.want)
			}
		}
	}
}
