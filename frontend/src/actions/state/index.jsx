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
        url: 'http://aurs7r4.local:8080/api/v1/state',
    })
        .then(response => dispatch(getRemoteStateSuccess(response.data)))
        .catch(error => error.response);
    return {
        type: 'FETCH_REMOTE_STATE',
        payload: request
    }
}

