package main

import (
    "context"
    "fmt"
    taskpb "github.com/zeddmathews/tasksync/proto"
)

type TaskServiceServer struct {
    taskpb.UnimplementedTaskServiceServer
}

func (s *TaskServiceServer) CreateTask(ctx context.Context, req *taskpb.TaskRequest) (tr *taskpb.TaskResponse, err error) {
    fmt.Println("Received CreateTask request:", req)

    // Simulate creating a task and generating an ID
    taskID := "task-123"
	tr = &taskpb.TaskResponse{
		Id: taskID,
		Status: "Created",
	}
	err = nil
	return
    // return &taskpb.TaskResponse{
    //     Id:     taskID,
    //     Status: "Created",
    // }, nil
}

