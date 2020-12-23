package models

import (
	"encoding/json"
	"github.com/gobuffalo/nulls"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gofrs/uuid"
	"time"
)
// Bateauid is used by pop to map your .model.Name.Proper.Pluralize.Underscore database table to your go code.
type Bateauid struct {
    ID uuid.UUID `json:"id" db:"id"`
    NombateauTitle nulls.String `json:"nombateau_title" db:"nombateau_title"`
    ModeleSubtitle nulls.String `json:"modele_subtitle" db:"modele_subtitle"`
    LocalisationportGeoLat nulls.Float64 `json:"localisationport_geo_lat" db:"localisationport_geo_lat"`
    LocalisationportGeoLong nulls.Float64 `json:"localisationport_geo_long" db:"localisationport_geo_long"`
    Spare1Str nulls.String `json:"spare1_str" db:"spare1_str"`
    Spare2Str nulls.String `json:"spare2_str" db:"spare2_str"`
    Spare3Str nulls.String `json:"spare3_str" db:"spare3_str"`
    Spare4Str nulls.String `json:"spare4_str" db:"spare4_str"`
    Spare5Str nulls.String `json:"spare5_str" db:"spare5_str"`
    Spare6Str nulls.String `json:"spare6_str" db:"spare6_str"`
    Spare7Str nulls.String `json:"spare7_str" db:"spare7_str"`
    Spare8Str nulls.String `json:"spare8_str" db:"spare8_str"`
    Spare9Str nulls.String `json:"spare9_str" db:"spare9_str"`
    Spare10Str nulls.String `json:"spare10_str" db:"spare10_str"`
    Spare11Int nulls.Int `json:"spare11_int" db:"spare11_int"`
    Spare12Int nulls.Int `json:"spare12_int" db:"spare12_int"`
    Spare13Int nulls.Int `json:"spare13_int" db:"spare13_int"`
    Spare14Int nulls.Int `json:"spare14_int" db:"spare14_int"`
    Spare15Int nulls.Int `json:"spare15_int" db:"spare15_int"`
    Spare16Int nulls.Int `json:"spare16_int" db:"spare16_int"`
    Spare17ID nulls.UUID `json:"spare17_id" db:"spare17_id"`
    Spare18ID nulls.UUID `json:"spare18_id" db:"spare18_id"`
    Spare19ID nulls.UUID `json:"spare19_id" db:"spare19_id"`
    Spare20ID nulls.UUID `json:"spare20_id" db:"spare20_id"`
    Spare21ID nulls.UUID `json:"spare21_id" db:"spare21_id"`
    Spare22ID nulls.UUID `json:"spare22_id" db:"spare22_id"`
    CreatedAt time.Time `json:"created_at" db:"created_at"`
    UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (b Bateauid) String() string {
	jb, _ := json.Marshal(b)
	return string(jb)
}

// Bateauids is not required by pop and may be deleted
type Bateauids []Bateauid

// String is not required by pop and may be deleted
func (b Bateauids) String() string {
	jb, _ := json.Marshal(b)
	return string(jb)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (b *Bateauid) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (b *Bateauid) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (b *Bateauid) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
