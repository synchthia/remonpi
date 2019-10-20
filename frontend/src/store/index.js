/*import { createStore, applyMiddleware  } from 'redux';
import thunk from 'redux-thunk';
import rootReducer from '../reducers';

const configureStore = initialState => {
    return createStore(
        rootReducer,
        initialState,
        applyMiddleware()
    );
}

export default configureStore*/


import { createStore, applyMiddleware } from 'redux';
import thunk from 'redux-thunk';
import logger from 'redux-logger';
import reducers from '../reducers';

const middlewares = [thunk];
if (process.env.NODE_ENV !== 'production') {
      const reduxImmutableStateInvariant = require('redux-immutable-state-invariant').default();
    middlewares.push(reduxImmutableStateInvariant, logger);
}

export default createStore(reducers, applyMiddleware(...middlewares));
