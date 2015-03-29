// Package taskqueue provides access to the TaskQueue API.
//
// See https://developers.google.com/appengine/docs/python/taskqueue/rest
//
// Usage example:
//
//   import "github.com/jfcote87/api/taskqueue/v1beta2"
//   ...
//   taskqueueService, err := taskqueue.New(oauthHttpClient)
package taskqueue

import (
	"errors"
	"fmt"
	"github.com/jfcote87/api/googleapi"
	"golang.org/x/net/context"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Always reference these packages, just in case the auto-generated code
// below doesn't.
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = io.Copy
var _ = url.Parse
var _ = googleapi.Version
var _ = errors.New
var _ = strings.Replace
var _ = context.Background

const apiId = "taskqueue:v1beta2"
const apiName = "taskqueue"
const apiVersion = "v1beta2"

var baseURL *url.URL = &url.URL{Scheme: "https", Host: "www.googleapis.com", Path: "/taskqueue/v1beta2/projects/"}

// OAuth2 scopes used by this API.
const (
	// Manage your Tasks and Taskqueues
	TaskqueueScope = "https://www.googleapis.com/auth/taskqueue"

	// Consume Tasks from your Taskqueues
	TaskqueueConsumerScope = "https://www.googleapis.com/auth/taskqueue.consumer"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client}
	s.Taskqueues = NewTaskqueuesService(s)
	s.Tasks = NewTasksService(s)
	return s, nil
}

type Service struct {
	client *http.Client

	Taskqueues *TaskqueuesService

	Tasks *TasksService
}

func NewTaskqueuesService(s *Service) *TaskqueuesService {
	rs := &TaskqueuesService{s: s}
	return rs
}

type TaskqueuesService struct {
	s *Service
}

func NewTasksService(s *Service) *TasksService {
	rs := &TasksService{s: s}
	return rs
}

type TasksService struct {
	s *Service
}

type Task struct {
	// EnqueueTimestamp: Time (in seconds since the epoch) at which the task
	// was enqueued.
	EnqueueTimestamp int64 `json:"enqueueTimestamp,omitempty,string"`

	// Id: Name of the task.
	Id string `json:"id,omitempty"`

	// Kind: The kind of object returned, in this case set to task.
	Kind string `json:"kind,omitempty"`

	// LeaseTimestamp: Time (in seconds since the epoch) at which the task
	// lease will expire. This value is 0 if the task isnt currently leased
	// out to a worker.
	LeaseTimestamp int64 `json:"leaseTimestamp,omitempty,string"`

	// PayloadBase64: A bag of bytes which is the task payload. The payload
	// on the JSON side is always Base64 encoded.
	PayloadBase64 string `json:"payloadBase64,omitempty"`

	// QueueName: Name of the queue that the task is in.
	QueueName string `json:"queueName,omitempty"`

	// Retry_count: The number of leases applied to this task.
	Retry_count int64 `json:"retry_count,omitempty"`

	// Tag: Tag for the task, could be used later to lease tasks grouped by
	// a specific tag.
	Tag string `json:"tag,omitempty"`
}

type TaskQueue struct {
	// Acl: ACLs that are applicable to this TaskQueue object.
	Acl *TaskQueueAcl `json:"acl,omitempty"`

	// Id: Name of the taskqueue.
	Id string `json:"id,omitempty"`

	// Kind: The kind of REST object returned, in this case taskqueue.
	Kind string `json:"kind,omitempty"`

	// MaxLeases: The number of times we should lease out tasks before
	// giving up on them. If unset we lease them out forever until a worker
	// deletes the task.
	MaxLeases int64 `json:"maxLeases,omitempty"`

	// Stats: Statistics for the TaskQueue object in question.
	Stats *TaskQueueStats `json:"stats,omitempty"`
}

type TaskQueueAcl struct {
	// AdminEmails: Email addresses of users who are "admins" of the
	// TaskQueue. This means they can control the queue, eg set ACLs for the
	// queue.
	AdminEmails []string `json:"adminEmails,omitempty"`

	// ConsumerEmails: Email addresses of users who can "consume" tasks from
	// the TaskQueue. This means they can Dequeue and Delete tasks from the
	// queue.
	ConsumerEmails []string `json:"consumerEmails,omitempty"`

	// ProducerEmails: Email addresses of users who can "produce" tasks into
	// the TaskQueue. This means they can Insert tasks into the queue.
	ProducerEmails []string `json:"producerEmails,omitempty"`
}

type TaskQueueStats struct {
	// LeasedLastHour: Number of tasks leased in the last hour.
	LeasedLastHour int64 `json:"leasedLastHour,omitempty,string"`

	// LeasedLastMinute: Number of tasks leased in the last minute.
	LeasedLastMinute int64 `json:"leasedLastMinute,omitempty,string"`

	// OldestTask: The timestamp (in seconds since the epoch) of the oldest
	// unfinished task.
	OldestTask int64 `json:"oldestTask,omitempty,string"`

	// TotalTasks: Number of tasks in the queue.
	TotalTasks int64 `json:"totalTasks,omitempty"`
}

type Tasks struct {
	// Items: The actual list of tasks returned as a result of the lease
	// operation.
	Items []*Task `json:"items,omitempty"`

	// Kind: The kind of object returned, a list of tasks.
	Kind string `json:"kind,omitempty"`
}

type Tasks2 struct {
	// Items: The actual list of tasks currently active in the TaskQueue.
	Items []*Task `json:"items,omitempty"`

	// Kind: The kind of object returned, a list of tasks.
	Kind string `json:"kind,omitempty"`
}

// method id "taskqueue.taskqueues.get":

type TaskqueuesGetCall struct {
	s             *Service
	project       string
	taskqueue     string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Get detailed information about a TaskQueue.

func (r *TaskqueuesService) Get(project string, taskqueue string) *TaskqueuesGetCall {
	return &TaskqueuesGetCall{
		s:             r.s,
		project:       project,
		taskqueue:     taskqueue,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/taskqueues/{taskqueue}",
	}
}

// GetStats sets the optional parameter "getStats": Whether to get
// stats.
func (c *TaskqueuesGetCall) GetStats(getStats bool) *TaskqueuesGetCall {
	c.params_.Set("getStats", fmt.Sprintf("%v", getStats))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TaskqueuesGetCall) Fields(s ...googleapi.Field) *TaskqueuesGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *TaskqueuesGetCall) Do() (*TaskQueue, error) {
	var returnValue *TaskQueue
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":   c.project,
		"taskqueue": c.taskqueue,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Get detailed information about a TaskQueue.",
	//   "httpMethod": "GET",
	//   "id": "taskqueue.taskqueues.get",
	//   "parameterOrder": [
	//     "project",
	//     "taskqueue"
	//   ],
	//   "parameters": {
	//     "getStats": {
	//       "description": "Whether to get stats. Optional.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "project": {
	//       "description": "The project under which the queue lies.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "taskqueue": {
	//       "description": "The id of the taskqueue to get the properties of.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/taskqueues/{taskqueue}",
	//   "response": {
	//     "$ref": "TaskQueue"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/taskqueue",
	//     "https://www.googleapis.com/auth/taskqueue.consumer"
	//   ]
	// }

}

// method id "taskqueue.tasks.delete":

type TasksDeleteCall struct {
	s             *Service
	project       string
	taskqueue     string
	task          string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Delete: Delete a task from a TaskQueue.

func (r *TasksService) Delete(project string, taskqueue string, task string) *TasksDeleteCall {
	return &TasksDeleteCall{
		s:             r.s,
		project:       project,
		taskqueue:     taskqueue,
		task:          task,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/taskqueues/{taskqueue}/tasks/{task}",
	}
}

func (c *TasksDeleteCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":   c.project,
		"taskqueue": c.taskqueue,
		"task":      c.task,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Delete a task from a TaskQueue.",
	//   "httpMethod": "DELETE",
	//   "id": "taskqueue.tasks.delete",
	//   "parameterOrder": [
	//     "project",
	//     "taskqueue",
	//     "task"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "The project under which the queue lies.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "task": {
	//       "description": "The id of the task to delete.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "taskqueue": {
	//       "description": "The taskqueue to delete a task from.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/taskqueues/{taskqueue}/tasks/{task}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/taskqueue",
	//     "https://www.googleapis.com/auth/taskqueue.consumer"
	//   ]
	// }

}

// method id "taskqueue.tasks.get":

type TasksGetCall struct {
	s             *Service
	project       string
	taskqueue     string
	task          string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Get a particular task from a TaskQueue.

func (r *TasksService) Get(project string, taskqueue string, task string) *TasksGetCall {
	return &TasksGetCall{
		s:             r.s,
		project:       project,
		taskqueue:     taskqueue,
		task:          task,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/taskqueues/{taskqueue}/tasks/{task}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TasksGetCall) Fields(s ...googleapi.Field) *TasksGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *TasksGetCall) Do() (*Task, error) {
	var returnValue *Task
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":   c.project,
		"taskqueue": c.taskqueue,
		"task":      c.task,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Get a particular task from a TaskQueue.",
	//   "httpMethod": "GET",
	//   "id": "taskqueue.tasks.get",
	//   "parameterOrder": [
	//     "project",
	//     "taskqueue",
	//     "task"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "The project under which the queue lies.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "task": {
	//       "description": "The task to get properties of.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "taskqueue": {
	//       "description": "The taskqueue in which the task belongs.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/taskqueues/{taskqueue}/tasks/{task}",
	//   "response": {
	//     "$ref": "Task"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/taskqueue",
	//     "https://www.googleapis.com/auth/taskqueue.consumer"
	//   ]
	// }

}

// method id "taskqueue.tasks.insert":

type TasksInsertCall struct {
	s             *Service
	project       string
	taskqueue     string
	task          *Task
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Insert: Insert a new task in a TaskQueue

func (r *TasksService) Insert(project string, taskqueue string, task *Task) *TasksInsertCall {
	return &TasksInsertCall{
		s:             r.s,
		project:       project,
		taskqueue:     taskqueue,
		task:          task,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/taskqueues/{taskqueue}/tasks",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TasksInsertCall) Fields(s ...googleapi.Field) *TasksInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *TasksInsertCall) Do() (*Task, error) {
	var returnValue *Task
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":   c.project,
		"taskqueue": c.taskqueue,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.task,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Insert a new task in a TaskQueue",
	//   "httpMethod": "POST",
	//   "id": "taskqueue.tasks.insert",
	//   "parameterOrder": [
	//     "project",
	//     "taskqueue"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "The project under which the queue lies",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "taskqueue": {
	//       "description": "The taskqueue to insert the task into",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/taskqueues/{taskqueue}/tasks",
	//   "request": {
	//     "$ref": "Task"
	//   },
	//   "response": {
	//     "$ref": "Task"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/taskqueue",
	//     "https://www.googleapis.com/auth/taskqueue.consumer"
	//   ]
	// }

}

// method id "taskqueue.tasks.lease":

type TasksLeaseCall struct {
	s             *Service
	project       string
	taskqueue     string
	numTasks      int64
	leaseSecs     int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Lease: Lease 1 or more tasks from a TaskQueue.

func (r *TasksService) Lease(project string, taskqueue string, numTasks int64, leaseSecs int64) *TasksLeaseCall {
	return &TasksLeaseCall{
		s:             r.s,
		project:       project,
		taskqueue:     taskqueue,
		numTasks:      numTasks,
		leaseSecs:     leaseSecs,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/taskqueues/{taskqueue}/tasks/lease",
	}
}

// GroupByTag sets the optional parameter "groupByTag": When true, all
// returned tasks will have the same tag
func (c *TasksLeaseCall) GroupByTag(groupByTag bool) *TasksLeaseCall {
	c.params_.Set("groupByTag", fmt.Sprintf("%v", groupByTag))
	return c
}

// Tag sets the optional parameter "tag": The tag allowed for tasks in
// the response. Must only be specified if group_by_tag is true. If
// group_by_tag is true and tag is not specified the tag will be that of
// the oldest task by eta, i.e. the first available tag
func (c *TasksLeaseCall) Tag(tag string) *TasksLeaseCall {
	c.params_.Set("tag", fmt.Sprintf("%v", tag))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TasksLeaseCall) Fields(s ...googleapi.Field) *TasksLeaseCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *TasksLeaseCall) Do() (*Tasks, error) {
	var returnValue *Tasks
	c.params_.Set("leaseSecs", fmt.Sprintf("%v", c.leaseSecs))
	c.params_.Set("numTasks", fmt.Sprintf("%v", c.numTasks))
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":   c.project,
		"taskqueue": c.taskqueue,
	})
	call := &googleapi.Call{
		Method: "POST",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Lease 1 or more tasks from a TaskQueue.",
	//   "httpMethod": "POST",
	//   "id": "taskqueue.tasks.lease",
	//   "parameterOrder": [
	//     "project",
	//     "taskqueue",
	//     "numTasks",
	//     "leaseSecs"
	//   ],
	//   "parameters": {
	//     "groupByTag": {
	//       "description": "When true, all returned tasks will have the same tag",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "leaseSecs": {
	//       "description": "The lease in seconds.",
	//       "format": "int32",
	//       "location": "query",
	//       "required": true,
	//       "type": "integer"
	//     },
	//     "numTasks": {
	//       "description": "The number of tasks to lease.",
	//       "format": "int32",
	//       "location": "query",
	//       "required": true,
	//       "type": "integer"
	//     },
	//     "project": {
	//       "description": "The project under which the queue lies.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "tag": {
	//       "description": "The tag allowed for tasks in the response. Must only be specified if group_by_tag is true. If group_by_tag is true and tag is not specified the tag will be that of the oldest task by eta, i.e. the first available tag",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "taskqueue": {
	//       "description": "The taskqueue to lease a task from.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/taskqueues/{taskqueue}/tasks/lease",
	//   "response": {
	//     "$ref": "Tasks"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/taskqueue",
	//     "https://www.googleapis.com/auth/taskqueue.consumer"
	//   ]
	// }

}

// method id "taskqueue.tasks.list":

type TasksListCall struct {
	s             *Service
	project       string
	taskqueue     string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: List Tasks in a TaskQueue

func (r *TasksService) List(project string, taskqueue string) *TasksListCall {
	return &TasksListCall{
		s:             r.s,
		project:       project,
		taskqueue:     taskqueue,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/taskqueues/{taskqueue}/tasks",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TasksListCall) Fields(s ...googleapi.Field) *TasksListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *TasksListCall) Do() (*Tasks2, error) {
	var returnValue *Tasks2
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":   c.project,
		"taskqueue": c.taskqueue,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "List Tasks in a TaskQueue",
	//   "httpMethod": "GET",
	//   "id": "taskqueue.tasks.list",
	//   "parameterOrder": [
	//     "project",
	//     "taskqueue"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "The project under which the queue lies.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "taskqueue": {
	//       "description": "The id of the taskqueue to list tasks from.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/taskqueues/{taskqueue}/tasks",
	//   "response": {
	//     "$ref": "Tasks2"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/taskqueue",
	//     "https://www.googleapis.com/auth/taskqueue.consumer"
	//   ]
	// }

}

// method id "taskqueue.tasks.patch":

type TasksPatchCall struct {
	s               *Service
	project         string
	taskqueue       string
	task            string
	newLeaseSeconds int64
	task2           *Task
	caller_         googleapi.Caller
	params_         url.Values
	pathTemplate_   string
}

// Patch: Update tasks that are leased out of a TaskQueue. This method
// supports patch semantics.

func (r *TasksService) Patch(project string, taskqueue string, task string, newLeaseSeconds int64, task2 *Task) *TasksPatchCall {
	return &TasksPatchCall{
		s:               r.s,
		project:         project,
		taskqueue:       taskqueue,
		task:            task,
		newLeaseSeconds: newLeaseSeconds,
		task2:           task2,
		caller_:         googleapi.JSONCall{},
		params_:         make(map[string][]string),
		pathTemplate_:   "{project}/taskqueues/{taskqueue}/tasks/{task}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TasksPatchCall) Fields(s ...googleapi.Field) *TasksPatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *TasksPatchCall) Do() (*Task, error) {
	var returnValue *Task
	c.params_.Set("newLeaseSeconds", fmt.Sprintf("%v", c.newLeaseSeconds))
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":   c.project,
		"taskqueue": c.taskqueue,
		"task":      c.task,
	})
	call := &googleapi.Call{
		Method:  "PATCH",
		URL:     u,
		Params:  c.params_,
		Payload: c.task2,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Update tasks that are leased out of a TaskQueue. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "taskqueue.tasks.patch",
	//   "parameterOrder": [
	//     "project",
	//     "taskqueue",
	//     "task",
	//     "newLeaseSeconds"
	//   ],
	//   "parameters": {
	//     "newLeaseSeconds": {
	//       "description": "The new lease in seconds.",
	//       "format": "int32",
	//       "location": "query",
	//       "required": true,
	//       "type": "integer"
	//     },
	//     "project": {
	//       "description": "The project under which the queue lies.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "task": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "taskqueue": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/taskqueues/{taskqueue}/tasks/{task}",
	//   "request": {
	//     "$ref": "Task"
	//   },
	//   "response": {
	//     "$ref": "Task"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/taskqueue",
	//     "https://www.googleapis.com/auth/taskqueue.consumer"
	//   ]
	// }

}

// method id "taskqueue.tasks.update":

type TasksUpdateCall struct {
	s               *Service
	project         string
	taskqueue       string
	task            string
	newLeaseSeconds int64
	task2           *Task
	caller_         googleapi.Caller
	params_         url.Values
	pathTemplate_   string
}

// Update: Update tasks that are leased out of a TaskQueue.

func (r *TasksService) Update(project string, taskqueue string, task string, newLeaseSeconds int64, task2 *Task) *TasksUpdateCall {
	return &TasksUpdateCall{
		s:               r.s,
		project:         project,
		taskqueue:       taskqueue,
		task:            task,
		newLeaseSeconds: newLeaseSeconds,
		task2:           task2,
		caller_:         googleapi.JSONCall{},
		params_:         make(map[string][]string),
		pathTemplate_:   "{project}/taskqueues/{taskqueue}/tasks/{task}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TasksUpdateCall) Fields(s ...googleapi.Field) *TasksUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *TasksUpdateCall) Do() (*Task, error) {
	var returnValue *Task
	c.params_.Set("newLeaseSeconds", fmt.Sprintf("%v", c.newLeaseSeconds))
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":   c.project,
		"taskqueue": c.taskqueue,
		"task":      c.task,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.task2,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Update tasks that are leased out of a TaskQueue.",
	//   "httpMethod": "POST",
	//   "id": "taskqueue.tasks.update",
	//   "parameterOrder": [
	//     "project",
	//     "taskqueue",
	//     "task",
	//     "newLeaseSeconds"
	//   ],
	//   "parameters": {
	//     "newLeaseSeconds": {
	//       "description": "The new lease in seconds.",
	//       "format": "int32",
	//       "location": "query",
	//       "required": true,
	//       "type": "integer"
	//     },
	//     "project": {
	//       "description": "The project under which the queue lies.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "task": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "taskqueue": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/taskqueues/{taskqueue}/tasks/{task}",
	//   "request": {
	//     "$ref": "Task"
	//   },
	//   "response": {
	//     "$ref": "Task"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/taskqueue",
	//     "https://www.googleapis.com/auth/taskqueue.consumer"
	//   ]
	// }

}
