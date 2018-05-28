import * as React from 'react';
import * as ReactDOM from 'react-dom';
// import App from './App';
import registerServiceWorker from './registerServiceWorker';

import TodoList from './components/TodoList';
import ITodo from './models/ITodo';

import './index.css';

const TODOS: ITodo[] = [
  {
    completed: false,
    id: '1',
    name: 'First',
  },
  {
    completed: true,
    id: '2',
    name: 'Second',
  },
  {
    completed: false,
    id: '3',
    name: 'Third',
  },
];


ReactDOM.render(
  <TodoList todos={TODOS} />,
  document.getElementById('root') as HTMLElement
);
registerServiceWorker();
