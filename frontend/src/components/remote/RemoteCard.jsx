import { Button, Navbar, Row, Toast } from 'react-bootstrap';
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
import ModalPopup from '../basement/Modal';
import Range from './Range';
import Shot from './Shot';
import Step from './Step';
import Toggle from './Toggle';

const mapStateToProps = state => {
    return {
        remote: {
            data: state.remote.payload,
            error: state.remote.error
        },
        //remoteState: {
        //    data: state.remoteState.payload,
        //    error: state.remoteState.error
        //},
        template: {
            data: state.template.payload,
            error: state.template.error
        }
    }
};
const mapDispatchToProps = (dispatch, props) => ({
    //getRemoteState() {
    //    dispatch(getRemoteState(dispatch))
    //},
    //saveRemoteState(entry) {
    //    dispatch(saveRemoteState(dispatch, entry))
    //},
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
        //this.props.getRemoteState()
        this.props.getRemote()
        library.add(fas)
        this.remote = {}
    }

    render() {
        //const remote = {...this.state.remote};
        //if (!this.props.remote.data || !this.props.template) {
        //    return (
        //        <div>
        //            <Navbar variant="light">
        //                <Navbar.Brand>
        //                    <FontAwesomeIcon spin icon={["fa", "asterisk"]} />
        //                    <span style={{margin: "0.5em"}}>{'LOADING...'}</span>
        //                </Navbar.Brand>
        //            </Navbar>
        //        </div>
        //    )
        //}

        console.log("error entries:")
        console.log(this.props.remote.error)
        console.log(this.props.template.error)
        if (this.props.template.error) {
            console.log("Something went wrong. Action will be reverted.")
            return (
                <ModalPopup
                    title={
                        <div>
                            <FontAwesomeIcon icon={["fas", "exclamation-triangle"]} />
                            <span style={{margin: "0.5em"}}>Error</span>
                        </div>
                    }
                    body={`${this.props.template.error.error}`}
                    closable={false}
                    type={'ERROR_MODAL'}
                    onReload={()=> {
                        window.location.reload()
                            //this.props.loadTemplate()
                            //this.props.getRemote()
                        }
                    }
                >
                </ModalPopup>
            )
            //return (
            //    <div>
            //        <p>Something went wrong...</p>
            //    </div>
            //)
        }

        if (this.props.remote.data != null) {
            this.remote = {...this.props.remote.data}
        }

        const remote = this.remote;
        //const remote = {...this.props.remote.data};
        console.log(remote)
        return (
            <div>
                {
                    this.props.remote.error &&
                        <ModalPopup
                            show={true}
                            title="Failed"
                            body={this.props.remote.error}
                        >
                        </ModalPopup>
                }

                <Navbar variant="light">
                    <Navbar.Brand>
                        <FontAwesomeIcon icon={["fas", "terminal"]} />
                        <span style={{margin: "0.5em"}}>{'Control'}</span>
                    </Navbar.Brand>
                    <Navbar.Collapse id="basic-navbar-nav">
                        <Button
                            onClick={() => {
                                //this.props.getRemoteState()
                                this.props.getRemote()
                            }}
                        >Sync</Button>
                        <Button
                            onClick={() => {
                                this.props.pushRemote(remote);
                            }}>Re-Send
                        </Button>
                    </Navbar.Collapse>
                </Navbar>
                {(remote != null && this.props.template.data) != null &&
                <CardBase>
                    <Row>
                        <Toggle
                            name="Operation"
                            status={remote.operation ? "ON" : "OFF"}
                            onClick={() => {
                                remote.operation = !remote.operation;
                                //this.props.saveRemoteState(this.dataToState(remote));
                                this.props.pushRemote(remote)

                            }}
                            icon={<FontAwesomeIcon icon={["fas", "power-off"]} />}
                        />

                        {remote.temp &&
                        <Range
                            name="Temp"
                            suffix="℃"
                            rangeFrom={this.props.template.data[remote.mode].temp.range.from}
                            rangeTo={this.props.template.data[remote.mode].temp.range.to}
                            status={remote.temp ? remote.temp : "--"}
                            onIncrement={() => {
                                remote.temp = remote.temp + this.props.template.data[remote.mode].temp.range.step;
                                //this.props.saveRemoteState(this.dataToState(remote));
                                this.props.pushRemote(remote)
                            }}
                            onDecrement={() => {
                                remote.temp = remote.temp - this.props.template.data[remote.mode].temp.range.step;
                                //this.props.saveRemoteState(this.dataToState(remote));
                                this.props.pushRemote(remote)
                            }}
                        />
                        }
                    </Row>
                    <hr />
                    <Row>
                        <Step
                            name="Mode"
                            status={remote.mode}
                            contents={Object.keys(this.props.template.data)}
                            onClick={(key) => {
                                if (Object.keys(this.props.template.data).includes(key)) {
                                    this.props.saveRemoteState(this.dataToState(remote));

                                    // Get RemoteData from local State
                                    const r = this.stateToData(key);
                                    this.props.saveRemoteState(this.dataToState(r));
                                    this.props.pushRemote(r)
                                }
                            }}
                        />
                        {remote.fan &&
                        <Step
                            name="Fan"
                            status={remote.fan}
                            contents={this.props.template.data[remote.mode].fan.step}
                            onClick={(key) => {
                                console.log('fan')
                                remote.fan = key
                                //this.props.saveRemoteState(this.dataToState(remote));
                                this.props.pushRemote(remote)
                            }}
                        />
                        }
                    </Row>
                    <hr />
                    <Row>
                        {remote.horizontal_vane &&
                        <Step
                            name="Horizontal Vane"
                            status={remote.horizontal_vane}
                            contents={this.props.template.data[remote.mode].horizontal_vane.step}
                            onClick={(key) => {
                                remote.horizontal_vane = key
                                //this.props.saveRemoteState(this.dataToState(remote));
                                this.props.pushRemote(remote)
                            }}
                        />
                        }
                        {remote.vertical_vane &&
                            <Shot
                                name="Vertical Vane"
                                status={remote.vertical_vane}
                                contents="Toggle"
                                onClick={(key) => {
                                    remote.vertical_vane = key
                                    //this.props.saveRemoteState(this.dataToState(remote));
                                    this.props.pushRemote(remote)
                                }}
                            />
                        }
                    </Row>
                </CardBase>
                }
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
