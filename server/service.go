package main

import (
    "context"
    "fmt"
    taskpb "github.com/zeddmathews/tasksync/proto"
)

// type TaskServiceServer struct {
//     taskpb.UnimplementedTaskServiceServer
// }

func (s *TaskServiceServer) CreateTask(ctx context.Context, req *taskpb.CreateRequest) (tr *taskpb.CreateResponse, err error) {
    fmt.Println("Received CreateTask request:", req)

    // Simulate creating a task and generating an ID
    var taskID int32 = 123
	tr = &taskpb.CreateResponse{
		    Res: taskID,
	}
	err = nil
	return
    // return &taskpb.TaskResponse{
    //     Id:     taskID,
    //     Status: "Created",
    // }, nil
}

