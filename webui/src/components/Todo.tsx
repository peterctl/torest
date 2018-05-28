import * as React from 'react';
import ITodo from 'src/models/ITodo';

export interface ITodoProps {
    todo: ITodo;
    onToggle: (id: string, completed: boolean) => void;
    onEdit: (id: string, name: string) => void;
    onDelete: (id: string) => void;
}

export interface ITodoState {
    editing: boolean;
    name: string;
}

export default class Todo extends React.Component<ITodoProps, ITodoState> {

    public constructor(props: ITodoProps) {
        super(props);
        this.state = {
            editing: false,
            name: props.todo.name,
        };
    }

    public render() {
        const className = this.state.editing ? 'todo editing' : 'todo';

        return (
            <li className={className}>
                <input type="checkbox" onChange={this.onToggle}
                    checked={this.props.todo.completed} />

                <label>{this.props.todo.name}</label>

                <input type="text" value={this.state.name}
                    onChange={this.updateName}
                    onKeyPress={this.handleKeyPress} />

                <button className="edit" onClick={this.onEdit}>
                    {this.state.editing ? 'Done' : 'Edit'}
                </button>

                <button className="delete" onClick={this.onDelete}>
                    Delete
                </button>
            </li>
        );
    }

    public handleKeyPress = (e: React.KeyboardEvent<HTMLInputElement>) => {
        if (e.key === 'Enter') {
            this.onEdit();
        }
    }

    public updateName = (e: React.ChangeEvent<HTMLInputElement>) => {
        this.setState({ name: e.target.value });
    }

    public onToggle = () => {
        this.props.onToggle(this.props.todo.id, !this.props.todo.completed);
    }

    public onEdit = () => {
        if (this.state.editing === false) {
            this.setState({ editing: true });
        } else {
            if (this.state.name !== this.props.todo.name) {
                this.props.onEdit(this.props.todo.id, this.state.name);
            }
            this.setState({ editing: false });
        }
    }

    public onDelete = () => {
        this.props.onDelete(this.props.todo.id);
    }
}
