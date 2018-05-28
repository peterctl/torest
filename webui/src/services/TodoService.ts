import axios from 'axios';

import ITodo from 'src/models/ITodo';

const BASE_URL = '/api/todos';

export default class TodoService {
    
    public async all(): Promise<ITodo[]> {
        const res = await axios.get<ITodo[]>(BASE_URL);
        return res.data;
    }

    public async find(id: string): Promise<ITodo> {
        const res = await axios.get<ITodo>(`${BASE_URL}/${id}`);
        return res.data;
    }

    public async create(name: string): Promise<ITodo> {
        const res = await axios.post<ITodo>(BASE_URL, { name });
        return res.data;
    }

    public async update(id: string, todo: Partial<ITodo>): Promise<ITodo> {
        const res = await axios.patch<ITodo>(`${BASE_URL}/${id}`, todo);
        return res.data;
    }

    public async delete(id: string): Promise<ITodo> {
        const res = await axios.delete(`${BASE_URL}/${id}`);
        return res.data as ITodo;
    }
}
