const initialState = {
    payload: null
}
export const remote = (state = initialState, action) => {
    switch (action.type) {
        case 'FETCH_REMOTE':
            return {
                ...state,
                payload: null
            }
        case 'FETCH_REMOTE_SUCCESS':
            return {
                ...state,
                payload: action.payload
            }
        case 'SAVE_REMOTE':
            return {
                ...state,
                payload: action.payload
            }
        case 'POST_REMOTE':
            return {
                ...state,
            }
        case 'POST_REMOTE_SUCCESS':
            return {
                ...state,
                payload: action.payload
            }
        default:
            return state
    }
}

