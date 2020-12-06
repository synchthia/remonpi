import { API_URL } from '../../config/env';
import httpClient from '../../httpClient';

export function fetchTemplateSuccess(payload) {
    return {
        type: 'FETCH_TEMPLATE_SUCCESS',
        payload: payload
    }
}

export function fetchTemplateError(error) {
    return {
        type: 'FETCH_TEMPLATE_FAILED',
        error: error,
    }
}

export function fetchTemplate(dispatch) {
    const request = httpClient({
        method: 'GET',
        url: `${API_URL}/api/v1/template`
    }).then(response => dispatch(fetchTemplateSuccess(response.data))
    ).catch(error => dispatch(fetchTemplateError(error.response)));
    return {
        type: 'FETCH_TEMPLATE',
        payload: request
    }
}
