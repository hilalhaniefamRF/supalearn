package controllers

import (
	"supalearn/internal/models"

	"github.com/sev-2/raiden"
)

type ProfilesController struct {
	raiden.ControllerBase
	Http  string `path:"/profiles" type:"rest"`
	Model models.Profiles
}
