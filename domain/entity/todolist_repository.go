package entity

import "context"

type TodoListRepository interface {
	List(ctx context.Context) ([]*TodoList, error)
	Find(ctx context.Context, id string) (*TodoList, error)
	Store(ctx context.Context, todo *TodoList) (*TodoList, error)
	Delete(ctx context.Context, id string) (int, error)
}
