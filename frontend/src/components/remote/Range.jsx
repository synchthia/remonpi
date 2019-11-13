import { Button, ButtonGroup } from 'react-bootstrap';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faChevronDown, faChevronUp } from '@fortawesome/free-solid-svg-icons';
import React from 'react';

class Range extends React.Component {
    render() {
        return (
            <div className="col-sm">
                <div className="col">
                    <h1 className="display-4">{this.props.status}{this.props.suffix ? this.props.suffix : ''}</h1>
                </div>
                <div className="col">
                    <p>{this.props.name}</p>
                    <ButtonGroup>
                        <Button
                            variant="primary"
                            onClick={() => {
                                if (this.props.status > this.props.rangeFrom) {
                                    this.props.onDecrement()
                                }
                            }}
                            disabled={this.props.status <= this.props.rangeFrom}
                        >
                            <FontAwesomeIcon icon={faChevronDown} />
                        </Button>
                        <Button
                            onClick={() => {
                                if (this.props.status < this.props.rangeTo) {
                                    this.props.onIncrement()
                                }
                            }}
                            disabled={this.props.status >= this.props.rangeTo}
                        >
                            <FontAwesomeIcon icon={faChevronUp} />
                        </Button>
                    </ButtonGroup>
                </div>
            </div>
        )
    }
}

export default Range;
