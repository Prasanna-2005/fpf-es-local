package db

import "testing"

func TestDecideVersionValue(t *testing.T) {
	priorEdit := &currentSnapshot{Value: "50", IsEdit: true, CreatedBy: "editor"}
	priorExtracted := &currentSnapshot{Value: "40", IsEdit: false, CreatedBy: "extractor"}
	same := "40"
	changed := "75"

	cases := []struct {
		name      string
		extracted string
		prior     *currentSnapshot
		latest    *string
		wantVal   string
		wantEdit  bool
		wantKeep  bool
	}{
		{"first version has no prior", "40", nil, nil, "40", false, false},
		{"prior was extracted", "75", priorExtracted, &same, "75", false, false},
		{"edit kept when file value unchanged", "40", priorEdit, &same, "50", true, true},
		{"edit dropped when file value changed", "75", priorEdit, &same, "75", false, false},
		{"edit dropped when no extracted ancestor", "40", priorEdit, nil, "40", false, false},
		{"string compare is exact", "40.0", priorEdit, &same, "40.0", false, false},
		{"changed pointer still exact", changed, priorEdit, &changed, "50", true, true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, isEdit, kept := DecideVersionValue(tc.extracted, tc.prior, tc.latest)
			if got != tc.wantVal || isEdit != tc.wantEdit || kept != tc.wantKeep {
				t.Fatalf("got (%q, edit=%v, keep=%v)", got, isEdit, kept)
			}
		})
	}
}
