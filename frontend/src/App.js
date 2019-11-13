import './components/themes/theme.css'
import './components/themes/ui.css';

import { Container } from 'react-bootstrap';
import { Provider } from 'react-redux';
import React from 'react';

import Navigation from './components/basement/Navbar';
import RemoteCard from './components/remote/RemoteCard';
import store from './store';

const App = () => {
    return (
        <Provider store={store}>
            <Navigation/>
            <Container>
                <RemoteCard/>
            </Container>
        </Provider>
    );
}

export default App;
