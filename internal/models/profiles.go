package models

import (
	"time"
	"github.com/sev-2/raiden"
)

type Profiles struct {
	raiden.ModelBase
	Id int64 `json:"id,omitempty" column:"name:id;type:bigint;primaryKey;autoIncrement;nullable:false"`
	CreatedAt time.Time `json:"created_at,omitempty" column:"name:created_at;type:timestampz;nullable:false;default:now()"`
	FullName *string `json:"full_name,omitempty" column:"name:full_name;type:varchar;nullable"`
	Phone *string `json:"phone,omitempty" column:"name:phone;type:varchar;nullable"`
	Email *string `json:"email,omitempty" column:"name:email;type:varchar;nullable"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" column:"name:updated_at;type:timestampz;nullable"`

	// Table information
	Metadata string `json:"-" schema:"public" rlsEnable:"true" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`
}
