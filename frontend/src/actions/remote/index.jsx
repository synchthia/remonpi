import httpClient from '../../httpClient';

export function fetchRemoteSuccess(payload) {
    return {
        type: 'FETCH_REMOTE_SUCCESS',
        payload: payload
    }
}

export function fetchRemote(dispatch) {
    const request = httpClient({
        method: 'GET',
        url: 'http://aurs7r4.local:8080/api/v1/remote',
    })
        .then(response => dispatch(fetchRemoteSuccess(response.data)))
        .catch(error => error.response);
    return {
        type: 'FETCH_REMOTE',
        payload: request
    }
}

export function fetchRemoteByMode(dispatch, mode) {
    const request = httpClient({
        method: 'GET',
        url: 'http://aurs7r4.local:8080/api/v1/remote',
        params: {
            mode: mode
        }
    })
        .then(response => dispatch(fetchRemoteSuccess(response.data)))
        .catch(error => error.response);
    return {
        type: 'FETCH_REMOTE',
        payload: request
    }
}

export function saveRemote(dispatch, payload) {
    dispatch(fetchRemoteSuccess(payload))
    return {
        type: 'SAVE_REMOTE',
        payload: payload
    }
}

export function postRemoteSuccess(payload) {
    return {
        type: 'POST_REMOTE_SUCCESS',
        payload: payload
    }
}

export function postRemote(dispatch, payload) {
    console.log('payload')
    console.log(payload)
    const request = httpClient({
        method: 'POST',
        url: 'http://aurs7r4.local:8080/api/v1/remote',
        headers: {'Content-Type': 'application/json'},
        data: payload,
    })
        .then(response => dispatch(postRemoteSuccess(response.data)))
        .catch(error => error.response);
    return {
        type: 'POST_REMOTE',
        payload: request
    }
}
