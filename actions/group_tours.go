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
// Model: Singular (GroupTour)
// DB Table: Plural (group_tours)
// Resource: Plural (GroupTours)
// Path: Plural (/group_tours)
// View Template Folder: Plural (/templates/group_tours/)

// GroupToursResource is the resource for the GroupTour model
type GroupToursResource struct {
	buffalo.Resource
}

// List gets all GroupTours. This function is mapped to the path
// GET /group_tours
func (v GroupToursResource) List(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.New("no transaction found")
	}

	groupTours := &models.GroupTours{}

	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())

	// Retrieve all GroupTours from the DB
	if err := q.All(groupTours); err != nil {
		return err
	}

	// Add the paginator to the context so it can be used in the template.
	c.Set("pagination", q.Paginator)

	return c.Render(200, r.Auto(c, groupTours))
}

// Show gets the data for one GroupTour. This function is mapped to
// the path GET /group_tours/{group_tour_id}
func (v GroupToursResource) Show(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.New("no transaction found")
	}

	// Allocate an empty GroupTour
	groupTour := &models.GroupTour{}

	// To find the GroupTour the parameter group_tour_id is used.
	if err := tx.Find(groupTour, c.Param("group_tour_id")); err != nil {
		return c.Error(404, err)
	}

	return c.Render(200, r.Auto(c, groupTour))
}

// New renders the form for creating a new GroupTour.
// This function is mapped to the path GET /group_tours/new
func (v GroupToursResource) New(c buffalo.Context) error {
	return c.Render(200, r.Auto(c, &models.GroupTour{}))
}

// Create adds a GroupTour to the DB. This function is mapped to the
// path POST /group_tours
func (v GroupToursResource) Create(c buffalo.Context) error {
	// Allocate an empty GroupTour
	groupTour := &models.GroupTour{}

	// Bind groupTour to the html form elements
	if err := c.Bind(groupTour); err != nil {
		return err
	}

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.New("no transaction found")
	}

	// Validate the data from the html form
	verrs, err := tx.ValidateAndCreate(groupTour)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the new.html template that the user can
		// correct the input.
		return c.Render(422, r.Auto(c, groupTour))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", T.Translate(c, "groupTour.created.success"))
	// and redirect to the group_tours index page
	return c.Render(201, r.Auto(c, groupTour))
}

// Edit renders a edit form for a GroupTour. This function is
// mapped to the path GET /group_tours/{group_tour_id}/edit
func (v GroupToursResource) Edit(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.New("no transaction found")
	}

	// Allocate an empty GroupTour
	groupTour := &models.GroupTour{}

	if err := tx.Find(groupTour, c.Param("group_tour_id")); err != nil {
		return c.Error(404, err)
	}

	return c.Render(200, r.Auto(c, groupTour))
}

// Update changes a GroupTour in the DB. This function is mapped to
// the path PUT /group_tours/{group_tour_id}
func (v GroupToursResource) Update(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.New("no transaction found")
	}

	// Allocate an empty GroupTour
	groupTour := &models.GroupTour{}

	if err := tx.Find(groupTour, c.Param("group_tour_id")); err != nil {
		return c.Error(404, err)
	}

	// Bind GroupTour to the html form elements
	if err := c.Bind(groupTour); err != nil {
		return err
	}

	verrs, err := tx.ValidateAndUpdate(groupTour)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the edit.html template that the user can
		// correct the input.
		return c.Render(422, r.Auto(c, groupTour))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", T.Translate(c, "groupTour.updated.success"))
	// and redirect to the group_tours index page
	return c.Render(200, r.Auto(c, groupTour))
}

// Destroy deletes a GroupTour from the DB. This function is mapped
// to the path DELETE /group_tours/{group_tour_id}
func (v GroupToursResource) Destroy(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.New("no transaction found")
	}

	// Allocate an empty GroupTour
	groupTour := &models.GroupTour{}

	// To find the GroupTour the parameter group_tour_id is used.
	if err := tx.Find(groupTour, c.Param("group_tour_id")); err != nil {
		return c.Error(404, err)
	}

	if err := tx.Destroy(groupTour); err != nil {
		return err
	}

	// If there are no errors set a flash message
	c.Flash().Add("success", T.Translate(c, "groupTour.destroyed.success"))
	// Redirect to the group_tours index page
	return c.Render(200, r.Auto(c, groupTour))
}
