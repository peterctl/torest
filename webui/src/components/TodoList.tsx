import * as React from 'react';
import ITodo from 'src/models/ITodo';
import Todo from './Todo';
import TodoInput from './TodoInput';

import TodoService from 'src/services/TodoService';

export interface ITodoListState {
    todos: ITodo[];
}

export interface ITodoListProps {
    todos: ITodo[];
}

const service = new TodoService();

export default class TodoList extends React.Component<ITodoListProps, ITodoListState> {

    public constructor(props: ITodoListProps) {
        super(props);
        this.state = { todos: [] };
    }

    public render() {
        const pending = this.state.todos.filter(t => !t.completed);
        const done = this.state.todos.filter(t => t.completed);

        return (
            <div className="container">
                <h3>Add Item</h3>
                <TodoInput onNewTodo={this.onNewTodo} />

                <h3>Pending</h3>
                <ul id="incomplete-tasks">
                    {pending.map(t =>
                        <Todo key={t.id} todo={t}
                            onToggle={this.onToggle}
                            onEdit={this.onEdit}
                            onDelete={this.onDelete} />
                    )}
                </ul>
        
                <h3>Completed</h3>
                <ul id="completed-tasks">
                    {done.map(t =>
                        <Todo key={t.id} todo={t}
                            onToggle={this.onToggle}
                            onEdit={this.onEdit}
                            onDelete={this.onDelete} />
                    )}
                </ul>
            </div>
        );
    }

    public componentDidMount() {
        this.fetchTodos();
    }

    public async fetchTodos() {
        const todos = await service.all();
        this.setState({ todos });
    }

    public onNewTodo = async (name: string) => {
        await service.create(name);
        await this.fetchTodos();
    }

    // public onNewTodoLocal = (name: string) => {
    //     const todo = {
    //         completed: false,
    //         id: Math.random().toString(),
    //         name,
    //     };

    //     const todos = [
    //         ...this.state.todos,
    //         todo,
    //     ];

    //     // TODO: Handle new todos on the server side.

    //     this.setState({ todos });
    // }

    public onToggle = async (id: string, completed: boolean) => {
        await service.update(id, { completed });
        await this.fetchTodos();
    }

    // public onToggleLocal = (id: string) => {
    //     const todos = this.state.todos.slice();
    //     const todo = todos.find(t => t.id === id)!;

    //     todo.completed = !todo.completed;

    //     // TODO: Handle toggles on the server side.

    //     this.setState({ todos });
    // }

    public onEdit = async (id: string, name: string) => {
        await service.update(id, { name });
        await this.fetchTodos();
    }

    // public onEditLocal = (id: string, name: string) => {
    //     const todos = this.state.todos.slice();
    //     const todo = todos.find(t => t.id === id)!;

    //     todo.name = name;

    //     // TODO: Handle edits on the server side.

    //     this.setState({ todos });
    // }

    public onDelete = async (id: string) => {
        await service.delete(id);
        await this.fetchTodos();
    }

    // public onDeleteLocal = (id: string) => {
    //     const todos = this.state.todos.slice();
    //     this.setState({ todos: todos.filter(t => t.id !== id) });
    //     // TODO: Handle deletes on the server side.
    // }
}
