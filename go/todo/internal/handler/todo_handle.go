package handler
import (
    "encoding/json"
    "net/http"
    "todo/internal/service"
)

type TodoHandler struct {
  svc service.TodoService
}

func NewServiceHandler (svc *service.TodoService) *TodoHandler { //will be called in main.go and gets initialised with svc
  return &TodoHandler{svc:svc}
}

func (h *TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
 var body struct {
        Title string `json:"title"`
    }

    if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    todo, err := h.svc.CreateTodo(r.Context(), body.Title) //pass the title to service layer createToda 
    if err != nil {
        http.Error(w, "Failed to create todo", http.StatusInternalServerError)
        return
    }
 w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(todo) 
}

