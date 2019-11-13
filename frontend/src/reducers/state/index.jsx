const initialState = {
    payload: null
}
export const remoteState = (state = initialState, action) => {
    switch (action.type) {
        case 'FETCH_REMOTE_STATE':
            return {
                ...state,
                payload: null
            }
        case 'FETCH_REMOTE_STATE_SUCCESS':
            return {
                ...state,
                payload: action.payload
            }
        case 'SAVE_REMOTE_STATE':
            return {
                ...state,
                payload: action.payload
            }
        default:
            return state
    }
}

