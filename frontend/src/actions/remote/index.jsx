import { API_URL } from '../../config/env';
import httpClient from '../../httpClient';

export function fetchRemoteSuccess(payload) {
    return {
        type: 'FETCH_REMOTE_SUCCESS',
        payload: payload
    }
}

export function fetchRemoteError(error) {
    return {
        type: 'FETCH_REMOTE_FAILED',
        error: error,
    }
}

export function fetchRemote(dispatch) {
    const request = httpClient({
        method: 'GET',
        url: `${API_URL}/api/v1/remote`,
    })
        .then(response => dispatch(fetchRemoteSuccess(response.data)))
        .catch(error => dispatch(fetchRemoteError(error.response)));
    return {
        type: 'FETCH_REMOTE',
        payload: request
    }
}

export function fetchRemoteByMode(dispatch, mode) {
    const request = httpClient({
        method: 'GET',
        url: `${API_URL}/api/v1/remote`,
        params: {
            mode: mode
        }
    })
        .then(response => dispatch(fetchRemoteSuccess(response.data)))
        .catch(error => dispatch(fetchRemoteError(error.response)));
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

export function postRemoteError(error) {
    console.log('failed')
    return {
        type: 'POST_REMOTE_FAILED',
        error: error,
    }
}

export function postRemote(dispatch, payload) {
    console.log('payload')
    console.log(payload)
    const request = httpClient({
        method: 'POST',
        url: `${API_URL}/api/v1/remote`,
        headers: {'Content-Type': 'application/json'},
        data: payload,
    })
        .then(response => dispatch(postRemoteSuccess(response.data)))
        .catch(error => dispatch(postRemoteError(error.response))
        );
    return {
        type: 'POST_REMOTE',
        payload: request
    }
}
