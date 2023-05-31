package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/AdjectiveAllison/tasks/app/model"
	"github.com/mailru/easyjson"
	"github.com/syumai/workers/cloudflare/d1"
	_ "github.com/syumai/workers/cloudflare/d1" // register driver
)

type taskHandler struct{}

var _ http.Handler = (*taskHandler)(nil)

func NewTaskHandler() http.Handler {
	return &taskHandler{}
}

func (h *taskHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// initialize DB.
	// D1 connector requires request's context to initialize DB.
	c, err := d1.OpenConnector(req.Context(), "taskDB")
	if err != nil {
		h.handleErr(w, http.StatusInternalServerError, fmt.Sprintf("failed to initialize DB: %v", err))
	}
	// use sql.OpenDB instead of sql.Open.
	db := sql.OpenDB(c)

	switch req.Method {
	case http.MethodGet:
		h.listTasks(w, req, db)
		return
		// TODO: Add a default to add an easteregg if someone attempts to use an alternative HTTP method.

	}
	w.WriteHeader(http.StatusNotFound)
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("not found"))
}

func (h *taskHandler) handleErr(w http.ResponseWriter, status int, msg string) {
	// TODO handle errors in a more standard JSON API style.
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(msg))
}

func (h *taskHandler) listTasks(w http.ResponseWriter, req *http.Request, db *sql.DB) {
	rows, err := db.Query(`
SELECT id, title, description, links, updated_at, completed FROM tasks
ORDER BY updated_at DESC;
   `)
	if err != nil {
		h.handleErr(w, http.StatusInternalServerError,
			"failed to load tasks")
		return
	}

	tasks := []model.Task{}
	for rows.Next() {
		var t model.Task
		var linksJson string
		err = rows.Scan(&t.ID, &t.Title, &t.Description, &linksJson, &t.UpdatedAt, &t.Completed)
		if err != nil {
			log.Println(err)
			h.handleErr(w, http.StatusInternalServerError,
				"failed to scan task")
			return
		}

		// Use the custom unmarshalLinks function.
		linksSlice, err := unmarshalLinks([]byte(linksJson))
		if err != nil {
			log.Println(err)
			h.handleErr(w, http.StatusInternalServerError,
				"failed to unmarshal links")
			return
		}

		t.Links = linksSlice

		tasks = append(tasks, t)
	}

	res := model.ListTasksResponse{
		Tasks: tasks,
	}

	acceptHeader := req.Header.Get("Accept")

	if strings.ToLower(acceptHeader) == "application/json" {
		w.Header().Set("Content-Type", "application/json")
		if _, err := easyjson.MarshalToWriter(&res, w); err != nil {
			h.handleErr(w, http.StatusInternalServerError, "failed to encode response")
		}
	} else {
		// HTML response
		html := getTaskHTML(res)
		w.Write([]byte(html))
	}
}

func getTaskHTML(tasksResponse model.ListTasksResponse) string {
	html := `<!DOCTYPE html>
<html>
<head>
	<title>Tasks</title>
	<style>
		table {
			width: 100%;
			border-collapse: collapse;
		}
		th, td {
			border: 1px solid #ddd;
			padding: 8px;
		}
	</style>
</head>
<body>
	<table>
		<thead>
			<tr>
				<th>ID</th>
				<th>Title</th>
				<th>Description</th>
				<th>Links</th>
				<th>Updated At</th>
				<th>Completed</th>
			</tr>
		</thead>
		<tbody>`
	for _, task := range tasksResponse.Tasks {
		html += fmt.Sprintf(`
			<tr>
				<td>%d</td>
				<td>%s</td>
				<td>%s</td>
				<td><ul>`, task.ID, task.Title, task.Description)
		for _, link := range task.Links {
			html += fmt.Sprintf("<li><a href=\"%s\">%s</a></li>", link, link)
		}
		html += `</ul></td>`

		updatedAtTime := time.Unix(int64(task.UpdatedAt), 0)
		html += fmt.Sprintf(`
			<td>%s</td>
			<td>%t</td>
		</tr>`, updatedAtTime.Format(time.RFC3339), task.Completed)
	}

	html += `</tbody></table></body></html>`

	return html
}

// This is only needed because I don't know how easyJson works, tinygo doesn't support reflection yet(version 28 on the way) and I didn't take the right approach in the SQL schema. There is a solution over in the schema file that includes multiple tables.
func unmarshalLinks(data []byte) ([]string, error) {
	// Prepare a slice to hold the unmarshaled links.
	var links []string

	// The current link being read.
	var link []byte

	// Parse the JSON manually.
	for i := 0; i < len(data); i++ {
		c := data[i]

		switch {
		// Ignore the opening and closing brackets and commas.
		case c == '[' || c == ']' || c == ',':
			continue
		// When a double quote is encountered, append the link (if any) to the links slice.
		case c == '"':
			if link != nil {
				str := string(link)
				// Only append the string to links if it isn't just spaces.
				if strings.TrimSpace(str) != "" {
					links = append(links, str)
				}
				link = nil
			}
		// Otherwise, append the character to the current link.
		default:
			link = append(link, c)
		}
	}

	return links, nil
}
