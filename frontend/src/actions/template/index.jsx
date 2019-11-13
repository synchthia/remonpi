import { API_URL } from '../../config/env';
import httpClient from '../../httpClient';

export function fetchTemplateSuccess(payload) {
    return {
        type: 'FETCH_TEMPLATE_SUCCESS',
        payload: payload
    }
}

export function fetchTemplate(dispatch) {
    const request = httpClient({
        method: 'GET',
        url: `${API_URL}/api/v1/template`
    }).then(response => dispatch(fetchTemplateSuccess(response.data))
    ).catch(error => error.response);
    return {
        type: 'FETCH_TEMPLATE',
        payload: request
    }
}
