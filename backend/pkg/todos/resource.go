package todos

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"golang.org/x/exp/maps"
)

type TodoItem struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

type CreateTodoData struct {
	Title string `json:"title"`
}

type UpdateTodoData struct {
	Title string `json:"title"`
}

var todoItems = make(map[int]TodoItem)

// HandleTodoList is a handler for the GET /todo endpoint.
func HandleTodoList(ctx echo.Context) error {
    values := maps.Values(todoItems)

	return ctx.JSONPretty(http.StatusOK, values, "    ")
}

// HandleTodoDetails is a handler for the GET /todo/:id endpoint.
func HandleTodoDetails(ctx echo.Context) error {
    id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
    if err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "Id must be an integer")
    }


    item, exists := todoItems[int(id)]
    if !exists {
        return echo.NewHTTPError(http.StatusNotFound, "Todo item not found")
    }

    return ctx.JSONPretty(http.StatusOK, item, "    ")
}

// HandleTodoCreate is a handler for the POST /todo endpoint.
func HandleCreateTodo(ctx echo.Context) error {
    var body CreateTodoData
    if err := ctx.Bind(&body); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, err)
    }

    id := len(todoItems) + 1
    todoItem := TodoItem{
        ID:    id,
        Title: body.Title,
        Done:  false,
    }

    todoItems[id] = todoItem

    return ctx.JSONPretty(http.StatusCreated, todoItem, "    ")
}

// HandleTodoUpdate is a handler for the PUT /todo/:id endpoint.
func HandleUpdateTodo(ctx echo.Context) error {
    id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
    if err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "Id must be an integer")
    }

    item, exists := todoItems[int(id)]
    if !exists {
        return echo.NewHTTPError(http.StatusNotFound, "Todo item not found")
    }

    var body UpdateTodoData
    if err := ctx.Bind(&body); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, err)
    }

    item.Title = body.Title

    todoItems[int(id)] = item

    return ctx.NoContent(http.StatusNoContent)
}

// HandleDoneTodo is a handler for the PUT /todo/:id/done endpoint.
func HandleDoneTodo(ctx echo.Context) error {
    id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
    if err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "Id must be an integer")
    }

    item, exists := todoItems[int(id)]
    if !exists {
        return echo.NewHTTPError(http.StatusNotFound, "Todo item not found")
    }

    item.Done = true

    todoItems[int(id)] = item

    return ctx.NoContent(http.StatusNoContent)
}

// HandleUndoneTodo is a handler for the PUT /todo/:id/undone endpoint.
func HandleUndoneTodo(ctx echo.Context) error {
    id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
    if err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "Id must be an integer")
    }

    item, exists := todoItems[int(id)]
    if !exists {
        return echo.NewHTTPError(http.StatusNotFound, "Todo item not found")
    }

    item.Done = false

    todoItems[int(id)] = item

    return ctx.NoContent(http.StatusNoContent)
}

// HandleTodoDelete is a handler for the DELETE /todo/:id endpoint.
func HandleDeleteTodo(ctx echo.Context) error {
    id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
    if err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "Id must be an integer")
    }

    _, exists := todoItems[int(id)]
    if !exists {
        return echo.NewHTTPError(http.StatusNotFound, "Todo item not found")
    }

    delete(todoItems, int(id))

    return ctx.NoContent(http.StatusNoContent)
}
