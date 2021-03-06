// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "cellar": Application User Types
//
// Command:
// $ goagen
// --design=go-playground/design
// --out=$(GOPATH)/src/go-playground
// --version=v1.3.1

package client

import (
	"github.com/goadesign/goa"
	"unicode/utf8"
)

// BottlePayload that creates bottles
type bottlePayload struct {
	// Name of bottle
	Name *string `form:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty" xml:"name,omitempty"`
	// Rating of bottle
	Rating *int `form:"rating,omitempty" json:"rating,omitempty" yaml:"rating,omitempty" xml:"rating,omitempty"`
	// Vintage of bottle
	Vintage *int `form:"vintage,omitempty" json:"vintage,omitempty" yaml:"vintage,omitempty" xml:"vintage,omitempty"`
}

// Validate validates the bottlePayload type instance.
func (ut *bottlePayload) Validate() (err error) {
	if ut.Name == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "name"))
	}
	if ut.Vintage == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "vintage"))
	}
	if ut.Rating == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "rating"))
	}
	if ut.Name != nil {
		if utf8.RuneCountInString(*ut.Name) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`request.name`, *ut.Name, utf8.RuneCountInString(*ut.Name), 2, true))
		}
	}
	if ut.Rating != nil {
		if *ut.Rating < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`request.rating`, *ut.Rating, 1, true))
		}
	}
	if ut.Rating != nil {
		if *ut.Rating > 5 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`request.rating`, *ut.Rating, 5, false))
		}
	}
	if ut.Vintage != nil {
		if *ut.Vintage < 1900 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`request.vintage`, *ut.Vintage, 1900, true))
		}
	}
	return
}

// Publicize creates BottlePayload from bottlePayload
func (ut *bottlePayload) Publicize() *BottlePayload {
	var pub BottlePayload
	if ut.Name != nil {
		pub.Name = *ut.Name
	}
	if ut.Rating != nil {
		pub.Rating = *ut.Rating
	}
	if ut.Vintage != nil {
		pub.Vintage = *ut.Vintage
	}
	return &pub
}

// BottlePayload that creates bottles
type BottlePayload struct {
	// Name of bottle
	Name string `form:"name" json:"name" yaml:"name" xml:"name"`
	// Rating of bottle
	Rating int `form:"rating" json:"rating" yaml:"rating" xml:"rating"`
	// Vintage of bottle
	Vintage int `form:"vintage" json:"vintage" yaml:"vintage" xml:"vintage"`
}

// Validate validates the BottlePayload type instance.
func (ut *BottlePayload) Validate() (err error) {
	if ut.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "name"))
	}

	if utf8.RuneCountInString(ut.Name) < 2 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`type.name`, ut.Name, utf8.RuneCountInString(ut.Name), 2, true))
	}
	if ut.Rating < 1 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`type.rating`, ut.Rating, 1, true))
	}
	if ut.Rating > 5 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`type.rating`, ut.Rating, 5, false))
	}
	if ut.Vintage < 1900 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`type.vintage`, ut.Vintage, 1900, true))
	}
	return
}
