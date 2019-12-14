import { Button, Navbar, Row, Spinner } from 'react-bootstrap';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { connect } from 'react-redux';
import { fas } from '@fortawesome/free-solid-svg-icons';
import { library } from '@fortawesome/fontawesome-svg-core'
import React from 'react';
import styled from 'styled-components';

import {
  fetchRemote,
  fetchRemoteByMode,
  postRemote
} from '../../actions/remote';
import { fetchTemplate } from '../../actions/template';
import { getRemoteState, saveRemoteState } from '../../actions/state';
import Range from './Range';
import Shot from './Shot';
import Step from './Step';
import Toggle from './Toggle';

const mapStateToProps = state => {
    return {
        remote: state.remote.payload,
        remoteState: state.remoteState.payload,
        template: state.template.payload
    }
};
const mapDispatchToProps = (dispatch, props) => ({
    getRemoteState() {
        dispatch(getRemoteState(dispatch))
    },
    saveRemoteState(entry) {
        dispatch(saveRemoteState(dispatch, entry))
    },
    getRemote() {
        dispatch(fetchRemote(dispatch))
    },
    getRemoteByMode(mode) {
        dispatch(fetchRemoteByMode(dispatch, mode))
    },
    pushRemote(entry) {
        dispatch(postRemote(dispatch, entry))
    },
    loadTemplate() {
        dispatch(fetchTemplate(dispatch))
    }
});

class RemoteCard extends React.Component {
    constructor(props) {
        super(props);
        this.props.loadTemplate()
        this.props.getRemoteState()
        library.add(fas)
    }

    // Convert RemoteData to State
    dataToState(entry) {
        let state = {...this.props.remoteState};

        // Must be deep-copy children nodes
        state = {
            operation: entry.operation,
            mode: entry.mode,
            mode_data: {...state.mode_data}
        }

        state.mode_data[entry.mode] = {
            temp: entry.temp,
            fan: entry.fan,
            horizontal_vane: entry.horizontal_vane,
            vertical_vane: entry.vertical_vane
        }

        return state;
    }

    // Convert State to Remote
    stateToData(m) {
        // m: mode
        const remoteState = {...this.props.remoteState};
        const operation = remoteState['operation'];
        const mode = m ? m : remoteState['mode'];
        return {
            operation: operation,
            mode: mode,
            ...remoteState['mode_data'][mode],
        }
    }

    render() {
        //const remote = {...this.state.remote};
        if (!this.props.remoteState || !this.props.template) {
            return (
                <div>
                    <Navbar variant="light">
                        <Navbar.Brand>
                            <FontAwesomeIcon spin icon={["fa", "asterisk"]} />
                            <span style={{margin: "0.5em"}}>{'LOADING...'}</span>
                        </Navbar.Brand>
                    </Navbar>
                </div>
            )
        }

        const remote = this.stateToData();

        return (
            <div>
                <Navbar variant="light">
                    <Navbar.Brand>
                        <FontAwesomeIcon icon={["fas", "terminal"]} />
                        <span style={{margin: "0.5em"}}>{'Control'}</span>
                    </Navbar.Brand>
                    <Navbar.Collapse id="basic-navbar-nav">
                        <Button
                            onClick={() => {
                                    this.props.getRemoteState()
                                }}
                        >Sync</Button>
                        {true ?
                        <Button
                            onClick={() => {
                                this.props.pushRemote(remote);
                            }}>Re-Send</Button> :
                        <p></p>
                        }
                    </Navbar.Collapse>
                </Navbar>
                <CardBase>
                    <Row>
                        <Toggle
                            name="Operation"
                            status={remote.operation ? "ON" : "OFF"}
                            onClick={() => {
                                remote.operation = !remote.operation;
                                this.props.saveRemoteState(this.dataToState(remote));
                                this.props.pushRemote(remote)

                            }}
                                icon={<FontAwesomeIcon icon={["fas", "power-off"]} />}
                            />

                            {remote.temp ?
                            <Range
                                name="Temp"
                                suffix="℃"
                                rangeFrom={this.props.template[remote.mode].temp.range.from}
                                rangeTo={this.props.template[remote.mode].temp.range.to}
                                status={remote.temp ? remote.temp : "--"}
                                onIncrement={() => {
                                    remote.temp = remote.temp + this.props.template[remote.mode].temp.range.step;
                                    this.props.saveRemoteState(this.dataToState(remote));
                                    this.props.pushRemote(remote)
                                }}
                                onDecrement={() => {
                                    remote.temp = remote.temp - this.props.template[remote.mode].temp.range.step;
                                    this.props.saveRemoteState(this.dataToState(remote));
                                    this.props.pushRemote(remote)
                                }}
                                /> : <p></p>
                            }
                        </Row>
                        <hr />
                        <Row>
                            <Step
                                name="Mode"
                                status={remote.mode}
                                contents={Object.keys(this.props.template)}
                                onClick={(key) => {
                                    if (Object.keys(this.props.template).includes(key)) {
                                        this.props.saveRemoteState(this.dataToState(remote));

                                        // Get RemoteData from local State
                                        const r = this.stateToData(key);
                                        this.props.saveRemoteState(this.dataToState(r));
                                        this.props.pushRemote(r)
                                    }
                                }}
                            />
                            {remote.fan ?
                                <Step
                                    name="Fan"
                                    status={remote.fan}
                                    contents={this.props.template[remote.mode].fan.step}
                                    onClick={(key) => {
                                        console.log('fan')
                                        remote.fan = key
                                        this.props.saveRemoteState(this.dataToState(remote));
                                        this.props.pushRemote(remote)
                                    }}
                                /> : <p></p>
                            }
                        </Row>
                        <hr />
                        <Row>
                            {remote.horizontal_vane ?
                            <Step
                                name="Horizontal Vane"
                                status={remote.horizontal_vane}
                                contents={this.props.template[remote.mode].horizontal_vane.step}
                                onClick={(key) => {
                                    remote.horizontal_vane = key
                                    this.props.saveRemoteState(this.dataToState(remote));
                                    this.props.pushRemote(remote)
                                }}
                            /> : <p></p>
                            }
                            {remote.vertical_vane ?
                            <Shot
                                name="Vertical Vane"
                                status={remote.vertical_vane}
                                contents="Toggle"
                                onClick={(key) => {
                                    remote.vertical_vane = key
                                    this.props.saveRemoteState(this.dataToState(remote));
                                    this.props.pushRemote(remote)
                                }}
                            /> : <p></p>
                            }
                        </Row>
                    </CardBase>
                </div>
        )
    }
}

const CardBase = styled.div`
    background-color: #FAFAFA;
`

//export default RemoteCard;
export default connect(
    mapStateToProps,
    mapDispatchToProps
)(RemoteCard);
