package company

import (
	"reflect"
	"testing"
	"time"
)

// mock
type makeRepository func(t *testing.T) Repository

type repository struct {
	company *Company
	err     error
	// flag
	create bool
	// test
	t *testing.T
}

func (r *repository) CompanyCreate(*Company) (*Company, error) {
	r.t.Helper()

	if r.create {
		return r.company, r.err
	}
	r.t.Fatal("invalid CompanyCreate")
	panic("invalid CompanyCreate")
}

func TestServer_Create(t *testing.T) {
	type test struct {
		name           string
		makeRepository makeRepository
		company        *Company
		want           *Company
		wantErr        bool
	}

	do := func(tt *test) {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewServer(tt.makeRepository(t)).Create(tt.company)
			if tt.wantErr != (err != nil) {
				t.Fatalf("want-erorr=%v, error=%v.", tt.wantErr, err)
			}

			if !reflect.DeepEqual(tt.want, got) {
				t.Fatalf("want=%v, got=%v.", tt.want, got)
			}
		})
	}

	tests := []*test{
		func() *test {
			updatedAt := time.Now()

			return &test{
				name:    "ok",
				company: New("GREATE COMPANY", 1),
				makeRepository: func(t *testing.T) Repository {
					return &repository{
						company: &Company{
							ID:        2,
							Name:      "GREATE COMPANY",
							OwnerID:   1,
							UpdatedAt: updatedAt,
						},
						err:    nil,
						create: true,
						t:      t,
					}
				},
				want: &Company{
					ID:        2,
					Name:      "GREATE COMPANY",
					OwnerID:   1,
					UpdatedAt: updatedAt,
				},
				wantErr: false,
			}
		}(),
		{
			name:    "invalid",
			company: New("", 1),
			makeRepository: func(t *testing.T) Repository {
				return &repository{
					create: false,
					t:      t,
				}
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}
