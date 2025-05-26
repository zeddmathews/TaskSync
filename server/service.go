package main

import (
    "context"
    "fmt"
    // "database/sql"
    taskpb "github.com/zeddmathews/tasksync/proto"
)

type TaskServiceServer struct {
    taskpb.UnimplementedTaskServiceServer
}

func (s *TaskServiceServer) CreateTask(ctx context.Context, req *taskpb.CreateRequest) (tr *taskpb.CreateResponse, err error) {
    fmt.Println("Received CreateTask request:", req)
    query := `INSERT INTO tasks (messages) VALUES ($1) RETURNING id`
    var id int32

    err = DB.QueryRowContext(ctx, query, req.Message).Scan(&id)
    if err != nil {
        return
    }
    // var taskID int32 = 123
	tr = &taskpb.CreateResponse{
		    Res: id,
	}
	err = nil
	return
    // return &taskpb.TaskResponse{
    //     Id:     taskID,
    //     Status: "Created",
    // }, nil
}

func (s * TaskServiceServer) GetTask(ctx context.Context, req *taskpb.GetRequest) (tr *taskpb.GetResponse, err error) {
    query := `SELECT messages FROM tasks WHERE id = $1`
    var message string

    err = DB.QueryRowContext(ctx, query, req.Req).Scan(&message)
    if err != nil {
        return
    }
    tr = &taskpb.GetResponse{
        Message: message,
    }
    return
}

