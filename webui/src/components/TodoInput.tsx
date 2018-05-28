import * as React from 'react';

export interface ITodoProps {
    onNewTodo: (name: string) => void;
}

export interface ITodoState {
    name: string;
}

export default class Todo extends React.Component<ITodoProps, ITodoState> {

    public state = {
        name: '',
    };

    public render() {
        return (
            <p>
                <input id="new-task" type="text"
                    value={this.state.name}
                    onChange={this.updateName}
                    onKeyPress={this.handleKeyPress} />
                <button onClick={this.onButtonClick}>Add</button>
            </p>
        );
    }

    public handleKeyPress = (e: React.KeyboardEvent<HTMLInputElement>) => {
        if (e.key === 'Enter') {
            this.finish();
        }
    }

    public updateName = (e: React.ChangeEvent<HTMLInputElement>) => {
        this.setState({ name: e.target.value });
    }

    public onButtonClick = () => {
        this.finish();
    }

    public finish() {
        const name = this.state.name;
        this.props.onNewTodo(name);
        this.setState({ name: '' });
    }
}
