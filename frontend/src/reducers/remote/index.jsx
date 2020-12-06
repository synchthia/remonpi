const initialState = {
    payload: null
}
export const remote = (state = initialState, action) => {
    switch (action.type) {
        case 'FETCH_REMOTE':
            return {
                ...state,
                payload: null,
                error: null
            }
        case 'FETCH_REMOTE_SUCCESS':
            return {
                ...state,
                payload: action.payload,
                error: null
            }
        case 'FETCH_REMOTE_FAILED':
            return {
                ...state,
                error: action.error
            }
        case 'SAVE_REMOTE':
            return {
                ...state,
                payload: action.payload,
                error: null
            }
        case 'POST_REMOTE':
            return {
                ...state,
                error: null
            }
        case 'POST_REMOTE_SUCCESS':
            return {
                ...state,
                payload: action.payload,
                error: null
            }
        case 'POST_REMOTE_FAILED':
            return {
                ...state,
                error: action.error
            }
        default:
            return state
    }
}

