package model

import (
	"time"

	"github.com/globalsign/mgo/bson"
)

// User is the person who completed a task
type User struct {
	ID        bson.ObjectId `bson:"_id" json:"id"`
	FirstName string        `bson:"firstname" json:"firstname"`
	LastName  string        `bson:"lastname" json:"lastname"`
	NickName  string        `bson:"nickname" json:"nickname"`
}

// Lists is a container for lists which will be returned in their preferred order
type Lists struct {
	Items []List
}

// List is a container for tasks which will be returned in priority order
type List struct {
	ID       bson.ObjectId `bson:"_id" json:"id"`
	ListName string        `bson:"listname" json:"listname"`
	Tasks    []Task        `bson:"tasks" json:"tasks"`
}

// Task represents the thing to be done
type Task struct {
	ID       bson.ObjectId `bson:"_id" json:"id"`
	Name     string        `bson:"taskname" json:"taskname"`
	Priority int           `bson:"priority" json:"priority"`
}

// TaskCompletionHistory is the list of events related to the completion of a single task
type TaskCompletionHistory struct {
	History []TaskCompletionEvent
}

// TaskCompletionEvent is a record of the date and time of a task completion and who
// completed that task
type TaskCompletionEvent struct {
	ID                bson.ObjectId `bson:"_id" json:"id"`
	DateTimeCompleted time.Time     `bson:"datetime_completed" json:"datetime_completed"`
	CompletedBy       User
}

// APIError is the error object from lists api
type ListsAPIError struct {
	// ErrorCode equal to the HTTP error coder
	ErrorCode int `json:"code" bson:",omitempty"`
	// ErrorMessage
	ErrorMessage string `json:"error" bson:",omitempty"`
}
