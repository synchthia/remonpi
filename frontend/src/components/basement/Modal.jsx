import { Button, Modal } from 'react-bootstrap';
import React from 'react';

class ModalPopup extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            show: true
        };
    }
    render() {
        const { title, body } = this.props;
        const handleClose = () => {
            this.setState({
                show: false
            })
        }
        return (
            <div>
                <Modal show={this.state.show} onHide={handleClose}>
                    <Modal.Header closeButton={this.props.closable}>
                        <Modal.Title>{title}</Modal.Title>
                    </Modal.Header>
                    <Modal.Body>{this.props.body}</Modal.Body>
                    {this.props.type === "ERROR_MODAL" &&
                        <Modal.Footer>
                            <Button variant="primary" onClick={this.props.onReload}>
                                Refresh
                            </Button>
                        </Modal.Footer>
                    }
                </Modal>
            </div>
        )
    }
}

export default ModalPopup;
