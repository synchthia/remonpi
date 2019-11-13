import { combineReducers } from 'redux';

import { template } from './template';
import { remote } from './remote';
import { remoteState } from './state';

export default combineReducers({
    template,
    remote,
    remoteState
});
