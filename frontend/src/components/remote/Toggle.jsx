import { Button } from 'react-bootstrap';
import React from 'react';
import styled from 'styled-components';

class Toggle extends React.Component {
    render() {
        return (
            <ColBase className="col-sm">
                <div className="col">
                    <h1 className="display-4">{this.props.status}</h1>
                </div>
                <div className="col">
                    <p>{this.props.name}</p>
                    <Button
                        variant="primary"
                        type="button"
                        onClick={this.props.onClick ? () => this.props.onClick() : null}
                    >
                        {this.props.icon !== undefined ? this.props.icon : this.props.name}
                    </Button>
                </div>
            </ColBase>
        )
    }
}

const ColBase = styled.div`
    margin: 12px;
`;

export default Toggle;
