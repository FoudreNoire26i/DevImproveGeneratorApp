package actions

import (

  "fmt"
  "github.com/gobuffalo/buffalo"
  "github.com/gobuffalo/pop"
  "generatedApps/appimprovegeneratorapp202012326cb7bdde2e44/models"
)

// This file is generated by Buffalo. It offers a basic structure for
// adding, editing and deleting a page. If your model is more
// complex or you need more than the basic implementation you need to
// edit this file.

// Following naming logic is implemented in Buffalo:
// Model: Singular (Placeportid)
// DB Table: Plural (placeportids)
// Resource: Plural (Placeportids)
// Path: Plural (/placeportids)
// View Template Folder: Plural (/templates/placeportids/)

// PlaceportidsResource is the resource for the Placeportid model
type PlaceportidsResource struct{
  buffalo.Resource
}

// List gets all Placeportids. This function is mapped to the path
// GET /placeportids
func (v PlaceportidsResource) List(c buffalo.Context) error {
  // Get the DB connection from the context
  tx, ok := c.Value("tx").(*pop.Connection)
  if !ok {
    return fmt.Errorf("no transaction found")
  }

  placeportids := &models.Placeportids{}

  // Paginate results. Params "page" and "per_page" control pagination.
  // Default values are "page=1" and "per_page=20".
  q := tx.PaginateFromParams(c.Params())

  // Retrieve all Placeportids from the DB
  if err := q.All(placeportids); err != nil {
    return err
  }

  // Add the paginator to the context so it can be used in the template.
  c.Set("pagination", q.Paginator)

  return c.Render(200, r.Auto(c, placeportids))
}

// Show gets the data for one Placeportid. This function is mapped to
// the path GET /placeportids/{placeportid_id}
func (v PlaceportidsResource) Show(c buffalo.Context) error {
  // Get the DB connection from the context
  tx, ok := c.Value("tx").(*pop.Connection)
  if !ok {
    return fmt.Errorf("no transaction found")
  }

  // Allocate an empty Placeportid
  placeportid := &models.Placeportid{}

  // To find the Placeportid the parameter placeportid_id is used.
  if err := tx.Find(placeportid, c.Param("placeportid_id")); err != nil {
    return c.Error(404, err)
  }

  return c.Render(200, r.Auto(c, placeportid))
}

// Create adds a Placeportid to the DB. This function is mapped to the
// path POST /placeportids
func (v PlaceportidsResource) Create(c buffalo.Context) error {
  // Allocate an empty Placeportid
  placeportid := &models.Placeportid{}

  // Bind placeportid to the html form elements
  if err := c.Bind(placeportid); err != nil {
    return err
  }

  // Get the DB connection from the context
  tx, ok := c.Value("tx").(*pop.Connection)
  if !ok {
    return fmt.Errorf("no transaction found")
  }

  // Validate the data from the html form
  verrs, err := tx.ValidateAndCreate(placeportid)
  if err != nil {
    return err
  }

  if verrs.HasAny() {
    // Make the errors available inside the html template
    c.Set("errors", verrs)

    // Render again the new.html template that the user can
    // correct the input.
    return c.Render(422, r.Auto(c, placeportid))
  }

  // and redirect to the placeportids index page
  return c.Render(201, r.Auto(c, placeportid))
}

// Update changes a Placeportid in the DB. This function is mapped to
// the path PUT /placeportids/{placeportid_id}
func (v PlaceportidsResource) Update(c buffalo.Context) error {
  // Get the DB connection from the context
  tx, ok := c.Value("tx").(*pop.Connection)
  if !ok {
    return fmt.Errorf("no transaction found")
  }

  // Allocate an empty Placeportid
  placeportid := &models.Placeportid{}

  if err := tx.Find(placeportid, c.Param("placeportid_id")); err != nil {
    return c.Error(404, err)
  }

  // Bind Placeportid to the html form elements
  if err := c.Bind(placeportid); err != nil {
    return err
  }

  verrs, err := tx.ValidateAndUpdate(placeportid)
  if err != nil {
    return err
  }

  if verrs.HasAny() {
    // Make the errors available inside the html template
    c.Set("errors", verrs)

    // Render again the edit.html template that the user can
    // correct the input.
    return c.Render(422, r.Auto(c, placeportid))
  }

  // and redirect to the placeportids index page
  return c.Render(200, r.Auto(c, placeportid))
}

// Destroy deletes a Placeportid from the DB. This function is mapped
// to the path DELETE /placeportids/{placeportid_id}
func (v PlaceportidsResource) Destroy(c buffalo.Context) error {
  // Get the DB connection from the context
  tx, ok := c.Value("tx").(*pop.Connection)
  if !ok {
    return fmt.Errorf("no transaction found")
  }

  // Allocate an empty Placeportid
  placeportid := &models.Placeportid{}

  // To find the Placeportid the parameter placeportid_id is used.
  if err := tx.Find(placeportid, c.Param("placeportid_id")); err != nil {
    return c.Error(404, err)
  }

  if err := tx.Destroy(placeportid); err != nil {
    return err
  }

  // Redirect to the placeportids index page
  return c.Render(200, r.Auto(c, placeportid))
}
