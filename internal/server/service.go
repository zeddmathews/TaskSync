package server

import (
	"context"

	taskpb "github.com/zeddmathews/tasksync/proto"
	"google.golang.org/grpc"
)

type TaskServiceServer struct {
    taskpb.UnimplementedTaskServiceServer
}
func RegisterServices(grpcServer *grpc.Server) {
	taskpb.RegisterTaskServiceServer(grpcServer, &TaskServiceServer{})
}
func (s *TaskServiceServer) CreateTask(ctx context.Context, req *taskpb.CreateRequest) (res *taskpb.CreateResponse, err error) {
    // fmt.Println("Received CreateTask request:", req)
    query := `INSERT INTO tasks (messages) VALUES ($1) RETURNING id`
    var id int32

    err = DB.QueryRowContext(ctx, query, req.Message).Scan(&id)
    if err != nil {
        return nil, err
    }

	res = &taskpb.CreateResponse{
		    Res: id,
	}
	err = nil
	return
}

func (s * TaskServiceServer) GetTask(ctx context.Context, req *taskpb.GetRequest) (res *taskpb.GetResponse, err error) {
    query := `SELECT messages FROM tasks WHERE id = $1`
    var message string

    err = DB.QueryRowContext(ctx, query, req.Req).Scan(&message)
    if err != nil {
        return nil, err
    }
    res = &taskpb.GetResponse{
        Message: message,
    }
    return
}

func (s * TaskServiceServer) UpdateTask (ctx context.Context, req *taskpb.UpdateRequest) (res *taskpb.Empty, err error) {
    query := `UPDATE tasks SET messages = $1 WHERE id = $2`
    _, err = DB.ExecContext(ctx, query, req.Message, req.Id)
    if err != nil {
        return nil, err
    }
    res = &taskpb.Empty{}
    return
}

func (s * TaskServiceServer) DeleteTask (ctx context.Context, req * taskpb.DeleteRequest) (res *taskpb.Empty, err error) {
    query := `DELETE FROM tasks WHERE id = $1`
    _, err = DB.ExecContext(ctx, query, req.Id)
    if err != nil {
        return nil, err
    }
    res = &taskpb.Empty{}
    return
}