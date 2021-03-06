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
// Model: Singular (Chat)
// DB Table: Plural (chats)
// Resource: Plural (Chats)
// Path: Plural (/chats)
// View Template Folder: Plural (/templates/chats/)

// ChatsResource is the resource for the Chat model
type ChatsResource struct{
  buffalo.Resource
}

// List gets all Chats. This function is mapped to the path
// GET /chats
func (v ChatsResource) List(c buffalo.Context) error {
  // Get the DB connection from the context
  tx, ok := c.Value("tx").(*pop.Connection)
  if !ok {
    return fmt.Errorf("no transaction found")
  }

  chats := &models.Chats{}

  // Paginate results. Params "page" and "per_page" control pagination.
  // Default values are "page=1" and "per_page=20".
  q := tx.PaginateFromParams(c.Params())

  // Retrieve all Chats from the DB
  if err := q.All(chats); err != nil {
    return err
  }

  // Add the paginator to the context so it can be used in the template.
  c.Set("pagination", q.Paginator)

  return c.Render(200, r.Auto(c, chats))
}

// Show gets the data for one Chat. This function is mapped to
// the path GET /chats/{chat_id}
func (v ChatsResource) Show(c buffalo.Context) error {
  // Get the DB connection from the context
  tx, ok := c.Value("tx").(*pop.Connection)
  if !ok {
    return fmt.Errorf("no transaction found")
  }

  // Allocate an empty Chat
  chat := &models.Chat{}

  // To find the Chat the parameter chat_id is used.
  if err := tx.Find(chat, c.Param("chat_id")); err != nil {
    return c.Error(404, err)
  }

  return c.Render(200, r.Auto(c, chat))
}

// Create adds a Chat to the DB. This function is mapped to the
// path POST /chats
func (v ChatsResource) Create(c buffalo.Context) error {
  // Allocate an empty Chat
  chat := &models.Chat{}

  // Bind chat to the html form elements
  if err := c.Bind(chat); err != nil {
    return err
  }

  // Get the DB connection from the context
  tx, ok := c.Value("tx").(*pop.Connection)
  if !ok {
    return fmt.Errorf("no transaction found")
  }

  // Validate the data from the html form
  verrs, err := tx.ValidateAndCreate(chat)
  if err != nil {
    return err
  }

  if verrs.HasAny() {
    // Make the errors available inside the html template
    c.Set("errors", verrs)

    // Render again the new.html template that the user can
    // correct the input.
    return c.Render(422, r.Auto(c, chat))
  }

  // and redirect to the chats index page
  return c.Render(201, r.Auto(c, chat))
}

// Update changes a Chat in the DB. This function is mapped to
// the path PUT /chats/{chat_id}
func (v ChatsResource) Update(c buffalo.Context) error {
  // Get the DB connection from the context
  tx, ok := c.Value("tx").(*pop.Connection)
  if !ok {
    return fmt.Errorf("no transaction found")
  }

  // Allocate an empty Chat
  chat := &models.Chat{}

  if err := tx.Find(chat, c.Param("chat_id")); err != nil {
    return c.Error(404, err)
  }

  // Bind Chat to the html form elements
  if err := c.Bind(chat); err != nil {
    return err
  }

  verrs, err := tx.ValidateAndUpdate(chat)
  if err != nil {
    return err
  }

  if verrs.HasAny() {
    // Make the errors available inside the html template
    c.Set("errors", verrs)

    // Render again the edit.html template that the user can
    // correct the input.
    return c.Render(422, r.Auto(c, chat))
  }

  // and redirect to the chats index page
  return c.Render(200, r.Auto(c, chat))
}

// Destroy deletes a Chat from the DB. This function is mapped
// to the path DELETE /chats/{chat_id}
func (v ChatsResource) Destroy(c buffalo.Context) error {
  // Get the DB connection from the context
  tx, ok := c.Value("tx").(*pop.Connection)
  if !ok {
    return fmt.Errorf("no transaction found")
  }

  // Allocate an empty Chat
  chat := &models.Chat{}

  // To find the Chat the parameter chat_id is used.
  if err := tx.Find(chat, c.Param("chat_id")); err != nil {
    return c.Error(404, err)
  }

  if err := tx.Destroy(chat); err != nil {
    return err
  }

  // Redirect to the chats index page
  return c.Render(200, r.Auto(c, chat))
}
