package utils_test

import (
	"testing"
	"time"

	"github.com/gera9/blog/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestPatchStruct(t *testing.T) {
	type U struct {
		X int
		Y string
	}
	type O struct {
		A       string
		B       *int
		C       float64
		d       string // unexported field
		OtherO  *O
		Date    time.Time
		U       U
		List    []string
		O2      *O
		IsValid *string
	}
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		dst     any
		src     any
		want    any
		wantErr bool
		err     error
	}{
		{
			name: "basic patch",
			dst: &O{
				A:    "old",
				B:    nil,
				C:    1.0,
				d:    "text",
				Date: time.Date(2025, time.October, 17, 0, 0, 0, 0, time.UTC),
				List: []string{"a"},
				O2: &O{
					List: []string{"1"},
				},
				IsValid: utils.Ptr(""),
			},
			src: O{
				A: "new",
				B: utils.Ptr(42),
				C: 0.0,       // zero value, should not overwrite
				d: "newtext", // unexported, should not overwrite
				OtherO: &O{
					A: "new A",
				},
				Date: time.Date(2025, time.October, 18, 0, 0, 0, 0, time.UTC),
				U: U{
					X: 1,
					Y: "1",
				},
				List: []string{"a", "b"},
				O2: &O{
					List: []string{"1", "2"},
				},
				IsValid: utils.Ptr("Okay"),
			},
			want: &O{
				A: "new",
				B: utils.Ptr(42),
				C: 1.0,
				d: "text",
				OtherO: &O{
					A: "new A",
				},
				Date: time.Date(2025, time.October, 18, 0, 0, 0, 0, time.UTC),
				U: U{
					X: 1,
					Y: "1",
				},
				List: []string{"a", "b"},
				O2: &O{
					List: []string{"1", "2"},
				},
				IsValid: utils.Ptr("Okay"),
			},
			wantErr: false,
		},
		{
			name:    "src not struct",
			dst:     &O{},
			src:     42,
			wantErr: true,
			err:     utils.ErrSrcNotStruct,
		},
		{
			name:    "dst not pointer",
			dst:     O{},
			src:     O{},
			wantErr: true,
			err:     utils.ErrDstNotPointer,
		},
		{
			name:    "dst not pointer to struct",
			dst:     utils.Ptr(42),
			src:     O{},
			wantErr: true,
			err:     utils.ErrDstNotPointerStruct,
		},
		{
			name:    "diff types",
			dst:     &O{},
			src:     U{},
			wantErr: true,
			err:     utils.ErrDiffTypes,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotErr := utils.PatchStruct(tt.dst, tt.src)
			if tt.wantErr {
				if assert.Error(t, gotErr) {
					assert.Equal(t, tt.err, gotErr)
				}
				return
			}

			if assert.NoError(t, gotErr) {
				assert.Equal(t, tt.want, tt.dst)
			}
		})
	}
}
