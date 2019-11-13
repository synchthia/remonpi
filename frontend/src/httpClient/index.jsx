import Axios from 'axios';

const axios = Axios.create({});
axios.interceptors.response.use(
    response => response,
    error => {
        if (error.message === 'Network Error') {
            return Promise.reject({ response: { status: 500, error } });
        }
        return Promise.reject(error);
    }
);

export default axios;
