// Code generated by raiden-cli; DO NOT EDIT.
package bootstrap

import (
	"github.com/sev-2/raiden/pkg/resource"
	"supalearn/internal/models"
)

func RegisterModels() {
	resource.RegisterModels(
		&models.Profiles{},
	)
}