package advertise

import (
	"reflect"
	"testing"

	"github.com/willhu-commit/pango/testdata"
)

func TestPanoNormalization(t *testing.T) {
	testCases := getTests()

	mc := &testdata.MockClient{}
	ns := PanoramaNamespace(mc)

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			mc.Reset()
			mc.Version = tc.version
			mc.AddResp("")
			err := ns.Set("tmpl", "", "vr", "ag", tc.conf)
			if err != nil {
				t.Errorf("Error in set: %s", err)
			} else {
				mc.AddResp(mc.Elm)
				r, err := ns.Get("tmpl", "", "vr", "ag", tc.conf.Name)
				if err != nil {
					t.Errorf("Error in get: %s", err)
				}
				if !reflect.DeepEqual(tc.conf, r) {
					t.Errorf("%#v != %#v", tc.conf, r)
				}
			}
		})
	}
}
