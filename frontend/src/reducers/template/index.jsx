const initialState = {
    payload: null,
    error: null
}
export const template = (state = initialState, action) => {
    switch (action.type) {
        case 'FETCH_TEMPLATE':
            return {
                ...state,
                payload: null
            }
        case 'FETCH_TEMPLATE_SUCCESS':
            return {
                ...state,
                payload: action.payload
            }
        case 'FETCH_TEMPLATE_FAILED':
            return {
                ...state,
                error: action.error
            }
        default:
            return state
    }
}
