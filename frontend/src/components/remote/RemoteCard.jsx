import { Button, Navbar, Row, Spinner } from 'react-bootstrap';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { connect } from 'react-redux';
import { faPowerOff, faTerminal } from '@fortawesome/free-solid-svg-icons'
import React from 'react';
import styled from 'styled-components';

import {
    fetchRemote,
    fetchRemoteByMode,
    postRemote,
    saveRemote
} from '../../actions/remote';
import { fetchTemplate } from '../../actions/template';
import { getRemoteState } from '../../actions/state';
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
    saveRemote(entry) {
        dispatch(saveRemote(dispatch, entry))
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
        this.state = {
            stashed: false
        }
        this.props.loadTemplate()
        this.props.getRemoteState()
        this.props.getRemote()
    }

    // Propsから複製して、編集した奴をDispatchすれば良いのでは？

    componentDidMount() {
        console.log('didmount')
        console.log(this.state)
        console.log(this.props)
    }

    componentDidUpdate(prevProps, prevState) {
        console.log('didupdate')
        console.log(this.state)
        console.log(this.props)
    }

    saveState(entry) {
        this.props.saveRemote(entry);
        this.setState({
            stashed: true
        });
    }

    render() {
        //fetchTemplate().then(r =>
        //    console.log(r)
        //)

        console.log("Rendering...")
        console.log(this.state)
        console.log(this.props)

        //const remote = {...this.state.remote};
        if (!this.props.remote || !this.props.template) {
            return (
                <div>
                    <Spinner animation="border" />
                    <span style={{fontSize: "2em"}}> Loading...</span>
                </div>
            )
        }
        //const remote = {
        //    operation: this.props.remoteState['operation'],
        //    mode: this.props.remoteState['mode'],
        //    ...this.props.remoteState['mode_data'][this.props.remoteState['mode']],
        //}
        const remote = {...this.props.remote};

        const remoteState = {...this.props.remoteState};

        return (
            <div>
                <Navbar variant="light">
                    <Navbar.Brand>
                        <FontAwesomeIcon icon={faTerminal} />
                        <span style={{margin: "0.5em"}}>{'Debug //'}</span>
                    </Navbar.Brand>
                    <Navbar.Collapse id="basic-navbar-nav">
                        <Button
                            status={"N/A"}
                            onClick={
                                () => {
                                    this.props.loadTemplate()
                                    this.props.getRemote()
                                }
                            }
                            icon={<FontAwesomeIcon icon={faPowerOff} />}
                        >Sync</Button>

                        {<Button>Send</Button>}
                    </Navbar.Collapse>
                </Navbar>
                <CardBase>
                    <Row>
                        <Toggle
                            name="Operation"
                            status={remote.operation ? "ON" : "OFF"}
                            onClick={() => {
                                remote.operation = !remote.operation;
                                //this.props.saveRemote(remote)
                                this.saveState(remote)
                                this.props.pushRemote(remote)

                            }}
                                icon={<FontAwesomeIcon icon={faPowerOff} />}
                            />

                            {remote.temp ?
                            <Range
                                name="Temp"
                                suffix="℃"
                                rangeFrom={this.props.template[remote.mode].temp.range.from}
                                rangeTo={this.props.template[remote.mode].temp.range.to}
                                status={remote.temp ? remote.temp : "--"}
                                onIncrement={() => {
                                    remote.temp = remote.temp + 1;
                                    this.props.pushRemote(remote)
                                }}
                                    onDecrement={() => {
                                        remote.temp = remote.temp - 1;
                                        this.props.pushRemote(remote)
                                    }}
                                />: <p></p>
                            }
                        </Row>
                        <hr />
                        <Row>
                            <Step
                                name="Mode"
                                status={remote.mode}
                                contents={["cool", "dry", "heat"]}
                                onClick={(key) => {
                                    if (["cool", "dry", "heat"].includes(key)) {
                                        //this.props.getRemoteByMode(key)
                                        const newRemote = {
                                            operation: remote.operation,
                                            mode: key,
                                            ...remoteState['mode_data'][key],
                                        }
                                        console.log('change mode')
                                        console.log(newRemote)
                                        this.props.saveRemote(newRemote)
                                        this.props.pushRemote(newRemote)
                                    }
                                }}
                            />
                            <Step
                                name="Fan"
                                status={remote.fan}
                                contents={this.props.template[remote.mode].fan.step}
                                onClick={(key) => {
                                    remote.Fan = key
                                    this.props.pushRemote(remote)
                                }}
                            />
                        </Row>
                        <hr />
                        <Row>
                            <Step
                                name="Vertical Vane"
                                status={remote.vertical_vane}
                                contents={this.props.template[remote.mode].vertical_vane.step}
                                onClick={(key) => {
                                    remote.vertical_vane = key
                                    this.props.pushRemote(remote)
                                }}
                            />
                            <Shot
                                name="Horizontal Vane"
                                status={remote.horizontal_vane}
                                contents="Toggle"
                                onClick={(key) => {
                                    remote.horizontal_vane = key
                                    this.props.pushRemote(remote)
                                }}
                            />
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
