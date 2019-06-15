package actions

import (
	"errors"
	"github.com/bopbi/travel_server/models"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
)

// This file is generated by Buffalo. It offers a basic structure for
// adding, editing and deleting a page. If your model is more
// complex or you need more than the basic implementation you need to
// edit this file.

// Following naming logic is implemented in Buffalo:
// Model: Singular (Featured)
// DB Table: Plural (featureds)
// Resource: Plural (Featureds)
// Path: Plural (/featureds)
// View Template Folder: Plural (/templates/featureds/)

// FeaturedsResource is the resource for the Featured model
type FeaturedsResource struct {
	buffalo.Resource
}

// List gets all Featureds. This function is mapped to the path
// GET /featureds
func (v FeaturedsResource) List(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.New("no transaction found")
	}

	featureds := &models.Featureds{}

	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())

	// Retrieve all Featureds from the DB
	if err := q.All(featureds); err != nil {
		return err
	}

	// Add the paginator to the context so it can be used in the template.
	c.Set("pagination", q.Paginator)

	return c.Render(200, r.Auto(c, featureds))
}

// Show gets the data for one Featured. This function is mapped to
// the path GET /featureds/{featured_id}
func (v FeaturedsResource) Show(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.New("no transaction found")
	}

	// Allocate an empty Featured
	featured := &models.Featured{}

	// To find the Featured the parameter featured_id is used.
	if err := tx.Find(featured, c.Param("featured_id")); err != nil {
		return c.Error(404, err)
	}

	return c.Render(200, r.Auto(c, featured))
}

// New renders the form for creating a new Featured.
// This function is mapped to the path GET /featureds/new
func (v FeaturedsResource) New(c buffalo.Context) error {
	return c.Render(200, r.Auto(c, &models.Featured{}))
}

// Create adds a Featured to the DB. This function is mapped to the
// path POST /featureds
func (v FeaturedsResource) Create(c buffalo.Context) error {
	// Allocate an empty Featured
	featured := &models.Featured{}

	// Bind featured to the html form elements
	if err := c.Bind(featured); err != nil {
		return err
	}

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.New("no transaction found")
	}

	// Validate the data from the html form
	verrs, err := tx.ValidateAndCreate(featured)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the new.html template that the user can
		// correct the input.
		return c.Render(422, r.Auto(c, featured))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", T.Translate(c, "featured.created.success"))
	// and redirect to the featureds index page
	return c.Render(201, r.Auto(c, featured))
}

// Edit renders a edit form for a Featured. This function is
// mapped to the path GET /featureds/{featured_id}/edit
func (v FeaturedsResource) Edit(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.New("no transaction found")
	}

	// Allocate an empty Featured
	featured := &models.Featured{}

	if err := tx.Find(featured, c.Param("featured_id")); err != nil {
		return c.Error(404, err)
	}

	return c.Render(200, r.Auto(c, featured))
}

// Update changes a Featured in the DB. This function is mapped to
// the path PUT /featureds/{featured_id}
func (v FeaturedsResource) Update(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.New("no transaction found")
	}

	// Allocate an empty Featured
	featured := &models.Featured{}

	if err := tx.Find(featured, c.Param("featured_id")); err != nil {
		return c.Error(404, err)
	}

	// Bind Featured to the html form elements
	if err := c.Bind(featured); err != nil {
		return err
	}

	verrs, err := tx.ValidateAndUpdate(featured)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the edit.html template that the user can
		// correct the input.
		return c.Render(422, r.Auto(c, featured))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", T.Translate(c, "featured.updated.success"))
	// and redirect to the featureds index page
	return c.Render(200, r.Auto(c, featured))
}

// Destroy deletes a Featured from the DB. This function is mapped
// to the path DELETE /featureds/{featured_id}
func (v FeaturedsResource) Destroy(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.New("no transaction found")
	}

	// Allocate an empty Featured
	featured := &models.Featured{}

	// To find the Featured the parameter featured_id is used.
	if err := tx.Find(featured, c.Param("featured_id")); err != nil {
		return c.Error(404, err)
	}

	if err := tx.Destroy(featured); err != nil {
		return err
	}

	// If there are no errors set a flash message
	c.Flash().Add("success", T.Translate(c, "featured.destroyed.success"))
	// Redirect to the featureds index page
	return c.Render(200, r.Auto(c, featured))
}
