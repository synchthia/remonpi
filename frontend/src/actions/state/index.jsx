import { API_URL } from '../../config/env';
import httpClient from '../../httpClient';

export function getRemoteStateSuccess(payload) {
    return {
        type: 'FETCH_REMOTE_STATE_SUCCESS',
        payload: payload
    }
}

export function getRemoteState(dispatch) {
    const request = httpClient({
        method: 'GET',
        url: `${API_URL}/api/v1/state`
    })
        .then(response => dispatch(getRemoteStateSuccess(response.data)))
        .catch(error => error.response);
    return {
        type: 'FETCH_REMOTE_STATE',
        payload: request
    }
}
