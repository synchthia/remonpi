const initialState = {
    payload: null
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
        default:
            return state
    }
}
